# Golang Key Value Database

A lightweight and intuitive, in-memory key-value store available through TCP written in Go.

## Features

- **SET**: Store a key-value pair in the database, with an optional time-to-live (TTL) value.
  - Example Command : `SET <key> <value> <ttl>`
- **GET**: Retrieve the value associated with a given key from the database.
  - Example Command: `GET <key>`
- **DELETE**: Remove a key-value pair from the database.
  - Example Command: `DELETE <key>`

## Getting Started

To get started, follow these steps:

1. Install `Go` version `1.23`
2. Clone the repository: `git clone https://github.com/sebastian-nunez/golang-key-value-db`
3. Start database server: `make run`

## Usage

Once the database server is up and running, you can interact with it using any `tcp` client like `telnet` or `nc`. Here are some examples:

```shell
$ nc localhost 3000
$ SET foo bar 300
$ OK ( Server Response )
$ GET foo
$ bar ( Server Response )
```
