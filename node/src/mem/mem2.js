const v8 = require('v8');
const fs = require('fs');

const fileName = `${Date.now()}.heapsnapshot`;
v8.writeHeapSnapshot(fileName);
console.log('Heap snapshot written to', fileName);
