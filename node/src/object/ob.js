let rows = [
  {
    tenant: 'localhost',
    id: '83631438fec240c4b47496403b9b68c3',
    status: 0,
    status_info: 0,
    created_at: 1744104969336,
    last_open_date: 1744104969336,
    user_index: 1,
    change_id: 65,
    callback: '\x05{"userIndex":1,"callback":"http://127.0.0.1:3302/ugreen/v1/office/file/callback?uid=1000"}',
    baseurl: 'https://192.168.79.27:9443',
    password: null,
    additional: null
  }
]

console.log('created_at' in rows[0]);
console.log('created_ats' in rows[0]);
console.log(rows[0].hasOwnProperty('created_at')); // Object.hasOwn
console.log(rows[0]?.created_at);
console.log(rows[0]?.created_ats);
console.log(rows[0]?.created_at === undefined);
console.log(rows[0]?.created_at !== undefined);
console.log(rows[0]?.created_ats === undefined);
console.log(rows[0]?.created_ats !== undefined);

console.log(typeof rows[0].created_at);
console.log(typeof rows[0].created_at === 'number');

for (let i = 0; i < rows.length; i++) {
  if (rows[i].hasOwnProperty('change_date')) {
    rows[i].change_date = new Date(rows[i].change_date);
  }
  if (rows[i].hasOwnProperty('created_at')) {
    rows[i].created_at = new Date(rows[i].created_at);
  }
  if (rows[i].hasOwnProperty('last_open_date')) {
    rows[i].last_open_date = new Date(rows[i].last_open_date);
  }
}

console.log(rows);

for (const key in rows[0]) {
  if (Object.prototype.hasOwnProperty.call(rows[0], key)) {
    console.log(key, rows[0][key]);
  }
}

for (const element of rows) {
  console.log(element);
}

let ob1 = {
  key1: 'value1'
}
// for (const element of ob1) {
//   console.log(element);
// }

for (const key of Object.keys(ob1)) {
  console.log(key);
}

for (const val of Object.values(ob1)) {
  console.log(val);
}
