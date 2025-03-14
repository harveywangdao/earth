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

  //let msg = Buffer.from('I am server');
  //server.send(msg);
});
server.on('listening', function () {
    console.info('listening');
});
server.on('error', function (err) {
    console.error(err);
});
server.bind(srvSockFile);

setTimeout(function () {
    server.close();
}, 30000);
