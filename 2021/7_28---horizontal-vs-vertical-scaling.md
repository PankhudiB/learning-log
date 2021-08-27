## July 28th 2021

Scaling

     Horizontal                         |       Vertical
     Buy more machine                   |       Buy bigger machine
     load balancing needed              |       N/A
     Resilient                          |       Single point of failure
     communication over n/w             |       communication over function call
     need to manage Data consistency    |       simpler
     scales well as user increase       |       handware limit   

Example 
    
    Scaling pizza parlour
    1. Start with 1 chef 
        -> Pay chef more -> work more ---> #Vertical scaling
        -> optimize process -> eg keep pizza base ready at low peak hours
        -> chef -> sick -> SPOF
        -> get more chef -> #Horizontal scaling

