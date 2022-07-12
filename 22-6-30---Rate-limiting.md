
##Rate limiting
#### June 30th 2022


### Rate limiting using Redis -> 

  Before this : Read on : https://github.com/PankhudiB/learning-log/blob/main/22-30-6---redis-sorted-sets.md

[no more than 10 messages per minute or 2 per 3 seconds] 

- Each user has a sorted set associated with them
- keys & values ->  identical -> timestamp of attempted action
- anytime user performs action -> drop all elements of the set which occurred before one interval ago
  - using `ZREMRANGEBYSCORE`
- fetch all elements of the set -> `ZRANGE(0, -1)`
- add the current timestamp to the set -> `ZADD`
  - set its `TTL` to rate-limiting interval
- After all operations are completed 
  - `count` the number of fetched elements.
  - If it `exceeds` the limit -> `don’t allow` the action. 
  - compare the largest fetched element to the current timestamp -> for verifying 2 per 3 seconds RULE
- all Redis operations can be performed as an `atomic` action, using the `MULTI` command

------------

### Rate Limiting capable of scaling to millions of domains :  

https://blog.cloudflare.com/counting-things-a-lot-of-different-things/#fnref4

- they use memcached as cache server + [ sliding window ]
- X -> req in prev minute 
- Y -> req so far in current minute
- algorithm assumes a constant rate of requests during the previous sampling period

```
- Eg : I did 18 requests in the current minute, which started 15 seconds ago,
 and 42 requests during the entire previous minute.
  
-> X * (60-15)/60 + Y
-> rate = 42 * ((60-15)/60) + 18
     = 42 * 0.75 + 18
     = 49.5 requests 
```

- `Tiny memory usage`: only 2 numbers per counter  
- Incrementing a counter can be done by sending a single `INCR` command
- Calculating the rate is reasonably easy: one `GET` command and fast math

#### They have handled the attack scenario very well
- in case of an attack -> 
    - the ultimate load on cache server is going to be huge !!
    - they slow down the legitimate requests
  

- To solve this they are doing the above **+** 
  - `increment jobs` are run `asynchronously` without slowing down the requests.
  - `asks all servers` in the Point of Preference to `start` applying the `mitigation` for that client
  - and they `store this info in-memory` -> that they have started mitigation
  - so `subsequent hits` from this attacker client -> will `not be allowed` to put load on cache server
