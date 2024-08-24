package manager

import "fmt"

type GroupContainerManager struct{}

func (m GroupContainerManager) GetUserKey(username string) string {
    return fmt.Sprintf("user:%s:groupContainers", username)
}

func (m GroupContainerManager) GetContainerKey(dockerNodeContainerID string) string {
    return fmt.Sprintf("groupContainer:%s", dockerNodeContainerID)
}

func (m GroupContainerManager) GetChallengeIDField() string {
    return "GroupChallengeID"
}
