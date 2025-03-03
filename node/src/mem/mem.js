const heapdump = require('heapdump');

heapdump.writeSnapshot('init.heapsnapshot');

let i = 0;
let theThing = null;
let replaceThing = function() {
  const newThing = theThing;
  let unused = function() {
    if (newThing) console.log("hi");
  };

  // 不断修改引用
  theThing = {
    longStr: new Array(1e8).join("*"),
    someMethod: function() {
      console.log("a");
    },
  };

  if (++i >= 1000) {
    heapdump.writeSnapshot('leak.heapsnapshot'); // 记录运行一段时间后内存的堆快照
    process.exit(0);
  }
};

setInterval(replaceThing, 100);
