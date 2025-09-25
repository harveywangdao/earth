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
    println!("{}", *gg)
}