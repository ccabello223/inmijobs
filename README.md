# inmijobs

inmijobs is an application consisting of a React frontend, an authentication backend (Node.js), and a core backend (Go), with Turso (libSQL) as the primary database.

## Prerequisites

- [Node.js](https://nodejs.org/) (v18 or higher)
- [Go](https://go.dev/) (v1.25 or higher)
- [Docker](https://www.docker.com/) (for the shared libSQL database)

## Project Structure

- **web**: React frontend (TanStack Router, Vite, Tailwind CSS)
- **backend-auth**: Authentication service (Node.js, Hono, Better Auth, Drizzle ORM)
- **backend-core**: Core application logic (Go, GORM)

## Getting Started

Follow these steps to run the application locally.

### 1. Start the Shared Database

A `docker-compose.yaml` file is provided to run the shared Turso (libSQL) instance for `backend-core`.

```bash
docker-compose up -d
```

This will start the database on port `8787`.

### 2. Set up Environment Variables

**Backend Auth**
Copy the example environment file and adjust if necessary.

```bash
cd backend-auth
cp example.env .env
```

The default `DATABASE_URL` points to the local `database.db` file.

### 3. Run the Backend Services

**Authentication Service (Backend Auth)**

```bash
cd backend-auth
npm install
npm run dev
```

Runs on `http://localhost:3000` (default Hono port).

**Core Service (Backend Core)**

```bash
cd backend-core
go mod tidy
go run ./cmd/server/main.go
```

Runs on `http://localhost:8080` (default).

### 4. Run the Frontend

**Web Application**

```bash
cd web
npm install
npm run dev
```

Runs on `http://localhost:5173` (default Vite port).

## Available Scripts

### Backend Auth
- `npm run dev`: Start development server with hot reload.
- `npm run build`: Compile TypeScript.

### Web
- `npm run dev`: Start Vite development server.
- `npm run build`: Build for production.
- `npm run lint`: Run ESLint.
- `npm run format`: Run Prettier.
