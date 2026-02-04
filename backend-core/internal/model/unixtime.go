package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UnixTime time.Time

func (ut *UnixTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	if v, ok := value.(int64); ok {
		*ut = UnixTime(time.Unix(v, 0))
		return nil
	}
	return fmt.Errorf("cannot scan type %T into UnixTime", value)
}

func (ut UnixTime) Value() (driver.Value, error) {
	return time.Time(ut).Unix(), nil
}

func (ut UnixTime) Time() time.Time {
	return time.Time(ut)
}
