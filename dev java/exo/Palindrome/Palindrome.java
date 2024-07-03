public class Palindrome {
    public static boolean isPalindrome(String str) {
        String reversed = new StringBuilder(str).reverse().toString();
        return str.equals(reversed);
    }

    public static void main(String[] args) {
        String word = "radar";
        if (isPalindrome(word)) {
            System.out.println(word + " est un palindrome.");
        } else {
            System.out.println(word + " n'est pas un palindrome.");
        }
    }
}
