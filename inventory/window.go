package inventory

// WindowType represents the type of container window
type WindowType string

const (
	WindowTypeChest         WindowType = "minecraft:chest"
	WindowTypeCraftingTable WindowType = "minecraft:crafting_table"
	WindowTypeFurnace       WindowType = "minecraft:furnace"
	WindowTypeDispenser     WindowType = "minecraft:dispenser"
	WindowTypeEnchanting    WindowType = "minecraft:enchanting_table"
	WindowTypeBrewingStand  WindowType = "minecraft:brewing_stand"
	WindowTypeAnvil         WindowType = "minecraft:anvil"
	WindowTypeBeacon        WindowType = "minecraft:beacon"
	WindowTypeHopper        WindowType = "minecraft:hopper"
	WindowTypeShulkerBox    WindowType = "minecraft:shulker_box"
)

// Window represents a container window (chest, furnace, etc.)
type Window struct {
	ID    int         // Window ID
	Type  WindowType  // Window type
	Title string      // Window title
	Slots []ItemStack // Window slots
	Size  int         // Number of slots
}

// NewWindow creates a new window
func NewWindow(id int, windowType WindowType, title string, size int) *Window {
	return &Window{
		ID:    id,
		Type:  windowType,
		Title: title,
		Slots: make([]ItemStack, size),
		Size:  size,
	}
}

// GetSlot gets an item stack from a window slot
func (w *Window) GetSlot(slot int) (*ItemStack, error) {
	if slot < 0 || slot >= w.Size {
		return nil, ErrSlotOutOfRange
	}
	return &w.Slots[slot], nil
}

// SetSlot sets an item stack in a window slot
func (w *Window) SetSlot(slot int, stack ItemStack) error {
	if slot < 0 || slot >= w.Size {
		return ErrSlotOutOfRange
	}
	w.Slots[slot] = stack
	return nil
}
