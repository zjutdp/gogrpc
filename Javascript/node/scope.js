/*
Hoisting
Understanding the concept of hoisting is fundamental to understanding how JavaScript works. 
JavaScript has two phases: a parsing phase—where all of the code is read by the JavaScript 
engine—followed by an execution phase in which the code that has been parsed is executed. 
It is during this second phase that most things happen; for example, when you use a console.log 
statement, the actual log message is printed to the console during the execution phase.

However, some important things happen during the parsing phase as well, including memory allocation
 for variables and scope creation. The term hoisting describes what happens when the JavaScript
  engine encounters an identifier, such as a variable or function declaration; when it does this,
   it acts as if it literally lifts (hoists) that declaration up to the top of the current scope.

Use keyword "let" will suppress the Hoisting of the varables in block

*/
var globalVar = "1"

if(true){
	var insideVar = "2"
	console.log('insider');
}

console.log(globalVar, insideVar);


function fn1() {
  var x = 'function scope';

  if (true) {
    var y = 'not block scope';
  }

  function innerFn() {
    console.log(x, y); // function scope not block scope
  }
  innerFn();
}

fn1();
