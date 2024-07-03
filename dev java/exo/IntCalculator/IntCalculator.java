import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Scanner;

public class IntCalculator {
    private List<String> history;
    private Map<String, Integer> variables;

    public IntCalculator() {
        history = new ArrayList<>();
        variables = new HashMap<>();
    }

    public void run() {
        Scanner scanner = new Scanner(System.in);
        String input;

        do {
            System.out.print("Enter an operation or command: ");
            input = scanner.nextLine();

            try {
                if (input.matches("\\d+")) {
                    int value = Integer.parseInt(input);
                    System.out.println("Value: " + value);
                    history.add("Value: " + value);
                } else if (input.matches("\\d+ [\\+\\-\\*/] \\d+")) {
                    String[] parts = input.split(" ");
                    int operand1 = Integer.parseInt(parts[0]);
                    int operand2 = Integer.parseInt(parts[2]);
                    char operator = parts[1].charAt(0);

                    int result = performOperation(operand1, operator, operand2);
                    System.out.println("Result: " + result);
                    history.add("Result: " + result);
                } else if (input.equals("h")) {
                    printHistory();
                } else if (input.matches("[\\+\\-\\*/]")) {
                    printHistoryForOperator(input.charAt(0));
                } else if (input.equals("exit")) {
                    System.out.println("Exiting...");
                } else {
                    System.out.println("Invalid input. Please try again.");
                }
            } catch (NumberFormatException e) {
                System.out.println("Invalid number format. Please try again.");
            } catch (ArithmeticException e) {
                System.out.println("Arithmetic error. Please try again.");
            }
        } while (!input.equals("exit"));
    }

    private int performOperation(int operand1, char operator, int operand2) {
        int result;

        switch (operator) {
            case '+':
                result = operand1 + operand2;
                break;
            case '-':
                result = operand1 - operand2;
                break;
            case '*':
                result = operand1 * operand2;
                break;
            case '/':
                result = operand1 / operand2;
                break;
            default:
                throw new IllegalArgumentException("Invalid operator: " + operator);
        }

        return result;
    }

    private void printHistory() {
        System.out.println("History:");
        for (String entry : history) {
            System.out.println(entry);
        }
    }

    private void printHistoryForOperator(char operator) {
        System.out.println("History for operator " + operator + ":");
        for (String entry : history) {
            if (entry.contains(" " + operator + " ")) {
                System.out.println(entry);
            }
        }
    }

    public static void main(String[] args) {
        IntCalculator calculator = new IntCalculator();
        calculator.run();
    }
}