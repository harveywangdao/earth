function* myGenerator() {
    yield 'Hello';
    yield 'World';
    yield '!';
}

const generator = myGenerator();

console.log(generator.next().value); // 输出: Hello
console.log(generator.next().value); // 输出: World
console.log(generator.next().value); // 输出: !
console.log(generator.next().done);  // 输出: true (表示生成器函数执行完毕)

