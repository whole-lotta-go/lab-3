#!/bin/bash

curl -X POST http://localhost:17000 -d "green"
curl -X POST http://localhost:17000 -d "bgrect 0.2 0.2 0.8 0.8"
curl -X POST http://localhost:17000 -d "update"
