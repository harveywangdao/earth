function* outer(){
    yield 'begin';
    yield inner();
    yield 'end';
}

function* inner(){
    yield 'inner';
}

var it = outer(), v;

v = it.next().value;

console.log(v); // begin

v = it.next().value;

console.log(v);
console.log(v.toString());

v = it.next().value;

console.log(v);
