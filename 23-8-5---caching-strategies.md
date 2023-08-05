## Caching strategies

### Write Through -->

- Make sure to write to both cache and data store --> || ey or sequentially
- increases latency
- guarantees consistency

### Write Back -->

- Write to cache --> and then asynchronously to datastore
- decreases latency
- no consistency guaranty

### Write Around -->

- Write to Datastore -> and when read comes -> it'll be cache miss -> go read in db -> hydrate cache 
- reading every recently written data entry will result in a cache miss.
- NOT favourable for reading recently written data
- Used this in ft-api --> since not write heavy use-case - people dont change toggle state so frequently

Quiz time: 
1. A system wants to write data and promptly read it back. At the same time, we want consistency between the cache and database
- - Write Through --> Reading will be prompt and data will be consistent in the database as well.

2. What is best and worst for write-heavy system ?
- Write Back -> best -- low latency
- Write through - worst -> very high latency
- Write-around -> bad -> why to write to db ? for heavy writes

## Cache Invalidation

Maintain metadata with each entry --> TTL

Active expiration : run a daemon / thread --> to actively check TTL and remove
Passive expiration : check TTL and remove --> while accessing ---> did during ft-api


