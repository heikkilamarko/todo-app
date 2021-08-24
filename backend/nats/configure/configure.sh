#!/bin/sh

NATS_URL=nats://${NATS_TOKEN}@${NATS_HOST}

nats -s $NATS_URL str add --config /streams/todo.json
nats -s $NATS_URL con add todo --config /consumers/todo.json
