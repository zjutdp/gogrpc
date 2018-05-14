let arr = [1, 2, 3].map(
    (value) => {
        return new Promise((resolve, reject) => {
            setTimeout(() => {
                resolve(value);
            }, value * value * 1000);
        });
    }
);

console.log(arr);

let promises = Promise.race(arr)
.then((result) => {
    console.log(result);
}).catch((err) => {
    console.log(err);
});
