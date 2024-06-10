package config

import "github.com/gocql/gocql"

func NewSession() (*gocql.Session, error){
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "todoapp"
	cluster.Consistency = gocql.Quorum

	return cluster.CreateSession()
}