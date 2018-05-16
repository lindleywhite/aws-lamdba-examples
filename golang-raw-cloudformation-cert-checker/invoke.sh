#!/bin/bash

aws lambda invoke --function-name cert_check out.txt

cat out.txt