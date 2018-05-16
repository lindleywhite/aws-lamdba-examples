
#!/bin/bash

# Set this to a bucket owned by you
export STACK_NAME=aws-lambda-golang-example
export BUCKET_NAME=python-examples


rm -rf build
mkdir build

GOOS=linux GOARCH=amd64 go build -ldflags="-d -s -w" -o build/cert_check cert_check/cert_check.go
chmod +x build/cert_check
zip -j build/cert_check.zip build/cert_check

aws cloudformation package \
    --template-file stack.yml \
    --s3-bucket $BUCKET_NAME \
    --output-template-file packaged-template.yml

aws cloudformation deploy \
    --template-file packaged-template.yml \
    --stack-name $STACK_NAME \
    --capabilities CAPABILITY_IAM \
    --parameter-overrides BucketName=$BUCKET_NAME DomainsToCheck="lindleywhite.com,www.slcdevopsdays.org"

