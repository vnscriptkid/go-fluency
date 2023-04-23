up:
	docker compose up -d

up_build:
	docker compose up -d --build

down:
	docker compose down

db:
	docker exec -it postgres1 createdb -U user shard
	docker exec -it postgres2 createdb -U user shard

table:
	docker exec -it postgres1 psql -U user -d shard -c "CREATE TABLE users (id serial PRIMARY KEY, name text);"
	docker exec -it postgres2 psql -U user -d shard -c "CREATE TABLE users (id serial PRIMARY KEY, name text);"

psql1:
	docker exec -it postgres1 psql -U user -d shard

psql2:
	docker exec -it postgres2 psql -U user -d shard
