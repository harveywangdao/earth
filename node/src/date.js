var d1 = new Date();
var d2 = new Date(1000*10); // 参数为毫秒
var d3 = new Date("2025-01-24");
var d4 = new Date(2025, 3, 24, 12, 59, 59, 0);
console.log(d1);
console.log(d2);
console.log(d3);
console.log(d4);

console.log(Date.now());
console.log(new Date().getTime());
