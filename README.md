# Messenger

Full-stack messenger application inspired by modern chat platforms.

The project is built as a monorepo with a Go gRPC backend and a Next.js web client.

---

# Features

## Authentication

- User registration
- Login with email and password
- JWT access/refresh tokens
- Email verification
- Password reset
- QR code login

## Users

- User profiles
- Username
- Avatar
- Bio
- Online status
- Last seen
- User search

## Chats

- Private chats
- Group chats
- Chat members
- Chat settings

## Messages

- Text messages
- Message history
- Delivery status
- Read status
- Attachments support

## Media

- File uploads
- Images
- Documents

## Notifications

- Message notifications
- User activity notifications

---

# Tech Stack

## Backend

- Go 1.25
- gRPC
- Protocol Buffers
- Buf
- PostgreSQL
- pgx v5
- sqlc

## Frontend

- Next.js 16
- React 19
- TypeScript
- Tailwind CSS
- shadcn/ui

## Infrastructure

- Docker
- Docker Compose

---

# Project Structure

```
Messanger/

├── apps
│   │
│   ├── server
│   │   ├── cmd
│   │   │   └── server
│   │   │       └── main.go
│   │   │
│   │   ├── internal
│   │   │   ├── app
│   │   │   ├── config
│   │   │   ├── database
│   │   │   ├── logger
│   │   │   └── service
│   │   │
│   │   ├── gen
│   │   │   └── generated protobuf files
│   │   │
│   │   ├── migrations
│   │   │
│   │   └── sql
│   │       └── queries
│   │
│   └── web
│       ├── app
│       ├── components
│       ├── lib
│       └── public
│
├── proto
│   ├── auth
│   ├── user
│   ├── common
│   ├── chat
│   ├── message
│   ├── media
│   └── notification
│
├── docker-compose.yml
├── Makefile
├── buf.yaml
└── buf.gen.yaml
```

---

# Architecture

The project follows a service-oriented architecture.

```
                 Web Client
                    |
                    |
              Next.js Application
                    |
                    |
              API Communication
                    |
                    |
              Go gRPC Server
                    |
        -------------------------
        |      |       |        |
      Auth   User    Chat   Message
        |
    PostgreSQL
```

---

# Backend Services

## Auth Service

Responsible for:

- Registration
- Login
- Logout
- Token refresh
- Email verification
- Password reset
- QR authentication

Proto:

```
proto/auth/auth.proto
```

---

## User Service

Responsible for:

- User profile
- Updating profile
- Avatar management
- User search

Proto:

```
proto/user/user.proto
```

---

## Chat Service

Responsible for:

- Creating chats
- Group chats
- Managing members

Proto:

```
proto/chat/chat.proto
```

---

## Message Service

Responsible for:

- Sending messages
- Message history
- Message status

Proto:

```
proto/message/message.proto
```

---

## Media Service

Responsible for:

- Uploading files
- Managing attachments

Proto:

```
proto/media/media.proto
```

---

## Notification Service

Responsible for:

- Notifications
- User activity events

Proto:

```
proto/notification/notification.proto
```

---

# Getting Started

## Requirements

Install:

- Go >= 1.25
- Node.js
- pnpm
- Docker
- Buf

Check:

```bash
go version
```

```bash
node -v
```

```bash
pnpm -v
```

```bash
buf --version
```

---

# Environment Setup

Create environment files.

Backend:

```
apps/server/.env
```

Example:

```env
DATABASE_URL=postgres://user:password@localhost:5432/messenger

SERVER_PORT=50051

JWT_SECRET=secret
```

Frontend:

```
apps/web/.env.local
```

Example:

```env
NEXT_PUBLIC_API_URL=http://localhost:3000
```

---

# Docker

Start infrastructure:

```bash
make up
```

Check containers:

```bash
make ps
```

View logs:

```bash
make logs
```

Stop containers:

```bash
make down
```

---

# Protocol Buffers

All API contracts are stored in:

```
proto/
```

Example:

```
proto/auth/auth.proto
proto/user/user.proto
proto/common/common.proto
```

---

## Validate protobuf

```bash
make lint
```

or:

```bash
buf lint
```

---

## Generate protobuf code

After changing `.proto` files:

```bash
make proto
```

or:

```bash
buf generate
```

Generated files:

```
apps/server/gen/
```

Do not edit generated files manually.

---

# Database

Database migrations:

```
apps/server/migrations
```

Example:

```
001_users.sql
002_password_resets.sql
```

SQL queries:

```
apps/server/sql/queries
```

---

# SQLC

SQLC generates type-safe Go database code.

After changing SQL queries:

```bash
sqlc generate
```

Generated files:

```
apps/server/internal/database
```

Do not edit generated files manually.

---

# Running Development

## Start backend

```bash
make server
```

or:

```bash
cd apps/server

go run ./cmd/server
```

---

## Start frontend

Install dependencies:

```bash
cd apps/web

pnpm install
```

Run:

```bash
pnpm run dev
```

Frontend:

```
http://localhost:3000
```

---

# Development Workflow

## Adding new gRPC method

1. Edit protobuf:

```
proto/**/*.proto
```

2. Validate:

```bash
buf lint
```

3. Generate:

```bash
buf generate
```

4. Implement service logic in Go.

---

## Adding database changes

1. Create migration:

```
apps/server/migrations
```

2. Update SQL queries:

```
apps/server/sql/queries
```

3. Generate:

```bash
sqlc generate
```

---

# Make Commands

Available commands:

| Command       | Description           |
| ------------- | --------------------- |
| `make up`     | Start docker services |
| `make down`   | Stop docker services  |
| `make logs`   | Docker logs           |
| `make ps`     | Show containers       |
| `make proto`  | Generate protobuf     |
| `make lint`   | Validate protobuf     |
| `make server` | Run backend           |
| `make web`    | Run frontend          |

---

# Frontend Architecture

Frontend is built with Next.js App Router.

Structure:

```
apps/web

├── app
│
├── components
│   └── ui
│
├── features
│
├── hooks
│
├── lib
│
└── services
```

UI components are built with:

- shadcn/ui
- Tailwind CSS

---

# Generated Files

Automatically generated:

```
apps/server/gen/
apps/server/internal/database/*.go
```

Do not modify manually.

---

# Roadmap

- [ ] Real-time messaging
- [ ] WebSocket gateway
- [ ] Message encryption
- [ ] Voice messages
- [ ] Video calls
- [ ] Push notifications
- [ ] Multi-device sessions
- [ ] Message reactions
- [ ] Message editing
- [ ] Message deletion
- [ ] Search messages

---

# License

Private project.
