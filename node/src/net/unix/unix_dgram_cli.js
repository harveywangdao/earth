const unix = require('unix-dgram');
const {Buffer} = require('node:buffer');
const fs = require('node:fs');

const srvSockFile = '/tmp/srv.sock';
const cliSockFile = '/tmp/cli.sock';

let client = unix.createSocket('unix_dgram');
let connected = false;
client.on('error', function (err) {
    console.error('unix socket client err:', err);
    connected = false;
    setTimeout(function () {
        client.connect(srvSockFile);
    }, 2000);
});
client.on('connect', function () {
    connected = true;
    console.info('client connect');
});
//client.send(msg, 0, msg.length, sockFile);
//client.close();
client.connect(srvSockFile);

setInterval(function () {
    if (!connected) {
        return;
    }
    console.log("client send data");
    let msg = Buffer.from('I am client');
    client.send(msg);
}, 2000);

try {
    fs.unlinkSync(cliSockFile);
} catch (err) {
    console.warn(`delete ${cliSockFile} fail, err: ${err}`);
}

let server = unix.createSocket('unix_dgram', function(buf, rinfo) {
  console.log('client recv:', buf.toString(), rinfo);
});
server.on('listening', function () {
    console.info('listening');
});
server.on('error', function (err) {
    console.error('unix socket server err:', err);
});
server.bind(cliSockFile);
