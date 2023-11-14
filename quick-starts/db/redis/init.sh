#!/bin/sh

read -p "Container Name? " container_name

docker build -t redis-temp .

docker run -d --name $container_name redis-temp
