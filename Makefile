.PHONY: run-test
run-test:
	docker-compose down
	docker-compose up todo_api_db -d
	docker-compose run --rm todo_api sh -c "go run todo migrate && go run todo seed && go test -race -v ./... "
	docker-compose down

.PHONY: run-lint-check
run-lint-check:
	docker-compose run --rm todo_api sh -c "gofmt -l ."

.PHONY: run-lint
run-lint:
	docker-compose run --rm todo_api sh -c "gofmt -w ."	
	

.PHONY: seed_db
seed_db:
	docker-compose run --rm todo_api sh -c "go run todo seed"