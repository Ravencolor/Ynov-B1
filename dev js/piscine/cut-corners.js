function round(num) {
    let i = 0;
    if (num >= 0) {
        for (i=0; i < num;) {
            i++
        }
        if (i-num > 0.5){
            return i - 1
        }
        if (i-num <= 0.5){ 
            return i
        }
    }
    if (num < 0) {
        i = 0
        for (i=0; i > num;) {
            i--
        }
        if (i-num > -0.5) {
            return i
        }
        if (i-num <= 0.5) {
            if (i + 1 === 0){
                return -0
            }
            return i + 1
        }
    }
}

  function ceil(num) {
    let i = 0;
    if (num >= 0) {
        for (; i < num;) {
            i++
        }
        if (i === num) {
            return i--
        }
        return i
    }
    if (num < 0) {
        i = 0
        for (; i > num;) {
            i--
        }
        if (i === num) {
            return i
        }
        if (i + 1 === 0) {
            return -0
        }
        return i + 1
    }
}

function floor(num) {
    let i = 0;
    if (num >= 0) {
        for (; i < num;) {
            i++
        }
        if (i === num) {
            return i
        }
        return i - 1
    }
    if (num < 0) {
        i = 0
        for (i = 0; i > num;) {
            i--
        }
        if (num === -0) {
            return -0
        }
        if (i === num) {
            return i
        }
        return i
    }
}

function findNearest(x) {
    let num = 0;
    if (x < 1 && x >= 0) return 0
    if (x < 0 && x >= -1) return 0
    for (let i = x; i > 1; i = i / 2) {
        num += 1;
    }
    return num - 1
}

function trunc(x) {
    let i = 2 ** findNearest(x);
    if (x >= 0) {
        for (; i < x;) {
            i++
        }
        if (i === x) return i
        return i - 1
    }
    if (x < 0) {
        i = 0
        for (; i > x;) {
            i--
        }
        if (x > -1) return -0
        if (i === x) return i
        return i + 1
    }
}