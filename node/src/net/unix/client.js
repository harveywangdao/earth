const net = require('node:net');

const sockFile = '/tmp/nodejs.sock';

let cliConn = net.createConnection(sockFile, function () {
    console.info('createConnection');
});

setTimeout(function () {
    cliConn.on('data', function (data) {
        // 有粘包问题
        console.info('client recv data:', data.toString());
    
        cliConn.write('I am client', function (err) {
            if (err) {
                console.error('client send data fail, err:', err);
            } else {
                console.info('client send data success');
            }
        });
    });
}, 2000);

cliConn.on('close', function (hadError) {
    console.info('close, hadError:', hadError);
});
cliConn.on('connect', function () {
    console.info('connect');
});
cliConn.on('connectionAttempt', function (ip, port, family) {
    console.info('connectionAttempt', ip, port, family);
});
cliConn.on('drain', function () {
    console.info('drain');
});
cliConn.on('end', function () {
    console.info('end callback');
});
cliConn.on('error', function (err) {
    console.info('error:', err);
});
cliConn.on('lookup', function (err, address, family, host) {
    console.info('lookup', err, address, family, host);
});
cliConn.on('ready', function () {
    console.info('ready');
});
cliConn.on('timeout', function () {
    console.info('timeout');
});

setTimeout(function () {
    cliConn.end(function () {
        console.info('conn end');
    });
}, 4000);
