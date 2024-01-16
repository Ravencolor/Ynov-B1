import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Locale;

public class FormatDate {

    public static String formatToFullText(LocalDateTime dateTime) {
        if (dateTime == null) return null;
        return dateTime.format(java.time.format.DateTimeFormatter.ofPattern("'Le' dd MMM 'de' 'l'\'an 'yyyy' à 'HH'h'mm'm et' ss's'", Locale.FRANCE));
    }

    public static String formatSimple(LocalDate date) {
        return date == null ? null : date.format(java.time.format.DateTimeFormatter.ofPattern("MMMM dd yy", Locale.ITALIAN));
    }

    public static String formatIso(LocalTime time) {
        if (time == null) return null;
        if (time.getNano() == 0) return time.format(java.time.format.DateTimeFormatter.ofPattern("HH:mm:ss"));
        return time.format(java.time.format.DateTimeFormatter.ofPattern("HH:mm:ss.SSSSSSSSS"));
    }

}