const co = require('co');

// function* myGenerator() {
//     yield 'Hello';
//     yield 'World';
//     yield '!';
// }

// const generator = myGenerator();

// console.log(generator.next().value); // 输出: Hello
// console.log(generator.next().value); // 输出: World
// console.log(generator.next().value); // 输出: !
// console.log(generator.next().done);  // 输出: true (表示生成器函数执行完毕)

function fly1() {
    console.log("fly1");
}

function fly2() {
    console.log("fly2");
    return 12;
}

function fly3() {
    console.log("fly3");
    //yield 13;
    return 14;
}

function* fly4() {
    console.log("fly4");
    //yield 15;
    return 16;
}

function fly5() {
    console.log("fly5");
    return new Promise(function (resolve, reject) {
        setTimeout(function () {
            resolve(17);
        }, 5000);
    });
}

function* test() {
    //let a1 = yield fly1();
    //console.log("a1:", a1);
    //let a2 = yield fly2();
    //console.log("a2:", a2);
    //let a3 = yield fly3();
    //console.log("a3:", a3);
    let a4 = yield fly4();
    console.log("a4:", a4);
    let a41 = yield* fly4();
    console.log("a41:", a41);
    let a5 = yield fly5();
    console.log("a5:", a5);
}

co(function* () {
    yield* test();
    yield test();
});
