var app = require('http').createServer(handler)
    , io = require('socket.io').listen(app)
    , receiveSocket = null
    
app.listen(4555);
    
function handler (req, res) {
    var body = '';
    req.on('data', function(chunk) {
        body += chunk.toString();
    });
    
    req.on('end', function() {
        console.log("Received message: " + body)
        try {
            var query = JSON.parse(body);
            if (query.command !== 'undefined') {
                if (receiveSocket !== null) {
                    receiveSocket.emit('command', body);
                }
            }
        } catch(e) {
            console.log("Error parsing '" + body + "' to JSON: " + e);
        }
        res.writeHead(200);
        res.end('OK');
    });
}

io.sockets.on('connection', function (socket) {
    if (receiveSocket === null) {
        console.log('Have receive socket');
        receiveSocket = socket;

        socket.on('disconnect', function () {
            receiveSocket = null;
        });
    }
});
