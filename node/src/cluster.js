const cluster = require('node:cluster');
const http = require('node:http');
const numCPUs = require('node:os').cpus().length;
const process = require('node:process');

if (cluster.isPrimary) {
  console.log(`Primary ${process.pid} is running`);

  // Fork workers.
  for (let i = 0; i < numCPUs; i++) {
    let worker = cluster.fork();
    worker.send(i);
  }

  console.log(cluster.workers);
  console.log(Object.keys(cluster.workers));

  cluster.on('exit', (worker, code, signal) => {
    console.log(`worker ${worker.process.pid} died`);
  });
} else {
  process.on("message", function (msg) {
    console.log(msg);

    let port = 8000 + msg;
    // Workers can share any TCP connection
    // In this case it is an HTTP server
    http.createServer((req, res) => {
      res.writeHead(200);
      res.end('hello world\n');
    }).listen(port);
    console.log(`Worker ${process.pid} started, ${port}`);
  })
}