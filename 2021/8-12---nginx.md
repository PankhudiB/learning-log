## August 12th 2021

NGINX

    Works both as Webserver and Proxy

Webserver

    Serves web content
    - static content / html / js

Proxy
    reverse proxy - public -> private
    load balancing 
    backend routing
    
Nginx can act as Layer 4 and Layer 7 proxy

    Layer 4 --> 
        The TCP/IP layer, where we are aware of the TCP handshake and src and destination ports attached. 
    
    Layer 7 -->
        The Application layer, where you have all the headers and content type etc.

    Nginx 4 proxy using stream context
    Nginx 7 proxy using http context

    nginx playground -> https://github.com/PankhudiB/learning-nginx

    
