const net = require('node:net');
const fs = require('node:fs');

const sockFile = '/tmp/nodejs.sock';
try {
    fs.unlinkSync(sockFile);
} catch (err) {
    console.warn(`delete ${sockFile} fail, err: ${err}`);
}

const server = net.createServer(function (socket) {
    console.log('createServer');
    handleConn(socket);
});

// TODO: listen失败怎么处理
server.listen(sockFile, function () {
    console.log(`listen ${sockFile}`);
});
server.on('connection', function (socket) {
    console.info('connection');
});
server.on('close', function () {
    console.info('close');
});
server.on('error', function (err) {
    console.error('error', err);
});
server.on('listening', function () {
    console.info('listening');
});
server.on('drop', function (data) {
    console.info('drop, data:', data);
});

function handleConn(conn) {
    conn.on('close', function (hadError) {
        console.info('handleConn close, hadError:', hadError);
    });
    conn.on('connect', function () {
        console.info('handleConn connect');
    });
    conn.on('connectionAttempt', function (ip, port, family) {
        console.info('handleConn connectionAttempt', ip, port, family);
    });
    conn.on('drain', function () {
        console.info('handleConn drain');
    });
    conn.on('end', function () {
        console.info('handleConn end callback');
    });
    conn.on('error', function (err) {
        console.info('handleConn error:', err);
    });
    conn.on('lookup', function (err, address, family, host) {
        console.info('handleConn lookup', err, address, family, host);
    });
    conn.on('ready', function () {
        console.info('handleConn ready');
    });
    conn.on('timeout', function () {
        console.info('handleConn timeout');
    });
    conn.on('data', function (data) {
        console.info('server recv data:', data.toString());
    });
    conn.write('I am server', function (err) {
        if (err) {
            console.error('server write data fail, err:', err);
        } else {
            console.info('server write data success');
        }
    });
    conn.write('I am server2', function (err) {
        if (err) {
            console.error('server write data fail, err:', err);
        } else {
            console.info('server write data success');
        }
    });
}

setTimeout(function () {
    server.close(function (err) {
        if (err) {
            console.error('server close fail, err:', err);
        } else {
            console.info('server close success');
        }
    });
}, 30000);
