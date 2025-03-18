const { createClient }  = require('redis');

// redis://alice:foobared@awesome.redis.server:6380
const client = createClient();

client.on('error', err => console.error('Redis Client Error:', err));

function sleep(ms) {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function produce() {
    await client.connect();

    let i = 0;
    while (true) {
        let msg = 'value'+i;
        console.log('produce msg:', msg);
        await client.LPUSH('listkey', msg);
        console.log('produce success msg:', msg);
        i++;

        await sleep(2000);
    }
}

produce();
