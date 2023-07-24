## Strangler Pattern

Inspiration - is derived from Banyan tree - that has roots which grow with support of host tree. but then over time the host is strangled.

### How do you migrate a legacy monolithic application to a microservice architecture?


1. Do full-blown cut-over replacement of old to new 

OR 

2. Gradually create a new system around the edges of the old, letting it grow slowly over several years until the old system is strangled.
i.e. One can Modernize a legacy application by incrementally developing a new (strangler) application around the legacy application. and slowly cutting down the older legacy.

Refer for more : https://martinfowler.com/bliki/StranglerFigApplication.html

### How to get our stragler running ? 

Through EventInterception

Basically there should be a way to get updates on the legacy by tapping into this stream of events.

Example you can intercept the 
1. kafka messages 
2. DB entries --> DB triggers
3. API built on top of DB
4. DB table initial snapshot --> then monitor updates
   
Once you've tapped into the event stream, you can then build application functionality in the strangler.

Refer : https://martinfowler.com/bliki/EventInterception.html
