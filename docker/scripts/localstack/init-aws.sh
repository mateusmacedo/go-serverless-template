#!/bin/bash
aws configure set aws_access_key_id key --profile=default
aws configure set aws_secret_access_key secret --profile=default
aws configure set region us-east-1 --profile=default

echo "########### AWS Configure ###########"
aws configure list --profile=default

# echo "########### Creating queues ###########"
# awslocal sqs create-queue --queue-name template-queue --profile=default