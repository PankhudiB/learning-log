### Recursive Database queries 

Came across this while I was trying to store a hierarchical data in relational database.
Something like a tree.

So lets say I have to store an N-ary tree in relational DB,

Requirements:
1. how do i add child given a 'parentid'
2. how to fetch the whole tree ?


From reading and digging up :
I found this 
``` 
WITH R AS (<base query>
            UNION ALL
            <recursive query involving R>)
<query involving R>
```

```
TREE
ID | Parent Paths | Label           | Children
1  |  1           | root            |  2              
2  |  1,2         | child           |  3,4 
3  |  1,2,3       | sub-child-1     |  5
4  |  1,2,4       | sub-child-2     |  -
5  |  1,2,4,5     | sub-child-2-1   |  -
```

Fetch a tree

ROOT_ID = 1

WITH CHILDREN AS <SELECT Children as c FROM Tree where id == ROOT_ID>
                    UNION ALL
                 SELECT  
SELECT * FROM TREE


-------------------
SYNTAX :
```sql

WITH [RECURSIVE] recursiveName AS
    (
        SELECT QUERY (non-recursive or base query) <--
        UNION
        SELECT QUERY (recursive query using recursiveName [with a termination condition])
    )
SELECT * from recursiveName;
```

when you execute 

-- 1st thing sql will understand that it is a recursive query 

-- then it ll search for base query (1st iteration) -- 1 record of output will be of base query

-- the 2nd iteration takes the output of 1st iteration as input

-- the 3nd iteration takes the output of 2nd iteration as input



1. Print 1 to 10 without builtin function
```sql
WITH [RECURSIVE] number AS
    (
    SELECT 1 as n
    UNION
    SELECT n+1 from number where n < 10
    )
SELECT * from number;
```

2. Print heihrarchy of employees under a given manager

![img_9.png](img_9.png)

-----------------
EXAMPLE TREE DATA =>

```
TREE
ID | Parent       | Node           | Label
1  |  0           | 1              |  root               
2  |  1           | 2              |  child1 
3  |  1           | 3              |  child2 
4  |  2           | 4              |  child-11 
5  |  2           | 5              |  child-12 
6  |  3           | 6              |  child-21 
7  |  3           | 7              |  child-22 
```

building 2 indexes on both columns

```
    WITH [RECURSIVE] tree_heihrarchy AS
       (
            ( SELECT * FROM TREE WHERE Parent = 0 )
            UNION            
            (
                SELECT * 
                FROM tree_heihrarchy TH
                JOIN TREE T on TH.node == T.parent
            )
       )
    SELECT * from tree_heihrarchy;
```


