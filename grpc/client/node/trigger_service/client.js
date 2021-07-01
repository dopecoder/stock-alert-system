const services = require("./trigger_service_grpc_pb");
const messages = require("./trigger_service_pb");

var async = require('async');
var fs = require('fs');
var parseArgs = require('minimist');
var path = require('path');
var _ = require('lodash');
var grpc = require('@grpc/grpc-js');

var client = new services.TriggerServiceClient('localhost:8000',
    grpc.credentials.createInsecure());

function runGetTrigger(callback) {
    var next = _.after(2, callback);
    function resCallback(error, res) {
        if (error) {
            callback(error);
            return;
        }
        console.log(res)
        next();
    }
    var getTriggerReq = new messages.GetTriggerReq();
    console.log(getTriggerReq);
    client.getTrigger(getTriggerReq, resCallback);
}

function main() {
    // async.series([runGetTrigger]);
    console.log(new messages.CreateTriggerReq());
    // console.log(client.getTriggerStatus(messages.createTriggerReq));
    
    var getTriggerReq = new messages.GetTriggerReq({"TriggerId": 1});
    // getTriggerReq.setTriggerid("1")
    client.getTrigger({"TriggerId": "1"}, function (err, response) {
        console.log('Err:', err);
        console.log('Res:', response);
    });
}

if (require.main === module) {
    main();
}

exports.runGetTrigger = runGetTrigger;