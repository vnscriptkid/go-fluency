## Setup
- Access spark dashboard: http://localhost:18080
- Run `docker logs jupyter`, use the link from log, sth like: `http://127.0.0.1:8888/lab?token=b98ace17fca2d371db7ac11cfd747e1b09a7bf46060da846`

## Concepts
- RDD (Resilient Distributed Datasets)
- DataFrames are built on top of RDDs in Spark
- Job splitting rules
    - Default: file > 128MB
    - Force partitioning: `text_rdd = spark.sparkContext.textFile(text_file, 10)` This will create 10 partitions