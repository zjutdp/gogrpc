promise = new Promise(function(resolve, reject){
  console.log('Executed immediately after Promise newed!')
  resolve("Return value when resolved!")
});

promise.then((resolvedResult) => {
  console.log(`Resolved with result: ${resolvedResult}, Executed after last sentence evaluated!`);
}, (rejectedResult)=>{console.log('rejected')});

console.log('Hi, I am here!')
