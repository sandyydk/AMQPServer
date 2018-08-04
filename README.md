# AMQPServer
Sample AMQP Server which publishes to an exchange called 'events' a general sample message

Note :-
By default AMQP default user guest is allowed to be connected only from the localhost and not remotely. This can come in the way if you launch using Docker. Hence create a new user and give appropriate rights and connect to it. Management UI has no such issues though.

Docker command to launch amqp :-
docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
