const p = new Promise((resolve, reject) => {
    setTimeout(()=>{
        let n = 10;
        if (n <= 30){
            resolve(n);//将promise对象的状态设置为「成功」,可以带出参数，供后面的then 回调函数 中使用
        } else {
            reject(n);//将promise对象的状态设置为「失败」,可以带出参数，供后面的then 回调函数 中使用
        }
    }, 1000);
});

p.then((value) => {
    console.log('恭喜,中奖号码为：' + value);
}, (reason) => {
    console.log('加油,您的号码为：' + reason);
})

let p2 = new Promise((resolve, reject) => {
    setTimeout(()=>{
        resolve('ok');
    }, 1000)
});

//链式调用
p2.then(value => {
    return new Promise((resolve, reject) => {
        resolve('success');
    })
}).then(value => {
    console.log(value)
}).then(value => {
    console.log(value)
}).catch(reason=>{
    console.warn(reason);
})

async function myFunction() {
    return "Hello, World!";
}

myFunction().then(result => {
    console.log(result); // 输出: Hello, World!
});

async function fetchData() {
    try {
        let response = await fetch('https://api.example.com/data');
        let data = await response.json();
        console.log(data);
    } catch (error) {
        console.error('Error:', error);
    }
}

fetchData();
