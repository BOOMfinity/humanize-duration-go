package i18n

import "strconv"

var LocalePolish = Locale{
	Days: func(i int64) string {
		if i == 1 {
			return "1 dzień"
		}
		return strconv.FormatInt(i, 10) + " dni"
	},

	Hours: func(i int64) string {
		if i == 1 {
			return "1 godzina"
		} else if i < 5 {
			return strconv.FormatInt(i, 10) + " godziny"
		}
		return strconv.FormatInt(i, 10) + " godzin"
	},

	Minutes: func(i int64) string {
		if i == 1 {
			return "1 minuta"
		} else if i < 5 {
			return strconv.FormatInt(i, 10) + " minuty"
		}
		return strconv.FormatInt(i, 10) + " minut"
	},

	Seconds: func(i int64) string {
		if i == 1 {
			return "1 sekunda"
		} else if i < 5 {
			return strconv.FormatInt(i, 10) + " sekundy"
		}
		return strconv.FormatInt(i, 10) + " sekund"
	},
}
