FROM mysql:8.0.28

# Install cron and a text editor (nano)
RUN apt-get update && \
    apt-get install -y cron nano && \
    rm -rf /var/lib/apt/lists/*