const url = require('node:url');
const URI = require("uri-js");
const querystring = require('node:querystring');

let up = url.parse('https://www.baud.com/sun/water');
console.log(up);
console.log(up.host);
console.log(up.hostname);
console.log(up.pathname);
console.log(up.protocol);
console.log(up.search);
console.log("query:", up.query);

let q1 = querystring.parse(up.query);
console.log(q1);
console.log(q1.hello);
console.log(q1.hessllo);

// let u0 = 'https://www.ugreen.com/sun/water?hello=world&kan=df';
// let u1 = URI.parse(u0);
// console.log(u1);
// let u2 = URI.serialize(u1);
// console.log(u2);
