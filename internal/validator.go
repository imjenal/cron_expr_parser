package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// ValidateCronExpression validates a cron expression
func ValidateCronExpression(cronExpression string) error {
	fields := strings.Fields(cronExpression)
	if len(fields) != 6 {
		return fmt.Errorf("Invalid cron expression: It should have exactly 6 fields")
	}

	fieldNames := []string{"minute", "hour", "day of month", "month", "day of week", "command"}

	for i, field := range fields {
		if err := validateField(field, fieldNames[i]); err != nil {
			return err
		}
	}

	return nil
}

func validateField(field, fieldName string) error {
	// Validate the field based on its position (fieldName)
	switch fieldName {
	case "minute":
		if err := validateMinuteField(field); err != nil {
			return err
		}
	case "hour":
		if err := validateHourField(field); err != nil {
			return err
		}
	case "day of month":
		if err := validateDayOfMonthField(field); err != nil {
			return err
		}
	case "month":
		if err := validateMonthField(field); err != nil {
			return err
		}
	case "day of week":
		if err := validateDayOfWeekField(field); err != nil {
			return err
		}
	case "command":
		// Implement validation logic for the command field (if needed)
	}

	return nil
}

func validateMinuteField(field string) error {
	if field == "*" || strings.HasPrefix(field, "*/") {
		return nil // "*" and "*/" formats are valid
	}

	// Check if the field is a valid number between 0 and 59
	minute, err := strconv.Atoi(field)
	if err != nil || minute < 0 || minute > 59 {
		return fmt.Errorf("Invalid minute field value: %s", field)
	}

	return nil
}

func validateHourField(field string) error {
	// Check if the field is a valid number between 0 and 23
	hour, err := strconv.Atoi(field)
	if err != nil || hour < 0 || hour > 23 {
		return fmt.Errorf("Invalid hour field value: %s", field)
	}
	return nil
}

func validateDayOfMonthField(field string) error {
	if field == "*" {
		return nil // "*" is a valid value
	}

	// Check if the field contains comma-separated values
	if strings.Contains(field, ",") {
		values := strings.Split(field, ",")
		for _, value := range values {
			dayOfMonth, err := strconv.Atoi(value)
			if err != nil || dayOfMonth < 1 || dayOfMonth > 31 {
				return fmt.Errorf("Invalid day of month field value: %s", value)
			}
		}
		return nil
	}

	// Check if the field is a valid number between 1 and 31
	dayOfMonth, err := strconv.Atoi(field)
	if err != nil || dayOfMonth < 1 || dayOfMonth > 31 {
		return fmt.Errorf("Invalid day of month field value: %s", field)
	}

	return nil
}

func validateMonthField(field string) error {
	if field == "*" {
		return nil // "*" is a valid value
	}

	// Check if the field is a valid number between 1 and 12
	month, err := strconv.Atoi(field)
	if err != nil || month < 1 || month > 12 {
		return fmt.Errorf("Invalid month field value: %s", field)
	}
	return nil
}

func validateDayOfWeekField(field string) error {
	if field == "*" {
		return nil // "*" is a valid value
	}

	// Check if the field contains a range (e.g., 1-4)
	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		start, err := strconv.Atoi(rangeParts[0])
		end, err2 := strconv.Atoi(rangeParts[1])
		if err != nil || err2 != nil || start < 0 || end > 6 || start > end {
			return fmt.Errorf("Invalid day of week field value: %s", field)
		}
		return nil
	}

	// Check if the field is a valid number between 0 and 6 (Sunday to Saturday)
	dayOfWeek, err := strconv.Atoi(field)
	if err != nil || dayOfWeek < 0 || dayOfWeek > 6 {
		return fmt.Errorf("Invalid day of week field value: %s", field)
	}

	return nil
}
