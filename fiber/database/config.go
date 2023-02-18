package database

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/table"
)

func cluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "store"
	return cluster
}

var messageTable = table.New(table.Metadata{
	Name:    "message",
	Columns: []string{"create_ts", "content"},
})
