    const retry = (count, callback, t = 0) => (...args) => {
        return callback(...args).catch(() => {
            return (t > count) ? Promise.reject(new Error(`x:${args}`)) : retry(count, callback, t += 1)(...args)
        })
    }
    
    const timeout = (delay, callback) => (...args) => {
        return Promise.race([callback(...args), new Promise(resolve => setTimeout(resolve, delay))])
            .then((value) => { return (!value) ? Promise.reject(new Error(`timeout`)) : value})
    }