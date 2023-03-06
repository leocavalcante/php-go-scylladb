package main

import (
	"context"
	"github.com/gocql/gocql"
	"net"
	"net/rpc"

	goridgeRpc "github.com/roadrunner-server/goridge/v3/pkg/rpc"
)

var (
	ctx  context.Context
	sess *gocql.Session
)

type ScyllaDB struct{}
type Todo struct {
	Id    string `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func (s *ScyllaDB) Query(cql string, r *[]Todo) error {
	*r = []Todo{}
	scanner := sess.Query(cql).WithContext(ctx).Iter().Scanner()

	for scanner.Next() {
		t := new(Todo)
		err := scanner.Scan(&t.Id, &t.Title)
		if err != nil {
			return err
		}
		*r = append(*r, *t)

	}

	return nil
}

func main() {
	ctx = context.Background()

	ln, err := net.Listen("unix", "/tmp/php-go-scylladb.sock")
	if err != nil {
		panic(err)

	}

	scylladb := gocql.NewCluster("scylladb:9042")
	scylladb.Keyspace = "php_go_scylladb"
	sess, err = scylladb.CreateSession()
	if err != nil {
		panic(err)
	}

	defer sess.Close()
	_ = rpc.Register(new(ScyllaDB))

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		_ = conn
		go rpc.ServeCodec(goridgeRpc.NewCodec(conn))
	}
}
