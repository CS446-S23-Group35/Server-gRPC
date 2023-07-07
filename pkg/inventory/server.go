package inventory

import (
	"context"
	"fmt"

	"github.com/CS446-S23-Group35/Server-gRPC/wastenot"
	"github.com/CS446-S23-Group35/Server-gRPC/wastenot/wastenotconnect"
	connect "github.com/bufbuild/connect-go"
)

var _ wastenotconnect.InventoryServiceHandler = (*Server)(nil)

type Server struct {
	wastenotconnect.UnimplementedInventoryServiceHandler
	itemMap map[string][]*wastenot.Item
}

func NewServer() *Server {
	return &Server{
		itemMap: make(map[string][]*wastenot.Item),
	}
}

func (s *Server) GetInventory(ctx context.Context, req *connect.Request[wastenot.GetInventoryRequest]) (*connect.Response[wastenot.GetInventoryResponse], error) {
	return &connect.Response[wastenot.GetInventoryResponse]{
		Msg: &wastenot.GetInventoryResponse{
			Items: s.itemMap[req.Msg.UserId],
		},
	}, nil
}

func (s *Server) AddItem(ctx context.Context, req *connect.Request[wastenot.AddItemRequest]) (*connect.Response[wastenot.AddItemResponse], error) {
	s.itemMap[req.Msg.UserId] = append(s.itemMap[req.Msg.UserId], req.Msg.Item)
	return &connect.Response[wastenot.AddItemResponse]{
		Msg: &wastenot.AddItemResponse{},
	}, nil
}

func (s *Server) RemoveItem(ctx context.Context, req *connect.Request[wastenot.RemoveItemRequest]) (*connect.Response[wastenot.RemoveItemResponse], error) {
	found := false
	for i, item := range s.itemMap[req.Msg.UserId] {
		if item.Name == req.Msg.ItemName {
			s.itemMap[req.Msg.UserId] = append(s.itemMap[req.Msg.UserId][:i], s.itemMap[req.Msg.UserId][i+1:]...)
			found = true
			break
		}
	}

	var err error
	if !found {
		err = connect.NewError(connect.CodeNotFound, fmt.Errorf("Item not found"))
	}

	return &connect.Response[wastenot.RemoveItemResponse]{
		Msg: &wastenot.RemoveItemResponse{},
	}, err
}
