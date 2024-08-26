# go-codebase

## Installation

### Requirements
- Go 1.22 or newer
- Docker
- Docker Compose

### Steps
1. Run docker compose to deploy postgres, redis, kafka, elasticsearch
    ```sh
    docker compose up -d
    ```
2. Copy .env.example to .env
    ```sh
    cp .env.example .env
    ```
3. Run go
    ```sh
    go run .
    ```