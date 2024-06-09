package config

import "github.com/gocql/gocql"

func NewSession() (*gocql.Session, error){
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "todoapp"
	cluster.Consistency = gocql.Quorum

	return cluster.CreateSession()
}