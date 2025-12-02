# Prismarine-Go Data System

## 概述

Prismarine-Go 使用 **JSON → Go 代碼生成** 系統來管理 Minecraft 遊戲數據（方塊、物品、實體）。

## 架構

```
data/
├── registry.go              # Registry 系統和多版本支援
├── blocks.go               # ⚙️ 自動生成 - 方塊數據
├── items.go                # ⚙️ 自動生成 - 物品數據
├── entities.go             # ⚙️ 自動生成 - 實體數據
│
├── minecraft_data/         # 📦 JSON 數據源
│   ├── 1.21.0/
│   │   ├── blocks.json
│   │   ├── items.json
│   │   └── entities.json
│   ├── 1.21.4/
│   ├── 1.21.8/
│   └── 1.21.10/           # 當前默認版本
│
└── tools/
    ├── generator.go        # 代碼生成器
    └── extractor.go        # 數據提取工具
```

## 支援版本

| Minecraft版本 | 協議版本 | 狀態 |
|--------------|---------|------|
| 1.21.0       | 767     | ✅   |
| 1.21.4       | 768     | ✅   |
| 1.21.5       | 769     | ✅   |
| 1.21.6       | 770     | ✅   |
| 1.21.7       | 771     | ✅   |
| 1.21.8       | 772     | ✅   |
| 1.21.9       | 773     | ✅   |
| 1.21.10      | 774     | ✅ 默認 |

## 使用方法

### 1. 獲取 Registry

```go
import "github.com/konjacbot/prismarine-go/data"

// 使用默認 Registry（1.21.10）
registry := data.DefaultRegistry

// 按版本獲取
registry := data.GetRegistryForVersion("1.21.8")

// 按協議版本獲取
registry := data.GetRegistryForProtocol(772) // 1.21.8

// 檢查版本支援
if data.IsVersionSupported("1.21.10") {
    // ...
}

if data.IsProtocolSupported(774) {
    // ...
}
```

### 2. 查詢遊戲數據

```go
// 獲取方塊信息
block, ok := registry.GetBlock(1) // stone
fmt.Printf("Block: %s, Solid: %v, Hardness: %.1f\n",
    block.Name, block.Solid, block.Hardness)

// 獲取物品信息
item, ok := registry.GetItem(276) // diamond_sword
fmt.Printf("Item: %s, Stack: %d, Durability: %d\n",
    item.Name, item.StackSize, item.Durability)

// 獲取實體信息
entity, ok := registry.GetEntity(10) // chicken
fmt.Printf("Entity: %s, Size: %.1fx%.1f\n",
    entity.Name, entity.Width, entity.Height)
```

### 3. 直接使用輔助函數

```go
import "github.com/konjacbot/prismarine-go/data"

// 方塊相關
blockID := data.BlockNameToID["stone"]     // 1
isSolid := data.IsSolid(1)                 // true
hardness := data.GetHardness(1)            // 1.5

// 物品相關
itemID := data.ItemNameToID["diamond_sword"] // 276
stackSize := data.GetMaxStackSize(276)       // 1
durability := data.GetMaxDurability(276)     // 1561

// 實體相關
entityID := data.EntityNameToID["chicken"]   // 10
entityName := data.EntityIDToName[10]        // "chicken"
```

## 開發工作流

### 修改遊戲數據

#### 方式 1: 修改 JSON（推薦）

1. 編輯 JSON 文件：
   ```bash
   # 修改 1.21.10 的方塊數據
   notepad data/minecraft_data/1.21.10/blocks.json
   ```

2. 重新生成 Go 代碼：
   ```bash
   cd data
   go generate
   ```

3. 編譯測試：
   ```bash
   go build ./...
   ```

#### 方式 2: 從頭提取數據

如果你有新的數據源（如 PrismarineJS minecraft-data）：

```bash
cd data/tools

# 修改 extractor.go 添加新數據源
# 然後運行：
go run extractor.go

# 生成 Go 代碼
go run generator.go
```

### 添加新版本

1. 創建新版本目錄：
   ```bash
   mkdir data/minecraft_data/1.22.0
   ```

2. 添加 JSON 文件：
   ```bash
   # 複製現有版本或從數據源提取
   cp data/minecraft_data/1.21.10/*.json data/minecraft_data/1.22.0/
   ```

3. 更新 `registry.go`：
   ```go
   var ProtocolToVersion = map[int32]string{
       // ... 現有版本
       775: "1.22.0", // 添加新版本
   }

   var SupportedVersions = []string{
       // ... 現有版本
       "1.22.0", // 添加新版本
   }
   ```

4. （可選）修改 `tools/generator.go` 支援多版本生成

5. 重新生成和測試：
   ```bash
   cd data
   go generate
   go build ./...
   ```

## go:generate 工作原理

在 `registry.go` 第一行：

```go
//go:generate go run tools/generator.go
```

當運行 `go generate` 時：
1. Go 工具會執行 `go run tools/generator.go`
2. `generator.go` 讀取 `minecraft_data/1.21.10/*.json`
3. 使用 Go 模板生成 `blocks.go`, `items.go`, `entities.go`
4. 生成的文件包含註釋：`// Code generated ... DO NOT EDIT`

## JSON 數據格式

### blocks.json
```json
{
  "stone": {
    "id": 1,
    "name": "stone",
    "solid": true,
    "hardness": 1.5
  }
}
```

### items.json
```json
{
  "diamond_sword": {
    "id": 276,
    "name": "diamond_sword",
    "stackSize": 1,
    "durability": 1561
  }
}
```

### entities.json
```json
{
  "chicken": {
    "id": 10,
    "name": "chicken",
    "width": 0.4,
    "height": 0.7
  }
}
```

## 優勢

### 相比硬編碼 Go 代碼

| 特性 | 硬編碼 | JSON + go:generate |
|------|--------|-------------------|
| **更新數據** | ❌ 需重新編寫代碼 | ✅ 編輯 JSON 即可 |
| **多版本支援** | ❌ 複雜 | ✅ 簡單目錄結構 |
| **從外部同步** | ❌ 手動轉換 | ✅ 直接複製 JSON |
| **可讀性** | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **維護性** | ⭐⭐ | ⭐⭐⭐⭐⭐ |
| **性能** | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ (相同) |

### 相比運行時 JSON 加載

| 特性 | 運行時 JSON | JSON + go:generate |
|------|------------|-------------------|
| **性能** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **二進制大小** | ⭐⭐ | ⭐⭐⭐⭐ |
| **啟動速度** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **類型安全** | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| **更新數據** | ✅ 直接改 JSON | ⚠️ 需 go generate |

## 疑難排解

### go generate 失敗

```bash
# 確認在正確目錄
cd E:\bot編寫\go-mc\prismarine-go\data

# 確認 minecraft_data 存在
ls minecraft_data/1.21.10

# 手動運行生成器（用於調試）
cd tools
go run generator.go
```

### JSON 格式錯誤

```bash
# 驗證 JSON 格式
python -m json.tool minecraft_data/1.21.10/blocks.json
```

### 編譯錯誤

```bash
# 確保生成的代碼是最新的
go generate

# 清理並重新編譯
go clean -cache
go build ./...
```

## 貢獻

添加新數據或新版本時：
1. 更新 JSON 文件
2. 運行 `go generate`
3. 確保編譯通過
4. 提交 JSON 和生成的代碼

---

**Created**: 2025-01-03
**System**: go:generate + JSON data source
**Default Version**: Minecraft 1.21.10 (Protocol 774)
