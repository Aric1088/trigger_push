package main

import (
	"context"
	"fmt"

	"github.com/containers/image/docker"
	"github.com/containers/image/types"
)

func main() {
	ref, err := docker.ParseReference("//gcr.io/aricz-notifiers-demo/notifiers/bigquery")
	if err != nil {
		panic(err)
	}
	stuff := &types.SystemContext{}
	stuff.DockerAuthConfig = &types.DockerAuthConfig{
		Username: "oauth2accesstoken",
		Password: "ya29.a0AfH6SMBKImLSzmd3faL98UyXpfCgKmK7TO9Oq6WPX1Z939tcpCXlLlS0-QssXijhvFYJupFm0XTvpa3yCCrasUgHQ-2PZiNc6MAOfmxyYtcTArg1ex3WG0SSS9GHSKonNyPQGPMna9NYnmeM51ybUozB3wH2zioA_vqjs9vlloLM",
	}
	ctx := context.Background()
	img, err := ref.NewImage(ctx, stuff)
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
