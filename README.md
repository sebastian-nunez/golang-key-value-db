# Golang key-value database (KVDB)

A lightweight and intuitive, in-memory, key-value store available through a [TCP](<https://www.fortinet.com/resources/cyberglossary/tcp-ip#:~:text=Transmission%20Control%20Protocol%20(TCP)%20is,exchange%20messages%20over%20a%20network.>) connection written in Go.

## Features/API

- **SET**: Store a key-value pair in the database, with an optional time-to-live (TTL) value.
  - Example : `SET <key> <value> <ttl>`
- **GET**: Retrieve the value associated with a given key from the database.
  - Example: `GET <key>`
- **DELETE**: Remove a key-value pair from the database.
  - Example: `DELETE <key>`

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
