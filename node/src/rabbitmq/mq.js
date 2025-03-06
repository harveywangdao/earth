var amqp = require('amqplib/callback_api');

// amqp://192.168.79.27:5672

amqp.connect('amqp://localhost', function(error0, connection) {
    if (error0) {
      throw error0;
    }
    connection.createChannel(function(error1, channel) {
      if (error1) {
        throw error1;
      }
      var queue = 'hello';
      var msg = 'Hello world';
  
      channel.assertQueue(queue, {
        durable: false
      });
  
      setInterval(function() {
        channel.sendToQueue(queue, Buffer.from(msg), {
            persistent: false
        });
        console.log(" [x] Sent %s", msg);
      }, 2000);
    });
});

// setTimeout(function() {
//     connection.close();
//     process.exit(0)
// }, 5000);

amqp.connect('amqp://localhost', function(error0, connection) {
    if (error0) {
      throw error0;
    }
    connection.createChannel(function(error1, channel) {
      if (error1) {
        throw error1;
      }
      var queue = 'hello';
  
      channel.assertQueue(queue, {
        durable: false
      });
      
      console.log(" [*] Waiting for messages in %s. To exit press CTRL+C", queue);
      channel.consume(queue, function(msg) {
        console.log(" [x] Received %s", msg.content.toString());
        //channel.ack(msg);
      }, {
        noAck: true
      });
    });
});

// sudo rabbitmqctl list_queues
