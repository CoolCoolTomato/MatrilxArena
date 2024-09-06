package containerManager

import "fmt"

type CTFContainerManager struct{}

func (m CTFContainerManager) GetUserKey(username string) string {
    return fmt.Sprintf("user:%s:ctfContainers", username)
}

func (m CTFContainerManager) GetContainerKey(dockerNodeContainerID string) string {
    return fmt.Sprintf("ctfContainer:%s", dockerNodeContainerID)
}

func (m CTFContainerManager) GetChallengeIDField() string {
    return "CTFChallengeID"
}
