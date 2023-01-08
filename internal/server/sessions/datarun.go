package sessions

import (
	"fmt"
	"sort"
	"time"
)

var drctr int

type DataRun struct {
	ID        string
	StartTime time.Time
	EndTime   time.Time
	data      []DataFrame
}

type DataFrame struct {
	Value float32
	Frame int
}

func NewDataRun() *DataRun {
	drctr++
	return &DataRun{
		ID:        fmt.Sprintf("%d", drctr),
		StartTime: time.Now(),
	}
}

func (d *DataRun) AddData(frame int, value float32) error {
	d.data = append(d.data, DataFrame{Frame: frame, Value: value})
	return nil
}

func (d *DataRun) Data() []DataFrame {
	sort.Slice(d.data, func(i, j int) bool {
		return d.data[i].Frame < d.data[j].Frame
	})
	return d.data
}
