# Go Prometheus Grafana Monitoring Stack

A comprehensive Go web application with integrated **Prometheus metrics collection** and **Grafana dashboards** for real-time monitoring and observability.

## 🚀 Features

- **RESTful API** built with [Gin](https://gin-gonic.com/) framework
- **Custom Prometheus Middleware** for HTTP request/error tracking
- **Go Runtime Metrics** monitoring (memory, GC, goroutines)
- **Grafana Dashboard** with 9 visualization panels
- **Load Testing Simulator** for generating realistic traffic
- **Docker Compose** setup for easy deployment
- **Error Simulation** endpoints for testing monitoring alerts

## 📊 Monitoring Capabilities

### HTTP Metrics
- **Request Counter**: Total HTTP requests by path, method, and status
- **Error Counter**: HTTP errors (4xx/5xx) tracking
- **Request Rate**: Real-time request frequency
- **Endpoint Distribution**: Pie chart of requests per endpoint
- **Error Distribution**: Bar gauge of errors by endpoint

### Go Runtime Metrics
- **Memory Usage**: Heap allocation, in-use memory, system memory
- **Garbage Collection**: Duration, frequency, and CPU impact
- **Goroutines**: Active goroutine monitoring
- **Heap Objects**: Object count tracking
- **Memory Allocation Rate**: Allocation and creation rates

## 🛠️ Tech Stack

- **Backend**: Go 1.24.4, Gin Web Framework
- **Monitoring**: Prometheus, Grafana
- **Containerization**: Docker, Docker Compose
- **Metrics**: Prometheus Client Library

## 📁 Project Structure

```
prometheus-grafana/
├── main.go                     # Main application server
├── middleware/
│   └── prometheus.go          # Prometheus metrics middleware
├── tests/
│   └── main.go               # Load testing simulator
├── grafana/
│   └── provisioning/
│       ├── dashboards/       # Grafana dashboard configurations
│       └── datasources/      # Prometheus datasource config
├── prometheus/
│   └── prometheus.yml        # Prometheus configuration
├── docker-compose.yml        # Docker orchestration
├── Dockerfile               # Application container
├── Makefile                # Build and run commands
└── go.mod                  # Go dependencies
```

## 🚦 Quick Start

### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)
- [Go 1.24+](https://golang.org/dl/) (for local development)
- [Make](https://www.gnu.org/software/make/) (optional, for convenience commands)

### 1. Clone and Start

```bash
git clone <repository-url>
cd prometheus-grafana

# Start the entire monitoring stack
make up-server
# OR
docker compose up -d
```

### 2. Access Services

| Service | URL | Credentials |
|---------|-----|-------------|
| **Go Application** | http://localhost:8080 | - |
| **Prometheus** | http://localhost:9090 | - |
| **Grafana** | http://localhost:3000 | admin/admin |

### 3. View Dashboard

1. Open Grafana at http://localhost:3000
2. Login with `admin/admin`
3. Navigate to **"Go Runtime Metrics - Load Testing Dashboard"**
4. Enjoy real-time monitoring! 📈

## 🔧 Usage

### API Endpoints

| Endpoint | Description |
|----------|-------------|
| `GET /` | Welcome message |
| `GET /users` | Users endpoint with error simulation |
| `GET /comments` | Comments endpoint with error simulation |
| `GET /posts` | Posts endpoint with error simulation |
| `GET /metrics` | Prometheus metrics endpoint |

### Error Simulation

Test different HTTP status codes using query parameters:

```bash
# Success responses (200)
curl http://localhost:8080/users
curl http://localhost:8080/comments
curl http://localhost:8080/posts

# Trigger 404 Not Found
curl http://localhost:8080/users?test=trigger-not-found

# Trigger 403 Forbidden
curl http://localhost:8080/users?test=trigger-forbidden

# Trigger 500 Internal Server Error
curl http://localhost:8080/users?test=trigger-server-error
```

### Load Testing

Generate realistic traffic to see monitoring in action:

```bash
# Start load testing simulator
make simulate
# OR
go run tests/main.go
```

The simulator creates:
- **30 concurrent goroutines** sending requests to different endpoints
- **Different error rates** (success, 404, 403, 500)
- **Varied request intervals** (1-4 seconds) for realistic patterns

Press `Ctrl+C` to stop the load testing.

## 📊 Dashboard Panels

The Grafana dashboard includes 9 comprehensive panels:

1. **Total Requests Error Endpoint** - Bar gauge showing error distribution
2. **Memory Usage** - Time series of heap and system memory
3. **Total Requests Endpoint** - Pie chart of request distribution
4. **Memory Allocation Rate** - Allocation and object creation rates
5. **Total All Requests** - Bar chart with gradient styling
6. **Heap Objects Count** - Object count monitoring
7. **Go Goroutines** - Active goroutine tracking
8. **Garbage Collection Duration** - GC performance metrics
9. **Request & GC Frequency** - Combined request and GC rates

### Dashboard Features
- **Auto-refresh**: 5-second intervals
- **Time range**: Last 5 minutes (configurable)
- **Dark theme** with modern visualizations
- **Interactive legends** and tooltips

## 🔨 Development Commands

```bash
# Run application locally
make run
# OR
go run main.go

# Start monitoring stack
make up-server

# Stop monitoring stack
make down-server

# Run load testing
make simulate
```

## 🐋 Docker Configuration

The `docker-compose.yml` orchestrates three services:

- **go-server**: The main application (port 8080)
- **prometheus**: Metrics collection (port 9090)
- **grafana**: Visualization dashboard (port 3000)

### Custom Configuration

- **Prometheus**: Scrapes metrics every 15 seconds
- **Grafana**: Auto-provisions dashboards and datasources
- **Network**: Isolated `monitoring` network for service communication

## 📈 Metrics Exposed

### Custom Application Metrics

```prometheus
# Total HTTP requests
http_requests_total{path="/users", method="GET", status="OK"}

# HTTP errors (4xx/5xx)
http_errors_total{path="/users", method="GET", status="Not Found"}
```

### Go Runtime Metrics (Auto-collected)

```prometheus
# Goroutines
go_goroutines

# Memory stats
go_memstats_heap_alloc_bytes
go_memstats_heap_inuse_bytes
go_memstats_sys_bytes

# Garbage collection
go_gc_duration_seconds
go_memstats_gc_cpu_fraction
```

## 🚀 Production Considerations

### Security
- Change default Grafana credentials
- Configure authentication for Prometheus
- Use HTTPS in production
- Implement rate limiting

### Scaling
- Configure Prometheus retention policies
- Set up alerting rules
- Consider Grafana data source clustering
- Implement log aggregation

### Monitoring
- Set up alerting for high error rates
- Monitor resource usage trends
- Configure notification channels
- Create SLA dashboards

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/amazing-feature`
3. Commit changes: `git commit -m 'Add amazing feature'`
4. Push to branch: `git push origin feature/amazing-feature`
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Prometheus](https://prometheus.io/) for metrics collection
- [Grafana](https://grafana.com/) for visualization
- [Gin](https://gin-gonic.com/) for the web framework
- [Go](https://golang.org/) community for excellent tooling

---

**Happy Monitoring!** 🎯📊

For questions or support, please open an issue in the repository. 