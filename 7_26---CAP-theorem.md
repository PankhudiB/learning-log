## July 26th 2021

CAP theorem

    Consistency -> all clients see same data at same time no matter which node they have connected to
    Availability  -> even if 1 node goes down, the other nodes in the cluster are available to serve the client
    Partitioning tolerence  -> server A thinks server B is dead -> either bcoz of heartbeat or it has actually gone down or
    -> networks are not reliable || GC pauses
    -> not practically possible to have a
    
    Theorem says : You can have any 2 of the properties..and not all three
    
    Amazon shopping cart
    They always make sure to make the cart is available .. the customer can always go and add an item.  ->
    Why we cannot have C and A system ? -> Martin fowler book for nodal
    
    A+P and tradeoff on C
    —>
    Read your own writes + session level consistency
    -> eg. if I a, placing order -> I need to see if my order is rightly placed..
    I don’t care if someone else’s or overall orders in system are consistenly placed or not
    
    CA ->
    not possible
    outliers:
    k8s -> etcd as control plain -> etcd is strongly consistent -> tolerate few failures
    -> why its not as available as dynamo -> bcoz if node goes down in etcd -> it will be available within max 1 second -> which is fine
    -> for shopping cart scenario -> u can’t tolerate on even 1 second delay.
    
    
    When we are scaling usually we fore-go consistency and discuss with business
    
    
    
    What kinds of applications work best with Availability-Partition Tolerance data stores?
    
    Applications like Facebook, Google+, Myspace, Instagram, LinkedIn, Pinterest, Snapchat, and Twitter can all benefit from such a data store. That's because there is no real property nor any operational consequence of getting an old value for an entry in your Facebook stream. Plus, there are almost no updates to such data anyway. Everything is an insert, so if you have a value it must be the latest.
    
    
    What kinds of applications require CP databases?
    
    Things that enterprises do like banking, billing, insurance, inventory, logistics, online e-commerce, manufacturing, risk management and trading are examples of applications that benefit from a CP-biased data store. The update heavy nature of these applications requires that they operate on correct and up-to-date data. This means that in the event of a network partition scenario, having an old copy of data is worse than getting an error when trying to access it if it is unavailable.

