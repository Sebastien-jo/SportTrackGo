#!/bin/bash

### INIT DDB
awslocal dynamodb create-table \
  --table-name prompt_paylaod \
  --key-schema AttributeName=uuid,KeyType=HASH \
  --attribute-definitions AttributeName=uuid,AttributeType=S \
  --billing-mode PAY_PER_REQUEST

awslocal dynamodb put-item \
  --table-name archiveRequest \
  --item '{"uuid": {"S": "e63a633a-42f6-4bb7-ae31-3c6716e667b0"},"user_uuid": {"S": "5dec90a7-ae81-44f7-b212-204efff641f9"},"sessions_per_week": {"N": "4"},"place": {"S": "gym"},"equipment": {"L": [{"S": "kettel-bell"},{"S": "weight"}]},"video": {"BOOL": true},"additional information": {"S": "add more information here"}}'
