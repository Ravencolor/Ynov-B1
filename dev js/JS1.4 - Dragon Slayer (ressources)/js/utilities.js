
/**
 * Tire un entier au hasard et le retourne
 * @param {number} min
 * @param {number} max
 * @returns {number}
 */
function getRandomInteger(min, max)
{
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

/**
 * Simule un lancé de dés (throw = jeter / dices = dés )
 * @param {number} dices - Nombre de dés que l'on souhaite lancer
 * @param {number} sides - Nombre de faces par dé
 * @returns {number} - Le total de la somme des dés
 */
function throwDices(dices, sides)
{
    // Déclaration des variables locales à la fonction throwDices()
    var index;
    var sum;

    // Initialisation de la somme des dés à 0
    sum = 0;

    /*
    dices = nombre de jets de dés
    Pour chaque jet de dé, le tirage renvoie un nombre entier compris entre 1 et le nombre de faces du dé
    */
    // Pour chaque jet de dé : un tirage aléatoire compris dans [1,sides]
    for(index = 0 ; index < dices ; index++){
        // ...on le lance et on ajoute sa valeur à la somme totale
        sum += getRandomInteger(1, sides);
    }

    // Retour en résultat de la somme totale des dés
    return sum;
}
