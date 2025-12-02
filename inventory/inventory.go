package inventory

import "errors"

var (
	ErrSlotOutOfRange = errors.New("slot index out of range")
	ErrInvalidSlot    = errors.New("invalid slot")
)

// Inventory represents a player's inventory
type Inventory struct {
	Slots    []ItemStack // All inventory slots (typically 46: 9 hotbar + 27 main + 4 armor + 1 offhand + 5 crafting)
	Size     int         // Total number of slots
	HeldSlot int         // Currently selected hotbar slot (0-8)
}

// NewInventory creates a new empty inventory
func NewInventory(size int) *Inventory {
	return &Inventory{
		Slots:    make([]ItemStack, size),
		Size:     size,
		HeldSlot: 0,
	}
}

// GetSlot gets an item stack from a slot
func (inv *Inventory) GetSlot(slot int) (*ItemStack, error) {
	if slot < 0 || slot >= inv.Size {
		return nil, ErrSlotOutOfRange
	}
	return &inv.Slots[slot], nil
}

// SetSlot sets an item stack in a slot
func (inv *Inventory) SetSlot(slot int, stack ItemStack) error {
	if slot < 0 || slot >= inv.Size {
		return ErrSlotOutOfRange
	}
	inv.Slots[slot] = stack
	return nil
}

// FindItem finds the first slot containing an item with the given ID
func (inv *Inventory) FindItem(itemID int) (int, bool) {
	for i, stack := range inv.Slots {
		if stack.Item.ID == itemID && !stack.IsEmpty() {
			return i, true
		}
	}
	return -1, false
}

// Count returns the total count of an item across all slots
func (inv *Inventory) Count(itemID int) int {
	total := 0
	for _, stack := range inv.Slots {
		if stack.Item.ID == itemID {
			total += stack.Count
		}
	}
	return total
}

// EmptySlots returns the number of empty slots
func (inv *Inventory) EmptySlots() int {
	count := 0
	for _, stack := range inv.Slots {
		if stack.IsEmpty() {
			count++
		}
	}
	return count
}
