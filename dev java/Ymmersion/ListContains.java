import java.util.List;

public class ListContains {
    public static boolean containsValue(List<Integer> list, Integer value) {
        return list == null || value == null ? false : list.contains(value);
    }
}