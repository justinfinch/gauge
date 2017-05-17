# gauge

## DB Setup

```
vagrant up

vagrant ssh

sudo su -

cd /vagrant

docker exec -it vagrant_roach1_1 ./cockroach user set gaugeapp --insecure

docker exec -it vagrant_roach1_1 ./cockroach sql --insecure -e 'CREATE DATABASE gaugedb'

docker exec -it vagrant_roach1_1 ./cockroach sql --insecure -e 'GRANT ALL ON DATABASE gaugedb TO gaugeapp'

exit
exit
```

## Run Local

go run main.go db-migrate
go run main.go serve -p 8080


