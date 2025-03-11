const fs = require('fs');
const http = require('http');

let server = http.createServer();
server.on("request", function(request, response) {
    //let url = request.url;
    let s1 = fs.createReadStream('big.bin');
    response.writeHead(200, {
        'Content-Type': 'application/octet-stream',
        'Content-Disposition': 'attachment; filename=big.bin'
    });
    s1.pipe(response);
})

let port = 15693;
let hostname = '0.0.0.0';
server.listen(port, hostname, function() {
    console.log(`Server running at http://${hostname}:${port}/`);
})
