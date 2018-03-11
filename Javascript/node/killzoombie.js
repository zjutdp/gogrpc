/*
1. os.cpus return the phsical CPU core number
2. process.pid will return the same process ID for all the fork() sub process/thread
*/

var cluster = require('cluster')
var http = require('http')
var numCPUS = require('os').cpus().length;

var rssWarn = (50 * 1024 * 1024), heapWarn = (50 * 1024 * 1024)

var workers = {}

if(cluster.isMaster){
	for(var i=0; i<numCPUS; i++){
		createWorker();
	}

	setInterval(function(){
		var time = new Date().getTime()
		for(pid in workers){
			if(workers.hasOwnProperty(pid) && workers[pid].lastCb + 5000 < time){
				console.log('Long running woekr ' + pid + ' killed')
				workers[pid].worker.kill()
				delete workers[pid]
				createWorker()
			}
		}
	}, 1000)
}else{
	http.Server(function(req, res){
		if(Math.floor(Math.random() * 200 ) === 4){
			console.log('Stopped ' + process.pid + ' from ever finishing')
			while(true) { continue }
		}

		res.writeHead(200);
		res.end('hellow world from ' + process.pid + '\n')
	}).listen(8000);

	setInterval(function report(){
		//process.send({cmd: "reportMem", memory: process.memoryUsage(), process: process.pid})
		process.send({cmd: "reportMem", memory: process.memoryUsage(), process: Math.floor(Math.random() * numCPUS) + 1})
	}, 1000)
}

function createWorker(){
	var worker = cluster.fork();
	console.log("created worker: " + worker.pid)
	workers[worker.id] = {worker:worker, lastCb: new Date().getTime() - 1000}
	worker.on('message', function(m){
		if(m.cmd === 'reportMem'){
			console.log('on message of m.process: ' + m.process)
			console.log('on message of workers[m.process]: ' + workers[m.process])			
			workers[m.process].lastCb = new Date().getTime()
			if(m.memory.rss > rssWarn){
				console.log('Worker ' + m.process + ' using too much memory.')
			}
		}
	})
}