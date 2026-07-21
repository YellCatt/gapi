# GAPI

一个基于 Go + GORM + SQLite 构建的 RESTful API 服务模板。

## 功能特性

- ✅ 自动配置文件创建：首次运行自动生成 `config.yaml`
- ✅ 纯 Go 实现：使用 `modernc.org/sqlite`，无需 CGO
- ✅ 跨平台：支持 Linux、Windows、macOS
- ✅ 用户管理 API：提供完整的用户 CRUD 操作
- ✅ 结构化日志：使用 Zap 日志库
- ✅ API 文档：集成 Swagger UI

## 快速开始

### 前置要求

- Go 1.22+

### 安装依赖

```bash
go mod tidy
```

### 构建

```bash
# 开发环境构建
go build -o gapi

# 生产环境构建（禁用 CGO）
CGO_ENABLED=0 go build -o gapi
```

### 运行

```bash
./gapi
```

首次运行时，系统会自动创建 `config/config.yaml` 配置文件。

### 访问服务

- API 服务：http://localhost:8080
- Swagger 文档：http://localhost:8080/swagger/index.html

## API 接口

| 方法 | 路径 | 描述 |
|------|------|------|
| GET | /users | 获取所有用户 |
| GET | /users/:id | 获取单个用户 |
| POST | /users | 创建用户 |
| PUT | /users/:id | 更新用户 |
| DELETE | /users/:id | 删除用户 |

### 示例请求

#### 创建用户

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

#### 获取所有用户

```bash
curl http://localhost:8080/users
```

#### 获取单个用户

```bash
curl http://localhost:8080/users/1
```

#### 更新用户

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "John Updated", "email": "john.updated@example.com"}'
```

#### 删除用户

```bash
curl -X DELETE http://localhost:8080/users/1
```

## 配置文件

配置文件 `config/config.yaml` 会在首次运行时自动创建：

```yaml
server:
  port: 8080

database:
  path: ./data.db

log:
  path: ./logs
  level: info
```

| 配置项 | 说明 | 默认值 |
|--------|------|--------|
| server.port | 服务端口 | 8080 |
| database.path | SQLite 数据库路径 | ./data.db |
| log.path | 日志目录 | ./logs |
| log.level | 日志级别 | info |

## 项目结构

```
gapi/
├── config/          # 配置管理
│   ├── config.go    # 配置加载
│   ├── config.yaml  # 配置文件
│   └── database.go  # 数据库初始化
├── controller/      # 控制器
│   └── user_controller.go
├── service/         # 服务层
│   └── user_service.go
├── repository/      # 数据访问层
│   └── user_repository.go
├── model/           # 数据模型
│   └── user.go
├── router/          # 路由配置
│   └── router.go
├── logger/          # 日志管理
│   └── logger.go
├── docs/            # API 文档
│   └── docs.go
├── main.go          # 入口文件
└── go.mod           # Go 模块依赖
```

## 技术栈

- **框架**: Go 1.22
- **ORM**: GORM
- **数据库**: SQLite (modernc.org/sqlite)
- **日志**: Zap
- **API 文档**: Swagger

## 编译说明

由于使用纯 Go 的 SQLite 驱动，可以在任何平台上编译：

```bash
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gapi-linux

# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o gapi.exe

# macOS
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o gapi-macos
```

## 许可证

MIT License