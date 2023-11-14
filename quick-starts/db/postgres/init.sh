#!/bin/bash

read -p "Container Name? " container_name

docker build -t postgres-temp .

docker run -d -e POSTGRES_PASSWORD=password --name $container_name postgres-temp
