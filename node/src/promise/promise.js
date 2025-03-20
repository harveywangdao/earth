// const p = new Promise((resolve, reject) => {
//     setTimeout(()=>{
//         let n = 10;
//         if (n <= 30){
//             resolve(n);//将promise对象的状态设置为「成功」,可以带出参数，供后面的then 回调函数 中使用
//         } else {
//             reject(n);//将promise对象的状态设置为「失败」,可以带出参数，供后面的then 回调函数 中使用
//         }
//     }, 1000);
// });

// p.then((value) => {
//     console.log('恭喜,中奖号码为：' + value);
// }, (reason) => {
//     console.log('加油,您的号码为：' + reason);
// })

// async function myFunction() {
//     return "Hello, World!";
// }

// myFunction().then(result => {
//     console.log(result); // 输出: Hello, World!
// });

// async function fetchData() {
//     try {
//         let response = await fetch('https://api.example.com/data');
//         let data = await response.json();
//         console.log(data);
//     } catch (error) {
//         console.error('Error:', error);
//     }
// }

// fetchData();

// let testPromise = new Promise((resolve, reject) => {
//     setTimeout(() => {
//         if (true) {
//             resolve('resolve1');
//         } else {
//             reject('reject1');
//         }
//     }, 1000);
// });

// testPromise.then(value => {
//     console.info('then1 fulfilled value:', value);
// }, reason => {
//     console.info('then1 rejected reason:', reason);
// }).then(value => {
//     console.info('then2 value:', value);
// }).then(value => {
//     console.info('then3 value:', value);
// }).catch(reason => {
//     console.info('catch reason:', reason);
// }).finally (() => {
//     console.info('finally');
// });

// let testPromise2 = new Promise((resolve, reject) => {
//     reject('fail');
//     resolve('success');
// });
// testPromise2.then(value => {
//     console.info('then  value:', value);
// }).catch(reason => {
//     console.info('catch reason:', reason);
// }).finally (() => {
//     console.info('finally');
// });

// new Promise(function (resolve, reject) {
//     console.log(1111);
//     resolve(2222);
// }).then(function (value) {
//     console.log(value);
//     return 3333;
// }).then(function (value) {
//     console.log(value);
//     throw "An error";
// }).catch(function (err) {
//     console.log(err);
// });

// catch只执行一个，finally执行多个，throw立即跳到第一个catch，return当成resolve(val)
// resolve/reject只执行第一个
// new Promise(function (resolve, reject) {
//     console.log(1111);
//     resolve(2222);
// }).then(function (value) {
//     console.log('then1:', value);
//     return 3333;
// }).then(function (value) {
//     console.log('then2:', value);
//     throw "An error";
// }).then(function (value) {
//     console.log('then3:', value);
//     return 444;
// }).catch(function (err) {
//     console.log('catch1', err);
// }).catch(function (err) {
//     console.log('catch2', err);
// }).finally(() => {
//     console.log('finally1');
// }).finally(() => {
//     console.log('finally2');
// });

new Promise((resolve, reject) => {
    resolve(12);
}).then((val) => {
    console.log('then:', val);
}).catch((err) => {
    console.error('catch:', err);
}).finally(() => {
    console.log('finally');
})
