#!/bin/bash

NATS_TOKEN=s3cr3t

nats -s nats://$NATS_TOKEN@localhost:4222 str add TODOS --config stream.json
nats -s nats://$NATS_TOKEN@localhost:4222 con add TODOS --config consumer.json
