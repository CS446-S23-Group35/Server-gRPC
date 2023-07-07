package inventory

import "github.com/CS446-S23-Group35/Server-gRPC/wastenot"

type InventoryQuerier interface {
	GetInventory(userId string) ([]*wastenot.Item, error)
	AddItem(userId string, item *wastenot.Item) error
	RemoveItem(userId string, itemName string) error
}
