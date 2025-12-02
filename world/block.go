package world

// Block represents a Minecraft block
type Block struct {
	ID    int    // Block ID (numeric)
	State int    // Block state
	Name  string // Block name (e.g., "minecraft:stone")
}

// IsAir returns true if this block is air
func (b *Block) IsAir() bool {
	return b.ID == 0 || b.Name == "minecraft:air"
}

// BlockInfo contains additional metadata about a block type
type BlockInfo struct {
	ID       int     // Block ID
	Name     string  // Block name
	Solid    bool    // Whether the block is solid
	Hardness float64 // Mining hardness
}
