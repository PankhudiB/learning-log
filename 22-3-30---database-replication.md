### Database replication

Today I came across a situation where we have application that is live
and uses a k8s managed postgres instance for storing some data.

Lets call k8s managed postgres -> Src

& externally managed postgress -> Dest

And we want to move away from k8s managed postgres db to --> db that is in external cluster.

Which means I need to move the data (DDL & DML) somehow to the external pg

    Solution1: 
    The simplest option could be to take a pg_dump on Src & pg_restore on Dest. 

Now the amount of data it has is ~300 GB. So pg_dump will take almost 5-6 hours.
But since the application is live, and writing to Src.

how do i take the data that was written while I was taking the dump ??
-> Only option is then bring the system down for those hours -> and do pg dump and restore.  

But I can't afford to take the entire system down !!

    Solution 2: 
    It expects us to take a snapshot of at timestamp T1.
    And keep streaming the changes from Src to Dest. 
    
    This is like physical replication -> being done at hard disc level.
    Being tranfered byte by byte.
    And stream will also handle DDL along with DML.
    
But we needed pg version upgrade for this.

    Solution 3:
    There is something called as logical replication in PG
    It needs you to take a snapshot (pg_dump) --of DDL alone.
    Restore it to the Dest.
    And then keeps publishing the DML changes alone.
    The Dest here has to act like a subscriber and done those changes on itself.

    The logical replication works only :
        - if the table has primary key or a unique key
            Because it has to somehow identify what has be migrated.
            in case subscriber misses to add data.
    




