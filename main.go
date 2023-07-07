package main

import (
	"context"
	"fmt"

	"github.com/CS446-S23-Group35/Server-gRPC/pkg/server"
)

func main() {
	srv := server.NewServer(":8080")

	go func() {
		fmt.Scanf("%s", "")
		srv.Shutdown(context.Background())
	}()

	fmt.Println("Server listening on :8080")
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
