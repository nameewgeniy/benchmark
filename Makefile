bench:
	@go test -bench=. -tags musl,go_tarantool_ssl_disable
	
run:
	@docker-compose -f docker/docker-compose.yaml up -d 

run.tarantool:
	@docker-compose -f docker/docker-compose.yaml up -d tarantool

run.kafka:
	@docker-compose -f docker/docker-compose.yaml up -d kafka

run.etcd:
	@docker-compose -f docker/docker-compose.yaml up -d etcd

run.mongodb:
	@docker-compose -f docker/docker-compose.yaml up -d mongodb

run.clickhouse:
	@docker-compose -f docker/docker-compose.yaml up -d clickhouse

run.postgres:
	@docker-compose -f docker/docker-compose.yaml up -d postgres

stop:
	@docker-compose -f docker/docker-compose.yaml stop

down:
	@docker-compose -f docker/docker-compose.yaml down