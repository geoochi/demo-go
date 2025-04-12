# Go Backend Demo

## install

```bash
go run cmd/server/main.go
```

### create user

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john.doe@example.com","password":"123456"}'
```

### get user list

```bash
curl http://localhost:8080/users
```

### get user by id

```bash
curl http://localhost:8080/users/1
```

### update user

```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john.doe@example.com","password":"654321"}'
```

### delete user

```bash
curl -X DELETE http://localhost:8080/users/1
```
