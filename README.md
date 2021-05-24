# needys-output-producer
An API micro-service for needys application which produces pdf output
when receiving a message from rabbitmq "hello" queue from / exchange.

Please contact <guillaume.penaud@gmail.com> if you have questions !

##### manage the compose stack

```
### to start it
docker-compose up --detach
OR
make start

### to stop it
docker-compose down
OR
make stop
```

##### to start only a part of the stack
```
### only the api part
docker-compose up needys-output-producer
OR
make api-only

### only the sidecars part
docker-compose up mariadb rabbitmq
OR
make sidecars-only
```
