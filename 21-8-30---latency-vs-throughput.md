## August 30th 2021

Latency --> [How Fast ?] 

    -> amount of time -> for a packet to be transfered across a network
    
    in terms of web application

        time taken for pkt or request from client to server

    Network latency 
        -> forming HTTP handshake
        -> multiple hops 
    
    Processing Latency 
        -> after server has received the request 
        -> deserilize the request
        -> other dependecny -> other service | DB | redis
        -> processing  
        -> return some data or some acknowledgement

    How do u measure latency ?
        - average -> not the best way
        - p95 
        - p99 -> good for alarm
    

