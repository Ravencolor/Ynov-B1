import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Random;
import java.util.Set;

public class Wedding {
    public static Map<String, String> createCouple(Set<String> set1, Set<String> set2) {
        List<String> list1 = new ArrayList<>(set1);
        List<String> list2 = new ArrayList<>(set2);
        Map<String, String> result = new HashMap<>();
        List<String> names1 = new ArrayList<>(list1);
        List<String> names2 = new ArrayList<>(list2);
        Random rand = new Random();
        int limit = Math.min(names1.size(), names2.size());
        for (int i = 0; i < limit; i++) {
            int index1 = rand.nextInt(names1.size());
            int index2 = rand.nextInt(names2.size());
            result.put(names1.get(index1), names2.get(index2));
            names1.remove(index1);
            names2.remove(index2);
        }
        return result;
    }
}
