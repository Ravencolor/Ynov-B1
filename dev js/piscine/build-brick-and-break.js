export function build(num){
    num++
    let counter = 1, start = 1, middle = false;
    let i = setInterval(function () {
        if (start === 4) start = 1;
        const newDiv = document.createElement('div');
        newDiv.id = 'brick-' + counter;
        if (start === 2) {
            newDiv.dataset.foundation = 'true'
        }
        document.body.append(newDiv);
        counter++
        start++
        if (counter === num) clearInterval(i)
    },100)
}


export function repair(...ids){
    for (let i = 0; i < ids.length; i++) {
        const brick = document.getElementById(ids[i]);
        if (brick === null) continue;
        if (brick.hasAttribute('data-foundation')) {
            brick.dataset.repaired = 'in progress';
        } else {
            brick.dataset.repaired = 'true';
        }
    }
}


export function destroy(){
    const bricks = document.querySelectorAll('div')
    console.log(bricks[bricks.length-1])
    bricks[bricks.length-1].remove()
}