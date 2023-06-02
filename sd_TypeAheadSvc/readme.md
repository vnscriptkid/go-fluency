## Components
- Data collection: assuming only source is from searches in the past
    - input: raw searches
    - output:
        - stored in distributed file system (S3)
        - file structure: `s3://your-bucket/YYYY/MM/DD/HH/logfile.txt`
        - write to s3
            - direct write to s3
            - buffer and batch

## Notes
- 09:00:00 -> 09_00/worker-1.txt
- 09:00:30 -> 09_00/worker-2.txt
- 09:01:00 -> 09_00/worker-1.txt
- 09:04:00 -> 09_00/worker-2.txt
- 09:04:59 -> 09_00/worker-2.txt
- 09:05:00 -> 09_05/worker-1.txt
