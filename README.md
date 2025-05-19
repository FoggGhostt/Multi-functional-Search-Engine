# Multi-functional Search Engine

Multi-functional Search Engine is a lightweight, high-performance search engine written in Go (Golang) with a Vue.js frontend. It supports indexing and searching across multiple text document formats by implementing text preprocessing (tokenization, normalization, stop-word removal, stemming), building an inverted index stored in MongoDB, and ranking results using TF-IDF and cosine similarity. The backend leverages Go’s goroutines for concurrent processing, and the application can be deployed via Docker and Nginx.

## Features

* **Text Preprocessing**: Tokenization, case normalization, removal of Russian and English stop words, and Porter stemming for both languages.
* **Inverted Index**: Builds and stores an inverted index mapping tokens to document IDs for fast lookup.
* **Ranking**: Calculates TF-IDF weights and ranks documents by cosine similarity.
* **Concurrency**: Uses goroutines for parallel file parsing, indexing, and query processing to maximize performance.
* **Frontend UI**: Vue.js-based interface for uploading documents, executing search queries, and displaying ranked results.
* **Persistence**: MongoDB stores both forward (document → token counts) and inverted (token → document list) indices.
* **Testing**: Unit tests for core packages (Parser, Indexer, Search).
* **Deployment**: Dockerized with Docker Compose and served via Nginx.

## Tech Stack

* **Backend**: Go (Golang)
* **Frontend**: Vue.js
* **Database**: MongoDB
* **Containerization**: Docker, Docker Compose
* **Web Server**: Nginx

## Getting Started

### Prerequisites

* Go 1.18+ installed
* Node.js and npm
* MongoDB instance
* Docker & Docker Compose (optional, for containerized deployment)

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/FoggGhostt/Multi-functional-Search-Engine.git
   cd Multi-functional-Search-Engine
   ```

2. **Configure**:

   * Edit `config.yaml` to set your MongoDB URI and database name.

3. **Build Backend**:

   ```bash
   go build -o search-engine main.go
   ```

4. **Install Frontend Dependencies**:

   ```bash
   cd frontend
   npm install
   ```

## Usage

1. **Start MongoDB** (if not already running).
2. **Run Backend**:

   ```bash
   ./search-engine
   ```
3. **Run Frontend**:

   ```bash
   cd frontend
   npm run serve
   ```
4. **Open** [http://localhost:8080](http://localhost:8080) in your browser to access the UI.

## Running Tests

To run unit tests for core packages:

```bash
cd pkg/parser && go test
cd pkg/indexer && go test
cd pkg/search && go test
```

## Docker Deployment

1. **Build and run containers**:

   ```bash
   docker-compose up --build
   ```
2. The frontend will be served by Nginx, which also proxies API requests to the Go backend.

## Project Structure

```plaintext
.
├── Dockerfile
├── docker-compose.yml
├── config.yaml                  # MongoDB and other settings
├── main.go                      # Application entrypoint (Go)
├── pkg                          # Go packages
│   ├── API                      # HTTP handlers
│   ├── parser                   # File reading & tokenization
│   ├── indexer                  # Inverted index builder
│   ├── search                   # TF-IDF ranking & query logic
│   ├── models                   # Data models
│   ├── mongodb                  # MongoDB connection & operations
│   └── config                   # Configuration loader
└── frontend                     # Vue.js application
    ├── public
    └── src
        ├── components           # UI components (SearchBar, FileUpload, Results)
        └── App.vue, main.js     # App entrypoint
```

## Contributing

Contributions and pull requests are welcome! Feel free to open issues for bugs or feature requests.

## References

* Porter stemming algorithm for Russian and English
* TF-IDF ranking and cosine similarity
* Go concurrency patterns (goroutines, sync.Map)
* Vue.js documentation

---

*Report and implementation details adapted from project documentation.*
