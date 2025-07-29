#!/bin/bash

echo "running script"

chmod +x scripts/setup.sh

if ! command -v go &>/dev/null; then
command sudo-apt install go -y
fi

if ! command -v docker &> /dev/null; then
command sudo-apt install docker -y
fi

if !command -v docker-compose &> /dev/null; then
command sudo-apt install docker-compose -y
fi