package physics

import "github.com/konjacbot/prismarine-go/world"

// AABB represents an Axis-Aligned Bounding Box
type AABB struct {
	MinX, MinY, MinZ float64
	MaxX, MaxY, MaxZ float64
}

// NewAABB creates a new AABB from two corner positions
func NewAABB(min, max world.Vec3d) *AABB {
	return &AABB{
		MinX: min.X, MinY: min.Y, MinZ: min.Z,
		MaxX: max.X, MaxY: max.Y, MaxZ: max.Z,
	}
}

// NewAABBFromBlock creates an AABB for a block at the given position
func NewAABBFromBlock(pos world.Position) *AABB {
	return &AABB{
		MinX: float64(pos.X), MinY: float64(pos.Y), MinZ: float64(pos.Z),
		MaxX: float64(pos.X + 1), MaxY: float64(pos.Y + 1), MaxZ: float64(pos.Z + 1),
	}
}

// Intersects checks if this AABB intersects with another
func (a *AABB) Intersects(b *AABB) bool {
	return a.MaxX > b.MinX && a.MinX < b.MaxX &&
		a.MaxY > b.MinY && a.MinY < b.MaxY &&
		a.MaxZ > b.MinZ && a.MinZ < b.MaxZ
}

// Contains checks if this AABB contains a point
func (a *AABB) Contains(pos world.Vec3d) bool {
	return pos.X >= a.MinX && pos.X <= a.MaxX &&
		pos.Y >= a.MinY && pos.Y <= a.MaxY &&
		pos.Z >= a.MinZ && pos.Z <= a.MaxZ
}

// Expand expands the AABB by a given amount in all directions
func (a *AABB) Expand(amount float64) *AABB {
	return &AABB{
		MinX: a.MinX - amount, MinY: a.MinY - amount, MinZ: a.MinZ - amount,
		MaxX: a.MaxX + amount, MaxY: a.MaxY + amount, MaxZ: a.MaxZ + amount,
	}
}

// Center returns the center point of the AABB
func (a *AABB) Center() world.Vec3d {
	return world.Vec3d{
		X: (a.MinX + a.MaxX) / 2,
		Y: (a.MinY + a.MaxY) / 2,
		Z: (a.MinZ + a.MaxZ) / 2,
	}
}

// Size returns the size of the AABB
func (a *AABB) Size() world.Vec3d {
	return world.Vec3d{
		X: a.MaxX - a.MinX,
		Y: a.MaxY - a.MinY,
		Z: a.MaxZ - a.MinZ,
	}
}

// Offset returns a new AABB offset by the given vector
func (a *AABB) Offset(offset world.Vec3d) *AABB {
	return &AABB{
		MinX: a.MinX + offset.X, MinY: a.MinY + offset.Y, MinZ: a.MinZ + offset.Z,
		MaxX: a.MaxX + offset.X, MaxY: a.MaxY + offset.Y, MaxZ: a.MaxZ + offset.Z,
	}
}
