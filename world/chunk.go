package world

// ChunkPos represents a chunk position (X, Z coordinates)
type ChunkPos struct {
	X, Z int
}

// Chunk represents a 16x16 column of the world
type Chunk struct {
	X, Z     int             // Chunk coordinates
	Sections []*ChunkSection // Y sections (typically 24 sections, -4 to 19)
	Biomes   []int           // Biome data
}

// ChunkSection represents a 16x16x16 section of blocks
type ChunkSection struct {
	Y      int     // Y coordinate of this section (section index * 16)
	Blocks []Block // 16x16x16 = 4096 blocks
}

// NewChunk creates a new empty chunk
func NewChunk(x, z int, sectionCount int) *Chunk {
	sections := make([]*ChunkSection, sectionCount)
	for i := range sections {
		sections[i] = &ChunkSection{
			Y:      i*16 - 64, // Minecraft world starts at Y=-64
			Blocks: make([]Block, 4096),
		}
	}
	return &Chunk{
		X:        x,
		Z:        z,
		Sections: sections,
		Biomes:   make([]int, 256), // 16x16 biomes per chunk
	}
}

// GetBlock gets a block at local chunk coordinates (0-15, y, 0-15)
func (c *Chunk) GetBlock(x, y, z int) *Block {
	sectionY := (y + 64) / 16 // Convert world Y to section index
	if sectionY < 0 || sectionY >= len(c.Sections) {
		return nil
	}

	section := c.Sections[sectionY]
	if section == nil {
		return nil
	}

	localY := y - section.Y
	index := (localY * 16 * 16) + (z * 16) + x

	if index < 0 || index >= len(section.Blocks) {
		return nil
	}

	return &section.Blocks[index]
}

// SetBlock sets a block at local chunk coordinates (0-15, y, 0-15)
func (c *Chunk) SetBlock(x, y, z int, block Block) {
	sectionY := (y + 64) / 16
	if sectionY < 0 || sectionY >= len(c.Sections) {
		return
	}

	section := c.Sections[sectionY]
	if section == nil {
		section = &ChunkSection{
			Y:      sectionY*16 - 64,
			Blocks: make([]Block, 4096),
		}
		c.Sections[sectionY] = section
	}

	localY := y - section.Y
	index := (localY * 16 * 16) + (z * 16) + x

	if index >= 0 && index < len(section.Blocks) {
		section.Blocks[index] = block
	}
}
