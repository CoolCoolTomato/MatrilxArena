package main

import (
	"fmt"
	"main/api"
)

func main() {
	test, _ := api.GetImageList()
	for _, i := range test {
		fmt.Println(i.RepoTags)
		fmt.Println(i.ID)
	}
}
