var arr1 = [1];
var allocbig = function () {
    arr1 = new Array(1024*1024*16);
    for (let i = 0; i < arr1.length; i++) {
        arr1[i] = i+1;
    }
    console.log(arr1);
}
setTimeout(allocbig, 20000);

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
}

setInterval(function () {
    print();
    console.log("arr1.length:", arr1.length);
}, 1000);
