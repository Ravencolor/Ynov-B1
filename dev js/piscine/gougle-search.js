    async function queryServers(serverName, q) {
        const uri1 = `/${serverName}?q=${encodeURI(q)}`
        const uri2 = `/${serverName}_backup?q=${encodeURI(q)}`
        return await Promise.race([getJSON(uri1), getJSON(uri2)])
    }
    
    async function gougleSearch(q) {
        return {
            web: await Promise.race([queryServers('web', q), (new Promise((resolve, reject) => setTimeout(reject, 80, new Error('timeout'))))]),
            image: await Promise.race([queryServers('image', q), (new Promise((resolve, reject) => setTimeout(reject, 80, new Error('timeout'))))]),
            video: await Promise.race([queryServers('video', q), (new Promise((resolve, reject) => setTimeout(reject, 80, new Error('timeout'))))]),
        }
    }