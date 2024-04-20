# KAR
__Kafka At Rest__
> Like Hasura, but for Kafka.


## Overview

KAR is a tool that allows you to expose your Kafka topics as a GraphQL/REST API.

You can connect the KAR server to a Kafka topic and instantly get a GraphQL/REST API for that topic.

## GraphQL API

```bash
$ kar --topic my-topic --port 8080
```

This will start a KAR server that listens on port 8080 and exposes the `my-topic` Kafka topic as a GraphQL/REST API.

Assuming the Kafka topic has the following messages:

```json
{"id": 1, "name": "Alice"}
{"id": 2, "name": "Bob"}
```

You can now query the API:

**Query**

You can use GraphQL query to query / filter messages from the Kafka topic.

```graphql
query {
  my_topic {
    id
    name
  }
}
```

**Response**

```json
{
  "data": {
    "my_topic": [
      {
        "id": 1,
        "name": "Alice"
      },
      {
        "id": 2,
        "name": "Bob"
      }
    ]
  }
}
```

You can also subscribe to the topic and get real-time updates:

**Subscription**

```graphql
subscription {
  my_topic {
    id
    name
  }
}
```

**Response**

```json
{
  "data": {
    "my_topic": {
      "id": 3,
      "name": "Charlie"
    }
  }
}
```

You can also use mutations to produce messages to the topic:

**Mutation**

```graphql
mutation {
  my_topic {
    produce(message: {id: 4, name: "David"}) {
      id
      name
    }
  }
}
```

**Response**

```json
{
  "data": {
    "my_topic": {
      "id": 4,
      "name": "David"
    }
  }
}
```

---

## REST API

```bash
$ kar --topic my-topic --port 8080 --rest
```

This will start a KAR server that listens on port 8080 and exposes the `my-topic` Kafka topic as a REST API.

Assuming the Kafka topic has the following messages:

```json
{"id": 1, "name": "Alice"}
{"id": 2, "name": "Bob"}
```

You can now query the API:

**Query**

You can use REST API to query / filter messages from the Kafka topic.

```bash
$ curl http://localhost:8080/my-topic
```

**Response**

```json
[
  {
    "id": 1,
    "name": "Alice"
  },
  {
    "id": 2,
    "name": "Bob"
  }
]
```

You can also use POST to produce messages to the topic:


**Mutation**

```bash
$ curl -X POST http://localhost:8080/my-topic -d '{"id": 3, "name": "Charlie"}'
```

**Response**

```json
{
  "id": 3,
  "name": "Charlie"
}
```

And you can use websockets to subscribe to the topic and get real-time updates:

**Subscription**

```bash
$ curl -i -N -H "Connection: Upgrade" -H "Upgrade: websocket" -H "Host: localhost:8080" -H "Origin: http://localhost:8080" http://localhost:8080/my-topic
```

**Response**

```json
{"id": 4, "name": "David"}
```

---

## Installation

```bash
$ go get github.com/adhaamehab/kar
```


---
