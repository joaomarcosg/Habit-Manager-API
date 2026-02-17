package domain

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (w *WeekDay) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("weekday must be a string: %w", err)
	}

	s = strings.ToLower(s)

	switch s {
	case "monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday":
		*w = WeekDay(s)
		return nil
	default:
		return fmt.Errorf("invalid weekday: %s", s)
	}
}
