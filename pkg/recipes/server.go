package recipes

import (
	"context"
	"encoding/json"
	"math/rand"
	"os"
	"strings"

	"github.com/CS446-S23-Group35/Server-gRPC/wastenot"
	"github.com/CS446-S23-Group35/Server-gRPC/wastenot/wastenotconnect"
	"github.com/bufbuild/connect-go"
)

type Server struct {
	wastenotconnect.UnimplementedRecipeServiceHandler
	recipes []*wastenot.Recipe
}

func NewServer() *Server {
	recipes := make([]*wastenot.Recipe, 0)
	file, err := os.Open("res/recipes.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&recipes)
	if err != nil {
		panic(err)
	}
	return &Server{
		recipes: recipes,
	}
}

func (s *Server) randomRecipes() []*wastenot.Recipe {
	indices := rand.Perm(len(s.recipes))
	numRecip := rand.Intn(4) + 2
	recip := make([]*wastenot.Recipe, numRecip)
	for i := 0; i < numRecip; i++ {
		recip[i] = s.recipes[indices[i]]
	}
	return recip
}

func (s *Server) SearchRecipesByName(ctx context.Context, req *connect.Request[wastenot.SearchRecipesByNameRequest]) (*connect.Response[wastenot.SearchRecipesByNameResponse], error) {
	arr := make([]*wastenot.Recipe, 0, len(s.recipes))
	for _, r := range s.recipes {
		if strings.HasPrefix(r.Name, req.Msg.Query) {
			arr = append(arr, r)
		}
	}
	return &connect.Response[wastenot.SearchRecipesByNameResponse]{
		Msg: &wastenot.SearchRecipesByNameResponse{
			Recipes: arr,
		},
	}, nil
}

func (s *Server) SearchRecipesByInventory(ctx context.Context, req *connect.Request[wastenot.SearchRecipesByInventoryRequest]) (*connect.Response[wastenot.SearchRecipesByInventoryResponse], error) {
	return &connect.Response[wastenot.SearchRecipesByInventoryResponse]{
		Msg: &wastenot.SearchRecipesByInventoryResponse{
			Recipes: s.randomRecipes(),
		},
	}, nil
}

func (s *Server) SearchRecipesByUserInventory(ctx context.Context, req *connect.Request[wastenot.SearchRecipesByInventoryRequest]) (*connect.Response[wastenot.SearchRecipesByInventoryResponse], error) {
	return &connect.Response[wastenot.SearchRecipesByInventoryResponse]{
		Msg: &wastenot.SearchRecipesByInventoryResponse{
			Recipes: s.randomRecipes(),
		},
	}, nil
}
