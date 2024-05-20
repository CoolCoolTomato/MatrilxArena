package api

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

func GetImageList() (images []image.Summary, err error){
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	
	ctx := context.Background()
	images, err = cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, err
	}
	
	return images, err
}

func PullImage(imageName string) (err error){
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return  err
	}
	defer cli.Close()
	
	ctx := context.Background()
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return err
	}
	defer out.Close()
	
	return err
}

func PullImageWithAuth(imageName string, username string, password string) (err error){
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()
	
	ctx := context.Background()
	authConfig := registry.AuthConfig{
		Username: username,
		Password: password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	out, err := cli.ImagePull(ctx, imageName, image.PullOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}
	defer out.Close()
	
	return err
}

func RemoveImage(imageID string) (err error){
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()
	
	ctx := context.Background()
	_, err = cli.ImageRemove(ctx, imageID, image.RemoveOptions{Force: true, PruneChildren: true})
	if err != nil {
		return err
	}
	return err
}