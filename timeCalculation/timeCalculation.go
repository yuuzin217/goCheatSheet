package timeCalculation

import "time"

type ManagedTime struct {
	Date time.Time
}

func GetCurrentTime() *ManagedTime {
	return &ManagedTime{
		Date: time.Now(),
	}
}

func NewTimeObject(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) *ManagedTime {
	return &ManagedTime{
		Date: time.Date(year, month, day, hour, min, sec, nsec, loc),
	}
}

// 時刻のフォーマット
const (
	DefaultFormat = "2006-01-02 15:04:05"
)

func (mt *ManagedTime) Format(t time.Time) string {
	return mt.Date.Format(DefaultFormat)
}

/*
	時刻加算
*/

func (mt *ManagedTime) AddYear(add int) {
	mt.Date.AddDate(add, 0, 0)
}

func (mt *ManagedTime) AddMonth(add int) {
	mt.Date.AddDate(0, add, 0)
}

func (mt *ManagedTime) AddDay(add int) {
	mt.Date.AddDate(0, 0, add)
}

func (mt *ManagedTime) AddSecond(add int) {
	mt.Date.Add(time.Duration(add) * time.Second)
}

func (mt *ManagedTime) AddMillisecond(add int) {
	mt.Date.Add(time.Duration(add) * time.Millisecond)
}

func (mt *ManagedTime) AddMicrosecond(add int) {
	mt.Date.Add(time.Duration(add) * time.Microsecond)
}

/*
	時刻減算
*/

func (mt *ManagedTime) SubYear(add int) {
	mt.AddYear(-add)
}

func (mt *ManagedTime) SubMonth(add int) {
	mt.AddMonth(-add)
}

func (mt *ManagedTime) SubDay(add int) {
	mt.AddDay(-add)
}

func (mt *ManagedTime) SubSecond(add int) {
	mt.AddSecond(-add)
}

func (mt *ManagedTime) SubMillisecond(add int) {
	mt.AddMillisecond(-add)
}

func (mt *ManagedTime) SubMicrosecond(add int) {
	mt.AddMicrosecond(-add)
}

func (mt *ManagedTime) CheckSameDate(targetDate time.Time) bool {
	return mt.Date.Equal(targetDate)
}
