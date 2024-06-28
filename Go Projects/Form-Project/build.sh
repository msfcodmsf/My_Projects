#!/bin/bash

# Build the Docker image
docker build -t forum .

# Run the Docker container
docker run -d -p 8065:8065 --name forum-container forum
