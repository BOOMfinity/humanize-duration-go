package humanize_duration_go

import (
	"fmt"
	"golang.org/x/text/language"
	"testing"
	"time"
)

var engtest = map[time.Duration]string{
	time.Hour * 10:  "10 hours",
	time.Hour:       "1 hour",
	time.Minute:     "1 minute",
	5 * time.Minute: "5 minutes",
	time.Second:     "1 second",
	2 * time.Second: "2 seconds",

	10*time.Hour + 30*time.Minute:                  "10 hours 30 minutes",
	24 * time.Hour:                                 "1 day",
	30*time.Hour + 10*time.Minute + 30*time.Second: "1 day 6 hours 10 minutes 30 seconds",
}

func TestDurationsEnglish(t *testing.T) {
	for duration, s := range engtest {
		h := HumanizeDuration(duration, language.English)
		if s != h {
			t.Fatal(fmt.Sprintf("Expected %v, got %v\n", s, h))
		}
	}
}

var pltest = map[time.Duration]string{
	time.Hour * 10:  "10 godzin",
	time.Hour:       "1 godzina",
	time.Minute:     "1 minuta",
	5 * time.Minute: "5 minut",
	time.Second:     "1 sekunda",
	2 * time.Second: "2 sekundy",

	10*time.Hour + 30*time.Minute:                  "10 godzin 30 minut",
	24 * time.Hour:                                 "1 dzień",
	30*time.Hour + 10*time.Minute + 30*time.Second: "1 dzień 6 godzin 10 minut 30 sekund",
}

func TestDurationsPolish(t *testing.T) {
	for duration, s := range pltest {
		h := HumanizeDuration(duration, language.Polish)
		if s != h {
			t.Fatal(fmt.Sprintf("Expected %v, got %v\n", s, h))
		}
	}
}
