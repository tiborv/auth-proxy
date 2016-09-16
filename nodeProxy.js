<<<<<<< HEAD
var http = require('http');

var server = http.createServer();
server.on('request', function(request, response) {
  console.log(request.headers)

});

server.listen(5000);
=======
var express    = require('express')
var bodyParser = require('body-parser')

var app = express()

// parse application/json
app.use(bodyParser.json())

app.use(function (req, res, next) {
  console.log(req.headers)
  console.log(req.body) // populated!
  next()
})

app.listen(5000);
>>>>>>> 3460fd86d3d6c1bd3ec76022f6e3ac884477016b
