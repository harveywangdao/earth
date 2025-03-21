
function closePromise() {
  return new Promise(function(resolve, reject) {
    console.log('Promise');
    resolve(12);
  });
}

function close() {
  let arr = [];
  arr.push(closePromise());
  arr.push(closePromise());
  console.log('arr:', arr);
  return Promise.all(arr);
}

async function testp() {
  let res = await close();
  console.log('res:', res);
}

testp();
