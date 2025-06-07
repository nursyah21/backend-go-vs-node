# README

an example backend implement in nodejs and golang use most faster and popular libary

## tech stack

-  nodejs: hono, drizzle
-  golang: fiber, sqlc
-  database: postgresql

## how to run

before you started it, please make sure you already run of docker compose for database.

```bash
docker compose up -d
```

and for golang, you need to install air for hot reload, sqlc for generate orm

```bash
go install github.com/air-verse/air@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### api-node

```bash
cd api-node
yarn
yarn drizzle:generate
yarn drizzle:migrate
yarn dev
```

### api-go

```bash
cd api-go
yarn
yarn drizzle:generate
yarn drizzle:migrate
yarn sqlc:generate
yarn dev
```

## how to see database

### api-go

```bash
cd api-go
yarn drizzle:studio
```

### api-node

```bash
cd api-node
yarn drizzle:studio
```

## seeding database

### api-go
```bash
cd api-go
yarn db:seed
```

### api-node
```bash
cd api-node
yarn db:seed
```

## drop database

### api-go
```bash
cd api-go
yarn db:drop
```

### api-node
```bash
cd api-node
yarn db:drop
```

## test

before you run test you need to install playwright and autocannon

please make sure you have already data, if not yet you can seeding

```bash
yarn
```

for run test
```bash
yarn test
```

## loadtest

for run loadtest
```bash
yarn loadtest
```

### spec loadtest

- os: windows 10
- ram: 8gb
- cpu: i5 6200u
- connection: 100
- duration: 10s

### benchmark json

- nodejs: 
    - latency: 16.24ms
    - req/s: 5,991
    - total: 60k in 10s

- golang:   
    - latency: 9.37ms
    - req/s: 10,195
    - total: 102k in 10s

**golang 1.7x faster than nodejs for return of json**

### benchmark database

- nodejs: 
    - latency: 120.09ms
    - req/s: 827
    - total: 8k in 10s

- golang:   
    - latency: 81.42ms
    - req/s: 1,217
    - total: 12k in 10s

**golang 1.5x faster than nodejs for retrieving data from postgresql**

### benchmark jwt (login process)

- nodejs: 
    - latency: 1558.83ms
    - req/s: 6
    - total: 74 in 10s

- golang:   
    - latency: 458.67ms
    - req/s: 21
    - total: 228 in 10s

**golang 3x faster than nodejs for retrieving data from postgresql**