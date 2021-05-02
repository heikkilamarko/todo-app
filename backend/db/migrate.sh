#!/bin/sh

migrate -path /migrations -database $POSTGRES_CONNECTIONSTRING up
