bench:
	@go test -bench=.
	
run:
	@docker-compose -f docker/docker-compose.yaml up -d 