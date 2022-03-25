###GO pointers

Address & Value 

![img_6.png](img_6.png)

-----------------------------------------------------------------------------------

Diff between pointer type declaration & pointer as an operator

![img_5.png](img_5.png)
-----------------------------------------------------------------------------------

Calling a function that has pointer receiver :
Left (long cut) vs Right (Short cut)

![img_7.png](img_7.png)
-----------------------------------------------------------------------------------

Gotcha !! 

![img_8.png](img_8.png)

    Slices are pass by reference.. 
    so even though u dont have a pointer receiver in updateSlice
     -> it ll still update the value

![img_10.png](img_10.png)

Slice -> is actually maintained like [length,capacity,ptr to head]

so even though - it looks like a copy is made.. we essentially copy the 3 things []. which inherently copies the address to head  

-----------------------------------------------------------------------------------

Value Types V/S Reference Types 

![img_9.png](img_9.png)
-----------------------------------------------------------------------------------

