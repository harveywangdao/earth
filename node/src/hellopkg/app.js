const format = function (bytes) {
    return (bytes / 1024 / 1024).toFixed(2) + ' MB';
};

const print = function() {
    const memoryUsage = process.memoryUsage();

    console.log(JSON.stringify({
        rss: format(memoryUsage.rss),
        heapTotal: format(memoryUsage.heapTotal),
        heapUsed: format(memoryUsage.heapUsed),
        external: format(memoryUsage.external),
    }));
}

function Quantity(num) {
    if (num) {
        return new Array(num * 1024 * 1024);
    }
    return num;
}

function Fruit(name, quantity) {
    this.name = name
    this.quantity = new Quantity(quantity)
}

let apple = new Fruit('apple');
print();
let banana = new Fruit('banana', 20);
print();

banana = null;

let cc = new Fruit('cc', 20);
let dd = new Fruit('dd', 20);
cc = null;
dd = null;
print();


// const total = [];
// setInterval(function() {
//     total.push(new Array(20 * 1024 * 1024)); // 大内存占用
//     print();
// }, 1000)

setInterval(function() {
    let xx = new Fruit('xx', 20);
    //xx = null;
    print();
    //global.gc();
}, 1000)

setInterval(function() {
    global.gc();
}, 1000)
