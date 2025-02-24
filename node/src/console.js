console.assert(true, "console.assert with true"); // 不打印
console.assert(false, "console.assert with false");
console.assert(false, "console.assert with false2");
// console.clear(); // 清屏

for (let index = 0; index < 5; index++) {
    console.log(`num: ${index}`); // 必须是``
}

const obj1 = { str: "Some text", id: 5 };
console.log(obj1);

const obj2 = {};
console.log(obj2);
obj2.prop = 123;

console.log(JSON.stringify(obj1));
console.log(JSON.parse(JSON.stringify(obj1)));

console.log("sec1", 'sec2');
console.log("%s, %d", "msg", 12);

console.log(
    "This is %cMy stylish message",
    "color: yellow; font-style: italic; background-color: blue;padding: 2px",
  );

console.log("This is the outer level");
console.group("First group");
console.log("In the first group");
console.group("Second group");
console.log("In the second group");
console.warn("Still in the second group");
console.groupEnd();
console.log("Back to the first group");
console.groupEnd();
console.debug("Back to the outer level");

console.time("answer time"); // 下面打印的内容必须一样
// alert("Click to continue");
console.timeLog("answer time");
// alert("Do a bunch of other stuff…");
console.timeEnd("answer time");

function foo() {
    function bar() {
        console.trace();
    }
    bar();
}
  
foo();

console.log("log");
console.debug("debug");
console.info("info");
console.warn("warn");
console.error("error");

console.table(["apples", "oranges", "bananas"]);

console.dir(obj1);

console.count("alice");
console.count("bob");
console.count("alice");
console.count("alice");
console.count("bob");
console.count("alice");
