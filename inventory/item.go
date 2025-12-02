package inventory

// Item represents a Minecraft item type
type Item struct {
	ID          int    // Item ID
	Name        string // Item name (e.g., "minecraft:diamond")
	DisplayName string // Display name (e.g., "Diamond")
	StackSize   int    // Maximum stack size
}

// ItemStack represents a stack of items in an inventory slot
type ItemStack struct {
	Item  Item                   // Item type
	Count int                    // Number of items in stack
	NBT   map[string]interface{} // NBT data (simplified)
}

// IsEmpty returns true if this stack is empty
func (s *ItemStack) IsEmpty() bool {
	return s.Count <= 0 || s.Item.ID == 0
}

// CanStack checks if this stack can be combined with another
func (s *ItemStack) CanStack(other *ItemStack) bool {
	if s.IsEmpty() || other.IsEmpty() {
		return false
	}
	// Simplified: check if same item (NBT comparison would be more complex)
	return s.Item.ID == other.Item.ID
}

// ItemInfo contains detailed information about an item type
type ItemInfo struct {
	ID         int    // Item ID
	Name       string // Item name
	StackSize  int    // Maximum stack size
	Durability int    // Maximum durability (0 if not applicable)
}
