## August 29th 2021

Systems need storage
    
u'll 
    
    store & retrieve data

need to take steps to persist the data

    If Write data to Disk -> takes time | file/data present even after server craches 
    If Write data to Memory -> faster | file/data lost after server craches

How DB stores data in Disk ??

    Lets say we have a table -> with 1 column -> integer datatype -> which mean 4 bytes 
    
Can do it in 2 ways :
    
    HDD || SSD

Hard Disk Drive (HDD)
    
    - platter(circles) -> track -> has sectors (triangle) -> blocks - tagged
    --> this int will stored in that block (some no of bytes)
    insert no 7 -> reserve 4 bytes , insert no 1000000 -> reserve 4 bytes 
    
    lets say block is 512 bytes 
    -> will lead to waste of bytes. 
    So DB smartly stores the rows in adjacently in the block

    DB stores metadata
    row no 1 -> exact sector - track - block
    
    now DB can will not need another I/O untill it reaches the last block -> 128th in our case (512/4)
    12th block will store where is the next block or address to look at

    u can write to same block multiple times

SSD
    
    they dont have sectors or track or anything 
    But they do have Blocks (whatever size)
    these blocks of memory have pages
    
    For SSD it just stores -> block no and page no
        
    this is much much faster than HDD -> bcox now its random access -> u r not spinning wheel 
    
    in ssd -> when u pull a block -> it pulls all pages 

    Write ->
        read block and page no of row
        even if you are modifying 
        SSD will write to a new page no matter what
        have "endurance" -> no of time u can keep writing to the sae block until it
    
    Btrees with ssd
    -> if u are adding new nodes -> ssd happy with it 
    -> wbut when indexes are restructered -> while rebalancing itself 
            -> this is update -> Worse for SSD -> Slows down

    Facebook came up with their own DB engine (to aovid above problem)
        -> LSM trees -> this is an append only tree
