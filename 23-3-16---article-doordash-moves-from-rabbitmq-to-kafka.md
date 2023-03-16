## Doordash moves their Backend to Apache Kafka from RabbitMQ

- Went through the article : https://doordash.engineering/2020/09/03/eliminating-task-processing-outages-with-kafka/

My notes

- Celery is Distributed Task Queue
    - similar to Airflow
- RabbitMQ
    - is a queue
- When they started using Celery & RabbitMQ
    - They started facing excessive conn churn
    - they got lots of orders -> increased celery jobs -> resulting in lots of jobs in rabbitmq
    - now they faced something called as High Connection Churn
    - High Connection Churn => When rate of opened conn in consistently high + rate of closed conn in consistently high
        - looks like Celery is creating and closing conns while speaking to RabbitMQ
        - should have been solved with http1.1 or keep-alive conn
        - they didn't mention if its behind TLS or not - but that would compound the problem
    - Anyways, they had 2 major options
        - move from rabbitmq to redis
        - move from rabbitmq to kafka
    - in the blog they listed down pros and cons of the options available
    - they chose to go ahead with kafka
    - while switching to kafka in production they needed to be backward compatible
        - Canary deployment !
        - so they kept writing to both rabbitmq and kafka
        - they had both set of worker nodes - rabbitmq worker and kafka workers running
        - once kafka was stable and running they got rid of rabbitmq
    - But again there were problems with Kafka too
    - HEAD OF LINE BLOCKING
        - a slow message in a partition blocks all messages behind it from getting processed.
        - what they did to solve is :
            - Inside 1 worker they had :
                - 1 kafka consumer process
                - multiple task execution process
            - the kafka consumer process would fetch messages from kafka and put them on local queue
            - the task execution processed would read from this local queue.
            - This solution allows messages in the partition to flow and only one task-execution process will be stalled
              by the slow message
      

