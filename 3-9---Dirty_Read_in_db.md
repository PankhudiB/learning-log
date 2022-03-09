Read phenomenas

Dirty Read
- change u read - could be part of some other transaction - and it is not commmited yet -> nit yet flushed -> therefore it might change !

Non Repeatable Read NRR
- in 1 transaction - u read twice - not like same query
  -> read value , select sum ---> what if 2nd time u read in your own transac --> u got inconsistent value

Phantom reads :
- in range query --> u read 10 entries , somebody else inserted in that range
- and do another select -> u get phantom new row
- not same as NRR - bcoz u quite didnt read in the first place

Lost update
- i wrote something - but before i commit - i tried to read what i wrote
- some other transaction can change what u wrote 

