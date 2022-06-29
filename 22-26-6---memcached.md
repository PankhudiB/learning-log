    Memcached & Facebook 

#### June 26th 2022

`Distributed key-value store for FB`

- Trillions of items.
- Billions req/sec
- n/w attached in memory hash table
- LRU based eviction
- Demand filled cache
  
- For FB usecase,
  - Read heavy
  - Delay of few seconds is okay - posts - but not days
  - If user themselves update - and dont see it reflected - extremely confusing

- get key - cache miss - DB lookup - set key

- For update flow -> web server will do -> DB update + delete the key in cache
  - WHY DELETE instead of actual UPDATE ??
  - bcoz delete are idempotent -> meaning u can replay them if its lost or gets delayed through the system






