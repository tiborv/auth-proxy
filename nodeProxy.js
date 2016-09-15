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
