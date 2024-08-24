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
    DockerNodeContainerID string        `json:"DockerNodeContainerID"`
    DockerNodeID          uint          `json:"DockerNodeID"`
    ChallengeID           uint          `json:"ChallengeID"`
    RemainingTime         time.Duration `json:"RemainingTime"`
    PortMaps              []PortMap     `json:"PortMaps"`
    Flag                  string        `json:"Flag"`
}

type ContainerManager interface {
    GetUserKey(username string) string
    GetContainerKey(dockerNodeContainerID string) string
    GetChallengeIDField() string
}

func GetUserContainerList(username string, manager ContainerManager) ([]ContainerInfo, error) {
    ctx := context.Background()
    userKey := manager.GetUserKey(username)

    dockerNodeContainerIDs, err := database.GetRedis().SMembers(ctx, userKey).Result()
    if err != nil {
        return nil, err
    }

    var containerList []ContainerInfo
    for _, dockerNodeContainerID := range dockerNodeContainerIDs {
        containerKey := manager.GetContainerKey(dockerNodeContainerID)

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

        containerInfo, err := getContainerInfo(ctx, containerKey, dockerNodeContainerID, manager)
        if err != nil {
            return nil, err
        }

        containerList = append(containerList, containerInfo)
    }

    return containerList, nil
}

func getContainerInfo(ctx context.Context, containerKey, dockerNodeContainerID string, manager ContainerManager) (ContainerInfo, error) {
    dockerNodeID, err := getUintField(ctx, containerKey, "DockerNodeID")
    if err != nil {
        return ContainerInfo{}, err
    }

    challengeID, err := getUintField(ctx, containerKey, manager.GetChallengeIDField())
    if err != nil {
        return ContainerInfo{}, err
    }

    ttl, err := database.GetRedis().TTL(ctx, containerKey).Result()
    if err != nil {
        return ContainerInfo{}, err
    }

    portMaps, err := getPortMaps(ctx, containerKey)
    if err != nil {
        return ContainerInfo{}, err
    }

    flag, err := database.GetRedis().HGet(ctx, containerKey, "Flag").Result()
    if err != nil {
        return ContainerInfo{}, err
    }

    return ContainerInfo{
        DockerNodeContainerID: dockerNodeContainerID,
        DockerNodeID:          dockerNodeID,
        ChallengeID:           challengeID,
        RemainingTime:         ttl,
        PortMaps:              portMaps,
        Flag:                  flag,
    }, nil
}

func getUintField(ctx context.Context, key, field string) (uint, error) {
    strValue, err := database.GetRedis().HGet(ctx, key, field).Result()
    if err != nil {
        return 0, err
    }
    value, err := strconv.ParseUint(strValue, 10, 64)
    if err != nil {
        return 0, err
    }
    return uint(value), nil
}

func getPortMaps(ctx context.Context, key string) ([]PortMap, error) {
    portMapsStr, err := database.GetRedis().HGet(ctx, key, "PortMaps").Result()
    if err != nil {
        return nil, err
    }
    var portMaps []PortMap
    if err := json.Unmarshal([]byte(portMapsStr), &portMaps); err != nil {
        return nil, err
    }
    return portMaps, nil
}

func AddUserContainer(username string, dockerNodeContainerID string, portMaps []PortMap, dockerNodeID uint, challengeID uint, flag string, manager ContainerManager) error {
    ctx := context.Background()
    userKey := manager.GetUserKey(username)
    containerKey := manager.GetContainerKey(dockerNodeContainerID)

    if err := database.GetRedis().SAdd(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
        return err
    }

    if err := setContainerFields(ctx, containerKey, dockerNodeID, challengeID, portMaps, flag, manager); err != nil {
        return err
    }

    if err := database.GetRedis().Expire(ctx, containerKey, time.Hour).Err(); err != nil {
        return err
    }

    return nil
}

func setContainerFields(ctx context.Context, key string, dockerNodeID, challengeID uint, portMaps []PortMap, flag string, manager ContainerManager) error {
    if err := database.GetRedis().HSet(ctx, key, "DockerNodeID", strconv.FormatUint(uint64(dockerNodeID), 10)).Err(); err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, key, manager.GetChallengeIDField(), strconv.FormatUint(uint64(challengeID), 10)).Err(); err != nil {
        return err
    }

    portMapsBytes, err := json.Marshal(portMaps)
    if err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, key, "PortMaps", string(portMapsBytes)).Err(); err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, key, "Flag", flag).Err(); err != nil {
        return err
    }

    return nil
}

func DeleteUserContainer(username, dockerNodeContainerID string, manager ContainerManager) error {
    ctx := context.Background()
    userKey := manager.GetUserKey(username)
    containerKey := manager.GetContainerKey(dockerNodeContainerID)

    if err := database.GetRedis().SRem(ctx, userKey, dockerNodeContainerID).Err(); err != nil {
        return err
    }
    if err := database.GetRedis().Del(ctx, containerKey).Err(); err != nil {
        return err
    }
    return nil
}

func ResetUserContainerTime(dockerNodeContainerID string, manager ContainerManager) error {
    ctx := context.Background()
    containerKey := manager.GetContainerKey(dockerNodeContainerID)

    ttl, err := database.GetRedis().TTL(ctx, containerKey).Result()
    if err != nil {
        return err
    }

    if ttl > 10*time.Minute {
        return fmt.Errorf("remaining time is greater than 10 minutes")
    }

    if err := database.GetRedis().Expire(ctx, containerKey, time.Hour).Err(); err != nil {
        return err
    }

    return nil
}
