import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Set;

public class KeepTheChange {
    public static List<Integer> computeChange(int amount, Set<Integer> existingCoins) {
        int[] coins = existingCoins.stream().mapToInt(i -> i).toArray();
        Arrays.sort(coins);
        List<Integer> result = new ArrayList<>();
        for (int i = coins.length - 1; i >= 0; i--) {
            while (amount >= coins[i]) {
                result.add(coins[i]);
                amount -= coins[i];
            }
        }
        return result;
    }
}
