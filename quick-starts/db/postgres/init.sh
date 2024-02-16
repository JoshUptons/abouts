#!/bin/bash

read -p "Container Name? " container_name

docker build -t postgres-temp .

docker run --rm -d -p 5432:5432 --name $container_name postgres-temp
