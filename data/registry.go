package data

import (
	"github.com/konjacbot/prismarine-go/entity"
	"github.com/konjacbot/prismarine-go/inventory"
	"github.com/konjacbot/prismarine-go/world"
)

// Registry holds all game data (blocks, items, entities)
type Registry struct {
	Blocks   map[int]*world.BlockInfo
	Items    map[int]*inventory.ItemInfo
	Entities map[int32]*entity.EntityInfo
	Version  string
}

// NewRegistry creates a new empty registry
func NewRegistry(version string) *Registry {
	return &Registry{
		Blocks:   make(map[int]*world.BlockInfo),
		Items:    make(map[int]*inventory.ItemInfo),
		Entities: make(map[int32]*entity.EntityInfo),
		Version:  version,
	}
}

// GetBlock gets block info by ID
func (r *Registry) GetBlock(id int) (*world.BlockInfo, bool) {
	info, ok := r.Blocks[id]
	return info, ok
}

// GetItem gets item info by ID
func (r *Registry) GetItem(id int) (*inventory.ItemInfo, bool) {
	info, ok := r.Items[id]
	return info, ok
}

// GetEntity gets entity info by type ID
func (r *Registry) GetEntity(typeID int32) (*entity.EntityInfo, bool) {
	info, ok := r.Entities[typeID]
	return info, ok
}

// Default registry for Minecraft 1.21
var DefaultRegistry *Registry

func init() {
	DefaultRegistry = NewRegistry("1.21")
	// TODO: Load from JSON files or embedded data
	initializeDefaultRegistry()
}

func initializeDefaultRegistry() {
	// Load all blocks from BlockNameToID
	for name, id := range BlockNameToID {
		DefaultRegistry.Blocks[id] = &world.BlockInfo{
			ID:       id,
			Name:     "minecraft:" + name,
			Solid:    IsSolid(id),
			Hardness: GetHardness(id),
		}
	}

	// Load all items from ItemNameToID
	for name, id := range ItemNameToID {
		DefaultRegistry.Items[id] = &inventory.ItemInfo{
			ID:         id,
			Name:       "minecraft:" + name,
			StackSize:  GetMaxStackSize(id),
			Durability: GetMaxDurability(id),
		}
	}

	// Load all entities from EntityNameToID
	// Note: We don't have width/height data in the extracted gamedata yet,
	// so we'll use default values. This can be enhanced later.
	for name, id := range EntityNameToID {
		DefaultRegistry.Entities[id] = &entity.EntityInfo{
			ID:     id,
			Name:   name,
			Width:  0.6, // Default width
			Height: 1.8, // Default height
		}
	}
}
