docker:
	docker run -d -e POSTGRES_PASSWORD=123456 -p 5438:5432 --name pg-gorm postgres:14.5

db:
	sleep 3 \
	&& docker exec -it pg-gorm createdb -U postgres test

psql:
	docker exec -it pg-gorm psql -U postgres -d test
# \l: list
# \c udemy: connect a db
# \dt: display tables
# \d test

up: docker db

down:
	docker stop pg-gorm \
	&& docker rm pg-gorm