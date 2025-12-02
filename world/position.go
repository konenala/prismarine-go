package world

import "math"

// Position represents a block position in the Minecraft world
type Position struct {
	X, Y, Z int
}

// Vec3d represents a 3D vector with floating-point coordinates
type Vec3d struct {
	X, Y, Z float64
}

// Vec2 represents a 2D vector (typically for rotation: yaw, pitch)
type Vec2 struct {
	X, Y float32
}

// ToVec3d converts a Position to a Vec3d
func (p Position) ToVec3d() Vec3d {
	return Vec3d{X: float64(p.X), Y: float64(p.Y), Z: float64(p.Z)}
}

// Add adds two positions
func (p Position) Add(other Position) Position {
	return Position{X: p.X + other.X, Y: p.Y + other.Y, Z: p.Z + other.Z}
}

// Distance calculates the distance between two Vec3d points
func (v Vec3d) Distance(other Vec3d) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

// Add adds two Vec3d vectors
func (v Vec3d) Add(other Vec3d) Vec3d {
	return Vec3d{X: v.X + other.X, Y: v.Y + other.Y, Z: v.Z + other.Z}
}

// ToPosition converts a Vec3d to a Position (floors to integers)
func (v Vec3d) ToPosition() Position {
	return Position{X: int(math.Floor(v.X)), Y: int(math.Floor(v.Y)), Z: int(math.Floor(v.Z))}
}
