#!/bin/bash
name="beekman9527/workload"
docker build -t $name .

hostPort=28080
memlimit="4096m"
echo "docker run --memory $memlimit -d -p $hostPort:8080 $name" 
