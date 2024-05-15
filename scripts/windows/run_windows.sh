#!/bin/bash

service="notification-service"
port=8084

echo "Running $service"

docker.exe run -d -e VAULT_TOKEN=${VAULT_TOKEN} -e VAULT_ADDRESS=${VAULT_ADDRESS} --name pickside-$service-dev -p $port:$port $service:latest
