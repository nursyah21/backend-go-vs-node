# README

read readme.md in each folder api to see how to run it

## benchmark

here's a result benchmark of golang vs nodejs

100 conection running in 60s, the api return a json

**golang (fiber)**:
- latency:
    - avg: 9.3ms
    - stdev:5.96ms
- req/s:
    - avg: 10,000
    - stdev: 1,603
- total: 616k in 60s

**nodejs (hono)**:
- latency:
    - avg: 11.8ms
    - stdev:5.7ms
- req/s:
    - avg: 8,140
    - stdev: 870
- total: 489k in 60s

olang use gin and nodejs use hono