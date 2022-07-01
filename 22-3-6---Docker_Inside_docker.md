I had fair idea of Docker inside Docker, but Today I learned much more about them.
1 learning was unusual and shocking

- How volume mount happens for DIND

For example, consider : 

    H - host
    D1 - Docker container running on H
    D2 - Child container running inside D1 

I thought if When I spin up container D2 from D1 and specify volume mount 

I considered it will from path-in-D1:path-in-D2.

But apparently that is NOT TRUE !! 

Docker expects it to be host path ..even though D1 is its parent container  

