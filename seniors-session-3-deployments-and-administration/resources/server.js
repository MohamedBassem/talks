var http = require('http');
var counter = 1;

http.createServer(function (req, res) {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  console.log("Got a new Request " + req.url);
  res.end("You are vistor number : " + counter);
  counter++;
}).listen(8080);
