'use strict';
const { createClient, commandOptions }  = require('redis');

var RECONNECT_TIMEOUT = 1000;

function connectPromise(closeCallback) {
  return new Promise(function(resolve, reject) {
    function startConnect() {
      let client = createClient();
      client.on('error', err => console.error('Redis Client Error:', err));
      client.on('end', () => {
        console.debug('[redis] conn close');
        closeCallback();
      });
      client.connect().then(() => {
        console.debug('[redis] connected');
        resolve(client);
      }).catch((err) => {
        console.error('[redis] connect fail, err:', err);
        setTimeout(startConnect, RECONNECT_TIMEOUT);
      });
    }
    startConnect();
  });
}

async function consumeMsg(client, queue, messageCallback) {
  try {
    while (true) {
      let blpopPromise = client.BRPOP(
        commandOptions({isolated: true}),
        queue,
        0
      );
      console.debug('redis consume wait');
      let res = await blpopPromise;
      console.debug('redis consume success, msg:', res.element);
      messageCallback(res.element);
    }
  } catch (err) {
    console.error('redis consume fail, err:', err);
    setTimeout(() => {
      console.warn('redis consume retry');
      consumeMsg(client, queue, messageCallback);
    }, 2000);
  }
}

function consumePromise(client, queue, messageCallback, options) {
  return new Promise(function(resolve, reject) {
    consumeMsg(client, queue, messageCallback);
    resolve();
  });
}

function closePromise(client) {
  return new Promise(function(resolve, reject) {
    client.quit().then(resolve, reject);
  });
}

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}

let i = 0;
async function produce(client, queue) {
  try {
    while (true) {
        let msg = 'value'+i;
        console.log('produce msg:', msg);
        await client.LPUSH(queue, msg);
        console.log('produce success msg:', msg);
        i++;
  
        await sleep(2000);
    }
  } catch (err) {
    console.error('produce fail:', err);
    setTimeout(() => {
      console.info('produce retry');
      produce(client, queue);
    }, 2000);
  }
}

let queue = 'testrmq';
async function consumeTask() {
  let client = await connectPromise(() => {
    console.log('closeCallback')
  });
  await consumePromise(client, queue, (msg) => {
    console.log('recv msg:', msg);
  });

  setTimeout(async () => {
    await closePromise(client);
  }, 9000);
}

async function produceTask() {
  let client = await connectPromise(() => {
    console.log('closeCallback')
  });

  produce(client, queue);
  setTimeout(async () => {
    await closePromise(client);
  }, 15000);
}

consumeTask();
produceTask();
