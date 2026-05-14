# Go RabbitMQ

[![Go Version](https://img.shields.io/badge/Go-1.25.5-blue)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/ggualbertosouza/go-rabbitMq/ci.yml)](https://github.com/ggualbertosouza/go-rabbitMq/actions)

Um projeto de estudo focado em mensageria e arquitetura orientada a eventos usando Go e RabbitMQ. Demonstra conceitos como exchanges (Direct, Topic, Fanout), produtores, consumidores e documentação com AsyncAPI.

## Índice
- [Recursos](#recursos)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Configuração](#configuração)
- [Execução](#execução)
- [Uso](#uso)
- [API Endpoints](#api-endpoints)
- [Documentação AsyncAPI](#documentação-asyncapi)
- [Arquitetura](#arquitetura)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Recursos
- Implementação de exchanges RabbitMQ (Direct, Topic, Fanout).
- Produtores e consumidores assíncronos.
- Servidor HTTP com endpoints REST.
- Documentação interativa via AsyncAPI.
- Infraestrutura containerizada com Docker Compose.

## Pré-requisitos
- [Go](https://golang.org/dl/) >= 1.25.5
- [Docker](https://www.docker.com/) e Docker Compose
- [Node.js](https://nodejs.org/) e npm (para geração de docs AsyncAPI)
- [Make](https://www.gnu.org/software/make/) (opcional, para comandos)

## Instalação
1. Clone o repositório:
   ```bash
   git clone https://github.com/ggualbertosouza/go-rabbitMq.git
   cd go-rabbitMq
   ```

2. Instale dependências Go:
   ```bash
   go mod download
   ```

## Configuração
1. Copie o arquivo de exemplo de ambiente:
   ```bash
   cp .env.example .env
   ```

2. Edite `.env` com suas configurações (ex.: URL RabbitMQ, porta do servidor).

## Execução
1. Inicie a infraestrutura (RabbitMQ):
   ```bash
   make infra-up
   ```

2. Execute o servidor:
   ```bash
   make run-server
   ```

3. Em outro terminal, execute os consumidores:
   ```bash
   make run-consumers
   ```

A aplicação estará disponível em `http://localhost:8080`.

Para parar a infraestrutura:
```bash
make infra-down
```

## Uso
- Acesse a documentação em `http://localhost:8080/docs`.
- Envie mensagens via API ou RabbitMQ para testar exchanges.

Exemplo de request para criar usuário:
```bash
curl -X POST http://localhost:8080/user \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

## API Endpoints
- `GET /health`: Verificação de saúde.
<<<<<<< HEAD
- `POST /rabbitmq/direct/user`: Cria um usuário.
- `PUT /rabbitmq/direct/user`: Atualiza um usuário.
=======
- `POST /user`: Cria um usuário.
- `PUT /user`: Atualiza um usuário.
>>>>>>> b4f2d76d7a14790dad584d900f33c2b43d36510f
- `GET /docs`: Documentação AsyncAPI (redireciona para `/docs`).

## Documentação AsyncAPI
O projeto usa AsyncAPI para documentar exchanges, canais, mensagens e payloads. O arquivo de especificação está em `asyncapi.yml`.

Para gerar/regenerar a documentação HTML:
```bash
make docs
```

Acesse em `http://localhost:8080/docs` quando o servidor estiver rodando.

## Arquitetura
On going

## Em andamento

### Implementar fanout
- [ ] Criar contexto novo - Notification
- [ ] Exchange fanout - Criar exchange novo
- [ ] Rotas - Criar nova rota
