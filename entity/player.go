package entity

// Player represents a player entity with additional properties
type Player struct {
	Entity             // Embed base entity
	Name       string  // Player name
	GameMode   int     // Game mode (0=survival, 1=creative, 2=adventure, 3=spectator)
	Health     float32 // Player health
	Food       int     // Food level
	Saturation float32 // Food saturation
}

// IsPlayer returns true if this entity is a player
func (e *Entity) IsPlayer() bool {
	return e.Type == TypePlayer
}

// NewPlayer creates a new player entity
func NewPlayer(eid int32, uuid [16]byte, name string) *Player {
	return &Player{
		Entity: Entity{
			EID:       eid,
			UUID:      uuid,
			Type:      TypePlayer,
			Metadata:  make(map[uint8]interface{}),
			Equipment: make(map[int8]interface{}),
		},
		Name:     name,
		GameMode: 0,
		Health:   20.0,
		Food:     20,
	}
}
