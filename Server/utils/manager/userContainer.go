package manager

import (
    "context"
    "fmt"
    "github.com/CoolCoolTomato/MatrilxArena/Server/database"
    "strconv"
    "time"
)

type ContainerInfo struct {
    ContainerID string        `json:"container_id"`
    DockerNodeID uint       `json:"docker_node_id"`
    RemainingTime time.Duration `json:"remaining_time"`
}

func GetUserContainers(username string) ([]ContainerInfo, error) {
    ctx := context.Background()
    userKey := "user:" + username + ":containers"

    containerIDs, err := database.GetRedis().SMembers(ctx, userKey).Result()
    if err != nil {
        return nil, err
    }

    var containers []ContainerInfo
    for _, containerID := range containerIDs {
        containerKey := "container:" + containerID

        exists, err := database.GetRedis().Exists(ctx, containerKey).Result()
        if err != nil {
            return nil, err
        }

        if exists == 0 {
            if err := database.GetRedis().SRem(ctx, userKey, containerID).Err(); err != nil {
                return nil, err
            }
            continue
        }

        dockerNodeIDStr, err := database.GetRedis().HGet(ctx, containerKey, "docker_node_id").Result()
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

        containers = append(containers, ContainerInfo{
            ContainerID:   containerID,
            DockerNodeID:  uint(dockerNodeID),
            RemainingTime: ttl,
        })
    }

    return containers, nil
}

func AddUserContainer(username, containerID string, dockerNodeID uint) error {
    ctx := context.Background()
    userKey := "user:" + username + ":containers"
    containerKey := "container:" + containerID

    if err := database.GetRedis().SAdd(ctx, userKey, containerID).Err(); err != nil {
        return err
    }

    if err := database.GetRedis().HSet(ctx, containerKey, "docker_node_id", strconv.FormatUint(uint64(dockerNodeID), 10)).Err(); err != nil {
        return err
    }

    if err := database.GetRedis().Expire(ctx, containerKey, time.Hour).Err(); err != nil {
        return err
    }

    return nil
}

func DeleteUserContainer(username, containerID string) error {
    ctx := context.Background()
    userKey := "user:" + username + ":containers"
    containerKey := "container:" + containerID

    if err := database.GetRedis().SRem(ctx, userKey, containerID).Err(); err != nil {
        return err
    }
    if err := database.GetRedis().Del(ctx, containerKey).Err(); err != nil {
        return err
    }
    return nil
}

func ResetUserContainerTime(containerID string) error {
    ctx := context.Background()
    containerKey := "container:" + containerID

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
