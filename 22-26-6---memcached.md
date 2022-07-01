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

  - uses `Leases` to solve `Stale set` problem
    - When a client encounters a cache miss for a key -> Memcache will give it a lease to set data into the cache after reading it from the DB.
    - `Lease` -> 64-bit token bound to the key
    - When the client wants to set the data -> Memcache verifies this key
    - When Memcache receives a `delete request` for the key, it will `invalidate` any existing lease tokens for that key.

- uses `Leases` to solve `Thundering herd problem`
  - Memcache servers `give leases only once every 10 seconds per key`
  
- `Memcache Client`
  - Each server constructs DAG of data dependencies 
  - Memcache Client can use this DAG to batch and fetch keys -> reduces n/w round trip
  - TCP for writes
  - UDP for reads
    - if actual cache miss -> talk to DB -> populate cache
    - if pkt drops -> talk to DB + NOT populate cache
    



