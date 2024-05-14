const all = async (obj) => {
    const result = {}
    for (const [key, value] of Object.entries(obj)) {
        result[key] = await Promise.resolve(value)
    }
    return result
}