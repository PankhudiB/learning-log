    Redis v/s Memcached

#### June 28th 2022
```
              Redis                                  Memcached         

- allows you to persist the data to disk   |  needs 3rd party tool to do this
- data upto 512 MB                         |  only 1 MB value
- Has support for many data-types          |  limited to string
  like Lists, Sets, SortedSets, Geo data,  
- Has Hashes -> key -> field - valued      | needs 3rd party tool to do this
-       okayish                            | has better pipelining

 Memcached is a subset of Redis
```

