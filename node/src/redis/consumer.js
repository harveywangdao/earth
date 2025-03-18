const { createClient, commandOptions }  = require('redis');

const client = createClient();

client.on('error', err => console.error('Redis Client Error:', err));

async function consumeMsg() {
    try {
        while (true) {
            let blpopPromise = client.BRPOP(
                commandOptions({isolated: true}),
                'listkey',
                0
            );
            console.log('consume wait');
            let listItem = await blpopPromise;
            console.log(listItem);
        }
    } catch (err) {
        console.error('consume fail, err:', err);
        setTimeout(() => {
            console.warn('retry consume');
            consumeMsg();
        }, 2000);
    }
}

async function consume() {
    await client.connect();
    consumeMsg();
}

consume();
