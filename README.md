# Go Commerce

## Description

A RESTful API for an e-commerce platform using Go.

## Features

- User authentication
- Product management
- Order management
- Payment processing

## Prerequisites

- Go >=1.23.2
- MySQL >= 8.0.41

## Installation

```bash
$ git clone https://github.com/DanielJohn17/go-commerce.git
$ cd go-commerce
```

## Usage

Create a `.env` file in the root directory and add the following environment variables

```bash
DB_USER=<DB_USER>
DB_PASSWORD=<DB_PASSWORD>
DB_HOST=<DB_HOST>
DB_PORT=:<DB_PORT>
DB_Name=<DB_NAME>
JWT_EXP=<JWT_EXP> //in seconds example: 3600 => 1hr
JWT_SECRET=<JWT_SECRET>
```

Run the following commands

```bash
$ go mod tidy
$ make migrate-up
$ make run
```

### Docker

Follow the instructions in [README.Docker.md](README.Docker.md) to run the application in a Docker container.

## Acknowledgements

- [Tiago Youtube Channel](https://www.youtube.com/@TiagoTaquelim)
