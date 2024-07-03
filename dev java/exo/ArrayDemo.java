///usr/bin/env Jbang "$0" "$@" ; exit $?

import java.util.random.RandomGenerator;

public class ArrayDemo {
    public static void main(String[] args) {

        System.out.println("Hello World");

        int[] numbers = new int[] {1, 2, 3, 4, 5};
        for (int i = 0; i < numbers.length; i++) {
            System.out.print(i);
            System.out.println("->" + numbers[i]);

        }
        ///Tablea uqui permet de stocker 5 elements et qui est initialement vide
        int[] othernumbers = new int[5];
        othernumbers[0] = 20;
        RandomGenerator random = RandomGenerator.getDefault();
        System.out.println("othernumbers ->");
        for (int i = 0; i < othernumbers.length; i++) {
            othernumbers[i] = random.nextInt(-5, 5);
            System.out.print(i);
            System.out.println("->" + othernumbers[i]);

        }

        ///Somme des elements de numbers
        int sum = 0;
        for (int i = 0; i < numbers.length; i++) {
            sum += numbers[i];
        }
        System.out.println("sum = " + sum);

        /// valeur max dans othernumbers
        int max = othernumbers[0];
        for (int i = 1; i < othernumbers.length; i++) {
            if (othernumbers[i] > max) {
                max = othernumbers[i];
            }
        }
        System.out.println("max = " + max);
    }

}