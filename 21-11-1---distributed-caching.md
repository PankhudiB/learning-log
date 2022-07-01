    Caching

- key value store

Use-case 
- to avoid n/w call
- to avoid re-computation 
- to avoid hitting a database several time

---> speeds up the response 

Cache ususally runs on SSDs

You cant keep tons of data -> that will increase the search time

Store info in cache - such that db has infinite info and 
cache has the most relevant info according to requests that are going to come in future

--------
2 imp question 

    When do i load data to cache ?
    When do i evict the data ?

usually ppl do :
    - LRU cache 
    - sliding window policy
--------
DONTs with Caching

- too small cache -> leads to more miss -> more n/w calls --> Thrashing -> inputing and outputing to cache without even using the content
- poor eviction policy 
- manage data inconsistency

---------

Where to keep the cache ?

    Close to server --> in-memory with server -> 
        -> if server fails - cache fails 
        -> what about sync between servers 
                -> if its just profile info - not big deal 
                -> if pwds or financial info -> then consistency matters
        -> faster and simpler

    Global cache (can be distributed)
        -> limited in size
                -> if miss - talk to db 
        -> so now even if server caches -> cache unimpacted 
        -> slightly slower but more accurate 
        -> scale independantly irrespective of server's resource
        
---------

How consistent is the data in the cache ?

when a requirement of update comes :

    Write through
        ->  you write to cache first and then to the DB

        Problem 
            in case of in memory cache + write through -> 
                if request hits server 1 
                    -> 1 server will get it updated 
                        -> also update the db but 2nd wont even know                 
        
    Write back
        -> hit the DB first ,then make sure to make the entry in cache  
    
        Problem
            

Spatial Cache 
    -> if i am picking up something , i ll pick up things that are nearby spatially located

Temporal Cache 
    -> similar to most recently data 

Distributed caching
    -> combination of both -> trick -> u need to keep cache in sync with main datastore
        -> redis , memcached
    







