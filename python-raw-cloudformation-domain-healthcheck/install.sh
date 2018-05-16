
#!/bin/bash

# Set this to a bucket owned by you
export STACK_NAME=aws-lambda-python-example
export BUCKET_NAME=python-examples


rm -rf build
mkdir build


zip -jr build/health_check.zip health_check/*

aws cloudformation package \
    --template-file stack.yml \
    --s3-bucket $BUCKET_NAME \
    --output-template-file packaged-template.yml

aws cloudformation deploy \
    --template-file packaged-template.yml \
    --stack-name $STACK_NAME \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides BucketName=$BUCKET_NAME DomainsToCheck="http://lindleywhite.com, https://slcdevopsdays.org"

