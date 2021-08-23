SYN Flood Attack

I was reading about TCP handshake which lead me reading about this attack. 

In a nutshell in a TCP handshake , 

SYN packets are least likely to be rejected by default.

The attacker sends repeated SYN packets to targeted server, [often using a fake IP address :D] 

The server, unaware of the attack, thinks them to be legitimate requests to establish communication.

It responds to each attempt with a SYN-ACK packet from each open port. And falls into the TRAP !!

During this time, the server cannot close down the connection by sending an RST packet, and the connection stays open.
Before the connection can time out, another SYN packet will arrive.

Read a good article on the attack and mitigation:
    
    https://www.imperva.com/learn/ddos/syn-flood/