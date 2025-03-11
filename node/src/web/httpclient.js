const request = require('request');
const fs = require('fs');
const crypto = require('crypto');

// request('https://www.baidu.com', function (error, response, body) {
//   console.error('error:', error); // Print the error if one occurred
//   console.log('statusCode:', response && response.statusCode); // Print the response status code if a response was received
//   // console.log('body:', body); // Print the HTML for the Google homepage.
// });

// request('https://www.baidu.com').pipe(fs.createWriteStream('doodle.html'))

// request.get('https://www.baidu.com').on('response', function(response) {
//     console.log(response.statusCode)
//     console.log(response.headers['content-type'])
// }).on('error', function(err) {
//     console.error(err)
// })

let hash = crypto.createHash('sha256');
let bufferLength = 0;

let fError = function(err) {
    console.log('error:', err);
}

let fData = function(chunk) {
    hash.update(chunk);
    bufferLength += chunk.length;
}

let fResponse = function(response) {
    var contentLength = response.caseless.get('content-length');
    console.log('response contentLength:', contentLength);
    console.log('response statusCode:', response.statusCode);
    console.log('response responseHeaders:', response.headers);
};

let options = {
    uri: 'http://127.0.0.1:15693'
};

// options.callback = function(err, response, body) {
//     if (err) {
//         console.log('callback error:', err);
//     } else {
//         var contentLength = response.caseless.get('content-length');
//         console.log('callback contentLength:', contentLength);

//         let sha256 = hash.digest('hex');
//         console.log('sha256:', sha256);
//         //console.log('response:', response);
//         //console.log('body:', body);
//     }
// };

let fileStream = fs.createWriteStream('download.bin')
let ro = request.get(options)
      .on('response', fResponse)
      .on('data', fData)
      .on('error', fError)
      .on('complete', function(response) {
        var contentLength = response.caseless.get('content-length');
        console.log('callback contentLength:', contentLength);
        let sha256 = hash.digest('hex');
        console.log('sha256:', sha256);
        console.log('bufferLength:', bufferLength);
      })
      .pipe(fileStream);

//request.get('http://127.0.0.1:15693', options).pipe(fileStream);
