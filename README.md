[![Build Status](https://travis-ci.org/Kalimaha/toyrobot-go.svg?branch=master)](https://travis-ci.org/Kalimaha/toyrobot-go)
[![Coverage Status](https://coveralls.io/repos/github/Kalimaha/toyrobot-go/badge.svg?branch=master)](https://coveralls.io/github/Kalimaha/toyrobot-go?branch=master)

# Toy Robot

Go implementation of the famous [Toy Robot coding challenge](./PROBLEM.md).

## Test

### With Docker

Build the container with `docker build -t toy-robot-go .`, then:

```
docker run -it toy-robot-go ./simon.says test
```

### Without Docker

```
./simon.says test
```

## Build

### With Docker

The following command will generate an executable `toyrobot` file inside the container:

```
docker run -it toy-robot-go ./simon.says build
```

### Without Docker

The following command will generate an executable `toyrobot` file:

```
./simon.says build
```

## Run

### With Docker

```
docker run -it toy-robot-go go run ./cmd/toyrobot.go resources/exampleA.txt
```

### Without Docker

The executable `toyrobot` file requires the absolute path to the file containing the instructions for the Toy Robot as 
input, e.g.:

```
./toyrobot /tmp/example.txt
```
