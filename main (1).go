package main

import (
	"context"
	"fmt"

	"github.com/containers/image/docker"
)

func main() {
	ref, err := docker.ParseReference("//gcr.io/aricz-notifiers-demo/notifiers/bigquery")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	img, err := ref.NewImage(ctx)
	if err != nil {
		panic(err)
	}
	defer img.Close()
	b, _, err := img.Manifest(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(b))
}
