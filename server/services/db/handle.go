package db

import (
	"IoT_Service/server/services/db/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func (d *DB) CreateDevice(device *models.Device) error {
	if err := d.db.Create(&device).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateDevice(device *models.Device) error {
	if err := d.db.Find(&models.Device{ID: device.ID}).Error; err != nil {
		return err
	}
	err := d.db.Model(&device).Updates(map[string]interface{}{
		"name":        &device.Name,
		"location":    &device.Location,
		"device_type": &device.DeviceType,
		"ip_address":  &device.IPAddress,
		"gateway":     &device.Gateway,
		"subnet_mask": &device.SubnetMask,
		"opc_url":     &device.OPCUrl,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteDevice(device *models.Device) error {
	if err := d.db.Find(&device).Error; err != nil {
		return err
	}
	if err := d.db.Model(&device).Delete(&device).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetDevice(device *models.Device) error {
	if err := d.db.Find(&device).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetDevices(devices *[]models.Device) error {
	if err := d.db.Find(&devices).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) CreateSensor(sensor *models.Sensor) error {
	if err := d.db.Create(&sensor).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateSensor(sensor *models.Sensor) error {
	if err := d.db.Find(&models.Sensor{ID: sensor.ID}).Error; err != nil {
		return err
	}
	err := d.db.Model(&sensor).Updates(map[string]interface{}{
		"device_id":        &sensor.DeviceID,
		"name":             &sensor.Name,
		"location":         &sensor.Location,
		"data_type":        &sensor.DataType,
		"gpio_num":         &sensor.GPIONum,
		"gpio_type":        &sensor.GPIOType,
		"pull_up_resistor": &sensor.PullUpResistor,
		"plc_data_area":    &sensor.PLCDataArea,
		"data_offset":      &sensor.DataOffset,
		"bool_index":       &sensor.BoolIndex,
		"db_num":           &sensor.DbNum,
		"node_id":          &sensor.NodeID,
	}).Error

	if err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteSensor(sensor *models.Sensor) error {
	if err := d.db.Find(&sensor).Error; err != nil {
		return err
	}
	if err := d.db.Model(&sensor).Delete(&sensor).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetSensor(sensor *models.Sensor) error {
	if err := d.db.Find(&sensor).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetSensors(device *models.Device, sensors *[]models.Sensor) error {
	if err := d.db.Find(&device).Error; err != nil {
		return err
	}
	if err := d.db.Find(&sensors, &models.Sensor{DeviceID: device.ID}).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) CreateRecord(record *models.Record) error {
	if err := d.db.Find(&models.Device{ID: record.DeviceID}).Error; err != nil {
		return err
	}
	if err := d.db.Create(&record).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) UpdateRecord(record *models.Record) error {
	if err := d.db.Where("record_id = ?", record.ID).Delete(&models.RecordSensor{}).Error; err != nil {
		return err
	}

	for _, sensor := range record.RecordSensor {
		if err := d.db.Create(&sensor).Error; err != nil {
			return err
		}
	}
	if err := d.db.Model(&models.Record{}).Where("id = ?", record.ID).
		Update("name", record.Name).Error; err != nil {
		return err
	}

	if err := d.db.Model(&models.Record{}).Where("id = ?", record.ID).
		Update("running", record.Running).Error; err != nil {
		return err
	}

	return nil
}

func (d *DB) DeleteRecord(record *models.Record) error {
	if err := d.db.Find(&record).Error; err != nil {
		return err
	}
	if err := d.db.Model(&record).Delete(&record).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetRecord(record *models.Record) error {
	if err := d.db.Find(&models.Device{ID: record.DeviceID}).Error; err != nil {
		return err
	}
	if err := d.db.Find(&record).Error; err != nil {
		return err
	}
	if err := d.db.Find(&record.RecordSensor, &models.RecordSensor{RecordID: record.ID}).Error; err != nil {
		return err
	}
	return nil
}

func (d *DB) GetRecords(device *models.Device) ([]*models.Record, error) {
	if err := d.db.Find(&device).Error; err != nil {
		return nil, err
	}

	var records []*models.Record
	err := d.db.Find(&records, &models.Record{DeviceID: device.ID}).Error
	if err != nil {
		return nil, err
	}

	for i, record := range records {
		err := d.db.Find(&records[i].RecordSensor, &models.RecordSensor{RecordID: record.ID}).Error
		if err != nil {
			return nil, err
		}
	}

	return records, nil
}

func (d *DB) CreateRecordData(data *models.RecordData) error {
	if err := d.db.Find(&models.Record{ID: data.RecordID}).Error; err != nil {
		return err
	}
	if err := d.db.Find(&models.Sensor{ID: data.SensorID}).Error; err != nil {
		return err
	}
	if err := d.db.Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (d *DB) GetRecordData(record *models.RecordDataFilter) ([]*models.RecordData, error) {
	if err := d.db.Find(&models.Record{ID: record.ID}).Error; err != nil {
		return nil, err
	}
	if err := d.db.Find(&models.RecordData{SensorID: record.SensorID}).Error; err != nil {
		return nil, err
	}

	sensor := models.Sensor{
		ID: record.SensorID,
	}
	if err := d.GetSensor(&sensor); err != nil {
		return nil, err
	}

	var recordData []*models.RecordData
	d.db.Scopes(filterTime(record.FilterTime), filterValue(sensor.DataType, record.FilterValue, record.MinValue, record.MaxValue)).
		Find(&recordData, &models.RecordData{
			RecordID: record.ID,
			SensorID: record.SensorID,
		})

	return recordData, nil
}

func (d *DB) DeleteRecordData(record *models.RecordData) error {
	if err := d.db.Find(&models.Record{ID: record.RecordID}).Error; err != nil {
		return err
	}
	if err := d.db.Find(&models.RecordData{SensorID: record.SensorID}).Error; err != nil {
		return err
	}
	if err := d.db.Where("record_id = ? AND sensor_id = ?", record.RecordID, record.SensorID).
		Delete(&record).Error; err != nil {
		return err
	}

	return nil
}

func filterTime(filterTime models.FilterTime) func(db *gorm.DB) *gorm.DB {
	var LAST24H = time.Now().UTC().Add(-(time.Hour * 24))
	var LASTWEEK = time.Now().UTC().Add(-((time.Hour * 24) * 7))
	var LASTMONTH = time.Now().UTC().Add(-(((time.Hour * 24) * 7) * 4))
	var LAST3MONTH = time.Now().UTC().Add(-((((time.Hour * 24) * 7) * 4) * 3))
	var LAST6MONTH = time.Now().UTC().Add(-((((time.Hour * 24) * 7) * 4) * 6))
	var LASTYEAR = time.Now().UTC().Add(-((((time.Hour * 24) * 7) * 4) * 12))

	return func(db *gorm.DB) *gorm.DB {
		switch filterTime {
		case models.All:
			return db.Where("")
		case models.Now:
			return db.Where("created_at = ?", time.Now().UTC())
		case models.Last24H:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LAST24H)
		case models.LastWeek:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LASTWEEK)
		case models.LastMonth:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LASTMONTH)
		case models.Last3MONTH:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LAST3MONTH)
		case models.Last6MONTH:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LAST6MONTH)
		case models.LastYear:
			return db.Where("created_at < ? AND created_at > ?", time.Now().UTC(), LASTYEAR)
		default:
			return db.Where("")
		}
	}
}

func castType(dataType models.SensorDataType) string {
	switch dataType {
	case models.SensorDataType_BOOL:
		return "INTEGER"
	case models.SensorDataType_BYTE:
		return "BYTEA"
	case models.SensorDataType_INT:
		return "INTEGER"
	case models.SensorDataType_FLOAT:
		return "FLOAT"
	case models.SensorDataType_DOUBLE:
		return "DOUBLE PRECISION"
	case models.SensorDataType_STRING:
		return "TEXT"
	case models.SensorDataType_DATETIME:
		return "DATE"
	default:
		return "TEXT"
	}
}

func filterValue(dataType models.SensorDataType, comp models.FilterValue, minValue string, maxValue string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch comp {
		case models.AllValue:
			return db.Where("")
		case models.EqualValue:
			return db.Where(fmt.Sprintf("CAST(value AS %s) = %s ", castType(dataType), minValue))
		case models.OddValue:
			return db.Where(fmt.Sprintf("CAST(value AS %s) != %s ", castType(dataType), minValue))
		case models.SmallerThenValue:
			return db.Where(fmt.Sprintf("CAST(value AS %s) < %s ", castType(dataType), minValue))
		case models.BiggerThenValue:
			return db.Where(fmt.Sprintf("CAST(value AS %s) > %s ", castType(dataType), minValue))
		case models.RangeValue:
			return db.Where(fmt.Sprintf("CAST(value AS %s) > %s AND CAST(value AS %s) < %s", castType(dataType), minValue, castType(dataType), maxValue))
		default:
			return db.Where("")
		}
	}
}
