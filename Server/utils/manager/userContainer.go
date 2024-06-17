package manager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/CoolCoolTomato/MatrilxArena/Server/database"
	"strconv"
	"time"
)

type PortMap struct {
    PortProtocol string
    Link         string
}

type ContainerInfo struct {
	DockerNodeContainerID string                   `json:"DockerNodeContainerID"`
	DockerNodeID          uint                     `json:"DockerNodeID"`
    ChallengeID           uint                     `json:"ChallengeID"`
	RemainingTime         time.Duration            `json:"RemainingTime"`
    PortMaps              []PortMap                 `json:"PortMaps"`
}

func GetUserContainers(username string) ([]ContainerInfo, error) {
	ctx := context.Background()
	userKey := "user:" + username + ":containers"

	dockerNodeContainerIDs, err := database.GetRedis().SMembers(ctx, userKey).Result()
	if err != nil {
		return nil, err
	}

	var containers []ContainerInfo
	for _, dockerNodeContainerID := range dockerNodeContainerIDs {
		containerKey := "container:" + dockerNodeContainerID

		exists, err := database.GetRedis().Exists(ctx, containerKey).Result()
		if err != nil {
			return nil, err
		}

		if exists == 0 {
			if err := database.GetRedis().SRem(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
				return nil, err
			}
			continue
		}

		dockerNodeIDStr, err := database.GetRedis().HGet(ctx, containerKey, "DockerNodeID").Result()
		if err != nil {
			return nil, err
		}

		dockerNodeID, err := strconv.ParseUint(dockerNodeIDStr, 10, 64)
		if err != nil {
			return nil, err
		}

		ttl, err := database.GetRedis().TTL(ctx, containerKey).Result()
		if err != nil {
			return nil, err
		}

		challengeIDStr, err := database.GetRedis().HGet(ctx, containerKey, "ChallengeID").Result()
		if err != nil {
			return nil, err
		}

		challengeID, err := strconv.ParseUint(challengeIDStr, 10, 64)
		if err != nil {
			return nil, err
		}

        portMapsStr, err := database.GetRedis().HGet(ctx, containerKey, "PortMaps").Result()
        if err != nil {
            return nil, err
        }

        var portMaps []PortMap
        if err := json.Unmarshal([]byte(portMapsStr), &portMaps); err != nil {
            return nil, err
        }

		containers = append(containers, ContainerInfo{
			DockerNodeContainerID: dockerNodeContainerID,
			DockerNodeID:          uint(dockerNodeID),
            ChallengeID:           uint(challengeID),
			RemainingTime:         ttl,
            PortMaps:              portMaps,
		})
	}

	return containers, nil
}

func AddUserContainer(username string, dockerNodeContainerID string, portMaps []PortMap, dockerNodeID uint, challengeID uint) error {
	ctx := context.Background()
	userKey := "user:" + username + ":containers"
	containerKey := "container:" + dockerNodeContainerID

	if err := database.GetRedis().SAdd(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
		return err
	}

	if err := database.GetRedis().HSet(ctx, containerKey, "DockerNodeID", strconv.FormatUint(uint64(dockerNodeID), 10)).Err(); err != nil {
		return err
	}

	if err := database.GetRedis().HSet(ctx, containerKey, "ChallengeID", strconv.FormatUint(uint64(challengeID), 10)).Err(); err != nil {
		return err
	}

    portMapsBytes, err := json.Marshal(portMaps)
    if err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, containerKey, "PortMaps", string(portMapsBytes)).Err(); err != nil {
        return err
    }

	if err := database.GetRedis().Expire(ctx, containerKey, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func DeleteUserContainer(username, dockerNodeContainerID string) error {
	ctx := context.Background()
	userKey := "user:" + username + ":containers"
	containerKey := "container:" + dockerNodeContainerID

	if err := database.GetRedis().SRem(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
		return err
	}
	if err := database.GetRedis().Del(ctx, containerKey).Err(); err != nil {
		return err
	}
	return nil
}

func ResetUserContainerTime(dockerNodeContainerID string) error {
	ctx := context.Background()
	containerKey := "container:" + dockerNodeContainerID

	ttl, err := database.GetRedis().TTL(ctx, containerKey).Result()
	if err != nil {
		return err
	}

	if ttl > 10*time.Minute {
		return fmt.Errorf("you can only extend the container life when it has less than 10 minutes remaining")
	}

	if err := database.GetRedis().Expire(ctx, containerKey, time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
