# prismarine-go Design Document

## Overview

prismarine-go is Layer 3 (L3) in the 5-layer go-mc bot architecture. It provides **pure, protocol-agnostic data models** for Minecraft game state.

**Critical Principle**: Zero protocol dependencies. All types must work across Minecraft versions without coupling to packet structures.

## Architecture Decisions

### 1. Vector Types

Use simple, self-contained vector types instead of external libraries:

```go
// Position - Integer block coordinates
type Position struct { X, Y, Z int }

// Vec3d - Floating-point 3D vector (entity positions, velocity)
type Vec3d struct { X, Y, Z float64 }

// Vec2 - 2D vector (rotation: yaw, pitch)
type Vec2 struct { X, Y float32 }
```

**Rationale**:
- No external dependencies (mgl64 is L2 concern)
- Simple serialization
- Protocol version independent

### 2. Entity Model

```go
type Entity struct {
    EID      int32      // Entity ID
    UUID     [16]byte   // Entity UUID
    Type     int32      // Entity type ID
    Position Vec3d      // Current position
    Velocity Vec3d      // Movement velocity
    Rotation Vec2       // Yaw, Pitch
    Metadata map[uint8]interface{}  // Generic metadata
    Equipment map[int8]*ItemStack   // Equipment slots
}
```

**Key Design Choices**:
- `Metadata` uses `interface{}` to support future metadata types
- `Equipment` maps slot → item (protocol-agnostic)
- Rotation uses Vec2 (not separate fields)

### 3. World Model

```go
type World interface {
    GetBlock(pos Position) (*Block, error)
    SetBlock(pos Position, block *Block) error
    GetChunk(x, z int) (*Chunk, error)
}

type Chunk struct {
    X, Z     int
    Sections []*ChunkSection  // Y sections (typically 24)
    Biomes   []int            // Biome data
}

type ChunkSection struct {
    Y      int
    Blocks []Block  // 16x16x16 = 4096 blocks
}

type Block struct {
    ID    int     // Block ID
    State int     // Block state
    Name  string  // Block name (e.g., "minecraft:stone")
}
```

**Design Rationale**:
- `World` is interface → implementations can use different storage
- Chunks store Y sections (handles different world heights)
- Block has both ID and Name (supports lookups)

### 4. Inventory Model

```go
type ItemStack struct {
    Item  Item
    Count int
    NBT   map[string]interface{}  // NBT data (simplified)
}

type Item struct {
    ID          int
    Name        string
    DisplayName string
    StackSize   int
}

type Inventory struct {
    Slots []ItemStack
    Size  int
}

type Window struct {
    ID    int
    Type  string
    Title string
    Slots []ItemStack
}
```

**Key Points**:
- `NBT` uses generic map (avoids protocol NBT types)
- Separate `Item` (definition) vs `ItemStack` (instance)
- `Window` for containers (chest, furnace, etc.)

### 5. Chat Model

```go
type Component struct {
    Text       string
    Bold       bool
    Italic     bool
    Underlined bool
    Color      string
    ClickEvent *ClickEvent
    HoverEvent *HoverEvent
    Extra      []Component  // Nested components
}

type Message struct {
    Components []Component
}

func (m *Message) ToJSON() string
func (m *Message) ToPlainText() string
func ParseJSON(json string) (*Message, error)
```

**Design Goals**:
- Matches Minecraft JSON chat format
- Can parse/serialize JSON
- Protocol-independent

### 6. Physics/Collision

```go
type AABB struct {
    MinX, MinY, MinZ float64
    MaxX, MaxY, MaxZ float64
}

func (a *AABB) Intersects(b *AABB) bool
func (a *AABB) Contains(pos Vec3d) bool
func (a *AABB) Expand(amount float64) *AABB
```

**Purpose**:
- Collision detection for pathfinding
- Entity bounding boxes

## Data Registry

`data/` package provides lookup tables:

```go
type Registry struct {
    Blocks   map[int]*BlockInfo
    Items    map[int]*ItemInfo
    Entities map[int]*EntityInfo
}

type BlockInfo struct {
    ID        int
    Name      string
    Solid     bool
    Hardness  float64
}

type ItemInfo struct {
    ID         int
    Name       string
    StackSize  int
}

type EntityInfo struct {
    ID     int
    Name   string
    Width  float64
    Height float64
}

func LoadRegistry(version string) (*Registry, error)
```

**Data Sources**:
- Extracted from 全能bot/internal/pkg/gamedata/
- Loaded from JSON files at runtime
- Version-specific registries supported

## Migration from nalago-mc

### What to Extract:

**From `pkg/game/world/entity.go`**:
- ✅ Entity struct → prismarine-go/entity/entity.go
- ✅ Use Vec3d/Vec2 instead of mgl64
- ✅ Metadata → generic map

**From `pkg/game/world/world.go`**:
- ✅ Chunk storage concept → prismarine-go/world/chunk.go
- ❌ Packet handlers stay in nalago-mc (L2)
- ✅ Block representation → prismarine-go/world/block.go

### What Stays in nalago-mc:

- All packet handling (`bot.AddHandler`)
- Protocol conversions
- Event emission
- Client references

## Interface Contracts

### L2 → L3 Data Flow

nalago-mc (L2) will:
1. Receive protocol packets
2. Convert to L3 data types
3. Update L3 models (World, Entity lists)
4. Emit events with L3 types

```go
// Example: Entity spawn packet → L3 Entity
func handleAddEntity(p *cp.AddEntity) {
    entity := &prismarine.Entity{
        EID: p.ID,
        UUID: p.UUID,
        Type: int32(p.Type),
        Position: prismarine.Vec3d{X: p.X, Y: p.Y, Z: p.Z},
        // ...
    }
    world.AddEntity(entity)
}
```

### L3 → L4 Usage

mineflayer-go (L4) will:
1. Call L3 interfaces (World, Inventory, etc.)
2. Access L3 data types directly
3. Never import protocol types

```go
// L4 example
func (b *Bot) GetNearbyBlocks(radius int) []prismarine.Block {
    pos := b.position
    blocks := []prismarine.Block{}
    // Query L3 World interface
    for x := -radius; x <= radius; x++ {
        for z := -radius; z <= radius; z++ {
            block, _ := b.world.GetBlock(prismarine.Position{
                X: int(pos.X) + x,
                Y: int(pos.Y),
                Z: int(pos.Z) + z,
            })
            blocks = append(blocks, *block)
        }
    }
    return blocks
}
```

## Testing Strategy

Each package will have:
- Unit tests for core types
- Serialization/deserialization tests
- Conversion tests (protocol → L3)

Example:
```go
// entity/entity_test.go
func TestEntityPosition(t *testing.T) {
    e := &Entity{Position: Vec3d{X: 1.5, Y: 64, Z: 3.7}}
    assert.Equal(t, 1.5, e.Position.X)
}
```

## Implementation Order

Phase 2 tasks:
1. ✅ World models (position, block, chunk, world) - **NEXT**
2. Entity models (entity, player, metadata, types)
3. Inventory models (inventory, item, slot, window)
4. Chat models (message, component, style)
5. Extract gamedata from 全能bot
6. Update nalago-mc imports

Estimated: 20-30 hours total

## Success Criteria

Phase 2 complete when:
- [ ] prismarine-go compiles independently
- [ ] No imports from go-mc or nalago-mc protocol packages
- [ ] All core types have tests
- [ ] nalago-mc imports and uses L3 types
- [ ] Game data registry loads correctly
