'use strict';

const _ = require('lodash')
const dynamodb = require('./dynamodb')
module.exports.randomFact = (event, context, callback) => {

  getFact().then((fact) => {
    const response = {
      statusCode: 200,
      body: JSON.stringify({
        fact: fact
      }),
    }
    callback(null, response);
  })
};


const getFact = () => {
    return new Promise((fulfill, reject) => {
      let params = {
        TableName: process.env.DB_FACT_TABLE,
      }

      dynamodb.scan(params, (error, data) => {

        if (error) {
          console.log('DynamoDB Error!:', error)
          reject(error)
        }

        let index = _.random(0, data.Items.length - 1)
        let fact = data.Items[index]
        fulfill(fact)
      });


    })
  }
