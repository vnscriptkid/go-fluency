start:
	docker-compose up -d

stop:
	docker-compose down

create-topic:
	docker exec -it kafka kafka-topics --create --topic $(topic) --bootstrap-server kafka:9092 --replication-factor 1 --partitions 1

delete-topic:
	docker exec -it kafka kafka-topics --delete --topic $(topic) --bootstrap-server kafka:9092

list-topics:
	docker exec -it kafka kafka-topics --list --bootstrap-server kafka:9092

consume:
	docker exec -it kafka kafka-console-consumer --topic $(topic) --from-beginning --bootstrap-server kafka:9092

produce:
	docker exec -it kafka kafka-console-producer --topic $(topic) --bootstrap-server kafka:9092

.PHONY: start stop create-topic delete-topic list-topics