
## Revising Multithreading
#### Aug 13th 2022

- Thread can be created by : 
  1. `extending Thread class`
     - obj init --> new theClassThatIsExtendingThreadClass();
     - comes with inbuilt methods like yield(), interrupt()
  2. `implementing Runnable interface`
     - obj init --> new Thread(new theClassThatIsImplementingRunnableInterface());
     - advantage - to extend more than 1 classes
- Thread `states` : 
  - New 
  - Runnable
  - Blocked
  - Waiting
  - Timed Waiting
  - Terminated
  
--------

- For each program, a Main thread is created by JVM
- `Thread.currentThread().join();` -> will get into DeadLock :P
- 