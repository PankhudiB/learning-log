
## Redis Sorted Sets & Rate limiting with Redis
#### June 30th 2022

Read the article : 

###What are Redis Sorted Sets ?

https://medium.com/analytics-vidhya/redis-sorted-sets-explained-2d8b6302525


- TC -> O(log(n))
- key -> member - score + (implicit attribute called Rank)
  - if same value -> then lexo order of member

```
    Eg.  Health Leaderboard based on footsteps count
         Key -> players/runners 
         members -> each person
         score -> step count in the day   
```

- Good for maintaining Real Time LeaderBoard , Priority queues, and secondary indexes.
- Also, good for doing rate limiting
- Values usually ordered by Ascending
- Operations :
  - ZADD 
  - ZINCRBY -> increment existing key's value by 
  - ZREVRANGE -> reverse order -> something like -> Get Top 10
  - ZCOUNT -> <key> <min-value> <max-value>
  - ZREM -> to remove member 
  