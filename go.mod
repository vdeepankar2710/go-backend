module todo-backend

go 1.22.4

require github.com/gocql/gocql v1.6.0

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.1
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.14.0
