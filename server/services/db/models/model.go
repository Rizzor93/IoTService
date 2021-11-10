package models

import (
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type DeviceType int32

const (
	DeviceType_ARDUINO      DeviceType = 0
	DeviceType_RASPBERRYPI  DeviceType = 1
	DeviceType_PLC_Siemens  DeviceType = 2
	DeviceType_PLC_Beckhoff DeviceType = 3
	DeviceType_PLC_WAGO     DeviceType = 4
	DeviceType_OPC          DeviceType = 5
)

type Device struct {
	ID         uint       `gorm:"primary_key" json:"ID"`
	Name       string     `json:"Name"`
	Location   string     `json:"Location"`
	DeviceType DeviceType `json:"DeviceType"`
	IPAddress  string     `json:"IPAddress"`
	Gateway    string     `json:"Gateway"`
	SubnetMask string     `json:"SubnetMask"`
	OPCUrl     string     `json:"OPCUrl"`
}

type SensorDataType int32

const (
	SensorDataType_BOOL     SensorDataType = 0
	SensorDataType_BYTE     SensorDataType = 1
	SensorDataType_INT      SensorDataType = 2
	SensorDataType_FLOAT    SensorDataType = 3
	SensorDataType_DOUBLE   SensorDataType = 4
	SensorDataType_STRING   SensorDataType = 5
	SensorDataType_DATETIME SensorDataType = 6
)

type GPIOType int32

const (
	GPIOType_Input        GPIOType = 0
	GPIOType_Output       GPIOType = 1
	GPIOType_AnalogInput  GPIOType = 2
	GPIOType_AnalogOutput GPIOType = 3
)

type PLCDataArea int32

const (
	PLCDataArea_DB PLCDataArea = 0
	PLCDataArea_MK PLCDataArea = 1
	PLCDataArea_PE PLCDataArea = 2
	PLCDataArea_PA PLCDataArea = 3
	PLCDataArea_CT PLCDataArea = 4
)

type Sensor struct {
	ID       uint           `gorm:"primary_key" json:"ID"`
	DeviceID uint           `json:"DeviceID"`
	Name     string         `json:"Name"`
	Location string         `json:"Location"`
	DataType SensorDataType `json:"DataType"`
	// Arduino RaspberryPi
	GPIONum        int32    `json:"GPIONum"`
	GPIOType       GPIOType `json:"GPIOType"`
	PullUpResistor bool     `json:"PullUpResistor"`
	// PLC
	PLCDataArea PLCDataArea `json:"PLCDataArea"`
	DataOffset  int32       `json:"DataOffset"`
	BoolIndex   int32       `json:"BoolIndex"`
	DbNum       int32       `json:"DbNum"`
	// OPC
	NodeID string `json:"NodeID"`
}

type RecordType int32

const (
	RecordType_EQUAL        RecordType = 0
	RecordType_ODD          RecordType = 1
	RecordType_SMALLER_THEN RecordType = 2
	RecordType_BIGGER_THEN  RecordType = 3
	RecordType_RANGE        RecordType = 4
	RecordType_ALWAYS       RecordType = 5
)

type Record struct {
	ID           uint           `gorm:"primary_key" json:"ID"`
	DeviceID     uint           `json:"DeviceID"`
	Name         string         `json:"Name"`
	Running      bool           `json:"Running"`
	Data         []RecordData   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	RecordSensor []RecordSensor `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type RecordSensor struct {
	RecordID        uint       `json:"RecordID"`
	SensorID        uint       `json:"SensorID"`
	RecordType      RecordType `json:"RecordType"`
	TriggerValueMin string     `json:"TriggerValueMin"`
	TriggerValueMax string     `json:"TriggerValueMax"`
	Interval        float64    `json:"Interval"` // Represent in seconds

}

type RecordData struct {
	RecordID  uint           `json:"RecordID"`
	SensorID  uint           `json:"SensorID"`
	DataType  SensorDataType `json:"DataType"`
	Value     string         `json:"Value"`
	CreatedAt time.Time      `json:"created_at"`
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
