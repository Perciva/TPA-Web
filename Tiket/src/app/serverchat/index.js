var io = require('socket.io')(1000);

console.log('Server Started');

io.on('connection', function(socket) {
    console.log('User Connected');
    socket.on('chat', function(msg) {
        io.emit('chat', msg);
    });
    socket.on('blog', function(msg) {
	        io.emit('blog', msg);
    });
    socket.on('train', function(msg) {
		        io.emit('train', msg);
    });
    socket.on('plane', function(msg) {
		        io.emit('plane', msg);
    });
    socket.on('event', function(msg) {
		        io.emit('event', msg);
    });
    socket.on('hotel', function(msg) {
		        io.emit('hotel', msg);
    });
});

