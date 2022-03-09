Atomicity 

- all queries in transaction must succeed
- if 1 query fails - all prior successful queries in transac should rollback
- if DB went down prior to commit - all prior successful queries in transac should rollback
  - as a part of its startup maybe
- another reason why long transactions are bad idea !