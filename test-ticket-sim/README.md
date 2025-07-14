# Redis Ticket Purchase Test Project

This is a ticket purchase system simulator designed to test the effectiveness of Redis atomic write operations, ensuring that purchase operations can be completed correctly and efficiently in high-concurrency environments.

## Project Overview

This project simulates a ticket purchase system using Redis atomic operations to ensure accurate inventory management. The main testing features include:

- Reliability of Redis atomic decrement operations (`DECR`)
- Inventory management under high concurrency
- Purchase result recording and tracking
- System stability and performance

## Project Architecture

```
test-ticket-sim/
├── main.go              # Main program entry, HTTP server
├── stock/
│   └── stock.go         # Inventory management logic
├── cache/
│   └── cache.go         # Redis client initialization
├── record/
│   └── record.go        # Purchase record management
├── public/              # Static assets and screenshots
├── docker-compose.yml   # Redis container configuration
├── go.mod               # Go module dependencies
└── go.sum               # Dependency checksums
```

## Core Features

### 1. Inventory Preload

- Uses `SET` command to initialize ticket inventory
- Supports custom inventory quantities

### 2. Purchase Processing

- Uses `DECR` atomic operation to reduce inventory
- Automatically checks if inventory is sufficient
- Returns purchase success or failure results

### 3. Record Tracking

- Records detailed information for each purchase
- Includes timestamp, result status, and unique ID
- Supports querying all purchase records

## Technical Features

### Redis Atomic Operations

```go
// Using DECR atomic decrement operation
result, err := rds.Decr(ctx, ticketKey).Result()
if result < 0 {
    // Insufficient inventory, purchase failed
    return false
}
```

### Concurrency Safety

- Uses Redis atomic operations to ensure data consistency
- Local records use mutex to protect shared data

### Unique Identification

- Uses UUID v7 to generate unique IDs for each purchase
- Facilitates tracking and debugging

## Environment Requirements

- Go 1.23.5 or higher
- Redis 6.0 or higher
- Docker (optional, for quick Redis startup)

## Quick Start

### 1. Start Redis

Using Docker Compose:

```bash
docker-compose up -d
```

Or start Redis directly:

```bash
redis-server
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Run the Program

```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Endpoints

### Purchase Ticket

```
POST /buy
```

**Response Example:**

```json
{
  "id": "018f1234-5678-9abc-def0-123456789abc",
  "timestamp": 1703123456789,
  "result": "true"
}
```

### Query Purchase Records

```
GET /records
```

**Response Example:**

```json
[
  {
    "id": "018f1234-5678-9abc-def0-123456789abc",
    "timestamp": 1703123456789,
    "result": "true"
  },
  {
    "id": "018f1234-5678-9abc-def0-123456789def",
    "timestamp": 1703123456790,
    "result": "false"
  }
]
```

## Performance Testing

### Load Test Screenshot

<img src="public/截圖%202025-07-14%20下午6.50.23.png" width="500"/>

The above screenshot shows the results of our load testing, demonstrating the system's performance under high concurrency conditions.

### Load Testing with oha

You can use the [`oha`](https://github.com/hatoo/oha) HTTP load testing tool to perform stress testing:

```bash
oha https://localhost:8080/buy
```

This command will send concurrent requests to the purchase endpoint and provide detailed performance metrics including:

- Request rate (requests per second)
- Response time statistics
- Success/failure rates
- Throughput information

## Monitoring and Debugging

### Redis Monitoring

```bash
# Connect to Redis CLI
redis-cli

# Check ticket inventory
GET ticket:stock

# Monitor Redis operations
MONITOR
```

### Program Logs

The program outputs detailed log information, including:

- Inventory preload status
- Purchase success/failure information
- Error messages

## Performance Considerations

### Optimization Recommendations

1. **Connection Pool Management**: Redis client is configured with connection pooling
2. **Atomic Operations**: Uses `DECR` to ensure operation atomicity
3. **Memory Management**: Records use in-memory storage, suitable for small-scale testing

### Scalability

- Can easily scale to distributed systems
- Supports multiple Redis nodes
- Can integrate database persistence for records

## Troubleshooting

### Common Issues

1. **Redis Connection Failure**

   - Check if Redis service is running
   - Verify connection address and port

2. **Inaccurate Purchase Results**

   - Check if Redis atomic operations are working properly
   - Confirm inventory initialization is successful

3. **Performance Issues Under High Concurrency**
   - Consider using Redis Cluster
   - Optimize network connection configuration

## Contributing

Welcome to submit Issues and Pull Requests to improve this project.

## License

This project is licensed under the MIT License.
