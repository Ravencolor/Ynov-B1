function nasa(N = 0){
    let str = ''
    for (let i = 1; i<=N; i++){
        if (i % 3 === 0 && i% 5 !== 0){
            str += 'NA' + ' '
        } else if (i % 3 != 0 && i % 5 === 0) {
            str += 'SA' + ' '
        } else if (i % 3 === 0 && i% 5 === 0){
            str += 'NASA' + ' '
        } else {
            str += String(i) + ' '
        }
    }
    return str.slice(0, -1)
}
