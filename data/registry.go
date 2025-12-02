//go:generate go run tools/generator.go

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

// ProtocolToVersion maps protocol version numbers to Minecraft version strings
var ProtocolToVersion = map[int32]string{
	767: "1.21.0",
	768: "1.21.4",
	769: "1.21.5",
	770: "1.21.6",
	771: "1.21.7",
	772: "1.21.8",
	773: "1.21.9",
	774: "1.21.10",
}

// SupportedVersions lists all supported Minecraft versions
var SupportedVersions = []string{
	"1.21.0", "1.21.4", "1.21.5", "1.21.6",
	"1.21.7", "1.21.8", "1.21.9", "1.21.10",
}

// Default registry for latest Minecraft version (1.21.10)
var DefaultRegistry *Registry

// registryCache stores created registries for reuse
var registryCache = make(map[string]*Registry)

func init() {
	// Initialize default registry with latest version
	DefaultRegistry = GetRegistryForVersion("1.21.10")
}

// GetRegistryForProtocol returns a Registry for the given protocol version
// If the protocol is not recognized, returns the default registry for latest version
func GetRegistryForProtocol(protocolVersion int32) *Registry {
	version, ok := ProtocolToVersion[protocolVersion]
	if !ok {
		// Unknown protocol, return latest version
		return DefaultRegistry
	}
	return GetRegistryForVersion(version)
}

// GetRegistryForVersion returns a Registry for the given Minecraft version
// Registries are cached for reuse. If the version is not supported, returns the default registry.
func GetRegistryForVersion(version string) *Registry {
	// Check cache first
	if cached, ok := registryCache[version]; ok {
		return cached
	}

	// Validate version is supported
	supported := false
	for _, v := range SupportedVersions {
		if v == version {
			supported = true
			break
		}
	}

	// If not supported, return default (latest)
	if !supported && DefaultRegistry != nil {
		return DefaultRegistry
	}

	// Create new registry
	registry := NewRegistry(version)
	loadRegistryData(registry)

	// Cache it
	registryCache[version] = registry

	return registry
}

// IsVersionSupported checks if a Minecraft version is supported
func IsVersionSupported(version string) bool {
	for _, v := range SupportedVersions {
		if v == version {
			return true
		}
	}
	return false
}

// IsProtocolSupported checks if a protocol version is supported
func IsProtocolSupported(protocolVersion int32) bool {
	_, ok := ProtocolToVersion[protocolVersion]
	return ok
}

// loadRegistryData populates a registry with game data
// Currently all versions 1.21.0-1.21.10 share the same data
// (future versions with different data can be handled here)
func loadRegistryData(registry *Registry) {
	// Load all blocks from BlockNameToID
	for name, id := range BlockNameToID {
		registry.Blocks[id] = &world.BlockInfo{
			ID:       id,
			Name:     "minecraft:" + name,
			Solid:    IsSolid(id),
			Hardness: GetHardness(id),
		}
	}

	// Load all items from ItemNameToID
	for name, id := range ItemNameToID {
		registry.Items[id] = &inventory.ItemInfo{
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
		registry.Entities[id] = &entity.EntityInfo{
			ID:     id,
			Name:   name,
			Width:  0.6, // Default width
			Height: 1.8, // Default height
		}
	}
}
