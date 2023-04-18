setup:
	docker run -d --name rabbitmq-demo -p 5672:5672 -p 15672:15672 --hostname my-rabbit rabbitmq:3.11.13-management

clean:
	docker stop rabbitmq-demo
	docker rm rabbitmq-demo

# http://localhost:15672
