const { createClient } = require('redis');

async function pub() {
  const client = createClient();
  await client.connect();

  const channel1 = 'chan1nel';
  
  for (let i = 0; i < 10000; i++) {
    // 1st channel created to publish 10000 messages.
    await client.publish(channel1, `channel1_message_${i}`);
    console.log(`publishing message on ${channel1}`);
  }
}

async function sub() {
  const client = createClient();
  await client.connect();
  //const channel1Sub = client.duplicate();
  //await channel1Sub.connect();

  await client.subscribe('chan1nel', (message) => {
    console.log(`Channel1 subscriber collected message: ${message}`);
  }, true);
}

pub();
sub();
