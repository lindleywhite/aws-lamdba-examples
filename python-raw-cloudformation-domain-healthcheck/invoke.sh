#!/bin/bash

aws lambda invoke --function-name HealthCheckFunction out.txt

cat out.txt