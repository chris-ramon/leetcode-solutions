// **Complexity Analysis**
// * Time: O(1), constant time as input grows.
// * Space: O(1), constant number of variables used.

var Months = map[string]string{
	"Jan": "01",
	"Feb": "02",
	"Mar": "03",
	"Apr": "04",
	"May": "05",
	"Jun": "06",
	"Jul": "07",
	"Aug": "08",
	"Sep": "09",
	"Oct": "10",
	"Nov": "11",
	"Dec": "12",
}

func reformatDate(date string) string {
	dateParts := strings.Split(date, " ")
	fullDay, rawMonth, year := dateParts[0], dateParts[1], dateParts[2]
	day := ""
	if len(fullDay) == 3 {
		day = "0" + string(fullDay[0])
	} else {
		day = string([]byte{fullDay[0], fullDay[1]})
	}
	return fmt.Sprintf("%s-%s-%s", year, Months[rawMonth], day)
}
