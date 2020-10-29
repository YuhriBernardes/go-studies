# Kafka

| Study         | Description                           |
| :------------ | :------------------------------------ |
| [Loop](#loop) | Produces a message in a infinite loop |

Before run any sample, start the docker images using the command below:

> You can also use any kafka broker with these samples

```shell
docker-compose -f ./docker-compose.kafka.yaml -d
```

## Running samples

To make easier to execute the commands, you can devine an alias:

```shell
alias kfk='go run ./main.go'
```

To run the samples you can use the command below:

```shell
kfk [GLOBAL-OPTIONS] <study> [OPTIONS]
```

### Global Options
| Required? | Option | Description                                 | Default        |
| --------- | ------ | ------------------------------------------- | -------------- |
| NO        | bs     | Comma separated list of kafka brokers hosts | localhost:9092 |

### Examples

```sh
kfk -bs="other-servers,comma-separated" loop
```

## Samples

### Loop Consumer

The app runs a consumer and a porducer related to the same topic. The produces will often send a new message to the topic and the consumer will log them.

> The topic is created before the app start and deleted during the app's shutdown

To run this study, execute the command below:

```sh
kfk [GLOBAL-OPTIONS] loop [OPTIONS]
```

#### Options

| Required? | Option | Description       | Default |
| --------- | ------ | ----------------- | ------- |
| NO        | topic  | Topic name to use | 3000    |