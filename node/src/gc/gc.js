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
        arrayBuffers: format(memoryUsage.arrayBuffers),
    }));
};

var allocbig = function () {
    let arr1 = new Array(1024*1024*16);
    for (let i = 0; i < arr1.length; i++) {
        arr1[i] = i+1;
    }
    console.log(arr1);
}
setInterval(allocbig, 2000);

// setInterval(function() {
//     logger.info("start gc");
//     global.gc();
//     logger.info("gc end");
// }, 10000);

setInterval(function () {
    print();
}, 1000);
