I am facing this in one of the golang microservice

The microservice is deployed in containerised environment across multiple pods.

We have used a remote distributed cache. 

And the service uses the cache and talks to a toggle server to get the real data.

Now everytime the clock hits the cache expiration time.

All the concurrent request to service -> suffer cache miss -> in-turn putting all the load on toggle server

This is cache thundering herd problem.

I found a nice blog - describing the problem and solution.
https://instagram-engineering.com/thundering-herds-promises-82191c8af57d

Simply put they recommend 
    
    "instead of caching the actual value, cache a Promise that will eventually provide the value."

    