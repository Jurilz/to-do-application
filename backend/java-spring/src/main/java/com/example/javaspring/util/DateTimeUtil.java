package com.example.javaspring.util;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class DateTimeUtil {

    public static String getCurrentDate() {
        DateTimeFormatter formatter = DateTimeFormatter.ISO_DATE;
        LocalDateTime now = LocalDateTime.now();
        return formatter.format(now);
    }
}
