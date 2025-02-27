const test1 = {
    x: 42,
    getX: function () {
      return this.x;
    },
  };
  
  const unboundGetX = test1.getX;
  console.log(unboundGetX()); // The function gets invoked at the global scope
  // Expected output: undefined
  
  const boundGetX = unboundGetX.bind(test1);
  console.log(boundGetX());
  // Expected output: 42

function Product(name, price) {
    this.name = name;
    this.price = price;
}

function Food(name, price) {
    Product.call(this, name, price);
    this.category = "food";
}

console.log(new Food("cheese", 5));

const numbers = [5, 6, 2, 3, 7];

const max = Math.max.apply(null, numbers);

console.log(max);
// Expected output: 7

const min = Math.min.apply(null, numbers);

console.log(min);
// Expected output: 2



const animals = ["ant", "bison", "camel", "duck", "elephant"];

console.log(animals.slice(2));
// Expected output: Array ["camel", "duck", "elephant"]

console.log(animals.slice(2, 4));
// Expected output: Array ["camel", "duck"]

console.log(animals.slice(1, 5));
// Expected output: Array ["bison", "camel", "duck", "elephant"]

console.log(animals.slice(-2));
// Expected output: Array ["duck", "elephant"]

console.log(animals.slice(2, -1));
// Expected output: Array ["camel", "duck"]

console.log(animals.slice());
// Expected output: Array ["ant", "bison", "camel", "duck", "elephant"]

const arrayLike = {
    length: 3,
    0: 2,
    1: 3,
    2: 4,
};
console.log(Array.prototype.slice.call(arrayLike, 1, 3));
