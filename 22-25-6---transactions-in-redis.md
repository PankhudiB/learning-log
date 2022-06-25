
###Transactions in Redis
#### June 26th 2022

https://redis.io/docs/manual/transactions/

- `MULTI`:
  - A Redis Transaction is entered using the MULTI command. The command always replies with OK.
  - `redis queues` all the commands u entered
  - Finally executed -> when `EXEC` is called
  

- `EXEC` : returns an array of replies 
- `DISCARD` : to abort a transaction  
- `WATCH`: to provide a check-and-set (CAS) behavior to Redis transactions.
  - `WATCH`ed keys are monitored in order to detect changes against them. If at least one watched key is modified before the EXEC command, the whole transaction aborts

