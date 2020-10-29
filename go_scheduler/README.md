# Go Scheduler

| Samples Name            | Description           |
| :---------------------- | :-------------------- |
| [countdown](#countdown) | Parallel countdowns   |
| [servers](#servers)     | Parallel HTTP servers |

## Running samples

To make easier to execute the commands, you can devine an alias:

```shell
alias gs='go run ./main.go'
```

To run the samples, use the command below:

```sh
gs <sample-name> [OPTIONS]
```

## Samples

### Countdown

An app with 2 countdowns in parallel

To run the app, use the command below:

```shell
gs countdown --c1=<counter-1-value> --c2=<counter-2-value>
```

#### Options

| Required? | Option | Description             | Default |
| --------- | ------ | ----------------------- | ------- |
| NO        | c1     | Counter 1 initial value | 10      |
| NO        | c2     | Counter 2 initial value | 10      |


### Servers

An app with 2 HTTP servers.

To run the app, use the command below:

> You can stop the app at any time by pressing `ctrl+c`

```shell
gs servers --port1=<server1-port> --port2=<server2-port>
```

#### Options

| Required? | Option | Description   | Default |
| --------- | ------ | ------------- | ------- |
| NO        | port1  | Server 1 port | 3000    |
| NO        | port2  | Server 2 port | 3001    |