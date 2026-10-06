# Notification Agent

## Features

- Azure PostgreSQL Integration
- Token Authentication
- Rate Limiting
- Health Check
- Logging

## Run

go run ./cmd/server

## Build

go build ./...

## Test

go test ./...

## Database

Server: poojanotificationdb29
Database: notificationdb

## Known Deployment Limitation

Azure Container Registry was created successfully.

Registry SKU was upgraded from Basic to Premium.

The following operations failed:

- az acr build
- az acr task create

with:

TasksOperationsNotAllowed

This indicates Azure Container Registry Tasks are restricted by the current subscription/tenant policy.

Consequently, container images could not be built and Azure Container Instance deployment could not be completed.