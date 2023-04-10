docker:
	docker run -d -e POSTGRES_PASSWORD=123456 -p 5434:5432 --name pg-example postgres:14.5

db:
	sleep 3 \
	&& docker exec -it pg-example createdb -U postgres udemy

psql:
	docker exec -it pg-example psql -U postgres
# \l: list
# \c udemy: connect a db
# \dt: display tables

up: docker db

down:
	docker stop pg-example \
	&& docker rm pg-example