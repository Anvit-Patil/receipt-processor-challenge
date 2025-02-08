# Receipt Processor API

A simple Go web service to process receipts and calculate points.

## Installation

1. Install Go 1.18+
2. Clone the repo:
   ```sh
   git clone https://github.com/your-username/receipt-processor.git
   cd receipt-processor
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```

## API Endpoints

### Process a Receipt

**POST** `/receipts/process`

#### Example Request:

```sh
curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{ "retailer": "Target", "total": "18.74" }'
```

#### Example Response:

```json
{ "id": "1234-5678" }
```

### Get Receipt Points

**GET** `/receipts/{id}/points`

#### Example Request:

```sh
curl -X GET http://localhost:8080/receipts/1234-5678/points
```

#### Example Response:

```json
{ "points": 32 }
```

## License

MIT License © 2025 Anvit Rajesh Patil <github.com/Anvit-Patil>

