# API Specification

## Calculate Point

### Request

- Method: POST
- Path: /api/v1/calculatePoint
- Content-Type: application/json
- Body:

  ```json
  {
    "value": 99
  }
  ```

### Response

- Status Code: 200
- Content-Type: application/json
- Body:

  ```json
  {
    "value": 0
  }
  ```

## Command line Workflow

- run unit tests

  ```sh
    go test ./...
  ```

- start application

  ```sh
  go run main.go
  ```

- run api tests

  ```sh
  newman run Point.postman_collection.json
  ```

- call api with `curl`

  ```sh
  curl --location --request POST 'http://localhost:3000/api/v1/calculatePoints' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "value": 99
  }'
  ```
