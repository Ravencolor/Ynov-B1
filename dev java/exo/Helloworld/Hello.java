import java.util.random.RandomGenerator;

class Hello {
    public static void main(String[] args) {
        System.out.println("Hello World");

        byte b = 7;
        short s = 2;
        int i = 10;
        long l = 100;

        double d = 10.85252;
        float f = 7.41745252f;
        boolean bool = true;

        System.err.println(i);
        System.err.println(f);

        Integer i2 = 10;

        String message = "Hello";
        System.out.println(message);
        System.out.println(message.length());

        RandomGenerator r = RandomGenerator.getDefault();
        int i3 = r.nextInt(-10, 10);
        System.out.println(i3);
        double d3 = r.nextDouble();
        if (i3 > 0){
            System.out.println("positif");
        } else if (i3 < 0) {
            System.out.println("négatif");
        } else {
            System.out.println("zéro");
        }
        int[] numbers = new int[]{-8, 10, 3, 4};
        for (int j = 0; j < numbers.length; j++) {
            System.out.print(numbers[j]+ "/");
        }

    }
}