const redis = require('redis');
const client = redis.createClient();

client.on('error', err => {
    console.log('redis client err:', err);
})

async function connect() {
    console.log("connect");
    await client.connect();
    console.log("connect done");
}
console.log("connect1");
connect();
console.log("connect2");

async function test1() {
    console.log("set");
    await client.set('key4', 'val4');
    console.log("set done");
    let val = await client.get('key4');
    console.log('val:', val);
}
console.log("test1");
test1();
console.log("end");
