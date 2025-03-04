setTimeout(() => {
    console.log('Timeout callback');
}, 1000);

setImmediate(() => {
    console.log('Immediate callback');
});

setInterval(() => {
    console.log('Interval callback');
}, 1000);

console.log('Main thread');
