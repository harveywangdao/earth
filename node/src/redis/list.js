const { createClient, commandOptions }  = require('redis');

// redis://alice:foobared@awesome.redis.server:6380
const client = createClient();

client.on('error', err => console.log('Redis Client Error', err));

async function testlist() {
    await client.connect();

    const blpopPromise = client.brPop(
        commandOptions({isolated: true}),
        'listkey',
        0
    );

    await client.lPush('listkey', 'v1');
    const listItem = await blpopPromise;
    console.log(listItem);
}

testlist();
