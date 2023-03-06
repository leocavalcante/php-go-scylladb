module github.com/leocavalcante/php-go-scylladb

go 1.19

require (
	github.com/gocql/gocql v1.3.1
	github.com/roadrunner-server/goridge/v3 v3.6.2
)

require (
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/cpuid/v2 v2.0.3 // indirect
	github.com/roadrunner-server/errors v1.2.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.10.0
