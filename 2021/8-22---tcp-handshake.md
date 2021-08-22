    TCP Handshake

if i do curl http --> we establish a TCP connection
-> i am sending pkts to server

How does client know that its packet is delivered ?
-> by sending ACK from server
If client doesnt get and ACK , then it retries 

Each packet has sequence number.
Server can now order these pkts

Client and Server both will have their own sequencing
-> for this you need Sync step

    Client                                                         Server
    
    starts with                  syn 700      
    pkt-700                 ------------------>
    
    .                  [syn 200 / ack 701 (700 + size of pkt)]
    client                  <------------------                server decides lets say 200 as pkt number
                                      
                                  ack 201 
    pkt-700                 ------------------>                 


[randomness of exists to avoid syn attack]








