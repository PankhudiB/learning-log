## Distributed job scheduler

went through the article 

https://www.linkedin.com/pulse/system-design-distributed-job-scheduler-keep-simple-stupid-ismail/

The gist is :
1. Once user submits a job --> the request can be stored in a DB
2. The cluster of job schedulers (Master - Worker) pick the entries from DB based on some partition
3. Put those in the Queue
4. The cluster of job executors (Master - Worker) pick from kafka 
5. Mark it done/ not done in local DB
6. Parallel health-checker || re-conciliation service 
   -> can look at status of failed jobs during execution and put them back in Queue
