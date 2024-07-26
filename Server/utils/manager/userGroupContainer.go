package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"strconv"
	"time"
)

type GroupContainerInfo struct {
	DockerNodeContainerID string            `json:"DockerNodeContainerID"`
	DockerNodeID          uint              `json:"DockerNodeID"`
    GroupChallengeID      uint              `json:"GroupChallengeID"`
	RemainingTime         time.Duration     `json:"RemainingTime"`
    PortMaps              []PortMap         `json:"PortMaps"`
    Flag                  string            `json:"Flag"`
}

func GetUserGroupContainerList(username string) ([]GroupContainerInfo, error) {
	ctx := context.Background()
	userKey := "user:" + username + ":groupContainers"

	dockerNodeContainerIDs, err := database.GetRedis().SMembers(ctx, userKey).Result()
	if err != nil {
		return nil, err
	}

	var containerList []GroupContainerInfo
	for _, dockerNodeContainerID := range dockerNodeContainerIDs {
		groupContainerKey := "groupContainer:" + dockerNodeContainerID

		exists, err := database.GetRedis().Exists(ctx, groupContainerKey).Result()
		if err != nil {
			return nil, err
		}

		if exists == 0 {
			if err := database.GetRedis().SRem(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
				return nil, err
			}
			continue
		}

		dockerNodeIDStr, err := database.GetRedis().HGet(ctx, groupContainerKey, "DockerNodeID").Result()
		if err != nil {
			return nil, err
		}

		dockerNodeID, err := strconv.ParseUint(dockerNodeIDStr, 10, 64)
		if err != nil {
			return nil, err
		}

		ttl, err := database.GetRedis().TTL(ctx, groupContainerKey).Result()
		if err != nil {
			return nil, err
		}

		groupChallengeIDStr, err := database.GetRedis().HGet(ctx, groupContainerKey, "GroupChallengeID").Result()
		if err != nil {
			return nil, err
		}

		groupChallengeID, err := strconv.ParseUint(groupChallengeIDStr, 10, 64)
		if err != nil {
			return nil, err
		}

        portMapsStr, err := database.GetRedis().HGet(ctx, groupContainerKey, "PortMaps").Result()
        if err != nil {
            return nil, err
        }

        var portMaps []PortMap
        if err := json.Unmarshal([]byte(portMapsStr), &portMaps); err != nil {
            return nil, err
        }

		flag, err := database.GetRedis().HGet(ctx, groupContainerKey, "Flag").Result()
		if err != nil {
			return nil, err
		}

		containerList = append(containerList, GroupContainerInfo{
			DockerNodeContainerID: dockerNodeContainerID,
			DockerNodeID:          uint(dockerNodeID),
            GroupChallengeID:      uint(groupChallengeID),
			RemainingTime:         ttl,
            PortMaps:              portMaps,
            Flag:                  flag,
		})
	}

	return containerList, nil
}

func AddUserGroupContainer(username string, dockerNodeContainerID string, portMaps []PortMap, dockerNodeID uint, groupChallengeID uint, flag string) error {
	ctx := context.Background()
	userKey := "user:" + username + ":groupContainers"
	groupContainerKey := "groupContainer:" + dockerNodeContainerID

	if err := database.GetRedis().SAdd(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
		return err
	}

	if err := database.GetRedis().HSet(ctx, groupContainerKey, "DockerNodeID", strconv.FormatUint(uint64(dockerNodeID), 10)).Err(); err != nil {
		return err
	}

	if err := database.GetRedis().HSet(ctx, groupContainerKey, "GroupChallengeID", strconv.FormatUint(uint64(groupChallengeID), 10)).Err(); err != nil {
		return err
	}

    portMapsBytes, err := json.Marshal(portMaps)
    if err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, groupContainerKey, "PortMaps", string(portMapsBytes)).Err(); err != nil {
        return err
    }

	if err := database.GetRedis().HSet(ctx, groupContainerKey, "Flag", flag).Err(); err != nil {
		return err
	}

	if err := database.GetRedis().Expire(ctx, groupContainerKey, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func DeleteUserGroupContainer(username, dockerNodeContainerID string) error {
	ctx := context.Background()
	userKey := "user:" + username + ":groupContainers"
	groupContainerKey := "groupContainer:" + dockerNodeContainerID

	if err := database.GetRedis().SRem(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
		return err
	}
	if err := database.GetRedis().Del(ctx, groupContainerKey).Err(); err != nil {
		return err
	}
	return nil
}

func ResetUserGroupContainerTime(dockerNodeContainerID string) error {
	ctx := context.Background()
	groupContainerKey := "groupContainer:" + dockerNodeContainerID

	ttl, err := database.GetRedis().TTL(ctx, groupContainerKey).Result()
	if err != nil {
		return err
	}

	if ttl > 10*time.Minute {
		return fmt.Errorf("remaining time is greater than 10 minutes")
	}

	if err := database.GetRedis().Expire(ctx, groupContainerKey, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
