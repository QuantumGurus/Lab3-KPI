#!/bin/bash

# Set the green screen
echo "Setting the green screen..."
curl -X POST http://localhost:17000 -d "reset"
curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "update"
echo "Green screen set complete."
