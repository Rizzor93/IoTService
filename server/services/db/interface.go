package db

import (
	"IoT_Service/server/internal"
	"IoT_Service/server/services/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Handler handles all DB function
type Handler interface {
	Initialize(cfg internal.DB) error
	InitializeModels() error
	Open() error
	Close() error
	// Device
	CreateDevice(device *models.Device) error
	UpdateDevice(device *models.Device) error
	DeleteDevice(device *models.Device) error
	GetDevice(device *models.Device) error
	GetDevices(devices *[]models.Device) error
	// Sensor
	CreateSensor(sensor *models.Sensor) error
	UpdateSensor(sensor *models.Sensor) error
	DeleteSensor(sensor *models.Sensor) error
	GetSensor(sensor *models.Sensor) error
	GetSensors(device *models.Device, sensors *[]models.Sensor) error
	// Record
	CreateRecord(record *models.Record) error
	UpdateRecord(record *models.Record) error
	DeleteRecord(record *models.Record) error
	GetRecord(record *models.Record) error
	GetRecords(device *models.Device) ([]*models.Record, error)
	// RecordData
	CreateRecordData(data *models.RecordData) error
	GetRecordData(record *models.RecordDataFilter) ([]*models.RecordData, error)
	DeleteRecordData(data *models.RecordData) error
}

// DB represents the DB
type DB struct {
	db     *gorm.DB
	Config internal.DB
}

// NewDB create new instance
func NewDB() *DB {
	return &DB{}
}

// Initialize configuration
func (d *DB) Initialize(cfg internal.DB) error {
	d.Config = cfg
	return nil
}

// InitializeModels create/modify the DB Tables
func (d *DB) InitializeModels() error {
	// Migration of the dbModel
	err := d.db.AutoMigrate(&models.Device{}).Error
	err = d.db.AutoMigrate(&models.Sensor{}).Error
	err = d.db.AutoMigrate(&models.Record{}).Error
	err = d.db.AutoMigrate(&models.RecordSensor{}).Error
	err = d.db.AutoMigrate(&models.RecordData{}).Error
	// Add foreignKeys
	err = d.db.Model(&models.Sensor{}).AddForeignKey("device_id", "devices(id)", "CASCADE", "CASCADE").Error
	err = d.db.Model(&models.Record{}).AddForeignKey("device_id", "devices(id)", "CASCADE", "CASCADE").Error
	err = d.db.Model(&models.RecordSensor{}).AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE").Error
	err = d.db.Model(&models.RecordSensor{}).AddForeignKey("record_id", "records(id)", "CASCADE", "CASCADE").Error
	err = d.db.Model(&models.RecordData{}).AddForeignKey("sensor_id", "sensors(id)", "CASCADE", "CASCADE").Error
	err = d.db.Model(&models.RecordData{}).AddForeignKey("record_id", "records(id)", "CASCADE", "CASCADE").Error

	if err != nil {
		return err
	}
	return nil
}

// Open the database
func (d *DB) Open() error {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s",
		d.Config.Host, d.Config.User, d.Config.DBName, d.Config.Password, d.Config.Port)
	db, err := gorm.Open(d.Config.Dialect, dbUri)
	d.db = db
	if err != nil {
		return err
	}
	return nil
}

// Close the database
func (d *DB) Close() error {
	if err := d.db.Close(); err != nil {
		return err
	}

	return nil
}
