import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Locale;

public class ParseDate {
    public static LocalDateTime parseIsoFormat(String stringDate) {
        if (stringDate == null) return null;
        return LocalDateTime.parse(stringDate);
    }
    public static LocalDate parseFullTextFormat(String stringDate) {   
        if (stringDate == null) return null;
        return LocalDate.parse(stringDate, java.time.format.DateTimeFormatter.ofPattern("EEEE dd MMMM yyyy", Locale.FRANCE));
    }
    public static LocalTime parseTimeFormat(String stringDate) {
        if (stringDate == null) return null;
        stringDate = stringDate.replace("soir", "PM");
        stringDate = stringDate.replace("matin", "AM");
        return LocalTime.parse(stringDate, java.time.format.DateTimeFormatter.ofPattern("hh 'heures du' a',' mm 'minutes et' ss 'secondes'"));
    }
}