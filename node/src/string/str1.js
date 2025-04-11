let str1 = 'Hello world';
console.log(str1.substring(0, 5));
console.log(str1.slice(0, 5));
console.log(str1.match('^Hello'));
console.log(str1.startsWith('Hello'));
console.log(str1.startsWith('hello'));

console.log(str1.substring(0, 5).toUpperCase() === 'HELLO');
console.log(str1.substring(0, 5).toUpperCase() === 'Hello');
console.log(str1.substring(0, 5).toUpperCase() === 'hello');

//let str2 = 'SeLECT 1 FROM DUAL';
let str2 = 'SeL';
console.log(str2.slice(0, 6).toUpperCase());
console.log(str2.substring(0, 6).toUpperCase());
console.log(str2.substring(0, 6).toUpperCase() === 'SELECT');

// function concatParams(firstParameter, secondParameter) {
//   return `CONCAT(COALESCE(${firstParameter}, ''), COALESCE(${secondParameter}, ''))`;
// }

function concatParams(val1, val2) {
  return `COALESCE(${val1}, '') || COALESCE(${val2}, '')`;
}

console.log(concatParams(undefined, 'bbb'));
