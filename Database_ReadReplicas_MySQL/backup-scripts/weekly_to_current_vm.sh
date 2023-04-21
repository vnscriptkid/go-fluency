#!/bin/bash

DATE=$(date +%Y-%m-%d)
BACKUP_DIR="/backups/mysql/weekly"
DB_USER="root"
DB_PASS="mysecretpassword"
DB_NAME="test"
HOST="db_master"

mysqldump -h $HOST -u $DB_USER -p$DB_PASS $DB_NAME > $BACKUP_DIR/$DB_NAME-$DATE.sql

# Optional: compress the backup file
gzip $BACKUP_DIR/$DB_NAME-$DATE.sql

# ### Make the script executable:
# chmod +x /backups/mysql/daily_backup.sh