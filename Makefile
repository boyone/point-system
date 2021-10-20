run_unit_tests:
	go test ./...
start_application:
	docker compose up -d
run_api_tests:
	newman run Point.postman_collection.json
stop_application:
	docker compose down