Simple GoLang RabbitMQ integration.
How to run:
1. Please ensure, that you've installed and running docker and run the following command: `docker run -d --hostname my-rabbit --name local-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:4-management`
2. Open the project in two separate terminals.
3. Run `go run receiver.go`, wait to start
4. Run `go run sender.go` and observe how the receiver reacts
