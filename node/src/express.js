const express = require('express');
const app = express();
const port = 3000;

// 定义一个 GET 路由
app.get('/', (req, res) => {
    res.send('Hello, World!');
});

// 定义一个 POST 路由
app.post('/submit', (req, res) => {
    res.send('Form submitted!');
});

// 启动服务器
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});

app.get('/users/:id', (req, res) => {
    const userId = req.params.id;
    res.send(`User ID: ${userId}`);
});

app.get('/search/:query', (req, res) => {
    const query = req.params.query;
    res.send(`Search query: ${query}`);
});

const logger = (req, res, next) => {
    console.log(`Request Type: ${req.method} ${req.url}`);
    next();
};

// 使用中间件
app.use(logger);

// 创建一个路由器实例
const userRouter = express.Router();

// 定义用户相关的路由
userRouter.get('/', (req, res) => {
    res.send('List of users');
});

userRouter.get('/:id', (req, res) => {
    const userId = req.params.id;
    res.send(`User ID: ${userId}`);
});

// 挂载用户路由器
app.use('/users', userRouter);
