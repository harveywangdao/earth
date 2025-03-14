const unix = require('unix-dgram');
const {Buffer} = require('node:buffer');
const fs = require('node:fs');

const srvSockFile = '/tmp/srv.sock';
const cliSockFile = '/tmp/cli.sock';
try {
    fs.unlinkSync(cliSockFile);
} catch (err) {
    console.warn(`delete ${cliSockFile} fail, err: ${err}`);
}

let client = unix.createSocket('unix_dgram');
client.on('error', function (err) {
    console.error(err);
});
client.on('connect', function () {
    console.info('client connect');

    client.on('congestion', function() {
        console.log('congestion');
        /* The server is not accepting data */
    });
    client.on('writable', function() {
        console.log('writable');
        /* The server can accept data */
    });

    for (let i = 0; i < 10; i++) {
        let msg = Buffer.from('I am client'+i);
        client.send(msg);
    }
    client.close();
});

//client.send(msg, 0, msg.length, sockFile);
client.connect(srvSockFile);
client.bind(cliSockFile);
