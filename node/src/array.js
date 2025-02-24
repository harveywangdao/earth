var cars = ["Saab", "Volvo", "BMW"];
var cars2 = ["BYD", "LI", "WEY"];
console.log(cars);
console.log(cars.length);
console.log(cars.toString());
console.log(typeof cars);
console.log(cars[0]);
console.log(cars.at(1));

console.log(cars.concat(cars2));
console.log(cars);
console.log(cars2);

console.log(cars.every((c) => {
    console.log(c);
    return c != "sLI";
}));

console.log(cars.constructor);

cars.forEach(function (v, i) {
    console.log(v, i);
});

console.log(cars.join("+"));

var points = [40,100,1,5,25,10];
// points.sort(); // 默认升序
points.sort(function(a,b){return b-a}); // 降序
console.log(points);

var fruits = ["Banana", "Orange", "Apple", "Mango"];
fruits.sort();
console.log(fruits);
fruits.reverse();
console.log(fruits);

var fruits2 = fruits.valueOf();
console.log(fruits2); // 指向同一个数组
fruits2[0] = "111";
console.log(fruits);
console.log(fruits2);

console.log(Array.of(1, 2, 3));
