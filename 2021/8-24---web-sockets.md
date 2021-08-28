## August 24th 2021

HTTP build on TCP

        -> req/response system
        -> the CLIENT has to initiate the req. Server doesn't randomly sends response
        -> open TCP -> GET call -> close TCP conn
        -> if u want many more images -> multiple TCP open/close

        -> Alternative -> keep TCP conn open unless u are done
            -> done using header "keep-alive" -> ephimeral header
                -> cant propogated through proxies
        
-> Works --> But there are use cases where server needs to send me information despite me not requesting it

        -> Solve this using -> Web Sockets

ws:// OR wss://    

        uses http 1.1 -> with persistent tcp conn as vehicle
        and both are aware of each other
        -> which mean it is STATEFUL
        -> 1st. thing -> Websocket handshake between client and server
        -> there is no order | u can do it even concurrently
        -> Websocket is binary protocol
        -> anybody can send data C <-> S

Websocket Handshake            
        
        C              -------------------------------->                                    S
                    normat http GET req with upgrade header

                                                                                     if server knows that
                    <--------------------------------                                client wants to upgrade           
                    server will send 101 - switiching protocol  
            
                    now they both can talk 2 each other

* and it will work with HTTP1.1 not 1.0 because of the persistence connection
  
Client headers
            
            GET /anything HTTP/1.1
            upgrade: websocket 
            connection : Upgrade 
            sec-websocket-key : blah        
            sec-websocket-protocol : chat, superchat 
            sec-websocket-version: 13

Server headers

            HTTP/1.1 101 Switching Protocol
            upgrade: websocket
            connection : Upgrade
            sec-websocket-accept : blah        
            sec-websocket-protocol : chat

Use case
    Multiplayer gaming
    

Pros 

    Full duplex -> if u r building an application that sends you updates -> u dont need to do the polling 
    Http compatible -> without that proxies will fail 
    Firewall friendly -> 

Cons 
    
    Proxying is tricky -> nginx very recently started to support
    -> most of the time proxy has to break tls to look for data. and when it does that 
        -> it itself has to establish connection with destination
            -> god knows how you handle web socket -> bidirectional between 
                client<-> proxy & proxy <-> server ! Yikes !
    -> its better we do layer 4 load balancing or proxying than at layer 7 and get into all this mess

    Layer 7 load balancing is challenging -> http usually has timeouts 
    but web socckets dont need to be terminate

    Stateful -> difficult to scale horizontally 

How do we solve horizontally scaling issues ? 

Issue #1 State ----> Sticky session

        Ususally in project we use Proxy with round robin rule
        Instead here we can use something like 
            1. balance leastconnection
                -> make sure that a new user is connected to the instance with the lowest overall number of connections.
            
            2. sign every request from particular user with a cookie -> cookie will contain name of server 
            3. use proxy to accordingly tell which server to use depending on cookie

Issue #2 Broadcasting ----> solve by Pub/Sub
        
        The WebSocket Server knows only about clients connected to this specific instance.
        This means while broadcasting we’re sending a message only to a set of clients, not all of them.

        Solun:
            Use something like Kafka / Redis to braodcast message among servers
            and then those servers can respectively talk to their sets of client
        
        
Good article with code : https://tsh.io/blog/how-to-scale-websocket/


        