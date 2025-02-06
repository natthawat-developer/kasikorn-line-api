
```
kasikorn-line-api
├─ .git
├─ .gitignore
├─ Dockerfile
├─ README.md
├─ cmd
│  └─ main.go
├─ config.yaml
├─ go.mod
├─ go.sum
├─ internal
│  ├─ config
│  │  └─ config.go
│  └─ user
│     ├─ handlers
│     │  └─ user_handler.go
│     ├─ models
│     │  └─ user.go
│     ├─ repositories
│     │  ├─ models
│     │  │  └─ user.go
│     │  └─ user_repository.go
│     ├─ routes
│     │  └─ routes.go
│     └─ services
│        └─ user_service.go
└─ pkg
   ├─ database
   │  └─ database.go
   ├─ error
   │  └─ error.go
   ├─ log
   │  └─ logger.go
   ├─ models
   └─ validator
      └─ validator.go

```

go run github.com/golang/mock/mockgen@latest -source=internal/user/repositories/user_repository.go -destination=internal/user/repositories/mock/mock_user_repository.go -package=repositories