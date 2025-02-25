const express = require('express')
const app = express()
const port = 3000

app.get('/*', (req, res) => {
  // res.send('Hello World!')
  console.log(req.baseUrl)
  console.log(req.originalUrl)
  console.log(req.params)
  console.log(req.path)
  console.log(req.query)
  console.log(req.hostname)
  req.statusCode
  res.json({name: "xiaoming", age: 12})
})

var srv = app.listen(port, () => {
  console.log(srv.address().address, srv.address().port)
  console.log(`Example app listening on port ${port}`)
})
