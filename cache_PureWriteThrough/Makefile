up:
	docker compose up -d
	sleep 2
	docker exec wt-pg createdb -U postgres mydb
	docker exec wt-pg psql -U postgres -d mydb -c "CREATE TABLE IF NOT EXISTS mytable (name text primary key, score int);"
	docker exec wt-pg psql -U postgres -d mydb -c "INSERT INTO mytable (name, score) VALUES ('james', 10);"
	docker exec wt-pg psql -U postgres -d mydb -c "INSERT INTO mytable (name, score) VALUES ('henry', 10);"
	docker exec wt-pg psql -U postgres -d mydb -c "INSERT INTO mytable (name, score) VALUES ('jerry', 10);"

down:
	docker compose down

