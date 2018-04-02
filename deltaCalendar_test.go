package deltaCalendar

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

const invalidDuration = "invalid duration"

// TestDelta30d adds 30 days to 2018-04-01 which is 2018-04-01.
// TODO test a matrix of before-after dates
func TestDelta30d(t *testing.T) {
	before := time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC)
	after := time.Date(2018, time.May, 1, 0, 0, 0, 0, time.UTC)
	result, err := DeltaDate(before,30)
	if err != nil {
		t.Errorf("error computing delta: %v", err)
	}
	if result != after {
		t.Errorf(fmt.Sprintf("%v does not equal %v", result, after))
	}
}

// TestDelta300y attempts to add 300 years to a date.
// The maximum time.Duration is ~290 years, so this must fail.
func TestDelta300y(t *testing.T) {
	before := time.Date(2018, time.April, 1, 0, 0, 0, 0, time.UTC)
	_, err := DeltaDate(before,300 * 365)
	if err == nil {
		t.Errorf("DeltaDate should have failed")
	}
	if !strings.Contains(err.Error(), invalidDuration) {
		t.Errorf("'%v' not found in '%v'", invalidDuration, err)
	}
}
