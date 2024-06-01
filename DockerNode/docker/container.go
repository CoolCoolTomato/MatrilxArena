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

func CreateContainer(imageID string) (containerID string, err error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return containerID, err
	}
	defer cli.Close()
	
	image, err := GetImage(imageID)
	if err != nil {
		return containerID, err
	}
	
	exposedPorts := image.Config.ExposedPorts
	randomPorts, err := freeport.GetFreePorts(len(exposedPorts))
	if err != nil {
		return containerID, err
	}

	portBindings := nat.PortMap{}
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
	
	ctx := context.Background()
	res, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        imageID,
		ExposedPorts: exposedPorts,
	}, &container.HostConfig{
		PortBindings: portBindings,
	}, nil, nil, "")
	if err != nil {
		return containerID, err
	}

	return res.ID, err
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
