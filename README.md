# prismarine-go

**Pure Minecraft data models for Go** - Layer 3 (L3) of the go-mc bot architecture.

[![Go Reference](https://pkg.go.dev/badge/github.com/user/prismarine-go.svg)](https://pkg.go.dev/github.com/user/prismarine-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## 概述

prismarine-go 提供協議無關的 Minecraft 數據結構和模型：

- ✅ **World, Chunk, Block** - 世界、區塊、方塊表示
- ✅ **Entity** - 實體數據結構與元數據
- ✅ **Inventory, Item** - 背包與物品模型
- ✅ **Chat** - 聊天訊息與格式化組件
- ✅ **Physics** - 碰撞檢測與 AABB
- ✅ **Data Registry** - 遊戲數據註冊表（方塊、物品、實體）
- ✅ **多版本支援** - 支援 Minecraft 1.21.0-1.21.10
- ✅ **零協議依賴** - 可跨版本重用

## 特性

### 協議無關設計

prismarine-go 作為 L3 層，完全獨立於協議實現：

```
L5: 全能bot (應用層)
    ↓
L4: mineflayer-go (高階 API)
    ↓ ↙
L2: nalago-mc (協議層)  ←→  L3: prismarine-go (數據模型) ← 本庫
    ↓
L1: go-mc-core (底層協議)
```

### 數據註冊表

自動從 JSON 生成的遊戲數據：

```go
import "github.com/user/prismarine-go/data"

// 獲取方塊資訊
blockID := data.BlockNameToID["stone"]
block := data.GetBlock(blockID)
fmt.Printf("硬度: %.1f, 固體: %v\n", block.Hardness, block.IsSolid())

// 獲取物品資訊
itemID := data.ItemNameToID["diamond_sword"]
item := data.GetItem(itemID)
fmt.Printf("堆疊大小: %d\n", item.MaxStackSize)
```

支援多個 Minecraft 版本：

```go
// 使用特定版本
registry := data.GetRegistryForVersion("1.21.10")
block := registry.GetBlock("minecraft:stone")

// 使用協議版本號
registry := data.GetRegistryForProtocol(774) // 1.21.10
```

## 安裝

```bash
go get github.com/user/prismarine-go
```

## 快速開始

### World 與 Block

```go
import (
    "github.com/user/prismarine-go/world"
    "github.com/user/prismarine-go/data"
)

// 創建世界
w := world.NewWorld()

// 設置方塊
pos := world.Position{X: 100, Y: 64, Z: 200}
stoneID := data.BlockNameToID["stone"]
w.SetBlock(pos, stoneID, 0)

// 獲取方塊
block, err := w.GetBlock(pos)
if err == nil {
    fmt.Printf("方塊: %s\n", block.Name())
}
```

### Entity

```go
import (
    "github.com/user/prismarine-go/entity"
    "github.com/google/uuid"
)

// 創建實體
e := entity.NewEntity(
    123,                    // 實體 ID
    uuid.New(),            // UUID
    entity.TypePlayer,     // 類型
    entity.Vec3{100, 64, 200}, // 位置
    entity.Vec2{0, 0},     // 旋轉
)

// 更新位置
e.SetPosition(entity.Vec3{101, 64, 201})

// 獲取位置
pos := e.Position()
```

### Inventory

```go
import "github.com/user/prismarine-go/inventory"

// 創建背包
inv := inventory.NewInventory(36) // 36 格

// 設置物品
item := inventory.ItemStack{
    ItemID: data.ItemNameToID["diamond"],
    Count:  64,
    NBT:    nil,
}
inv.SetSlot(0, item)

// 獲取物品
slot, err := inv.GetSlot(0)
```

### Chat

```go
import "github.com/user/prismarine-go/chat"

// 解析 JSON 聊天訊息
msg := chat.ParseJSON(`{"text":"Hello","color":"green"}`)
fmt.Println(msg.String()) // "Hello"

// 創建格式化訊息
msg := chat.Component{
    Text:  "Error!",
    Color: "red",
    Bold:  true,
}
```

## 套件文檔

### world

世界狀態與方塊管理。

**主要類型**:
- `World` - 世界容器
- `Chunk` - 區塊 (16x16x256)
- `Block` - 方塊資訊
- `Position` - 方塊座標

**範例**: 見 [examples/world](examples/world)

### entity

實體表示與元數據。

**主要類型**:
- `Entity` - 基礎實體
- `Metadata` - 實體元數據
- `EntityType` - 實體類型常量

**範例**: 見 [examples/entity](examples/entity)

### inventory

背包與物品系統。

**主要類型**:
- `Inventory` - 背包容器
- `ItemStack` - 物品堆疊
- `Slot` - 物品槽

**範例**: 見 [examples/inventory](examples/inventory)

### chat

聊天訊息與格式化。

**主要類型**:
- `Message` - 聊天訊息接口
- `Component` - 格式化組件
- `Style` - 樣式（顏色、粗體等）

**範例**: 見 [examples/chat](examples/chat)

### data

遊戲數據註冊表（自動生成）。

**主要功能**:
- `BlockNameToID` - 方塊名稱到 ID 映射
- `ItemNameToID` - 物品名稱到 ID 映射
- `EntityNameToID` - 實體名稱到 ID 映射
- `GetBlock()` - 獲取方塊資訊
- `GetItem()` - 獲取物品資訊
- 多版本支援 (1.21.0-1.21.10)

**數據來源**: `data/minecraft_data/`

### physics

物理計算與碰撞檢測。

**主要類型**:
- `AABB` - 軸對齊邊界框
- `Collision` - 碰撞檢測

## 數據生成系統

prismarine-go 使用 `go:generate` 從 JSON 數據生成 Go 代碼：

```bash
# 提取現有數據到 JSON
cd data/tools
go run extractor.go

# 從 JSON 生成 Go 代碼
cd data
go generate

# 生成的檔案
data/blocks_gen.go   # 方塊數據
data/items_gen.go    # 物品數據
data/entities_gen.go # 實體數據
```

查看 [data/README.md](data/README.md) 了解詳情。

## 架構

### 5 層架構中的位置

```
┌──────────────────────────────────────┐
│ L5: 全能bot (應用邏輯)                │
├──────────────────────────────────────┤
│ L4: mineflayer-go (高階 API)         │
├──────────────────────────────────────┤
│ L2: nalago-mc (協議層) ←┐            │
├────────────────────────┐│            │
│ L3: prismarine-go      ││ (數據模型) │ ← 本庫
│ - 零協議依賴           ││            │
│ - 可跨版本重用          ││            │
└────────────────────────┴┴────────────┘
│ L1: go-mc-core (底層)                │
└──────────────────────────────────────┘
```

### 設計原則

1. **協議無關**: 不依賴任何協議實現
2. **版本獨立**: 數據結構可跨 Minecraft 版本重用
3. **類型安全**: 充分利用 Go 的類型系統
4. **高性能**: 最小化記憶體分配
5. **可測試**: 獨立單元測試

## 版本支援

當前支援的 Minecraft 版本：

| Minecraft | 協議版本 | 狀態 |
|-----------|---------|------|
| 1.21.0    | 767     | ✅   |
| 1.21.4    | 768     | ✅   |
| 1.21.5    | 769     | ✅   |
| 1.21.6    | 770     | ✅   |
| 1.21.7    | 771     | ✅   |
| 1.21.8    | 772     | ✅   |
| 1.21.9    | 773     | ✅   |
| 1.21.10   | 774     | ✅   |

## 開發

### 目錄結構

```
prismarine-go/
├── world/           # 世界與方塊
├── entity/          # 實體系統
├── inventory/       # 背包系統
├── chat/            # 聊天訊息
├── physics/         # 物理引擎
├── data/            # 遊戲數據註冊表
│   ├── minecraft_data/  # JSON 數據源
│   ├── tools/           # 代碼生成工具
│   └── *.go            # 生成的代碼
├── examples/        # 範例代碼
└── docs/            # 文檔
```

### 編譯

```bash
go build ./...
```

### 測試

```bash
go test ./...
```

### 生成數據

```bash
cd data
go generate
```

## 相關項目

- [nalago-mc](https://github.com/user/nalago-mc) - Minecraft 協議實現 (L2)
- [PrismarineJS](https://github.com/PrismarineJS) - JavaScript 版本參考
- [minecraft-data](https://github.com/PrismarineJS/minecraft-data) - 遊戲數據源

## 授權

MIT License

## 貢獻

歡迎提交 Issue 和 Pull Request！

查看 [CONTRIBUTING.md](CONTRIBUTING.md) 了解貢獻指南。
