console.log('hello world');
console.log("0 == 0:", 0 == 0);        // true
console.log("0 === 0:", 0 === 0);      // true

console.log("'0' == 0:", '0' == 0);    // true
console.log("'0' === 0", '0' === 0);   // false

console.log("'0' == \"0\"", '0' == "0"); // true
console.log("'0' === \"0\"", '0' === "0"); // true

let a1;
let a2 = 12;
let a3 = '12';
let a4 = null;
let a5 = undefined;
let a6 = {
    name: "xiaoming"
};
let a7 = [1,2];
let a8 = [1,"s"];
let a9 = false;
let a10 = function() {
    a9 = true;
}
console.log(a1); // undefined
console.log(typeof a1); // undefined
console.log(a2);
console.log(typeof a2); // number
console.log(a3);
console.log(typeof a3); // string
console.log(a4);
console.log(typeof a4); // object
console.log(typeof null); // object
console.log(a5);
console.log(typeof a5); // undefined
console.log(a6);
console.log(typeof a6); // object
console.log(a7);
console.log(typeof a7); // object
console.log(a8);
console.log(typeof a8); // object
console.log(a9);
console.log(typeof a9); // boolean
console.log(a10); // [Function: a10]
console.log(typeof a10); // function

//console.log(a1.constructor);
console.log(a2.constructor); // [Function: Number]
console.log(a3.constructor); // [Function: String]
//console.log(a4.constructor);
//console.log(a5.constructor);
console.log(a6.constructor); // [Function: Object]
console.log(a7.constructor); // [Function: Array]
console.log(a8.constructor); // [Function: Array]
console.log(a9.constructor); // [Function: Boolean]
console.log(a10.constructor); // [Function: Function]

let emptyStr = '';
console.log(emptyStr == ''); // true
console.log(emptyStr === ''); // true
console.log(emptyStr == 0); // true
console.log(emptyStr == false); // true
