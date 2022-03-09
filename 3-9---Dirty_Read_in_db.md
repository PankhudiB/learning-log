Read phenomenas

Dirty Read
- change u read - could be part of some other transaction - and it is not commmited yet -> nit yet flushed -> therefore it might change !

```
  begin tx1                   begin tx2
      |   select                    |  
      |                             | update
      |                             | rollback
      |  select sum                 |
      | O.o got != want             |
      and now u cant even derive how u got that value.. bcoz tx2 is rolledback  
```

Non Repeatable Read NRR
- in 1 transaction - u read twice - not like same query
  -> read value , select sum ---> what if 2nd time u read in your own transac --> u got inconsistent value

```
  begin tx1                   begin tx2
      |   select                    |  
      |                             | update
      |                             | commit
      |  select sum                 |
      | O.o got != want             |  
```


Phantom reads :
- in range query --> u read 10 entries , somebody else inserted in that range
- and do another select -> u get phantom new row
- not same as NRR - bcoz u quite didnt read in the first place


```
  begin tx1                              begin tx2
      |   select SOME RANGE QUERY            |  
      |                                      | insert
      |                                      | commit
      |  select sum                          |
      | O.o got != want                      |  
      
      diff from NRR bcoz - in the shared read lock u acquired in the beginning of txn 
      u cant even figure the new rows have been added. 
      its not even there in the lock to be tracked  
```

Lost update
- i wrote something - but before i commit - i tried to read what i wrote
- some other transaction can change what u wrote 


```
    DB table
    [product - quantity 10]
        
  begin tx1                                begin tx2
      | this txn read 10                       | this txn read 10  
      | update qty=qty+10                      |   
      |                                        |   update qty=qty+5  
      |                                        | commit
      |commit                                  |  15                                              15    
      want  25                                    
      got   20                       
```
