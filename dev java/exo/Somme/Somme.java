import java.util.Scanner;

public class Somme {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        
        System.out.print("Entrez un entier : ");
        int number = scanner.nextInt();
        
        int sum = 0;
        int temp = number;
        
        while (temp != 0) {
            int digit = temp % 10;
            sum += digit;
            temp /= 10;
        }
        
        System.out.println("La somme des chiffres de " + number + " est : " + sum);
        
        scanner.close();
    }
}