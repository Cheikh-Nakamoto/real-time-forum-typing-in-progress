package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseArrayInt(a []string) (tab []int, err error) {
	for _, v := range a {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		tab = append(tab, r)
	}
	return tab, nil
}

func FormatDuration(duration time.Duration) string {
	// Get the individual components of the duration
	days := duration / (24 * time.Hour)
	duration = duration % (24 * time.Hour)
	hours := duration / time.Hour
	duration = duration % time.Hour
	minutes := duration / time.Minute
	duration = duration % time.Minute
	seconds := duration / time.Second

	// Build the formatted string
	var parts []string
	if days > 0 {
		parts = append(parts, fmt.Sprintf("%d day", days))
	}
	if hours > 0 {
		parts = append(parts, fmt.Sprintf("%d hour", hours))
	}
	if minutes > 0 {
		parts = append(parts, fmt.Sprintf("%d mn", minutes))
	}
	if seconds > 0 {
		parts = append(parts, fmt.Sprintf("%d s", seconds))
	}

	// Join the parts with ", " to create the final string
	return strings.Join(parts, ", ")
}

func IsOrderParam(orderBy string) bool {
	var orderParams = []string{
		"TIME-ASC",
		"TIME-DESC",
		"MOSTLIKED-ASC",
		"MOSTLIKED-DESC",
	}

	for _, v := range orderParams {
		if v == orderBy {
			return true
		}
	}
	return false
}

func IsAlphanumeric(input string) bool {
	pattern := "^[a-zA-Z0-9]*$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func IsAlpha(input string) bool {
	pattern := "^[a-zA-Z]*$"
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(input)
}

func VerifyUsername(username string) error {
	pattern := `^[a-zA-Z][a-zA-Z0-9_]{6,15}$`
	regex := regexp.MustCompile(pattern)
	ok := regex.MatchString(username)
	if !ok {
		return fmt.Errorf("%v : not a valid username", username)
	}
	return nil
}

func VerifyName(s string) error {
	var maxChars = 25
	if len(s) > maxChars || !IsAlpha(s) {
		return fmt.Errorf("%v: invalid name", s)
	}
	return nil
}

func IsValidEmail(email string) error {
	emailRegex := `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`

	re := regexp.MustCompile(emailRegex)
	ok := re.MatchString(email)
	if !ok {
		return fmt.Errorf("%v: invalid email", email)

	}
	return nil
}
