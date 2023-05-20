from pyspark.sql import SparkSession
from pyspark.sql.functions import split, explode, lower, col, count

# Create a Spark session:
spark = SparkSession.builder \
    .appName("Word Count") \
    .getOrCreate()

# Load the data into a DataFrame. Replace <your_text_file> with the path to your text file:
text_file = "./work/textfile.txt"
text_df = spark.read.text(text_file)

# Force partitioning
# text_df = spark.sparkContext.textFile(text_file, 10).toDF()

# Repartition the DataFrame to have 10 partitions:
text_df = text_df.repartition(10)

print("Number of partitions: ", text_df.rdd.getNumPartitions())

# Process the data and count word occurrences:
word_counts = text_df.select(explode(split(lower(col("value")), r'\W+')).alias('word')) \
    .groupBy('word') \
    .agg(count('*').alias('count')) \
    .sort(col('count').desc())

word_counts.show()