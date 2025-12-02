package entity

// Metadata type constants
// These match Minecraft's metadata type IDs
const (
	MetadataByte        uint8 = 0
	MetadataVarInt      uint8 = 1
	MetadataVarLong     uint8 = 2
	MetadataFloat       uint8 = 3
	MetadataString      uint8 = 4
	MetadataChat        uint8 = 5
	MetadataOptChat     uint8 = 6
	MetadataSlot        uint8 = 7
	MetadataBoolean     uint8 = 8
	MetadataRotation    uint8 = 9
	MetadataPosition    uint8 = 10
	MetadataOptPosition uint8 = 11
	MetadataDirection   uint8 = 12
	MetadataOptUUID     uint8 = 13
	MetadataBlockID     uint8 = 14
	MetadataNBT         uint8 = 15
)

// Common entity metadata field indices
const (
	MetadataIndexFlags             uint8 = 0 // Entity flags (on fire, crouching, etc.)
	MetadataIndexAir               uint8 = 1 // Air time
	MetadataIndexCustomName        uint8 = 2 // Custom name
	MetadataIndexCustomNameVisible uint8 = 3 // Custom name visible
	MetadataIndexSilent            uint8 = 4 // Silent
	MetadataIndexNoGravity         uint8 = 5 // No gravity
	MetadataIndexPose              uint8 = 6 // Pose (standing, sleeping, etc.)
)

// Entity flags
const (
	FlagOnFire       byte = 0x01
	FlagCrouching    byte = 0x02
	FlagSprinting    byte = 0x08
	FlagSwimming     byte = 0x10
	FlagInvisible    byte = 0x20
	FlagGlowing      byte = 0x40
	FlagFlyingElytra byte = 0x80
)
