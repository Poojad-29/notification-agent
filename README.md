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

ACR Build could not be completed because the Azure subscription returned:

TasksOperationsNotAllowed

This prevents Azure Container Registry Tasks from executing build operations in the current subscription.