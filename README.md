# gauge


## Run Local

go run main.go serve -p 8080


## DB Setup

```
vagrant up

sudo su -

cd /vagrant

docker compose up -d

docker exec -it vagrant_roach1_1 ./cockroach user set gaugeapp --insecure

docker exec -it vagrant_roach1_1 ./cockroach sql --insecure -e 'CREATE DATABASE gaugedb'

docker exec -it vagrant_roach1_1 ./cockroach sql --insecure -e 'GRANT ALL ON DATABASE gaugedb TO gaugeapp'
```