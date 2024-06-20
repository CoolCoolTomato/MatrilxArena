package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/phayes/freeport"
)

func GetContainerList() (containers []types.Container, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return containers, err
	}
	defer cli.Close()

	ctx := context.Background()
	containers, err = cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		return containers, err
	}

	return containers, err
}

func GetContainer(containerID string) (containerJSON types.ContainerJSON, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return containerJSON, err
	}
	defer cli.Close()

	ctx := context.Background()
	containerJSON, err = cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return containerJSON, err
	}

	return containerJSON, err
}

func CreateContainer(imageID string, specifiedPorts []string) (containerID string, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return containerID, err
	}
	defer cli.Close()

	image, err := GetImage(imageID)
	if err != nil {
		return containerID, err
	}

	exposedPorts := nat.PortSet{}
	portBindings := nat.PortMap{}

	var randomPorts []int
	if len(specifiedPorts) > 0 {
		randomPorts, err = freeport.GetFreePorts(len(specifiedPorts))
		if err != nil {
			return containerID, err
		}

		for i, port := range specifiedPorts {
			exposedPort := nat.Port(port)
			exposedPorts[exposedPort] = struct{}{}
			portBindings[exposedPort] = []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", randomPorts[i]),
				},
			}
		}
	} else {
		exposedPorts = image.Config.ExposedPorts
		randomPorts, err = freeport.GetFreePorts(len(exposedPorts))
		if err != nil {
			return containerID, err
		}
		i := 0
		for port := range exposedPorts {
			portBindings[port] = []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: fmt.Sprintf("%d", randomPorts[i]),
				},
			}
			i++
		}
	}

	config := &container.Config{
		Image:        imageID,
		ExposedPorts: exposedPorts,
	}

	ctx := context.Background()
	res, err := cli.ContainerCreate(ctx, config, &container.HostConfig{
		PortBindings: portBindings,
	}, nil, nil, "")
	if err != nil {
		return containerID, err
	}
	containerID = res.ID

	return containerID, err
}

func StartContainer(containerID string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()
	err = cli.ContainerStart(ctx, containerID, container.StartOptions{})
	if err != nil {
		return err
	}

	return err
}

func StopContainer(containerID string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()
	timeout := 10
	stopOptions := container.StopOptions{Timeout: &timeout}
	err = cli.ContainerStop(ctx, containerID, stopOptions)
	if err != nil {
		return err
	}

	return err
}

func RemoveContainer(containerID string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx := context.Background()
	removeOptions := container.RemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         true,
	}
	err = cli.ContainerRemove(ctx, containerID, removeOptions)
	if err != nil {
		return err
	}

	return err
}

func ExecuteCommand(containerID string, command []string) (err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	execConfig := types.ExecConfig{
		Cmd:          command,
		User:         "root",
		AttachStdout: true,
		AttachStderr: true,
	}

	ctx := context.Background()
	res, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return err
	}

	err = cli.ContainerExecStart(ctx, res.ID, types.ExecStartCheck{})
	if err != nil {
		return err
	}

	inspectResp, err := cli.ContainerExecInspect(ctx, res.ID)
	if err != nil {
		return err
	}

	if inspectResp.ExitCode != 0 {
		return fmt.Errorf("command failed with exit code %d", inspectResp.ExitCode)
	}

	return err
}
