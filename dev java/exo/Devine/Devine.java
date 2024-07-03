import java.util.Scanner;

public class Devine {
    public static void main(String[] args) {
        int nombreAleatoire = (int) (Math.random() * 20) + 1;
        int nombreEssais = 0;
        boolean trouve = false;

        Scanner scanner = new Scanner(System.in);

        while (!trouve) {
            System.out.print("Devinez un nombre entre 1 et 20 : ");
            int proposition = scanner.nextInt();
            nombreEssais++;

            if (proposition < nombreAleatoire) {
                System.out.println("Le nombre choisi est plus grand.");
            } else if (proposition > nombreAleatoire) {
                System.out.println("Le nombre choisi est plus petit.");
            } else {
                trouve = true;
                System.out.println("Bravo, vous avez trouvé le nombre en " + nombreEssais + " essais !");
            }
        }

        scanner.close();
    }
}
