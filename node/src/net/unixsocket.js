const net = require('net')
const fs = require('fs')

const pipeFile = '/tmp/unix.sock'

const server = net.createServer(connection => {
  console.log('server socket connected.')
  connection.on('close', () => console.log('server disconnected.'))
  connection.on('data', data => {
    console.log(`server receive: ${data}`)
    connection.write(data+' ack')
    console.log(`server send: ${data}`)
  })
  connection.on('error', err => console.error(err.message))
})

try {
  fs.unlinkSync(pipeFile)
} catch (error) {}

server.listen(pipeFile)

const client = net.connect(pipeFile)
client.on('connect', () => console.log('client connected.'))
client.on('data', data => console.log(`client receive: ${data}`))
client.on('end', () => console.log('client disconnected.'))
client.on('error', err => console.error(err.message))

setInterval(() => {
  const msg = 'hello'
  console.log(`client send: ${msg}`)
  client.write(msg)
}, 2000)
