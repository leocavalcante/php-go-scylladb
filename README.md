# PHP Go ScyllaDB
:ghost: Proof of concept on making CQL queries to ScyllaDB from PHP using Spiral's Goridge and Go's shard-aware driver.

## Getting started

### The development environment

#### Host

##### Build the image with PHP and Go
```shell
make image
```

##### Run the services
```shell
make up
```

##### Start an interactive shell
```shell
make sh
```

#### Container

##### Install PHP dependencies
```shell
composer install
```

##### Install Go dependencies
```shell
go mod tidy
```
