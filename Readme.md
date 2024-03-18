# GoLang Microservices with Echo, gRPC, PostgreSQL, Entgo, Redis, and Docker Compose

This repository contains a set of microservices built with GoLang using Echo for API Gateway, gRPC for communication between services, PostgreSQL for the database, Entgo for ORM, Redis for caching, and Docker Compose for easy deployment.

## Services

1. **API Gateway**: Built with Echo Router.
2. **User Service**: Handles user-related functionalities using gRPC.
3. **Cart Service**: Manages user shopping carts using gRPC.
4. **Order Service**: Deals with order processing using gRPC.

## Prerequisites

- GoLang installed on your machine
- Docker and Docker Compose installed
- PostgreSQL and Redis installed or accessible via Docker

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/Sri2103/services.git

```bash
cd services
```

Start the services using Docker Compose:

```bash
docker-compose up --build
```

Once the services are up and running, you can access the API Gateway and interact with the microservices.

## Configuration

Make sure to update the configuration files for each service according to your environment. You may need to modify the database credentials, Redis connection details, and any other configurations specific to your setup.

## Usage

Once the services are running, you can start making requests to the API Gateway or directly to the individual services using gRPC.
