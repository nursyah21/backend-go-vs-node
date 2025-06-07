# README

## how to run api-node

before you run please make sure you already run of docker compose for database.

**run api-node**
```bash
cd api-node
yarn
yarn dev
```

**run api-go**

before  please, make sure you already install air for hot reload in golang, sqlc for generate migrate, migrate for push your migrate sql to your database

```bash
cd api-go
yarn generate
yarn migrate
yarn dev
```


## benchmark

spec: windows 10, ram 8gb, cpu i5 6200u

tech stack:
- nodejs   : hono, drizzle
- golang   : fiber, sqlc
- database : postgresql

### benchmark json

**100 conection running in 30s, and return of json**

**nodejs**:
- latency:
    - avg:   13.81ms
    - stdev: 6.2ms
- req/s:
    - avg:   6,987
    - stdev: 1,132
- total:     210k in 30s

**golang**:
- latency:
    - avg:   8.4ms
    - stdev: 4.38ms
- req/s:
    - avg:   11,198
    - stdev: 1,461
- total:     336k in 30s

**golang 1.6x faster than nodejs for return of json**

### benchmark database

**100 conection running in 30s, and return of data from database**

**nodejs**:
- latency:
    - avg:   113.7ms
    - stdev: 31.31ms
- req/s:
    - avg:   875
    - stdev: 142
- total:     26k in 30s

**golang**:
- latency:
    - avg:   10.0ms
    - stdev: 11.84ms
- req/s:
    - avg:   9,607
    - stdev: 2,391
- total:     288k in 30s

**golang 11x faster than nodejs for retrieving data from postgresql**