import java.util.Scanner;

public class Premier {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Entrez un entier : ");
        int number = scanner.nextInt();
        boolean siEntier = true;

        if (number <= 1) {
            siEntier = false;
        } else {
            for (int i = 2; i <= Math.sqrt(number); i++) {
                if (number % i == 0) {
                    siEntier = false;
                    break;
                }
            }
        }

        if (siEntier) {
            System.out.println(number + " est un nombre premier.");
        } else {
            System.out.println(number + " n'est pas un nombre premier.");
        }
    }
}