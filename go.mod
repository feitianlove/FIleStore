module github.com/feitianlove/FIleStore

go 1.13

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

require (
	github.com/feitianlove/golib v0.0.0-20210113134317-5a1362a4ac36
	github.com/garyburd/redigo v1.6.2
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/streadway/amqp v0.0.0-20190827072141-edfb9018d271
	github.com/ugorji/go v1.2.3 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210113131315-ba0562f347e0 // indirect
	gopkg.in/amz.v1 v1.0.0-20150111123259-ad23e96a31d2
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
