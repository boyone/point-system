run_code_analysis:
	golangci-lint run --issues-exit-code=0 --out-format junit-xml ./... > golangci-lint-junit-report.xml
run_unit_tests:
	go test -v -coverprofile=coverage.out ./... 2>&1 | go-junit-report > coverage.xml
start_application:
	docker compose up -d
run_api_tests:
	newman run Point.postman_collection.json
stop_application:
	docker compose down