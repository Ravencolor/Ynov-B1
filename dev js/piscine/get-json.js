async function getJSON(path = '', params = {}) {
    const url = path + '?' + Object.keys(params).map((key) => { return key.replace(' ', '+') + '=' + params[key].toString().replace(' ', '+') }).join('&')
    const result = await fetch(url).then((response) => { return (response.ok) ? response.json() : Promise.reject(new Error(response.statusText)) })
    return (result.error) ? Promise.reject(new Error(result.error)) : result.data
}