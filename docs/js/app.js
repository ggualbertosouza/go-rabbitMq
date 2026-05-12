
    const schema = {
  "asyncapi": "3.0.0",
  "info": {
    "title": "RabbitMQ Topology Example",
    "version": "1.0.0",
    "description": "Exemplo de topologia RabbitMQ usando:\n- **Direct** Exchange\n- **Topic** Exchange\n- **Fanout** Exchange\n"
  },
  "servers": {
    "rabbitmq": {
      "host": "localhost:5672",
      "protocol": "amqp"
    }
  },
  "channels": {
    "userCreated": {
      "address": "user.created",
      "bindings": {
        "amqp": {
          "is": "routingKey",
          "exchange": {
            "name": "direct.exchange",
            "type": "direct",
            "durable": true,
            "autoDelete": false,
            "vhost": "/"
          }
        }
      },
      "messages": {
        "userCreatedMessage": {
          "name": "UserCreated",
          "title": "User Created Event",
          "payload": {
            "type": "object",
            "properties": {
              "userId": {
                "type": "string",
                "x-parser-schema-id": "<anonymous-schema-2>"
              },
              "email": {
                "type": "string",
                "x-parser-schema-id": "<anonymous-schema-3>"
              }
            },
            "x-parser-schema-id": "<anonymous-schema-1>"
          },
          "x-parser-unique-object-id": "userCreatedMessage"
        }
      },
      "x-parser-unique-object-id": "userCreated"
    }
  },
  "operations": {
    "publishUserCreated": {
      "action": "send",
      "channel": "$ref:$.channels.userCreated",
      "x-parser-unique-object-id": "publishUserCreated"
    },
    "consumeUserQueue1": {
      "action": "receive",
      "channel": "$ref:$.channels.userCreated",
      "bindings": {
        "amqp": {
          "ack": true
        }
      },
      "x-parser-unique-object-id": "consumeUserQueue1"
    },
    "consumeUserQueue2": {
      "action": "receive",
      "channel": "$ref:$.channels.userCreated",
      "bindings": {
        "amqp": {
          "ack": true
        }
      },
      "x-parser-unique-object-id": "consumeUserQueue2"
    }
  },
  "components": {
    "messages": {
      "UserCreated": "$ref:$.channels.userCreated.messages.userCreatedMessage"
    }
  },
  "x-parser-spec-parsed": true,
  "x-parser-api-version": 3,
  "x-parser-spec-stringified": true
};
    const config = {"show":{"sidebar":true},"sidebar":{"showOperations":"byDefault"}};
    const appRoot = document.getElementById('root');
    AsyncApiStandalone.render(
        { schema, config, }, appRoot
    );
  