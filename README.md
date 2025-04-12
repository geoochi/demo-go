# Go Backend Demo

这是一个简单的 Go 后端 demo，实现了基本的用户 CRUD 操作。

## 功能

- 健康检查接口 (`GET /ping`)
- 用户管理接口
  - 创建用户 (`POST /users`)
  - 获取用户列表 (`GET /users`)
  - 获取单个用户 (`GET /users/:id`)
  - 更新用户 (`PUT /users/:id`)
  - 删除用户 (`DELETE /users/:id`)

## 运行项目

1. 确保已安装 Go 1.16 或更高版本
2. 克隆项目
3. 进入项目目录
4. 运行服务器：

```bash
go run cmd/server/main.go
```

服务器将在 http://localhost:8080 启动

## API 示例

### 创建用户

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com","password":"123456"}'
```

### 获取用户列表

```bash
curl http://localhost:8080/users
```

### 获取单个用户

```bash
curl http://localhost:8080/users/1
```

### 更新用户

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"张三改","email":"zhangsan@example.com","password":"123456"}'
```

### 删除用户

```bash
curl -X DELETE http://localhost:8080/users/1
``` 