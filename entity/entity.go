package entity

import "github.com/konjacbot/prismarine-go/world"

// Entity represents a basic entity in the Minecraft world
type Entity struct {
	EID       int32                 // Entity ID
	UUID      [16]byte              // Entity UUID
	Type      int32                 // Entity type ID
	Position  world.Vec3d           // Current position
	Velocity  world.Vec3d           // Movement velocity
	Rotation  world.Vec2            // Yaw, Pitch
	Metadata  map[uint8]interface{} // Entity metadata
	Equipment map[int8]interface{}  // Equipment slots (will be ItemStack from inventory package)
}

// ID returns the entity ID
func (e *Entity) ID() int32 {
	return e.EID
}

// Pos returns the entity position
func (e *Entity) Pos() world.Vec3d {
	return e.Position
}

// SetPosition updates the entity position
func (e *Entity) SetPosition(pos world.Vec3d) {
	e.Position = pos
}

// SetRotation updates the entity rotation
func (e *Entity) SetRotation(rot world.Vec2) {
	e.Rotation = rot
}

// Distance calculates distance to another entity
func (e *Entity) Distance(other *Entity) float64 {
	return e.Position.Distance(other.Position)
}
