import java.util.Random;

public class RandomArrayStats {
    public static void main(String[] args) {
        int[] array = new int[10];
        Random rand = new Random();

        // Génération du tableau aléatoire
        for (int i = 0; i < array.length; i++) {
            array[i] = rand.nextInt(21) - 10; // Génère des entiers entre -10 et 10
        }

        // Afficher le tableau généré
        System.out.println("Tableau généré:");
        for (int value : array) {
            System.out.print(value + " ");
        }
        System.out.println();

        // Initialisation des variables pour les statistiques
        int max = Integer.MIN_VALUE;
        int min = Integer.MAX_VALUE;
        int sum = 0;
        int positiveCount = 0;
        int negativeCount = 0;
        int evenCount = 0;
        int oddCount = 0;

        // Calcul des statistiques
        for (int value : array) {
            if (value > max) {
                max = value;
            }
            if (value < min) {
                min = value;
            }
            sum += value;

            if (value > 0) {
                positiveCount++;
            } else if (value < 0) {
                negativeCount++;
            }

            if (value % 2 == 0) {
                evenCount++;
            } else {
                oddCount++;
            }
        }

        // Calcul de la moyenne
        double mean = (double) sum / array.length;

        // Calcul de l'écart type
        double varianceSum = 0;
        for (int value : array) {
            varianceSum += Math.pow(value - mean, 2);
        }
        double standardDeviation = Math.sqrt(varianceSum / array.length);

        // Affichage des résultats
        System.out.println("La plus grande valeur : " + max);
        System.out.println("La plus petite valeur : " + min);
        System.out.println("La moyenne : " + mean);
        System.out.println("L'écart type : " + standardDeviation);
        System.out.println("Le nombre de valeurs positives : " + positiveCount);
        System.out.println("Le nombre de valeurs négatives : " + negativeCount);
        System.out.println("Le nombre de valeurs paires : " + evenCount);
        System.out.println("Le nombre de valeurs impaires : " + oddCount);
    }
}
