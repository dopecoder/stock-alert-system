
var PROTO_PATH = __dirname + '/../../trigger_service/trigger_service.proto';

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
const promisify = require('grpc-promisify');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
var trigger_service = grpc.loadPackageDefinition(packageDefinition).trigger_service;

async function main() {
    var client = new trigger_service.TriggerService("localhost:80",
        grpc.credentials.createInsecure());
    promisify(client);

    // console.log(client);
    createTriggerRes = await client.createTrigger({
        id: '3',
        tAttrib: 'LTP',
        operator: 'GTE',
        tPrice: 302,
        tNearPrice: 292,
        scrip: 'SBIN',
        kiteToken: '',
        exchangeToken: '3045',
        Exchange: 'NSE'
      });
    console.log(createTriggerRes)

    getTriggerRes = await client.getTrigger({ triggerId: "3" });
    console.log(getTriggerRes)

    getTriggerStatusRes = await client.getTriggerStatus({ triggerId: "3" });
    console.log(getTriggerStatusRes)
}

main().then(()=>console.log("Completed!"));