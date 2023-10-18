package datetime

import (
	"fmt"
	"strconv"
	"time"
)

func FloatToInt(f float64) (int, error) {
	t := int(f)
	s := fmt.Sprintf("%.0f", f)
	i, err := strconv.Atoi(s)
	if err != nil {
		return t, err
	}
	return i, nil
}

func GetMinuteDuration(expiredAtUtc time.Time) (int, error) {
	duration := time.Since(expiredAtUtc)
	minuteDuration := duration.Minutes() * -1
	minuteDurationInt, err := FloatToInt(minuteDuration)
	if err != nil {
		return minuteDurationInt, err
	}
	if minuteDurationInt < 0 {
		minuteDurationInt = 0
	}
	return minuteDurationInt, nil
}
