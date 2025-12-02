# prismarine-go

Pure Minecraft data models for Go - Layer 3 (L3) of the go-mc bot architecture.

## Overview

prismarine-go provides protocol-agnostic data structures and models for Minecraft:
- World, Chunk, and Block representations
- Entity data structures
- Inventory and Item models
- Chat message types

This layer has **zero protocol dependencies** and can be reused across different Minecraft versions.

## Packages

- `world/` - World state, chunks, blocks, positions
- `entity/` - Entity representations and metadata
- `inventory/` - Inventory, items, slots
- `chat/` - Chat messages and components
- `physics/` - Collision detection and AABB
- `data/` - Game data registry (blocks, items, entities)

## Status

🚧 Under active development - Phase 2 of architecture refactoring.
