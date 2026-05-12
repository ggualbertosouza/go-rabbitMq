# Go RabbitMQ

Study project focused on messaging and event-driven architecture using Go and RabbitMQ.

The project demonstrates:
- Direct Exchange
- Topic Exchange
- Fanout Exchange
- Producers and Consumers
- AsyncAPI documentation

---

# Getting Started

## Clone repository

```bash
git clone https://github.com/ggualbertosouza/go-rabbitMq
```

---

## Create environment file

```bash
cp .env.example .env
```

---

## Start required services

```bash
docker compose up -d
```

This starts required infrastructure services such as:
- RabbitMQ

---

## Run application

```bash
make run
```

Application will start at:

```text
http://localhost:8080
```

---

# AsyncAPI Documentation

This project uses AsyncAPI to document:
- Exchanges
- Channels
- Producers
- Consumers
- Message payloads

The AsyncAPI specification file is located at:

```text
asyncapi.yml
```

Generated static documentation is available in:

```text
/docs
```

When the application is running, documentation can be accessed at:

```text
http://localhost:8080/docs
```

---

# Updating Documentation

Whenever `asyncapi.yml` is updated, regenerate the documentation:

```bash
make docs
```

This command:
- Regenerates the AsyncAPI HTML documentation
- Recreates the `/docs` directory
- Overwrites old generated files

Documentation generation uses:
- Node.js
- npx
- AsyncAPI CLI
- AsyncAPI HTML Template