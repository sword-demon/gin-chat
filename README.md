# gin-chat

一个基于 **Gin** + **GORM(MySQL)** 的 Go 入门项目（当前包含最小可运行的 HTTP 接口与数据库模型/迁移示例）。

## 功能概览

- HTTP 服务启动（端口：`8081`）
- 示例接口：`GET /index`
- 用户表模型：`models.UserBasic`（表名：`user_basic`）
- GORM + MySQL 迁移/插入示例：`test/testGorm.go`

## 环境要求

- Go（`go.mod` 中为 `go 1.24.0`）
- 可选：MySQL（如需运行 `test/testGorm.go`）

## 快速开始

启动 HTTP 服务：

```bash
go run .
```

服务默认监听：

- `http://localhost:8081`

测试接口：

```bash
curl "http://localhost:8081/index"
```

预期输出示例：

```json
{"message":"index"}
```

## 路由与接口

目前路由注册在 `router/app.go`：

- `GET /index`：返回 `{"message":"index"}`

## 数据库（GORM + MySQL）示例

项目中 `test/testGorm.go` 提供了一个最小示例，包含：

- 连接 MySQL
- `AutoMigrate(&models.UserBasic{})` 自动迁移表结构
- 插入一条用户数据

### 运行示例

1. 确保本地 MySQL 可用，并创建数据库 `gin-chat`
2. 根据你本地情况修改 `test/testGorm.go` 中的 DSN（请不要在仓库里写入真实账号密码；下面仅为示例格式）：

```text
<USER>:<PASSWORD>@tcp(<HOST>:<PORT>)/gin-chat?charset=utf8mb4&parseTime=True&loc=Local
```

3. 运行：

```bash
go run ./test/testGorm.go
```

> 建议：使用环境变量/配置文件读取 DSN，并将 `.env` 等本地配置加入 `.gitignore`，避免敏感信息提交到仓库。

## 项目结构

```text
.
├── main.go            # 程序入口，启动 Gin
├── router/            # 路由注册
├── service/           # 业务处理（handler）
├── models/            # 数据模型（GORM）
├── test/              # 测试/实验代码（如 GORM 连接示例）
├── go.mod
└── go.sum
```

## 常见问题

- **端口怎么改？**
  - 修改 `main.go` 中 `r.Run(":8081")`
- **为什么数据库连不上？**
  - 先检查 MySQL 地址/端口/账号密码是否正确，以及数据库 `gin-chat` 是否已创建

