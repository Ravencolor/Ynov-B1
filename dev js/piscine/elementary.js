function multiply (a, b) {
    let result = 0
    if (b > 0) {
        for (let i = 0; i < b; i++) {
            result += a
        } 
    } else {
        for (let i = b; i < 0; i++) {
            result -= a
        }
    }
    return result
}
console.log(multiply(123, -22))

function divide(a, b) {
    let result = 0;
    if (b > 0) {
      if (a > 0) {
        while (a - b > 0) {
          result++;
          a -= b;
          console.log(a);
        }
      } else {
        while (a + b < 0) {
          result--;
          a += b;
        }
      }
    } else {
      if (a > 0) {
        while (a + b > 0) {
          a += b;
          result--;
        }
      } else {
        while (a - b < 0) {
          a -= b;
          result++;
        }
      }
    }
    return result;
  }
console.log(divide(10, 2))

function modulo(a, b) {
    return a - multiply(divide(a, b), b);
  }