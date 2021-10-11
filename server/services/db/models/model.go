package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Device struct {
	ID         uint   `gorm:"primary_key" json:"ID"`
	Name       string `json:"Name"`
	Location   string `json:"Location"`
	DeviceType uint   `json:"DeviceType"`
	IPAddress  string `json:"IPAddress"`
	Gateway    string `json:"Gateway"`
	SubnetMask string `json:"SubnetMask"`
	OPCUrl     string `json:"OPCUrl"`
}

type Sensor struct {
	ID       uint   `gorm:"primary_key" json:"ID"`
	DeviceID uint   `json:"DeviceID"`
	Name     string `json:"Name"`
	Location string `json:"Location"`
	DataType uint   `json:"DataType"`
	// Arduino RaspberryPi
	GPIONum        int32 `json:"GPIONum"`
	GPIOType       uint  `json:"GPIOType"`
	PullUpResistor bool  `json:"PullUpResistor"`
	// PLC
	PLCDataArea uint  `json:"PLCDataArea"`
	DataOffset  int32 `json:"DataOffset"`
	BoolIndex   int32 `json:"BoolIndex"`
	DbNum       int32 `json:"DbNum"`
	// OPC
	NodeID string `json:"NodeID"`
}

type Record struct {
	ID           uint           `gorm:"primary_key" json:"ID"`
	DeviceID     uint           `json:"DeviceID"`
	Name         string         `json:"Name"`
	Running      bool           `json:"Running"`
	Data         []RecordData   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	RecordSensor []RecordSensor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type RecordSensor struct {
	RecordID        uint    `json:"RecordID"`
	SensorID        uint    `json:"SensorID"`
	RecordType      uint    `json:"RecordType"`
	TriggerValueMin string  `json:"TriggerValueMin"`
	TriggerValueMax string  `json:"TriggerValueMax"`
	Interval        float64 `json:"Interval"` // Represent in seconds

}

type RecordData struct {
	RecordID  uint      `json:"RecordID"`
	SensorID  uint      `json:"SensorID"`
	DataType  uint      `json:"DataType"`
	Value     string    `json:"Value"`
	CreatedAt time.Time `json:"created_at"`
}

type FilterValue int32

const (
	EqualValue       FilterValue = 0
	OddValue         FilterValue = 1
	SmallerThenValue FilterValue = 2
	BiggerThenValue  FilterValue = 3
	RangeValue       FilterValue = 4
	AllValue         FilterValue = 5
)

type FilterTime int32

const (
	Now        FilterTime = 0
	Last24H    FilterTime = 1
	LastWeek   FilterTime = 2
	LastMonth  FilterTime = 3
	Last3MONTH FilterTime = 4
	Last6MONTH FilterTime = 5
	LastYear   FilterTime = 6
	All        FilterTime = 7
)

type RecordDataFilter struct {
	ID          uint        `json:"ID"`
	SensorID    uint        `json:"SensorID"`
	FilterValue FilterValue `json:"FilterValue"`
	MinValue    string      `json:"MinValue"`
	MaxValue    string      `json:"MaxValue"`
	FilterTime  FilterTime  `json:"FilterTime"`
}
