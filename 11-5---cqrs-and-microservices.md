CQRS
Command Query 

Microservices dealing with DB 
    - Read 
    - Write

    1st problem :
        we usually write entities to DB
        & read aggreagated data from DB

    If you index 
        read will become fast
        write will slow down 
    
    and this can backfire -> and you need to have the right balance between the two

    CQRS says have 2 separate models for 
        queries -> aggreagation ~ datagrid
        commands -> entity based 

    DB -> middle layer -> UI
    
    u can have a middle layer that will read data from DB 
        -> and aggreagate data for you 
            -> into presentation model 
                -> send to UI
    
    This doesnt really optimise the DB 
    
    --> Split the DB into Read DB and Write DB
        
        So in case of write path -> the DB -> will not have indexes what so ever
        In read path -> fully optimised DB - with indexes 
        And you ll have DB synchronization from your write DB to the read DB

        sync -> is error prone & slow

    What if we split the middle layer / app into 2 
        microservice that queries the DB
        microservice that writes/updates the DB
    
    -> and this will not just help in splitting the model
        but also helps in scalability 
            -> how many reads are you going to have -> XYZ/second
            -> how many write are you going to have -> XYZ/day

    but then WHAT ABOUT THE DATA ??

    -> we completely get rid of the DB interaction from the read microservice
    -> on the write service we expose a cache
            -> we replicate or distribute this cache to the read service
        ->& when update happens -> we can sync the cache

    NOW both service are optimized ! :D

    https://martinfowler.com/bliki/CQRS.html
    
        
        
    

    
    
    