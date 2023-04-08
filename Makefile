bench:
	@go test -bench=. -benchmem -benchtime=1x -tags musl,go_tarantool_ssl_disable
	# @benchstat /tmp/o
	
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

benchstat:
	@go get golang.org/x/perf/cmd/benchstat
	@go install  golang.org/x/perf/cmd/benchstat