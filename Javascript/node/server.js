/*
1. os.cpus return the phsical CPU core number
2. process.pid will return the same process ID for all the fork() sub process/thread
*/


var cluster = require('cluster')
var http = require('http')
var numCPUS = require('os').cpus().length;

if(cluster.isMaster){
	for(var i=0; i<numCPUS; i++){
		let pid = cluster.fork();
		console.log('Created thread with ID: ', pid.id);
	}

	cluster.on('death', function(work){
		console.log('worker ' + worker.pid + ' died');
	})
}else{
	http.Server(function(req, res){
		res.writeHead(200);
		let msg = "hello world from process: " + process.pid + "!\n"
		res.end(msg);
	}).listen(8000);
}
