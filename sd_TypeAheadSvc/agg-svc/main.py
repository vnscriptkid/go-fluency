import subprocess

def schedule_spark_job():
    try:
        # Use subprocess to call the Spark job script
        result = subprocess.run([
            'spark-submit', 
            '--master', 'spark://spark-master:7077',
            '--packages', 'org.apache.hadoop:hadoop-aws:3.3.2,com.amazonaws:aws-java-sdk-bundle:1.11.1026',
            'spark_job.py'], check=True)
        
        print("Spark job completed successfully.")
    except subprocess.CalledProcessError as e:
        print("Spark job failed with error: ", e)

if __name__ == "__main__":
    schedule_spark_job()