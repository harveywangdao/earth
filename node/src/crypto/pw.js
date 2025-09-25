const util = require('util');
const crypto = require('crypto');

async function decryptPassword (password) {
  const pbkdf2Promise = util.promisify(crypto.pbkdf2);
  const tenSecret = 'verysecretstring';
  const tenAESConfig = {
    "keyByteLength": 32,
    "saltByteLength": 64,
    "initializationVectorByteLength": 16,
    "iterationsByteLength": 5
  };

  const {
    keyByteLength = 32,
    saltByteLength = 64,
    initializationVectorByteLength = 16,
  } = tenAESConfig;

  const [iterations, dataHex] = password.split(':');
  const data = Buffer.from(dataHex, 'hex');
  // authTag in node.js equals 16 bytes(128 bits), see https://stackoverflow.com/questions/33976117/does-node-js-crypto-use-fixed-tag-size-with-gcm-mode
  const delta = [saltByteLength, initializationVectorByteLength, 16];
  const pointerArray = [];

  for (let byte = 0, i = 0; i < delta.length; i++) {
    const deltaValue = delta[i];
    pointerArray.push(data.subarray(byte, byte + deltaValue));
    byte += deltaValue;

    if (i === delta.length - 1) {
      pointerArray.push(data.subarray(byte));
    }
  }

  const [
    salt,
    initializationVector,
    authTag,
    encryptedData
  ] = pointerArray;

  const decryptionKey = await pbkdf2Promise(tenSecret, salt, parseInt(iterations, 16), keyByteLength, 'sha512');
  const decipher = crypto.createDecipheriv('aes-256-gcm', decryptionKey, initializationVector);
  decipher.setAuthTag(authTag);

  return Buffer.concat([decipher.update(encryptedData, 'binary'), decipher.final()]).toString();
};

let pw1
let pw2

async function testDecryptPassword() {
  pw1 = '2a9b:f25cc35407fc90dd4cbc53e3219245ba651f7f100f407b43cc4e215f97b2cb257711d7002783ebe85b8b3b308589cdda3fd8a0511277cbfd7995588153036b832ae1b17c91ea43e4cf2fe23ca8575f67bca1c732abbe0f3595bb24836d66e283da8245251c'
  pw2 = '91e9:2507a89674d657e35a1471f3c47e5f204faba776bd10a9cf763f3288c0c0a85f5993bba7b5564331271a20534290558573cac5122f6943f9b4f0cb03eb1ad4cd0fc3fd5fdf7bbf4639ea490aa6409f5aa93d3fc1e8bcfc65171021b70584ce5209ab2d6dfdcf2c9e3563ff'
    
  let str = await decryptPassword(pw1);
  console.log(str);

  str = await decryptPassword(pw2);
  console.log(str);
}

testDecryptPassword();
