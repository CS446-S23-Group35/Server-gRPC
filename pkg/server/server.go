package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/CS446-S23-Group35/Server-gRPC/pkg/inventory"
	"github.com/CS446-S23-Group35/Server-gRPC/pkg/recipes"
	"github.com/CS446-S23-Group35/Server-gRPC/wastenot/wastenotconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func NewServer(addr string) *http.Server {
	mux := http.NewServeMux()

	invServ := inventory.NewServer()
	invPath, invHandler := wastenotconnect.NewInventoryServiceHandler(invServ)
	mux.Handle(invPath, invHandler)

	recipeServ := recipes.NewServer()
	recipePath, recipeHandler := wastenotconnect.NewRecipeServiceHandler(recipeServ)
	mux.Handle(recipePath, recipeHandler)

	fmt.Printf("Services Available:\n%s\n%s\n", invPath, recipePath)
	return &http.Server{
		Addr:              addr,
		Handler:           h2c.NewHandler(mux, &http2.Server{}),
		ReadHeaderTimeout: 5 * time.Second,
	}
}
