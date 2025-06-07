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

and for golang, you need to install air for hot reload, sqlc for generate migrate

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

## test

for running test
```
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