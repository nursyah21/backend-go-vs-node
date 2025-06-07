# README

## how to run api-node

before you started it, please make sure you already run of docker compose for database.

```bash
docker compose up -d
```

and for golang, you need to install air for hot reload, sqlc for generate migrate, and migrate for push your migrate sql to your database

**run api-node**
```bash
cd api-node
yarn
yarn dev
```

**run api-go**


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

**100 conection running in 15s, and return of json**

**nodejs**:
- latency:
    - avg:   13.45ms
    - stdev: 7ms
- req/s:
    - avg:   7,183
    - stdev: 960
- total:     108k in 15s

**golang**:
- latency:
    - avg:   9.93ms
    - stdev: 8.41ms
- req/s:
    - avg:   9,660
    - stdev: 2,164
- total:     145k in 15s

**golang 1.3x faster than nodejs for return of json**

### benchmark database

**100 conection running in 15s, and return of data from database**

**nodejs**:
- latency:
    - avg:   149.1ms
    - stdev: 58.76ms
- req/s:
    - avg:   674
    - stdev: 158
- total:     10k in 15s

**golang**:
- latency:
    - avg:   19.78ms
    - stdev: 18.26ms
- req/s:
    - avg:   5,013
    - stdev: 957
- total:     288k in 15s

**golang 11x faster than nodejs for retrieving data from postgresql**
