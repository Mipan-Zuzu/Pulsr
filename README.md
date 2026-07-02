# Pulsr

Stop tab-hopping. Just check.

Dashboard buat monitoring status service dalam satu tempat. Gak perlu buka banyak tab buat cek Vercel, Supabase, Upstash, sama Docker satu-satu.

## Fitur

- Cek status web (HTTP/HTTPS)
- Cek status database (Postgres)
- Cek status Redis (Upstash)
- Cek status container Docker
- Response time & uptime history
- Notifikasi ke Telegram/Discord kalo service down

## Tech stack

**Backend**
- Go
- Fiber
- PostgreSQL
- robfig/cron buat scheduler

**Frontend**
- Vue
- Chart.js buat grafik uptime

## Struktur project

```
pulsr/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── checker/
│   │   └── model/
│   ├── go.mod
│   └── go.sum
├── frontend/
│   ├── src/
│   ├── package.json
│   └── vite.config.js
├── docker-compose.yml
└── README.md
```

## Cara jalanin

### Backend

```
cd backend
cp .env.example .env
go mod tidy
go run cmd/server/main.go
```

### Frontend

```
cd frontend
npm install
npm run dev
```

### Docker Compose

```
docker compose up -d
```

## Environment variables

```
DATABASE_URL=
REDIS_URL=
JWT_SECRET=
TELEGRAM_BOT_TOKEN=
TELEGRAM_CHAT_ID=
```

## Roadmap

- [x] MVP: cek status manual
- [ ] Scheduler otomatis
- [ ] Auth (email/password)
- [ ] OAuth Google/GitHub
- [ ] Realtime update via WebSocket
- [ ] Notifikasi Telegram/Discord
- [ ] Grafik uptime history
- [ ] Multi-user + admin page

## Lisensi

MIT 