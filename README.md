# SportTrackGo

## Description

## Init project
````shell
git clone git@github.com:Sebastien-jo/SportTrackGo.git
cd SportTrackGo
cp .envrc.dist .env
docker compose up --build -d
docker compose exec -ti localstack bash
awslocal dynamodb get-item --table-name prompt_payload --key '{"uuid":{"S":"e63a633a-42f6-4bb7-ae31-3c6716e667b0"}}'
```
