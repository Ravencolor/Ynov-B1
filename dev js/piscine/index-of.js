function indexOf(searchElement, target, index = 0){
    for (let i=index; i< searchElement.length; i++) {
        if (searchElement[i] === target){
            return i
        }
    }
    return -1
}

function lastIndexOf(searchElement, target, index = 0){
    if (index != 0) {
        index
        } else {
            index = searchElement.length - 1 
        }
            for (let i = index; i >= 0; i--) {
                if (searchElement[i] === target){
                    return i
            }
        }
    return -1
}

function includes(searchElement, target){
    for (let i= searchElement.length-1; i>=0; i--) {
        if (searchElement[i] === target){
            return true
        }
    }
    return false
}