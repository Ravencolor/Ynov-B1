import java.util.*;

public class WeddingComplex {
    public static Map<String, String> createBestCouple(Map<String, List<String>> preferences1,
                                                        Map<String, List<String>> preferences2) {
        Map<String, String> couples = new HashMap<>();
        Map<String, Integer> preferencesIndex2 = new HashMap<>();
        for (String name : preferences2.keySet()) {
            preferencesIndex2.put(name, 0);
        }
        List<String> freeMen = new ArrayList<>(preferences1.keySet());
        while (!freeMen.isEmpty()) {
            String man = freeMen.get(0);
            List<String> manPreferences = preferences1.get(man);
            Integer index = preferencesIndex2.getOrDefault(man, 0);
            String woman = manPreferences.get(index);
            preferencesIndex2.put(man, index + 1);
            if (!couples.containsKey(woman)) {
                couples.put(woman, man);
                freeMen.remove(0);
            } else {
                String currentPartner = couples.get(woman);
                List<String> womanPreferences = preferences2.get(woman);
                if (womanPreferences.indexOf(man) < womanPreferences.indexOf(currentPartner)) {
                    couples.put(woman, man);
                    freeMen.remove(0);
                    freeMen.add(currentPartner);
                }
            }
        }
        Map<String, String> revertedCouples = new HashMap<>();
        for (Map.Entry<String, String> entry : couples.entrySet()) {
            revertedCouples.put(entry.getValue(), entry.getKey());
        }
        return revertedCouples;
    }
}
