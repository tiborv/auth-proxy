var http = require('http');

var server = http.createServer();
server.on('request', function(request, response) {
  console.log(request.headers)

});

server.listen(5000);
