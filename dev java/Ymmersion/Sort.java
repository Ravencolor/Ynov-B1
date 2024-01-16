import java.util.List;
import java.util.Objects;
import java.util.Comparator;

public class Sort {

    public static List<Integer> sort(List<Integer> list) {
        if(Objects.isNull(list)) return null;
        return list.stream().sorted((Comparator.naturalOrder())).toList();
    }

    public static List<Integer> sortReverse(List<Integer> list) {
        if(Objects.isNull(list)) return null;
        return list.stream().sorted((Comparator.reverseOrder())).toList();
    }
}