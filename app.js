'use strict';

var bodyParser = require('body-parser');
var express = require('express');
var app = express();
var fs = require('fs');
var http = require('http');
var path = require('path');
var router = express.Router();

//Defining host and port number
var host = '127.0.0.1';
var port = 3000;

//Including the dependency into our node
app.use(bodyParser.json());
app.use(bodyParser.urlencoded());

//starting up the server
var Ibc1 = require('ibm-blockchain-js');
var ibc = new Ibc1();
var peer = JSON.parse(fs.readFileSync('peer_address.json', 'utf8'));
var options = {
  network: {
    peers: [peer.peers[0]],
    users: peer.users,
    options: {
      quiet: true,
        tls: false,
        maxRetry: 1
    }
  },
  chaincode: {
    zip_url: '',
    unzip_dir: '',
    git_url: ''
  }
}

var chaincode = null;
ibc.load(options, function(err, cc){
  if(err != null) console.log("Error while loading!");
  else{
    chaincode = cc;
    cc.deploy('init', ['99'], { delay_ms: 30000}, function(e){
      check_deployed(e, 1);
      console.log("I'm here yo!!");
    });
  }
});

function check_deployed(err, attempt){
  if (e) console.log("Error in deploying the chaincode.");
  else{
    console.log("Chaincode is ready");
    chaincode.query.read(['book_name'], function(err, resp){
      var cc_deployed = false;
      try{
        if (err == null && resp === 'null') cc_deployed = true;
        var json = JSON.parse(resp);
        if (json.constructor === Array) cc_deployed = true;
      } catch(e){
        console.log("Error in check_deployed: ", e);
      }

      if(cc_deployed) console.log("Successfully deployed.");
      else console.log("Error still exists.");
    });
  }
}

app.post("/login", function(req, res){
  //call to the chaincode function here

  //creating the response object
  res.setHeader('Content-Type', 'application/json');
  res.send(JSON.stringify({
    enrollId: req.body.enrollId,
    enrollSecret: req.body.enrollSecret,
    statusCode: 200
  }));
  console.log("Cool!");
});

app.post("/createBook", function(req, res){
  //call here for the creation of the book chaincode function


  //creating the response object
  res.setHeader('Content-Type', 'application/json');
  res.send(JSON.stringify({
    book_name: req.body.bookName,
    user_name: req.body.userName
  }));
  console.log("Was in /createBook with Book Name: ", req.body.bookName, "and User Name: ", req.body.userName);
});

/*router.use(function(req, res, next){
  console.log(req.url);
  next();
});

router.get('/sample', function(req, res){
  res.send("Hi there Boy!");
  console.log("Hi boy!");
});

app.use('/', router);
*/
/*app.listen(port, function(){
  console.log("Browse to localhost:3000 : ");
});*/
