#!/bin/bash

DATE=$(date +%Y-%m-%d)
BACKUP_DIR="/backups/mysql/daily"
DB_USER="root"
DB_PASS="mysecretpassword"
DB_NAME="test"
HOST="db_master"

mysqldump -h $HOST -u $DB_USER -p$DB_PASS $DB_NAME > $BACKUP_DIR/$DB_NAME-$DATE.sql

# Optional: compress the backup file
gzip $BACKUP_DIR/$DB_NAME-$DATE.sql

# ### Make the script executable:
# chmod +x /backups/mysql/daily_backup.sh

# gzip -d test-2023-04-21.sql.gz
# From master_db: mysql -u root -pmysecretpassword -e "CREATE DATABASE test_bkp;"
# From bkp_VM: mysql -h db_master -u root -pmysecretpassword test_bkp < /backups/mysql/daily/test-2023-04-21.sql
