module crontab

go 1.13

replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc v1.29.1 => google.golang.org/grpc v1.26.0

require (
	github.com/coreos/bbolt v1.3.4 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/gorhill/cronexpr v0.0.0-20180427100037-88b0669f7d75
	github.com/prometheus/client_golang v1.9.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.mongodb.org/mongo-driver v1.4.6
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	google.golang.org/grpc v1.29.1 // indirect
)
