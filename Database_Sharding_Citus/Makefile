up:
	docker compose up -d

down:
	docker compose down

psql:
	docker exec -it citus_master psql -U user -d postgres

psql1:
	docker exec -it citus_worker1 psql -U user -d postgres

psql2:
	docker exec -it citus_worker2 psql -U user -d postgres

config:
	docker exec -it citus_master psql -U user -d postgres -c "CREATE EXTENSION IF NOT EXISTS citus;"
	docker exec -it citus_worker1 psql -U user -d postgres -c "CREATE EXTENSION IF NOT EXISTS citus;"
	docker exec -it citus_worker2 psql -U user -d postgres -c "CREATE EXTENSION IF NOT EXISTS citus;"
	docker exec -it citus_master psql -U user -d postgres -c "SELECT * from master_add_node('worker1', 5432);"
	docker exec -it citus_master psql -U user -d postgres -c "SELECT * from master_add_node('worker2', 5432);"

table:
	docker exec -it citus_master psql -U user -d postgres -c "CREATE TABLE users (id bigint, name text, PRIMARY KEY (id));"
	docker exec -it citus_master psql -U user -d postgres -c "SELECT create_distributed_table('users', 'id');"


SELECT part.logicalrelid, part.replication_factor
FROM pg_dist_partition part
JOIN pg_dist_shard shard ON part.logicalrelid = shard.logicalrelid
WHERE part.logicalrelid::regclass::text = 'users'
GROUP BY part.logicalrelid, part.replication_factor;