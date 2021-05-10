#!/bin/sh

nats -s nats://$NATS_TOKEN@nats:4222 str add --config /streams/todo.json

nats -s nats://$NATS_TOKEN@nats:4222 con add todo --config /consumers/todo_created.json
nats -s nats://$NATS_TOKEN@nats:4222 con add todo --config /consumers/todo_completed.json
