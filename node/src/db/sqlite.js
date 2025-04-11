const sqlite3 = require('sqlite3').verbose();

const db = new sqlite3.Database('mydatabase.db');

db.serialize(function () {
  db.run(`CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    age INTEGER
  )`);

  db.run(`UPDATE users SET age=? WHERE name=?`, [26, 'John Doe'], function (err) {
    console.log('err:', err);
    console.log('this:', this);
  });

  db.run(`INSERT INTO users (name, age) VALUES ('John Doe', 25)`, function (err) {
    console.log('err:', err);
    console.log('this:', this);
  });

  db.run(`SELECT * FROM users`, function (err) {
    if (err) {
      console.error(err);
    } else {
      console.log('this:', this);
    }
  });

  db.all(`UPDATE users SET age=? WHERE name=?`, [27, 'John Doe'], function (err, rows) {
    console.log('err:', err);
    console.log('rows:', rows);
    console.log('this:', this);
  });

  db.all(`SELECT * FROM users`, (err, rows) => {
    if (err) {
      console.error(err);
    } else {
      console.log(rows);
    }
  });

  db.close();
})
