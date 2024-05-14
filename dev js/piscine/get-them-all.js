// getArchitects function
export const getArchitects = () => {
    const architects = Array.from(document.getElementsByTagName('a'));
    const nonArchitects = Array.from(document.querySelectorAll(':not(a)'));
    return [architects, nonArchitects];
}

// getClassical function
export const getClassical = () => {
    const classicalArchitects = Array.from(document.getElementsByClassName('classical'));
    const nonClassicalArchitects = Array.from(document.querySelectorAll('a:not(.classical)'));
    return [classicalArchitects, nonClassicalArchitects];
}

// getActive function
export const getActive = () => {
    const activeClassicalArchitects = Array.from(document.querySelectorAll('.classical.active'));
    const nonActiveClassicalArchitects = Array.from(document.querySelectorAll('.classical:not(.active)'));
    return [activeClassicalArchitects, nonActiveClassicalArchitects];
}

// getBonannoPisano function
export const getBonannoPisano = () => {
    const bonannoPisano = document.getElementById('BonannoPisano');
    const remainingActiveClassicalArchitects = Array.from(document.querySelectorAll('.classical.active:not(#BonannoPisano)'));
    return [bonannoPisano, remainingActiveClassicalArchitects];
}