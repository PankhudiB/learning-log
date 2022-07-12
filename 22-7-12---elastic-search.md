
## Elastic Search
#### July 12th 2022

- Full text search engine
- query structured data
- analyse app logs / system metrics
- anomaly detection
- alerting
---------------------
#### Written in JAVA and built on Apache Lucene
---------------------

#### Elastic Stack 

- Kibana -> Web interface for ES & for visualisation & dashboards
- Logstash -> process logs from app and send to ES + Data processor pipeline 
- X-Pack 
  - -> add security -> authorisation and authentication for both ES and Kibana
  - -> Enables monitoring the performance of ELK stack -> es,kibana,
  - Reporting -> pdf and csv   
  - ElasticSearch SQL

- Beats
  - data shippers
  - light-weight agents that u install on servers -> with single purpose of sending data to logstash || ES
  - Examples:
    - FileBeat -> collecting log files
    - MetricBeat -> for system and service metrics - CPU & memory usage
    - many other

```
[Data Ingestion] -----> [ Search,Analyse and Store data] -----> [Visualise]
  Logstash                          ElasticSearch                 Kibana
  Beats
```
