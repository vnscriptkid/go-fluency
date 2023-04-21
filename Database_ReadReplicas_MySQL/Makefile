up:
	docker compose up -d

up_build:
	docker compose up --build -d

down:
	docker compose down

cli-master:
	docker exec -it db_master mysql -u root -pmysecretpassword

# Replica with the 'root' user, which has the SUPER privilege. 
# Users with the SUPER privilege can still create tables and modify data, 
# even when the read_only mode is enabled.
cli-replica-root:
	docker exec -it db_replica_1 mysql -u root -pmysecretpassword

cli-bkp:
	docker exec -it db_bkp /bin/bash

replica-user:
	docker exec -it db_replica_1 mysql -u root -pmysecretpassword -e "CREATE USER 'test_user'@'%' IDENTIFIED BY 'test_password'; GRANT SELECT ON *.* TO 'test_user'@'%'; FLUSH PRIVILEGES;"

cli-replica:
	docker exec -it db_replica_1 mysql -u test_user -ptest_password

# SHOW DATABASES;
# CREATE DATABASE test;
# USE test;
# DROP DATABASE test;
# SELECT DATABASE();

# ERROR 1290 (HY000): The MySQL server is running with the --read-only option so it cannot execute this statement