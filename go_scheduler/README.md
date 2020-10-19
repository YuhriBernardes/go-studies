# Go Scheduler

To run the samples, use the command above

```shell
cd go_scheduler
go run ./main.go <sample-name> <sample-specific-parameters>
```

You can also build the cli as described below:

```shell
go build -o gs main.go
```

So you can use as:
```shell
gs <sample-name> <sample-specific-parameters>
```

> Some samples doesn't have parameters

|      Samples Name       |
| :---------------------: |
| [countdown](#countdown) |

## Samples

### Countdown

An app with 2 countdowns in parallel

To run the app, use the command below:

```shell
go run ./main.go countdown --c1=<counter-1-value> --c2=<counter-2-value>
```

or

```shell
gs countdown --c1=<counter-1-value> --c2=<counter-2-value>
```

#### Parameters

| Required? | Parameter | Description             | Default |
| --------- | --------- | ----------------------- | ------- |
| NO        | c1        | Counter 1 initial value | 10      |
| NO        | c2        | Counter 2 initial value | 10      |