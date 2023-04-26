start:
	docker rm $(docker ps -aqf "name=replica_1") || true
	docker volume rm $(docker volume ls -qf "name=sd_productlisting_replica_data") || true
	docker compose up db -d
	docker exec db bash -c "echo 'host replication repuser all trust' >> /var/lib/postgresql/data/pg_hba.conf"
	docker-compose restart db
	docker compose up db -d
	docker compose up replica_1 -d
	docker compose up catalog_app_1 catalog_app_2 catalog_app_3 -d
	docker compose up load_balancer -d

start_build:
	docker rm $(docker ps -aqf "name=replica_1") || true
	docker volume rm $(docker volume ls -qf "name=sd_productlisting_replica_data") || true
	docker compose up db -d
	docker exec db bash -c "echo 'host replication repuser all trust' >> /var/lib/postgresql/data/pg_hba.conf"
	docker-compose restart db
	docker compose up db -d
	docker compose up replica_1 -d
	docker compose up catalog_app_1 catalog_app_2 catalog_app_3 --build -d
	docker compose up load_balancer -d


up:
	docker compose up -d

up_build:
	docker compose up -d --build

psql:
	docker exec -it pg-product-listing psql -U postgres

down:
	docker compose down

psql_master:
	docker exec -it db psql -U postgres

psql_replica:
	docker exec -it replica_1 psql -U postgres
