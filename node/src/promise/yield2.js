function* outer(){
    yield 'begin';
    var ret = yield* inner();
    console.log(ret);
    yield 'end';
}

function* inner(){
    yield 'inner';
    return 'return from inner';
}

var it = outer(), v;

v = it.next().value;

console.log(v);

v = it.next().value;

console.log(v);

v = it.next().value;

console.log(v)
