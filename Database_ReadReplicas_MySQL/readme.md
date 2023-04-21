## Setup read replicas
1. On master
```sql
CREATE USER 'replication_user'@'%' IDENTIFIED WITH mysql_native_password BY 'replica_password';

GRANT REPLICATION SLAVE ON *.* TO 'replication_user'@'%';

SHOW MASTER STATUS;
```

2. On replica
```sql
CHANGE MASTER TO MASTER_HOST='db_master', MASTER_USER='replication_user', MASTER_PASSWORD='replica_password',MASTER_LOG_FILE='filename_from_master', MASTER_LOG_POS=position_from_master;

START SLAVE;

SHOW SLAVE STATUS\G
-- Slave_IO_Running: Yes
-- Slave_SQL_Running: Yes

SHOW VARIABLES LIKE 'read_only'; -- Can set manually: SET GLOBAL read_only = ON;
```

## Doing backups-restores
- A common approach to managing backup retention is to implement a tiered retention policy, which involves keeping more recent backups for a shorter period and older backups for a longer period. For example:
    - Keep daily backups for 7-14 days
    - Keep weekly backups for 4-8 weeks
    - Keep monthly backups for 3-12 months
    - Keep yearly backups indefinitely or for a specific number of years

- Step by step
    - `From bkp_VM`: mysqldump -h $HOST -u $DB_USER -p$DB_PASS $DB_NAME > $BACKUP_DIR/$DB_NAME-$DATE.sql
    - `From master_db`: mysql -u root -pmysecretpassword -e "CREATE DATABASE test_bkp;"
    - `From bkp_VM`: mysql -h db_master -u root -pmysecretpassword test_bkp < /backups/mysql/daily/test-2023-04-21.sql

## Schedule backups
- schedule the daily backup at 2:00 AM and the weekly backup at 3:00 AM every Sunday:
    - crontab -e
    - 0 2 * * * /backups/mysql/daily_backup.sh
    - 0 3 * * 0 /backups/mysql/weekly_backup.sh

## Delete backups
- Ones older than 30 days: `find /backups/mysql/daily -type f -mtime +30 -exec rm {} \;`