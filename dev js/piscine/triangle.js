function triangle(str = "*", num = 0){
    let final = ""
    for (let i = 0; i < num ; i++) {
        final += "" +str.repeat(i + 1)
        if (i < num - 1) {
            final += '\n'
        }
    }
    return final 
}
