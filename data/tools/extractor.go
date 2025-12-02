// extractor.go - One-time tool to extract existing Go data to JSON format
// Usage: go run extractor.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/konjacbot/prismarine-go/data"
)

// BlockData represents a block in JSON format
type BlockData struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Solid    bool    `json:"solid"`
	Hardness float64 `json:"hardness"`
}

// ItemData represents an item in JSON format
type ItemData struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	StackSize  int    `json:"stackSize"`
	Durability int    `json:"durability"`
}

// EntityData represents an entity in JSON format
type EntityData struct {
	ID     int32   `json:"id"`
	Name   string  `json:"name"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func main() {
	versions := []string{"1.21.0", "1.21.4", "1.21.8", "1.21.10"}

	for _, version := range versions {
		fmt.Printf("Extracting data for Minecraft %s...\n", version)

		versionDir := filepath.Join("..", "minecraft_data", version)
		if err := os.MkdirAll(versionDir, 0755); err != nil {
			fmt.Printf("Error creating directory: %v\n", err)
			continue
		}

		// Extract blocks
		if err := extractBlocks(versionDir); err != nil {
			fmt.Printf("Error extracting blocks: %v\n", err)
		}

		// Extract items
		if err := extractItems(versionDir); err != nil {
			fmt.Printf("Error extracting items: %v\n", err)
		}

		// Extract entities
		if err := extractEntities(versionDir); err != nil {
			fmt.Printf("Error extracting entities: %v\n", err)
		}

		fmt.Printf("✅ Extracted %s\n", version)
	}

	fmt.Println("\n✅ All data extracted successfully!")
}

func extractBlocks(versionDir string) error {
	blocks := make(map[string]BlockData)

	for name, id := range data.BlockNameToID {
		blocks[name] = BlockData{
			ID:       id,
			Name:     name,
			Solid:    data.IsSolid(id),
			Hardness: data.GetHardness(id),
		}
	}

	return writeJSON(filepath.Join(versionDir, "blocks.json"), blocks)
}

func extractItems(versionDir string) error {
	items := make(map[string]ItemData)

	for name, id := range data.ItemNameToID {
		items[name] = ItemData{
			ID:         id,
			Name:       name,
			StackSize:  data.GetMaxStackSize(id),
			Durability: data.GetMaxDurability(id),
		}
	}

	return writeJSON(filepath.Join(versionDir, "items.json"), items)
}

func extractEntities(versionDir string) error {
	entities := make(map[string]EntityData)

	for name, id := range data.EntityNameToID {
		entities[name] = EntityData{
			ID:     id,
			Name:   name,
			Width:  0.6, // Default values
			Height: 1.8,
		}
	}

	return writeJSON(filepath.Join(versionDir, "entities.json"), entities)
}

func writeJSON(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
