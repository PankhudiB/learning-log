
### WhatsApp: Engineering for Success at Scale 

#### Jan 4th 2023

Watching : https://developers.facebook.com/videos/f8-2016/a-look-at-whatsapp-engineering-for-success-at-scale/

Its a 2016 talk,  at which point of time there were 1 Billion whatsapp user
with 46 billion messages concurrently.

What they say is when the receiver user is offline - the messages are stored on transient server in whatsapp

In case of sharing media -> when the sender sends it - the receiver gets a url of where its transiently stored in WA server.
and when all have recieved - the media is deleted from WQ server.

Language - they are using Erlang.

And they say good reason being the feature of hot-swap in Erlang
- which means you can make changes to application code 
- and follow few rules 
- u can take that code and load into running application without affecting anything
    - no breaking connection - no killing processes
- OH my GOD whatt !! - need to read further 


- Erlang - the founder of the lang - used it for telephone switches.
- Erlang was designed for concurrency , scaling, low latency ,
- and fault tolerance is the first class feature of lang. 
- and it was used on firmware of telephone equipment. (non-stop support need)
- functional lang

Builtin Database - Mnesia 

- Mnesia is a distributed, soft real-time database management system written in the Erlang programming language.
- It is distributed as part of the Open Telecom Platform

---------

Watching : https://youtu.be/LJx6mUEFAqQ

- They run the db - literally alongside application node in the same Erlang VM
- 





