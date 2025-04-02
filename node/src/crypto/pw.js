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

async function testDecryptPassword() {
  let str = await decryptPassword('db28:361be2123f2747db9439145cc7ac50a5f57b5e565ff79d15332bc9e35727bcdfd745f14e5edeb8d8f73a16e9b0a3fb36413c6d1ceda678b8759b28ce4c2b62d985c9eceba8cddd012d1b4fc9257197ec55879b399f88d6a1e705d52405e3eb7a2f3643897c38');
  console.log(str);
}

testDecryptPassword();
