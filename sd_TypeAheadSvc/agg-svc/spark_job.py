from pyspark.sql import SparkSession
from operator import add

def map_prefixes_to_word(word):
    # word: `hello`
    prefixes = [word[:i] for i in range(1, len(word) + 1)]
    # prefixes: ['h', 'he', 'hel', 'hell', 'hello']
    return [(prefix, word) for prefix in prefixes]
    # return: [('h', 'hello'), ('he', 'hello'), ('hel', 'hello'), ('hell', 'hello'), ('hello', 'hello')]

def topK_searches(iterable, k=3):
    iterable = list(iterable)
    iterable.sort(key=lambda x: x[1], reverse=True)  # sort by frequency
    return iterable[:k]  # return top k

if __name__ == "__main__":
    spark = SparkSession.builder \
        .appName('Autocomplete') \
        .master('spark://spark-master:7077') \
        .config('spark.jars.packages', 'org.apache.hadoop:hadoop-aws:3.3.2,com.amazonaws:aws-java-sdk-bundle:1.11.1026') \
        .getOrCreate()

    sc = spark.sparkContext

    spark._jsc.hadoopConfiguration().set("fs.s3a.endpoint", "http://minio:9000")
    spark._jsc.hadoopConfiguration().set("fs.s3a.access.key", "minio")
    spark._jsc.hadoopConfiguration().set("fs.s3a.secret.key", "minio123")
    spark._jsc.hadoopConfiguration().set("fs.s3a.path.style.access", "true")
    spark._jsc.hadoopConfiguration().set("fs.s3a.connection.ssl.enabled", "false")

    # Print the debug string
    # print("Loaded configs: "+ spark.sparkContext._jsc.sc().getConf().toDebugString())

    # Now you can read files from Minio
    data = spark.read.text("s3a://test/data1.txt").rdd.map(lambda row: row[0])

    # Map phase: For each word, generate tuples of prefix and word
    result = (data
        # ['hi', 'hey', 'hi]
        .flatMap(map_prefixes_to_word)
        # [('h', 'hi'), ('hi', 'hi'), ('h', 'hey'), ('he', 'hey'), ('hey', 'hey'), ('h', 'hi'), ('hi', 'hi')]
        .map(lambda x: (x, 1))  # add 1 for each word occurrence
        # [(('h', 'hi'), 1), (('hi', 'hi'), 1), (('h', 'hey'), 1), (('he', 'hey'), 1), (('hey', 'hey'), 1), (('h', 'hi'), 1), (('hi', 'hi'), 1)]
        .reduceByKey(add)  # sum occurrences for each word
        # [(('h', 'hi'), 2), (('hi', 'hi'), 2), (('h', 'hey'), 1), (('he', 'hey'), 1), (('hey', 'hey'), 1)]
        .map(lambda x: (x[0][0], (x[0][1], x[1])))  # restructure for reduce phase
        # [('h', ('hi', 2)), ('hi', ('hi', 2)), ('h', ('hey', 1)), ('he', ('hey', 1)), ('hey', ('hey', 1))]
        .groupByKey()
        # [('h', [('hi', 2), ('hey', 1)]), ('hi', [('hi', 2)]), ('he', [('hey', 1)]), ('hey', [('hey', 1)])]
        .mapValues(topK_searches))  # Reduce phase: For each prefix, find top k words
        # [('h', [('hi', 2), ('hey', 1)]), ('hi', [('hi', 2)]), ('he', [('hey', 1)]), ('hey', [('hey', 1)])]

    result.saveAsTextFile("s3a://test/output/")
