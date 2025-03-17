const unix = require('unix-dgram');
const fs = require('node:fs');

const srvSockFile = '/tmp/srv.sock';
const cliSockFile = '/tmp/cli.sock';
try {
    fs.unlinkSync(srvSockFile);
} catch (err) {
    console.warn(`delete ${srvSockFile} fail, err: ${err}`);
}

let server = unix.createSocket('unix_dgram', function(buf, rinfo) {
  console.log('server recv:', buf.toString(), rinfo);
});
server.on('listening', function () {
    console.info('listening');
});
server.on('error', function (err) {
    console.error('unix socket server err:', err);
});
server.bind(srvSockFile);
//server.close();

let client = unix.createSocket('unix_dgram');
client.on('error', function (err) {
    console.error('unix socket client err:', err);
    setTimeout(function () {
        client.connect(cliSockFile);
    }, 5000);
});
client.on('connect', function () {
    console.info('client connect');
    for (let i = 0; i < 2; i++) {
        let msg = Buffer.from('I am server'+i);
        client.send(msg);
    }
});
client.connect(cliSockFile);
