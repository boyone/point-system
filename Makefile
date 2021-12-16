run_code_analysis:
	golangci-lint run
run_unit_tests:
	go test -v -coverprofile=coverage.out ./... 2>&1 | go-junit-report > coverage.xml
start_application:
	docker compose up -d
run_api_tests:
	newman run Point.postman_collection.json
stop_application:
	docker compose down