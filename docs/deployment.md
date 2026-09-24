# Deployment Steps

## Build Image

az acr build --registry <acr-name> --image notification-agent:v1 .

## Deploy to ACI

az container create \
  --resource-group <rg-name> \
  --name notification-agent \
  --image <acr-name>.azurecr.io/notification-agent:v1