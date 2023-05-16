## High-level design
- S3 to store file in one app bucket:
    - Organize files like this: `{user_id}/{file_id}`
- Expire files: 
    - cronjob
    - file store `expired_at`
- Postgres to store file metadata

## Setup S3 bucket + upload
- create s3 bucket in a region `ap-southeast-1`, note down arn `arn:aws:s3:::paste-bin-app-demo`
- create user with attached policy `AmazonS3FullAccess`
- go to user details > security credentials, create access key
- TODO: try other alternatives that involve create policies, then attach to user
    - https://youtu.be/NZElg91l_ms

## Misc
- Response headers
    - `Content-Type: application/octet-stream`:
        - binary data
        - not sth to display but sth downloadable 
        - imply that UI has sth like
        ```html
        <a href="/get/{{fileID}}" download>Download File</a>
        ```
    - `Content-Type: text/plain`:
        - can be displayed to UI