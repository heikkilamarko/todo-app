#!/bin/sh

nats -s nats://$NATS_TOKEN@nats:4222 str add TODOS --config /stream.json
nats -s nats://$NATS_TOKEN@nats:4222 con add TODOS --config /consumer.json
