# README

read readme.md in each folder api to see how to run it

## benchmark

spec: windows 10, ram 8gb, cpu i5 6200u

here's a result of golang vs nodejs

**100 conection running in 30s, the return of json**

**nodejs (hono)**:
- latency:
    - avg:   13.81ms
    - stdev: 6.2ms
- req/s:
    - avg:   6,987
    - stdev: 1,132
- total:     210k in 30s

**golang (fiber)**:
- latency:
    - avg:   8.4ms
    - stdev: 4.38ms
- req/s:
    - avg:   11,198
    - stdev: 1,461
- total:     336k in 30s

**golang 1.6x faster than nodejs for return of json**

**100 conection running in 30s, the return of data from database postgresql**

**nodejs (hono, drizzle)**:
- latency:
    - avg:   113.7ms
    - stdev: 31.31ms
- req/s:
    - avg:   875
    - stdev: 142
- total:     26k in 30s

**golang (fiber, sqlc)**:
- latency:
    - avg:   10.0ms
    - stdev: 11.84ms
- req/s:
    - avg:   9,607
    - stdev: 2,391
- total:     288k in 30s

**golang 11x faster than nodejs for retrieving data from postgresql**