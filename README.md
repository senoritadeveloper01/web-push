# Web Push Service using Firebase written in Go

A sample Firebase web push service written in go1.19 darwin/amd64.

<img src="https://img.shields.io/badge/Language-Go-orange.svg">

## Dependencies / Libraries

For development:\
[mysql](https://github.com/go-sql-driver/mysql) MySQL is used for database operations.\
[logrus](https://github.com/sirupsen/logrus) library for logging purposes,\
[viper](https://github.com/spf13/viper) for application properties management,\
[firebase messaging v4](firebase.google.com/go/v4/messaging) for registration of devices and sending notifications,\
[sarama](https://github.com/Shopify/sarama) for Apache Kafka,\
[gorilla/mux](https://github.com/gorilla/mux) for request routing,\
[gorilla/handlers](https://github.com/gorilla/handlers) as an HTTP middleware\

For testing:\
[go-sqlmock](https://github.com/DATA-DOG/go-sqlmock) as an sql mock library,\
[testify](https://github.com/stretchr/testify) for assertions and mocks

## Installation

You can either use Docker with Docker Compose using docker-compose.yml file under deployments/docker.
Run the command below after navigating to deployments/docker directory:

```
docker-compose up -d
```

After you see the containers are up and running, you can run the commands below in order to get Kafka ready:

```
docker exec docker_kafka-1_1 kafka-topics --create --bootstrap-server localhost:29092 --replication-factor 1 --partitions 1 --topic client-event-topic
docker exec -it --tty docker_kafka-1_1 kafka-console-producer --broker-list localhost:29092 --topic client-event-topic
docker exec -it --tty docker_kafka-1_1 kafka-console-consumer --bootstrap-server localhost:29092 --topic client-event-topic --from-beginning
```

For MySQL:

```
docker exec -it db mysql -u testDB -p password -h 127.0.0.1
```

Now, navigate to cmd folder and there you can run the application:

```
go run main.go
```

To test service register user API call, you can perform a cURL request (or import it to Postman):

```
curl -L 'http://localhost:8080/api/users/create' \
--data-raw '{
    "userId": "senorita@developer.com",
    "clientId": "54aa74b5-96b2-4d08-b2af-a97decbfb9c4",
    "credentials": "c2DviqDe-R7dJblMnC2v-qHRWCAjttgEJTZ0Qu-4NKazKxjfx"
}'
```

## Screenshots

![ScreenShot](https://raw.github.com/senoritadeveloper01/web-push/master/screenshots/web-push-screenshots.png)

## Authors / Contributors / Credits

**Nil Seri**

[Github 1](https://github.com/senoritadeveloper01)

[Github 2](https://github.com/nilseri01)

You can visit [my Medium profile](https://senoritadeveloper.medium.com/)

You can specifically visit [this post](https://senoritadeveloper.medium.com/firebase-web-push-notification-with-go-and-angular-811698cffe70) where you can find a post giving details about implementing this project.

## Copyright & Licensing Information

This project is licensed under the terms of the MIT license.
