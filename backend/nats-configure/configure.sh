#!/bin/sh
set -e

nats stream add --config "/streams/todo.json"
nats consumer add todo --config "/consumers/todo.json"
