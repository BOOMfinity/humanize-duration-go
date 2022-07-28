package humanize_duration_go

import (
	"github.com/boomfinity/humanize-duration-go/internal/i18n"
	"golang.org/x/text/language"
	"math"
	"strings"
	"time"
)

func HumanizeDuration(duration time.Duration, lang language.Tag) string {
	loc, ok := i18n.Locales[lang]
	if !ok {
		loc = i18n.LocaleEnglish
	}

	days := int64(duration.Hours() / 24)
	hours := int64(math.Mod(duration.Hours(), 24))
	minutes := int64(math.Mod(duration.Minutes(), 60))
	seconds := int64(math.Mod(duration.Seconds(), 60))

	chunks := []struct {
		fun    func(int64) string
		amount int64
	}{
		{loc.Days, days},
		{loc.Hours, hours},
		{loc.Minutes, minutes},
		{loc.Seconds, seconds},
	}

	var parts []string
	for _, chunk := range chunks {
		if chunk.amount == 0 {
			continue
		}
		parts = append(parts, chunk.fun(chunk.amount))
	}

	return strings.Join(parts, " ")
}

// https://gist.github.com/harshavardhana/327e0577c4fed9211f65?permalink_comment_id=2366908#gistcomment-2366908
