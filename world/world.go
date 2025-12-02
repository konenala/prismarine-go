package world

import (
	"errors"
	"sync"
)

var (
	ErrChunkNotLoaded   = errors.New("chunk not loaded")
	ErrBlockOutOfBounds = errors.New("block position out of bounds")
)

// World interface represents a Minecraft world
type World interface {
	GetBlock(pos Position) (*Block, error)
	SetBlock(pos Position, block *Block) error
	GetChunk(x, z int) (*Chunk, error)
	IsChunkLoaded(x, z int) bool
}

// SimpleWorld is a basic implementation of the World interface
type SimpleWorld struct {
	chunks map[ChunkPos]*Chunk
	mu     sync.RWMutex
}

// NewSimpleWorld creates a new simple world implementation
func NewSimpleWorld() *SimpleWorld {
	return &SimpleWorld{
		chunks: make(map[ChunkPos]*Chunk),
	}
}

// GetBlock gets a block at the given world position
func (w *SimpleWorld) GetBlock(pos Position) (*Block, error) {
	chunkX := pos.X >> 4 // Divide by 16
	chunkZ := pos.Z >> 4

	localX := pos.X & 15 // Modulo 16
	localZ := pos.Z & 15

	w.mu.RLock()
	chunk, exists := w.chunks[ChunkPos{X: chunkX, Z: chunkZ}]
	w.mu.RUnlock()

	if !exists {
		return nil, ErrChunkNotLoaded
	}

	block := chunk.GetBlock(localX, pos.Y, localZ)
	if block == nil {
		return nil, ErrBlockOutOfBounds
	}

	return block, nil
}

// SetBlock sets a block at the given world position
func (w *SimpleWorld) SetBlock(pos Position, block *Block) error {
	chunkX := pos.X >> 4
	chunkZ := pos.Z >> 4

	localX := pos.X & 15
	localZ := pos.Z & 15

	w.mu.Lock()
	defer w.mu.Unlock()

	chunkPos := ChunkPos{X: chunkX, Z: chunkZ}
	chunk, exists := w.chunks[chunkPos]

	if !exists {
		// Auto-create chunk if it doesn't exist
		chunk = NewChunk(chunkX, chunkZ, 24)
		w.chunks[chunkPos] = chunk
	}

	chunk.SetBlock(localX, pos.Y, localZ, *block)
	return nil
}

// GetChunk gets a chunk at the given chunk coordinates
func (w *SimpleWorld) GetChunk(x, z int) (*Chunk, error) {
	w.mu.RLock()
	chunk, exists := w.chunks[ChunkPos{X: x, Z: z}]
	w.mu.RUnlock()

	if !exists {
		return nil, ErrChunkNotLoaded
	}

	return chunk, nil
}

// IsChunkLoaded checks if a chunk is loaded
func (w *SimpleWorld) IsChunkLoaded(x, z int) bool {
	w.mu.RLock()
	_, exists := w.chunks[ChunkPos{X: x, Z: z}]
	w.mu.RUnlock()
	return exists
}

// LoadChunk loads a chunk into the world
func (w *SimpleWorld) LoadChunk(chunk *Chunk) {
	w.mu.Lock()
	w.chunks[ChunkPos{X: chunk.X, Z: chunk.Z}] = chunk
	w.mu.Unlock()
}

// UnloadChunk unloads a chunk from the world
func (w *SimpleWorld) UnloadChunk(x, z int) {
	w.mu.Lock()
	delete(w.chunks, ChunkPos{X: x, Z: z})
	w.mu.Unlock()
}

// GetNearbyBlocks gets blocks within a radius of a position
func (w *SimpleWorld) GetNearbyBlocks(center Position, radius int) []Block {
	blocks := []Block{}

	for x := center.X - radius; x <= center.X+radius; x++ {
		for y := center.Y - radius; y <= center.Y+radius; y++ {
			for z := center.Z - radius; z <= center.Z+radius; z++ {
				block, err := w.GetBlock(Position{X: x, Y: y, Z: z})
				if err == nil {
					blocks = append(blocks, *block)
				}
			}
		}
	}

	return blocks
}
