
### SSL ciphers 

#### Oct 14st 2022

Got intrigued to read about this when I came across vulnerability issue :
https://sweet32.info/ 

Thought i will read about ciphers first and then go about understanding the issue.

When 

```
Client       <------accessing secure website --->  Web server         
Web broswer 
```
Creating a secure, encrypted connection via the SSL/TLS protocol needs SSL/TLS Handshake

To negotiate a secure communication, client & server need to come on terms with four crucial elements:

```
1. Key exchange algorithm
2. Bulk encryption algorithm
3. Data encryption algorithm
4. Hash function.
```

whenever secure connection happens -it can be done via several protocols
    
    Protocol -> TLS 1.2 , TLS 1.1, TLS 1.0, SSLv3, SSLv2 (in descending order)
    Key Exchange -> DH, RSA, 

    
