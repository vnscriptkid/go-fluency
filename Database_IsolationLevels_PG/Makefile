docker:
	docker run -d -e POSTGRES_PASSWORD=123456 -p 5435:5432 --name pg-example-isolation postgres:14.5

db:
	sleep 3 \
	&& docker exec -it pg-example-isolation createdb -U postgres test \
	&& docker exec -it pg-example-isolation psql -U postgres -d test -c "CREATE TABLE mytable (id SERIAL, name TEXT);"

psql:
	docker exec -it pg-example-isolation psql -U postgres -d test
# \l: list
# \c udemy: connect a db
# \dt: display tables
# \d udemy
# SHOW TRANSACTION ISOLATION LEVEL;
# SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;

up: docker db

down:
	docker stop pg-example-isolation \
	&& docker rm pg-example-isolation