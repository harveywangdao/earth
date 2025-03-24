const { createClient } = require('redis');

let str1;
let str2 = '';
let str3 = 'sddd';

console.log('str1:', str1);
console.log('str2:', str2);
console.log('str3:', str3);

if (str1) {
  console.log('str1');
}
if (str2) {
  console.log('str2');
}
if (str3) {
  console.log('str3');
}

let option = {};
option.socket = {
    //path: '/run/redis/redis-server.sock'
    host: '127.0.0.1',
    port: 6379
};
console.log('option:', option);
const client = createClient(option);

client.on('error', err => console.log('Redis Client Error', err));

async function test () {
    await client.connect();

    await client.set('key', 'value');
    const value = await client.get('key');
    console.log(value);
}

test();
