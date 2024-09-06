package containerManager

import "fmt"

type StandardContainerManager struct{}

func (m StandardContainerManager) GetUserKey(username string) string {
    return fmt.Sprintf("user:%s:containers", username)
}

func (m StandardContainerManager) GetContainerKey(dockerNodeContainerID string) string {
    return fmt.Sprintf("container:%s", dockerNodeContainerID)
}

func (m StandardContainerManager) GetChallengeIDField() string {
    return "ChallengeID"
}
