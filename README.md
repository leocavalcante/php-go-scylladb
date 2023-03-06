# PHP Go ScyllaDB
:ghost: Proof of concept on making CQL queries to ScyllaDB from PHP using Spiral's Goridge and Go's shard-aware driver.

## Getting started

### The development environment

#### :computer: On your host machine
| Step | Command       | Description |
| --- |---------------| --- |
| 1. | `make image`  | Build the image with PHP and Go |
| 2. | `make up` | Run services |
| 3. | `make sh` | Start an interactive shell |
| 4. | `make cqlsh` | Start an interactive CQL shell inside ScyllaDB |


#### :whale: Insde the container
| Step | Command | Description                                                            |
| --- | --- |------------------------------------------------------------------------|
| 1. | `composer install` | Install PHP dependencies                                               |
| 2. | `go mod tidy` | Install Go dependencies                                                |
| 3. | `go run -mod=readonly .` | Run Go app (`-mod=readonly` is to avoid `vendor/` directory confusion) |
| 4. | `php app.php start` | Run PHP app                                                            |

### CQLs (for after `make cqlsh`)
```cassandraql
CREATE KEYSPACE php_go_scylladb
  WITH REPLICATION = { 
   'class' : 'SimpleStrategy', 
   'replication_factor' : 1 
  };
```

```cassandraql
CREATE TABLE php_go_scylladb.todos ( 
   id UUID PRIMARY KEY, 
   title text);
```

```cassandraql
INSERT INTO php_go_scylladb.todos (id, title)
  VALUES (465b5ec8-a580-4a8d-ba3d-69b8a66ed50c, 'PHP 🤝 Go');
```
