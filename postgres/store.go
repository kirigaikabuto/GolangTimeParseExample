package postgres

import (
	"fmt"
	"strings"
	"time"
)
const layout = "2006-01-02"
type SpecialTime struct {
	time.Time
}
func (ct *SpecialTime) UnmarshalJSON(b []byte) (err error){
	s := strings.Trim(string(b),"\"")
	if s == "null" {
		ct.Time = time.Time{}
	}
	ct.Time, err = time.Parse(layout,s)
	return
}
func (ct *SpecialTime) MarshalJSON() ([]byte, error){
	return []byte(fmt.Sprintf("\"%s\"",ct.Time.Format(layout))),nil
}
type Data struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DateOfBirth SpecialTime `json:"date_of_birth"`
}
type DataRepo interface {
	Create (obj *Data) (*Data,error)
	Get(id string) (*Data,error)
}