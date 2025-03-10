// sudo apt-get install pkg-config
// sudo apt-get install libdbus-1-dev

var DBus = require('dbus');

var bus = DBus.getBus('session');

bus.getInterface('nodejs.dbus.ExampleService', '/nodejs/dbus/ExampleService', 'nodejs.dbus.ExampleService.Interface1', function(err, iface) {

	iface.Hello({ timeout: 1000 }, function(err, result) {
		if (err) {
			return console.log(err);
		}

		console.log(result);
	});
});
