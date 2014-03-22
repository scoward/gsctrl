function gsClosure() {
    var socket = io.connect('http://localhost:4555')
    socket.on('connect', function() {
        console.log('connected')
    });
    socket.on('command', function(data) {
        console.log(data)
        try {
            var command = JSON.parse(data);
        } catch (e) {
            console.log("Error parsing '" + data + "' to JSON: " + e);
            return;
        }

        if (command.command === 'undefined') {
            return;
        } else if (command.command === 'next') {
            Grooveshark.next();
        } else if (command.command === 'previous') {
            Grooveshark.previous();
        } else if (command.command === 'stop') {
            Grooveshark.stop();
        } else if (command.command === 'start') {
            Grooveshark.start();
        } else if (command.command === 'toggle') {
            Grooveshark.togglePlayPause();
        } else if (command.command === 'volume') {
            Grooveshark.setVolume(command.volume);
        }
    })
}

window.addEventListener('load', function() {
    var socket_script = document.createElement('script');
    socket_script.src = "http://localhost:4555/socket.io/socket.io.js";
    document.body.appendChild(socket_script);
    
    setTimeout(function() {
        var script = document.createElement('script');
        script.textContent = ";(" + gsClosure.toString() + ")();";
        document.body.appendChild(script);
    }, 500);
}, false);
