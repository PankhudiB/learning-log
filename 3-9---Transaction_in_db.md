Transaction 

- collection of queries
- 1 unit of work ofc
-----
- TRANSACTION BEGIN
- TRANSACTION COMMIT --> persist to disks
  - 2 ways 
    1. do u write each query result to disk during the transac ? 
      - makes commit faster
    2. OR do u write it to memory -> & finally write to Disk ?
      - makes commit slower
- TRANSACTION ROLLBACK
  - forget all changes - dont persist
    - if taken 1 st way --> lot of work to UNDO
    - 2nd way --> delete in memory ! poof ! done !

Nature of Transaction
 - usually modify data
 - read only transaction --> why ??
   - in case u want to be isolated of changes in what u r reading