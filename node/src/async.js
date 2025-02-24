const fs = require('fs').promises;

async function readFiles() {
    try {
        const data1 = await fs.readFile('hello.js', 'utf8');
        const data2 = await fs.readFile('hello.js', 'utf8');
        const data3 = await fs.readFile('hello.js', 'utf8');

        console.log('Data from all files:', data1, data2, data3);
    } catch (err) {
        console.error('Error reading files:', err);
    }
}

readFiles();
