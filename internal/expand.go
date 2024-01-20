package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// ExpandCronExpression expands a cron expression
func ExpandCronExpression(cronExpression string) ([]string, [][]int, string, error) {
	fields := strings.Fields(cronExpression)

	if len(fields) != 6 {
		return nil, nil, "", fmt.Errorf("Invalid cron expression. It should have exactly 6 fields.")
	}

	minuteField := fields[0]
	hourField := fields[1]
	dayOfMonthField := fields[2]
	monthField := fields[3]
	dayOfWeekField := fields[4]
	command := fields[5]

	minuteValues, err := expandField(minuteField, 0, 59, "minute")
	if err != nil {
		return nil, nil, "", err
	}
	hourValues, err := expandField(hourField, 0, 23, "hour")
	if err != nil {
		return nil, nil, "", err
	}
	dayOfMonthValues, err := expandField(dayOfMonthField, 1, 31, "dayofmonth")
	if err != nil {
		return nil, nil, "", err
	}
	monthValues, err := expandField(monthField, 1, 12, "month")
	if err != nil {
		return nil, nil, "", err
	}
	dayOfWeekValues, err := expandField(dayOfWeekField, 0, 6, "dayofweek")
	if err != nil {
		return nil, nil, "", err
	}

	fieldNames := []string{"minute", "hour", "day of month", "month", "day of week", "command"}
	fieldValues := [][]int{minuteValues, hourValues, dayOfMonthValues, monthValues, dayOfWeekValues}
	fieldValues = append(fieldValues, []int{}) // Add an empty slice for the command
	fieldValues[5] = append(fieldValues[5], 0) // Append a value  0) for the command

	return fieldNames, fieldValues, command, nil
}

func expandField(field string, minValue, maxValue int, name string) ([]int, error) {
	var result []int

	if field == "*" {
		for i := minValue; i <= maxValue; i++ {
			result = append(result, i)
		}
		return result, nil
	}

	if strings.Contains(field, "/") {
		parts := strings.Split(field, "/")
		step, err := strconv.Atoi(parts[1])
		if err != nil || step <= 0 {
			return nil, fmt.Errorf("Invalid step value in %s field: %s", name, field)
		}
		field = parts[0]

		if field == "*" {
			// Handle cases like */15
			for i := minValue; i <= maxValue; i += step {
				result = append(result, i)
			}
			return result, nil
		}
	}

	if strings.Contains(field, ",") {
		values := strings.Split(field, ",")
		for _, value := range values {
			if value == "" {
				continue // Skip empty values
			}
			v, err := strconv.Atoi(value)
			if err != nil || v < minValue || v > maxValue {
				return nil, fmt.Errorf("Invalid %s field value: %s", name, value)
			}
			result = append(result, v)
		}
		return result, nil
	}

	if strings.Contains(field, "-") {
		rangeParts := strings.Split(field, "-")
		if len(rangeParts) != 2 {
			return nil, fmt.Errorf("%v Invalid range: %s", name, field)
		}
		start, err := strconv.Atoi(rangeParts[0])
		end, err2 := strconv.Atoi(rangeParts[1])
		if err != nil || err2 != nil || start < minValue || end > maxValue || start > end {
			return nil, fmt.Errorf("%v Invalid range: %s", name, field)
		}
		for i := start; i <= end; i++ {
			result = append(result, i)
		}
		return result, nil
	}

	// Handle single values like "5"
	value, err := strconv.Atoi(field)
	if err != nil || value < minValue || value > maxValue {
		return nil, fmt.Errorf("Invalid %s field value: %s", name, field)
	}
	return []int{value}, nil
}
