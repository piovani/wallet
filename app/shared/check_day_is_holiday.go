package shared

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/piovani/wallet/infra/http"
)

type ReponseBody struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}

var holidayCommum []string = []string{"Saturday", "Sunday"}

type CheckDayIsHoliday struct{}

func NewCheckDayIsHoliday() *CheckDayIsHoliday {
	return &CheckDayIsHoliday{}
}

func (c *CheckDayIsHoliday) Execute(date time.Time) bool {
	if c.checkWeekday(date.Weekday().String()) {
		return true
	}

	return c.checkIntegration(date)
}

func (c *CheckDayIsHoliday) checkWeekday(day string) bool {
	for _, c := range holidayCommum {
		if day == c {
			return true
		}
	}
	return false
}

func (c *CheckDayIsHoliday) checkIntegration(date time.Time) bool {
	url := fmt.Sprintf("https://brasilapi.com.br/api/feriados/v1/%d", date.Year())
	http := http.NewHttp()

	content, err := http.Get(url)
	if err != nil {
		return false
	}

	var body []ReponseBody
	if err = json.Unmarshal(content, &body); err != nil {
		return false
	}

	dateFormated := date.Format("2006-01-02")
	for _, h := range body {
		if h.Date == dateFormated {
			return true
		}
	}

	return false
}
