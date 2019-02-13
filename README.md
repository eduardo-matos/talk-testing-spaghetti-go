# Testing Spaghetti Go

## Installing

```
go get -t ./...
```

## Running RabbitMQ

Run
```sh
# It may take a while to be accessible
docker run -p 15672:15672 -p 5672:5672 -p 25676:25676 rabbitmq:3-management
```

Head to `http://localhost:15672` and login (user/password is `guest`/`guest`).


Create a queue named `emails`.


You may publish messages so this app will be able to consume it.


Example message: `{"to": "foo@bar.spam", "body": "Spamming you", "subject": "You're our 9.999.999th visitor!"}`

![rabbitmq-creating-queue-and-publishing-message](https://user-images.githubusercontent.com/483681/52729932-f268ce00-2f98-11e9-89e8-609b019fedcd.gif)

## Running

```
go run main.go
```
