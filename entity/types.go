package entity

// Entity type constants
// These are common entity types in Minecraft
const (
	TypePlayer        int32 = 125 // Player
	TypeZombie        int32 = 121 // Zombie
	TypeSkeleton      int32 = 105 // Skeleton
	TypeCreeper       int32 = 50  // Creeper
	TypeSpider        int32 = 106 // Spider
	TypeEnderman      int32 = 58  // Enderman
	TypeCow           int32 = 48  // Cow
	TypePig           int32 = 100 // Pig
	TypeSheep         int32 = 104 // Sheep
	TypeChicken       int32 = 41  // Chicken
	TypeVillager      int32 = 118 // Villager
	TypeItemFrame     int32 = 78  // Item Frame
	TypeArmorStand    int32 = 10  // Armor Stand
	TypeMinecart      int32 = 90  // Minecart
	TypeBoat          int32 = 21  // Boat
	TypeItem          int32 = 76  // Dropped Item
	TypeExperienceOrb int32 = 63  // Experience Orb
	TypeArrow         int32 = 11  // Arrow
	TypeFireball      int32 = 65  // Fireball
)

// EntityInfo contains metadata about entity types
type EntityInfo struct {
	ID     int32   // Entity type ID
	Name   string  // Entity name
	Width  float64 // Entity width
	Height float64 // Entity height
}

// Common entity dimensions (simplified)
var EntityDimensions = map[int32]EntityInfo{
	TypePlayer:    {ID: TypePlayer, Name: "player", Width: 0.6, Height: 1.8},
	TypeZombie:    {ID: TypeZombie, Name: "zombie", Width: 0.6, Height: 1.95},
	TypeSkeleton:  {ID: TypeSkeleton, Name: "skeleton", Width: 0.6, Height: 1.99},
	TypeCreeper:   {ID: TypeCreeper, Name: "creeper", Width: 0.6, Height: 1.7},
	TypeCow:       {ID: TypeCow, Name: "cow", Width: 0.9, Height: 1.4},
	TypePig:       {ID: TypePig, Name: "pig", Width: 0.9, Height: 0.9},
	TypeVillager:  {ID: TypeVillager, Name: "villager", Width: 0.6, Height: 1.95},
	TypeItemFrame: {ID: TypeItemFrame, Name: "item_frame", Width: 0.5, Height: 0.5},
	TypeItem:      {ID: TypeItem, Name: "item", Width: 0.25, Height: 0.25},
}
