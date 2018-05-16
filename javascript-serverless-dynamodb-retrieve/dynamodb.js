'use strict';

const AWS = require('aws-sdk'); // eslint-disable-line import/no-extraneous-dependencies

let options = {
	region: 'us-west-2'
};

// connect to local DB if running offline
if (process.env.IS_OFFLINE) {
  options = {
    region: 'localhost',
    endpoint: 'http://localhost:9000',
  };
}

const client = new AWS.DynamoDB.DocumentClient(options);

module.exports = client;
