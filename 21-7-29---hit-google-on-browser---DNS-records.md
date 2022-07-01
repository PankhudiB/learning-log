## July 29th 2021

CS75 (Summer 2012) Lecture 0 HTTP Harvard Web Development David Malan

https://youtu.be/8KuO4r5CHjM

What happens when you hit www.google.com ?

- Translate name of site -> to IP [that identifies a server] [32 bits ipv4 / 128 ipv6]
- How does computer know what maps to what ? -> it does a lookup in DNS
    - there is a hierarchy of DNS servers. There are some root servers to that know who has the ip address of .coms / .nets/ .co
    - u as host of server -> has to provide this information when you buy the domain name
- Browser send msg / request to goggle.com -> GET / http/ with some headers -> this packet has 2 ips sender and receivers
- from router to router it finally reaches google
- sees this pkt is for itself -> sees a / -> gets the resource .. flips the ips on the packet 
- sends it back via routers -> my computer sees the resource html -> renders it and u see page in browser
- for home internet connection / org -> u get private ip address something like 
  - 192.168.x.y || 172.16.y.y || 10.x.y.z 
    
Records in DNS    
- CNAME -> canonical name -> alias of domain -> typically used to map a subdomain such as www or mail to the domain hosting that subdomain's content
- MX -> record directs email to a mail server
- NS -> record that tell the Internet where to go to find out a domain's IP address.



