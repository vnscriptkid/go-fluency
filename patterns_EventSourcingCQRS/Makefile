docker_pg:
	docker run -d -e POSTGRES_PASSWORD=123456 -p 5435:5432 --name pg-example-event-sourcing postgres:14.5

kafka:
	docker network create kafka-network \
	&& docker run -d --network=kafka-network --name=zookeeper -p 2181:2181 -e ZOOKEEPER_CLIENT_PORT=2181 confluentinc/cp-zookeeper \
	&& docker run -d --network=kafka-network -p 9092:9092 --name=kafka -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
	-e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092 confluentinc/cp-kafka

topic:
	docker exec kafka kafka-topics --create --topic test-topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092

del:
	docker exec kafka kafka-topics --bootstrap-server kafka:9092 --delete --topic test-topic

list:
	docker exec kafka kafka-topics --bootstrap-server localhost:9092 --list

desc:
	docker exec kafka kafka-topics --bootstrap-server localhost:9092 --describe --topic test-topic

read_topic:
	docker exec kafka kafka-console-consumer --topic test-topic --bootstrap-server localhost:9092 --from-beginning

send_msg:
	docker exec -it kafka kafka-console-producer --topic test-topic --bootstrap-server localhost:9092


db_init:
	sleep 3 \
	&& docker exec -it pg-example-event-sourcing createdb -U postgres readdb \
	&& docker exec -it pg-example-event-sourcing createdb -U postgres writedb

psql1:
	docker exec -it pg-example-event-sourcing psql -U postgres -d bankread

psql2:
	docker exec -it pg-example-event-sourcing psql -U postgres -d bankwrite
# \l: list
# \c udemy: connect a db
# \dt: display tables
# \d udemy

up: docker db_init

down:
	docker stop pg-example-event-sourcing \
	&& docker rm pg-example-event-sourcing

down_kafka:
	docker stop kafka zookeeper \
	&& docker rm kafka zookeeper \
	&& docker network rm kafka-network \
	