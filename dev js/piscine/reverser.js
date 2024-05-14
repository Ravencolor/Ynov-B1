function reverse(input) {
    let reversed = [];
    for (let i = input.length - 1; i >= 0; i--) {
        reversed.push(input[i]);
    }
    return Array.isArray(input) ? reversed : reversed.join('');
}