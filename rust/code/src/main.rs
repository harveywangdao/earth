fn main() {
    let a = 12;
    println!("a is {}", a);
    println!("a is {}, a again is {}", a, a);
    println!("a is {0}, a again is {0}", a);
    println!("{{}}");
    test2();
}

fn test1() {
    // 不可变变量
    let a = 1;
    // 可变变量
    let mut b: i32 =2;

    println!("a = {}, b = {}", a, b);
    //a = 3;
    b = 4;
    
    println!("a = {}, b = {}", a, b);

    // let c = 1; // 允许
    let c = 2;
    println!("c = {}", c);

    const D: i32 = 1;
    //const d: i32 = 2;
    //d = 2;
    //let d = 2;
    println!("d = {}", D);

    let e8: i8 = 8;
    let e16: i16 = 16;
    let e32: i32 = 32;
    let e64: i64 = 64;
    let e128: i128 = 128;
    println!("e8 = {}, e16 = {}, e32 = {}, e64 = {}, e128 = {}", e8, e16, e32, e64, e128);

    let f8: u8 = 8;
    let f16: u16 = 16;
    let f32: u32 = 32;
    let f64: u64 = 64;
    let f128: u128 = 128;
    println!("f8 = {}, f16 = {}, f32 = {}, f64 = {}, f128 = {}", f8, f16, f32, f64, f128);

    //let g1: f16 = 1.2;
    let g2: f32 = 1.2;
    let g3: f64 = 1.200;
    //let g4: f128 = 1.2;
    println!("g2 = {}, g3 = {}", g2, g3);

    let bo: bool = true;
    let c1: char = 'A';
    println!("bo = {}, char = {}", bo, c1);

    println!("add(1,2)={}", add(1,2));

    let h1: i64 = 5;
    if h1 >= 5 {
        println!("h1 >= 5")
    } else {
        println!("h1 < 5")
    }

    let mut j1 = 0;
    loop {
        j1 += 1;
        if j1 > 10 {
            println!("break, j1 = {}", j1);
            break;
        }
    }

    let mut k1 = 0;
    while k1 < 10 {
        k1 += 1;
    }
    println!("while, k1 = {}", k1);

    for n in 0..2 {
        println!("for, n = {}", n);
    }

    let s1 = String::from("hello");
    let s2 = s1;
    //println!("s1 = {}", s1);
    println!("s2 = {}", s2);

    // 借用 引用数据而不获取所有权
    let s3 = String::from("hello");
    let len = calculate_length(&s3);
    println!("s3 len = {}", len);

    let u1 = User{
        name: String::from("xiaoming"),
        age: 12,
        active: true,
    };
    println!("u1.name = {} {} {}", u1.name, u1.age, u1.active);

    let l1 = Level::Debug;
    //println!("l1 = {}", l1);
    println!("GetLevel(l1) = {}", get_level(l1));
    println!("GetLevel(l1) = {}", get_level(Level::Info));

    let _ret1 = divide(1,0);
    //println!("divide(1/0) = {}", ret1);
    let _ret2 = divide(6,2);
    //println!("divide(6/2) = {}", ret2);

    let s1 = "123";
    println!("s1 = {}", s1);
    let s1 = s1.len();
    println!("s1 = {}", s1);
}

fn _get_element(index: usize, vec: &Vec<i32>) -> Option<i32> {
    if index < vec.len() {
        Some(vec[index])
    } else {
        None
    }
}

// enum Result<T, E> {
//     Ok(T),
//     Err(E),
// }

fn divide(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err(String::from("divide 0"))
    } else {
        Ok(a/b)
    }
}

fn get_level(level: Level) -> u8 {
    match level {
        Level::Debug => 1,
        Level::Info => 2,
    }
}

struct User {
    name: String,
    age: i32,
    active: bool,
}

enum Level {
    Debug,
    Info,
}

fn calculate_length(s: &String) -> usize {
    s.len()
}

fn add(a: i32, b: i32) -> i32 {
    //a - b
    a + b
    //return a + b;
}

