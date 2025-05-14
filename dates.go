package xero

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type NetDate struct {
	Time time.Time
}

func (s *NetDate) UnmarshalJSON(data []byte) error {
	// The regular expression supports formats like /Date(-1234567890000+0000)/ or /Date(1234567890000+0000)/
	// Capture only milliseconds, ignore offset
	re := regexp.MustCompile(`\\/Date\((-?\d+)(?:\+\d+)?\)\\/`)
	match := re.FindStringSubmatch(string(data))

	// match[0] - the whole string, match[1] - milliseconds
	if len(match) != 2 {
		return fmt.Errorf("invalid date format: %v", string(data))
	}

	// this conversion is only for Xero, which returns time in UTC
	// so we don't care about second part of the date e.g. +0500
	// thus using only the first matching group

	milliseconds, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return fmt.Errorf("error parsing milliseconds %v: %w", match[1], err)
	}

	// Convert to time.Time, truncate to seconds and force to UTC
	s.Time = time.Unix(int64(milliseconds/1000), 0).UTC()

	return nil
}

func (s *NetDate) MarshalJSON() ([]byte, error) {
	if s == nil {
		return []byte(""), nil
	}

	val := fmt.Sprintf("\"/Date(%v000+0000)/\"", s.Time.Unix())

	return []byte(val), nil
}
