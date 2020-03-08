# Handling DateTime in programming

As a general rule of thumb, DateTimes in programming should be in the [ISO-8601 format](https://en.wikipedia.org/wiki/ISO_8601). It formats datetimes from most significant detail to least significant.
For example, it is a String that might look like:
```
YYYY-MM-DD HH:MM:SS.<milliseconds>
```
Hours are in a 24-format and have max two "digits" for 10 or 12 and can be "01" to represent 1.
Many databases, like MySQL, support storing DateTime information in this format and can even support operations like sorting and converting between different date/time formats.

## What are the other date/time formats?
You can choose to store DateTime as just `Date` (no time), or `Time` (no date), or `TIMESTAMP` (similar to a long which is the number of milliseconds from 1/1/1970 - see epoch time). Some of these formats have different pros and cons - usually it is related to trade-offs between accuracy in precision vs. store space. For example, maybe we could use `Date` instead of `DateTime` because we don't care about the time of that date and it would take up less space.

## What is epoch time?
The Unix epoch (or Unix timestamp) is the number of milliseconds since January 1, 1970 (midnight 00:00).
It starts at this date because that is when the first Unix computer was created.
More precisely, the epoch time starts at 0 to represent January 1st, 1970 UTC.
UTC is a timezone which stands for [Coordinated Universal Time](https://en.wikipedia.org/wiki/Coordinated_Universal_Time.

Computers often like to store in epoch time so that it doesn't get mixed up by timezones or complicated date-time formats. For instance, certain strings like `12/30/1990 01:00pm EST` can be very painful to parse.
It might be easier to use a numeric value and convert it to an appropriate String. It would also be more precise.

You can use this to manually convert between epoch times and "regular human" times.
https://www.epochconverter.com/

Also note, a long has a limit in how big the number it can get and so longs and `TIMESTAMP` actually do not support dates after the year 2038.

## Using DateTime in Go
Golang offers a super easy conversion/support using the `time` and `database/sql` libraries.
In Go, to grab a DateTime of the current time, use:
```
time.Now().UTC()
```
This will grab a `time.Time` struct that represents the current moment and automatically convert it into the UTC timezone. In addition, there are many methods to convert this struct into DateTime or TIMESTAMP formats and even do operations like adding Days or Weeks. This is very useful when you need to perform scheduling operations.
[https://golang.org/pkg/time/](https://golang.org/pkg/time/)