fn test2() {
    let a: i32 = 2; // 显示i32
    let b = a.pow(3);
    println!("2^3 = {}", b);

    for n in 0..=2 {
        println!("{}", n);
    }

    let c = 3;
    let f = c as f32;
    println!("f = {}", f);

    let d = 10;
    let gg = &d;
    println!("{}", *gg);

    testVec();
    testStr();
    test_iter();
    test_closure();
}

fn testVec() {
    let mut v1: Vec<i32> = Vec::new();
    let mut v2 = vec![1,2,3,4,6];
    let v3: Vec<i32> = Vec::with_capacity(10);
    let v4 = Vec::from([3,5,8]);
    let v5 = vec![2; 5];

    println!("v3: {:?}", v3);
    println!("v4: {:?}", v4);
    println!("v5: {:?}", v5);

    v1.push(13);
    v1.push(3);
    println!("v1 = {:?}", v1);

    let a1: i32 = v1[0];
    let a2: &i32 = &v1[0];
    //v1.push(3);
    //let a3 = &v1[10];
    println!("a1 = {}", a1);
    println!("a2 = {}", a2);
    //println!("a3 = {}", a3);

    match v1.get(19) {
        Some(a3) => println!("a3 = {}", a3),
        None => println!("no key"),
    }

    for n in &v2 {
        println!("{}", n)
    }
    for n in &mut v2 {
        *n *= 10;
    }
    for n in &v2 {
        println!("{}", n)
    }

    let e1 = v2.pop();
    match e1 {
        Some(element) => println!("element = {}", element),
        None => println!("no element"),
    }
    for n in &v2 {
        println!("{}", n)
    }
    let e2 = v2.remove(1);
    println!("e2 = {}", e2);
    for n in &v2 {
        println!("{}", n)
    }
    println!("len: {}, cap: {}", v2.len(), v2.capacity());
    v2.clear();
    println!("is_empty: {}", v2.is_empty());

    trait Animal {
        fn speak(&self);
    }
    
    struct Dog;
    struct Cat;
    
    impl Animal for Dog {
        fn speak(&self) { println!("Woof!"); }
    }
    impl Animal for Cat {
        fn speak(&self) { println!("Meow!"); }
    }
    
    let animals: Vec<Box<dyn Animal>> = vec![
        Box::new(Dog),
        Box::new(Cat),
    ];
    
    for animal in animals {
        animal.speak();
    }

    // #[derive(Debug)]
    // enum SpreadsheetCell {
    //     Int(i32),
    //     Float(f64),
    //     Text(String),
    // }

    // let row = vec![
    //     SpreadsheetCell::Int(3),
    //     SpreadsheetCell::Text(String::from("蓝色")),
    //     SpreadsheetCell::Float(10.12),
    // ];
}

fn testStr() {
    let mut s1 = String::new();
    let s2: &str = "initial contents";
    let s3 = s2.to_string();
    let s4 = 123.to_string();
    let s5 = String::from("value");
    let s6: String = "asdasd".into();
    s1.push_str("ssssd");
    println!("s1 = {}", s1);
    println!("s2 = {}", s2);
    println!("s3 = {}", s3);
    println!("s4 = {}", s4);
    println!("s5 = {}", s5);
    println!("s6 = {}", s6);
    let s7 = s5 + &s6;
    println!("s7 = {}", s7);

    let s8 = format!("{s1}{s2}{s1}");
    println!("s8 = {}", s8);

    let s10: &str = "Hello, world!";
    println!("s10 = {}", s10);

    let s11: String = String::from("value11");
    let s12: &str = &s11;
    let s13: String = s11[0..2].to_string();
    let s14 = &s11[0..2];
    println!("s11 = {}", s11);
    println!("s12 = {}", s12);
    println!("s13 = {}", s13);
    println!("s14 = {}", s14);

    let s15: &'static str = "我是abc";
    for b in s15.bytes() {
        println!("{}", b);
    }
    for c in s15.chars() {
        println!("{}", c);
    }
}

// fn test_lifetime() {
//     let r;
//     {
//         let x = 5;
//         r = &x;
//     }
//     println!("{}", r);
// }

// fn longest(x: &str, y: &str) -> &str {
//     if x.len() > y.len() {
//         x
//     } else {
//         y
//     }
// }

fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn test_iter() {
    let v1: Vec<i32> = vec![1,2,3,4,5];
    let iter1= v1.iter();
    for n in iter1 {
        println!("{}", n);
    }

    let mut v2 = vec![4, 5, 6];
    let iter2 = v2.iter_mut();
    for n in iter2 {
        *n *= 10;
    }
    println!("v2 = {:?}", v2);

    let v3 = vec![7, 8, 9];
    let iter3 = v3.into_iter();
    for n in iter3 {
        println!("{}", n);
    }
    //println!("v3 = {:?}", v3);

    let v4 = vec![1, 2, 3, 4, 5];
    let doubled: Vec<i32> = v4.iter().map(|x| x * 2).collect();
    println!("{:?}", doubled); // [2, 4, 6, 8, 10]

    let evens: Vec<&i32> = v4.iter().filter(|x| *x % 2 == 0).collect();
    println!("{:?}", evens); // [2, 4]

    let first_three: Vec<&i32> = v4.iter().take(3).collect();
    println!("{:?}", first_three); // [1, 2, 3]

    for (index, value) in v4.iter().enumerate() {
        println!("索引 {}: 值 {}", index, value);
    }

    let numbers = vec![1, 2, 3, 4, 5];
    let total: i32 = numbers.iter().sum();
    println!("总和: {}", total); // 15
    let count = numbers.iter().count();
    println!("数量: {}", count); // 5
    let has_even = numbers.iter().any(|x| x % 2 == 0);
    let all_positive = numbers.iter().all(|x| x > &0);
    println!("有偶数: {}, 全部为正数: {}", has_even, all_positive);

    let first_even = numbers.iter().find(|x| *x % 2 == 0);
    println!("第一个偶数: {:?}", first_even); // Some(2)

    let sum = numbers.iter().fold(0, |acc, x| acc + x);
    let product = numbers.iter().fold(1, |acc, x| acc * x);
    println!("和: {}, 积: {}", sum, product); // 和: 15, 积: 120

    let max = numbers.iter().max();
    let min = numbers.iter().min();
    println!("最大值: {:?}, 最小值: {:?}", max, min);

    let sentence = "hello world this is rust";
    let capitalized: String = sentence
        .split_whitespace() // 分割成单词迭代器
        .map(|word| {
            let mut chars = word.chars();
            match chars.next() {
                None => String::new(),
                Some(first) => first.to_uppercase().chain(chars).collect(),
            }
        })
        .collect::<Vec<String>>() // 收集到 Vec
        .join(" "); // 再连接成字符串
    
    println!("{}", capitalized);

    struct Person {
        name: String,
        age: u32,
    }
    
    let people = vec![
        Person { name: "Alice".to_string(), age: 30 },
        Person { name: "Bob".to_string(), age: 25 },
        Person { name: "Charlie".to_string(), age: 35 },
    ];
    
    // 找出年龄大于25岁的人的名字，并排序
    let names: Vec<String> = people
        .into_iter() // 获取所有权
        .filter(|p| p.age > 25) // 过滤
        .map(|p| p.name) // 提取名字
        .collect(); // 收集
    println!("{:?}", names);
}

fn apply_to_value<F>(val: i32, f: F) -> i32
where
    F: Fn(i32) -> i32,
{
    f(val)
}

// fn make_adder(x: i32) -> impl Fn(i32) -> i32 {
//     move |y| x + y
// }

fn make_adder(x: i32) -> Box<dyn Fn(i32) -> i32> {
    Box::new(move |y| x + y)
}

fn call_closure<F>(f: F)
where
    F: FnOnce(),
{
    f(); // 只调用一次
}

fn test_closure() {
    let double = |x| x * 2;
    let result = apply_to_value(5, double);
    println!("Result: {}", result); // 输出: Result: 10

    let add_five = make_adder(5);
    println!("5 + 3 = {}", add_five(3)); // 输出: 5 + 3 = 8

    let add_ten = make_adder(10);
    println!("10 + 2 = {}", add_ten(2)); // 输出: 10 + 2 = 12

    let name = String::from("Rust");

    // 使用 move 强制捕获所有权
    let print_name = move || println!("Hello, {}!", name);

    call_closure(print_name);
    call_closure(print_name);
}