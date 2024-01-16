import java.util.ArrayList;
import java.util.List;

public class ListSearchIndex {
    public static Integer findLastIndex(List<Integer> list, Integer value) {
        if (list == null) return null;
        if (value == null) return null;
        return list.lastIndexOf(value) == -1 ? null : list.lastIndexOf(value);
    }
    public static Integer findFirstIndex(List<Integer> list, Integer value) {
        if (list == null) return null;
        if (value == null) return null;
        return list.indexOf(value) == -1 ? null : list.indexOf(value);
    }
    public static List<Integer> findAllIndexes(List<Integer> list, Integer value) {
        List<Integer> indexes = new ArrayList<>();
        if (list == null) {
            return indexes;
        }
        for (int i = 0; i < list.size(); ++i) {
            if (list.get(i).equals(value)) {
                indexes.add(i);
            }
        }
        return indexes;
    }
}