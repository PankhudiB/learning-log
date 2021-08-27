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

