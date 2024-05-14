function chunk(arr, size) {
    let myArr = []
    if (size === 0){
        return myArr;
    } for (let i = 0; i < arr.length/size; i++) {
        myArr.push(arr.slice(size*i, size*i+size))
    }
    return myArr
}