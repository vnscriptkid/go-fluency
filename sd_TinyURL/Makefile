up:
	docker compose up -d

psql1:
	docker exec -it postgres1 psql -U user -d shard1

psql2:
	docker exec -it postgres2 psql -U user -d shard2

truncate:
	docker exec -it postgres1 psql -U user -d shard1 -c "TRUNCATE ranges;"
	docker exec -it postgres2 psql -U user -d shard2 -c "TRUNCATE ranges;"

drop:
	docker exec -it postgres1 psql -U user -d shard1 -c "DROP TABLE ranges;"
	docker exec -it postgres2 psql -U user -d shard2 -c "DROP TABLE ranges;"