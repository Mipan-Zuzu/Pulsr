## System
- Healty checker 
- Grafik status simple fastest than Prometheus
- sistem check status Avalible:
- Web/HTTP = status code, response time
- Database (Postgres) = connection ping
- Docker container = status running/health
- Vercel deployment = via Vercel API
- Netlify deployment = via Netlify API
- Domain = DNS resolve, SSL expiry
- Cloudflare tunnel = via Cloudflare API


## Structure File
pulsr/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── handler/
│   │   ├── monitor_handler.go
│   │   ├── auth_handler.go
│   │   └── health_handler.go
│   ├── service/
│   │   ├── monitor_service.go
│   │   └── auth_service.go
│   ├── repository/
│   │   ├── monitor_repository.go
│   │   └── user_repository.go
│   ├── checker/
│   │   ├── web_checker.go
│   │   ├── db_checker.go
│   │   ├── redis_checker.go
│   │   └── docker_checker.go
│   ├── model/
│   │   ├── monitor.go
│   │   └── user.go
│   ├── middleware/
│   │   └── auth_middleware.go
│   └── scheduler/
│       └── scheduler.go
├── pkg/
│   └── utils/
├── .env
├── go.mod
├── go.sum
└── README.md