package gRPC

import (
	"IoT_Service/proto/service"
	"IoT_Service/server/internal"
	"IoT_Service/server/services/db"
	"IoT_Service/server/services/db/models"
	gRPCModel "IoT_Service/server/services/gRPC/models"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoutesHandler struct {
	IoT.UnimplementedIoTServiceServer
	DB db.Handler
}

func NewHandler() *RoutesHandler {
	return &RoutesHandler{DB: db.NewDB()}
}

func (r *RoutesHandler) Initialize(cfg internal.DB) error {
	if err := r.DB.Initialize(cfg); err != nil {
		return internal.WrapError(" Error on 'DB-Initialize' ", err)
	}

	if err := r.DB.Open(); err != nil {
		return internal.WrapError(" Error on 'DB-Open' ", err)
	}

	if err := r.DB.InitializeModels(); err != nil {
		return internal.WrapError(" Error on 'DB-InitializeModels' ", err)
	}

	return nil
}

func (r *RoutesHandler) CreateDevice(ctx context.Context, in *IoT.Device) (*IoT.Empty, error) {
	d := models.Device{
		Name:       in.Name,
		Location:   in.Location,
		DeviceType: models.DeviceType(in.DeviceType),
		IPAddress:  in.IPAddress,
		Gateway:    in.Gateway,
		SubnetMask: in.SubnetMask,
		OPCUrl:     in.OPCUrl,
	}

	if err := r.DB.CreateDevice(&d); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) UpdateDevice(ctx context.Context, in *IoT.Device) (*IoT.Empty, error) {
	d := models.Device{
		ID:         uint(in.ID),
		Name:       in.Name,
		Location:   in.Location,
		DeviceType: models.DeviceType(in.DeviceType),
		IPAddress:  in.IPAddress,
		Gateway:    in.Gateway,
		SubnetMask: in.SubnetMask,
		OPCUrl:     in.OPCUrl,
	}
	if err := r.DB.UpdateDevice(&d); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) DeleteDevice(ctx context.Context, in *IoT.Device) (*IoT.Empty, error) {
	d := models.Device{
		ID: uint(in.ID),
	}
	if err := r.DB.DeleteDevice(&d); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetDevice(ctx context.Context, in *IoT.Device) (*IoT.Device, error) {
	d := models.Device{
		ID: uint(in.ID),
	}
	if err := r.DB.GetDevice(&d); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Device{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Device{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	return &IoT.Device{
		ID:         uint32(d.ID),
		Name:       d.Name,
		Location:   d.Location,
		DeviceType: IoT.DeviceType(d.DeviceType),
		IPAddress:  d.IPAddress,
		Gateway:    d.Gateway,
		SubnetMask: d.SubnetMask,
		OPCUrl:     d.OPCUrl,
	}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetDevices(in *IoT.Empty, stream IoT.IoTService_GetDevicesServer) error {
	streamError := make(chan error)

	go func() {
		defer close(streamError)
		var devices []models.Device

		if err := r.DB.GetDevices(&devices); err != nil {
			streamError <- err
			return
		}
		for _, device := range devices {
			if err := stream.Send(&IoT.Device{
				ID:         uint32(device.ID),
				Name:       device.Name,
				Location:   device.Location,
				DeviceType: IoT.DeviceType(device.DeviceType),
				IPAddress:  device.IPAddress,
				Gateway:    device.Gateway,
				SubnetMask: device.SubnetMask,
				OPCUrl:     device.OPCUrl,
			}); err != nil {
				streamError <- err
				return
			}
		}
	}()

	err := <-streamError
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	return gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}

func (r *RoutesHandler) CreateSensor(ctx context.Context, in *IoT.Sensor) (*IoT.Empty, error) {
	sensor := models.Sensor{
		DeviceID:       uint(in.DeviceID),
		Name:           in.Name,
		Location:       in.Location,
		DataType:       models.SensorDataType(in.DataType),
		GPIONum:        in.GPIONum,
		GPIOType:       models.GPIOType(in.GPIOType),
		PullUpResistor: in.PullUpResistor,
		PLCDataArea:    models.PLCDataArea(in.PLCDataArea),
		DataOffset:     in.DataOffset,
		BoolIndex:      in.BoolIndex,
		DbNum:          in.DbNum,
		NodeID:         in.NodeID,
	}
	if err := r.DB.CreateSensor(&sensor); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) UpdateSensor(ctx context.Context, in *IoT.Sensor) (*IoT.Empty, error) {
	sensor := models.Sensor{
		ID:             uint(in.ID),
		DeviceID:       uint(in.DeviceID),
		Name:           in.Name,
		Location:       in.Location,
		DataType:       models.SensorDataType(in.DataType),
		GPIONum:        in.GPIONum,
		GPIOType:       models.GPIOType(in.GPIOType),
		PullUpResistor: in.PullUpResistor,
		PLCDataArea:    models.PLCDataArea(in.PLCDataArea),
		DataOffset:     in.DataOffset,
		BoolIndex:      in.BoolIndex,
		DbNum:          in.DbNum,
		NodeID:         in.NodeID,
	}
	if err := r.DB.UpdateSensor(&sensor); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) DeleteSensor(ctx context.Context, in *IoT.Sensor) (*IoT.Empty, error) {
	sensor := models.Sensor{
		ID: uint(in.ID),
	}
	if err := r.DB.DeleteSensor(&sensor); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetSensor(ctx context.Context, in *IoT.Sensor) (*IoT.Sensor, error) {
	sensor := models.Sensor{
		ID: uint(in.ID),
	}
	if err := r.DB.GetSensor(&sensor); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Sensor{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Sensor{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Sensor{
		ID:             uint32(sensor.ID),
		DeviceID:       uint32(sensor.DeviceID),
		Name:           sensor.Name,
		Location:       sensor.Location,
		DataType:       IoT.SensorDataType(sensor.DataType),
		GPIONum:        sensor.GPIONum,
		GPIOType:       IoT.GPIOType(sensor.GPIOType),
		PullUpResistor: sensor.PullUpResistor,
		PLCDataArea:    IoT.PLCDataArea(sensor.PLCDataArea),
		DataOffset:     sensor.DataOffset,
		BoolIndex:      sensor.BoolIndex,
		DbNum:          sensor.DbNum,
		NodeID:         sensor.NodeID,
	}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetSensors(in *IoT.Device, stream IoT.IoTService_GetSensorsServer) error {
	streamError := make(chan error)
	device := models.Device{
		ID: uint(in.ID),
	}

	go func(device models.Device) {
		defer close(streamError)

		var sensors []models.Sensor
		if err := r.DB.GetSensors(&device, &sensors); err != nil {
			streamError <- err
			return
		}

		for _, sensor := range sensors {
			if err := stream.Send(&IoT.Sensor{
				ID:             uint32(sensor.ID),
				DeviceID:       uint32(sensor.DeviceID),
				Name:           sensor.Name,
				Location:       sensor.Location,
				DataType:       IoT.SensorDataType(sensor.DataType),
				GPIONum:        sensor.GPIONum,
				GPIOType:       IoT.GPIOType(sensor.GPIOType),
				PullUpResistor: sensor.PullUpResistor,
				PLCDataArea:    IoT.PLCDataArea(sensor.PLCDataArea),
				DataOffset:     sensor.DataOffset,
				BoolIndex:      sensor.BoolIndex,
				DbNum:          sensor.DbNum,
				NodeID:         sensor.NodeID,
			}); err != nil {
				streamError <- err
				return
			}
		}

	}(device)

	err := <-streamError
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	return gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}

func (r *RoutesHandler) CreateRecord(ctx context.Context, in *IoT.Record) (*IoT.Empty, error) {
	var sensorRecord []models.RecordSensor

	for i := range in.RecordSensor {
		sensorRecord = append(sensorRecord, models.RecordSensor{
			SensorID:        uint(in.RecordSensor[i].SensorID),
			RecordType:      models.RecordType(in.RecordSensor[i].RecordType),
			TriggerValueMin: in.RecordSensor[i].TriggerValueMin,
			TriggerValueMax: in.RecordSensor[i].TriggerValueMax,
			Interval:        in.RecordSensor[i].Interval,
		})
	}

	record := models.Record{
		DeviceID:     uint(in.DeviceID),
		Name:         in.Name,
		Running:      in.Running,
		RecordSensor: sensorRecord,
	}

	if err := r.DB.CreateRecord(&record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) UpdateRecord(ctx context.Context, in *IoT.Record) (*IoT.Empty, error) {
	var sensorRecord []models.RecordSensor

	for i := range in.RecordSensor {
		sensorRecord = append(sensorRecord, models.RecordSensor{
			RecordID:        uint(in.ID),
			SensorID:        uint(in.RecordSensor[i].SensorID),
			RecordType:      models.RecordType(in.RecordSensor[i].RecordType),
			TriggerValueMin: in.RecordSensor[i].TriggerValueMin,
			TriggerValueMax: in.RecordSensor[i].TriggerValueMax,
			Interval:        in.RecordSensor[i].Interval,
		})
	}

	record := models.Record{
		ID:           uint(in.ID),
		DeviceID:     uint(in.DeviceID),
		Name:         in.Name,
		Running:      in.Running,
		RecordSensor: sensorRecord,
	}

	if err := r.DB.UpdateRecord(&record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) DeleteRecord(ctx context.Context, in *IoT.Record) (*IoT.Empty, error) {
	record := models.Record{
		ID:       uint(in.ID),
		DeviceID: uint(in.DeviceID),
	}
	if err := r.DB.DeleteRecord(&record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetRecord(ctx context.Context, in *IoT.Record) (*IoT.Record, error) {
	record := models.Record{
		ID:       uint(in.ID),
		DeviceID: uint(in.DeviceID),
	}
	if err := r.DB.GetRecord(&record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Record{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Record{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	var recordSensors []*IoT.RecordSensor
	for _, sensor := range record.RecordSensor {
		recordSensors = append(recordSensors, &IoT.RecordSensor{
			SensorID:        uint32(sensor.SensorID),
			RecordType:      IoT.RecordType(sensor.RecordType),
			TriggerValueMin: sensor.TriggerValueMin,
			TriggerValueMax: sensor.TriggerValueMax,
			Interval:        sensor.Interval,
		})
	}

	return &IoT.Record{
		ID:           in.ID,
		DeviceID:     in.DeviceID,
		Name:         record.Name,
		Running:      record.Running,
		RecordSensor: recordSensors,
	}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetRecords(in *IoT.Device, stream IoT.IoTService_GetRecordsServer) error {
	streamError := make(chan error)
	device := models.Device{
		ID: uint(in.ID),
	}

	go func(device models.Device) {
		defer close(streamError)

		records, err := r.DB.GetRecords(&device)
		if err != nil {
			streamError <- err
			return
		}

		for _, record := range records {
			var recordSensors []*IoT.RecordSensor
			for _, sensor := range record.RecordSensor {
				recordSensors = append(recordSensors, &IoT.RecordSensor{
					SensorID:        uint32(sensor.SensorID),
					RecordType:      IoT.RecordType(sensor.RecordType),
					TriggerValueMin: sensor.TriggerValueMin,
					TriggerValueMax: sensor.TriggerValueMax,
					Interval:        sensor.Interval,
				})
			}
			if err := stream.Send(&IoT.Record{
				ID:           uint32(record.ID),
				DeviceID:     uint32(record.DeviceID),
				Name:         record.Name,
				Running:      record.Running,
				RecordSensor: recordSensors,
			}); err != nil {
				streamError <- err
				return
			}
		}

	}(device)

	err := <-streamError
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	return gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}

func (r *RoutesHandler) CreateRecordData(ctx context.Context, in *IoT.RecordData) (*IoT.Empty, error) {
	recordData := models.RecordData{
		RecordID: uint(in.ID),
		SensorID: uint(in.SensorID),
		DataType: models.SensorDataType(in.DataType),
		Value:    in.Value,
	}

	if err := r.DB.CreateRecordData(&recordData); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}

	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) DeleteRecordData(ctx context.Context, in *IoT.RecordData) (*IoT.Empty, error) {
	record := models.RecordData{
		RecordID: uint(in.ID),
		SensorID: uint(in.SensorID),
	}
	if err := r.DB.DeleteRecordData(&record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return &IoT.Empty{}, gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
func (r *RoutesHandler) GetRecordData(in *IoT.RecordDataFilter, stream IoT.IoTService_GetRecordDataServer) error {
	streamError := make(chan error)
	recordDataFilter := models.RecordDataFilter{
		ID:          uint(in.ID),
		SensorID:    uint(in.SensorID),
		FilterValue: models.FilterValue(in.FilterValue),
		MinValue:    in.MinValue,
		MaxValue:    in.MaxValue,
		FilterTime:  models.FilterTime(in.FilterTime),
	}

	go func(recordDataFilter models.RecordDataFilter) {
		defer close(streamError)

		recordData, err := r.DB.GetRecordData(&recordDataFilter)
		if err != nil {
			streamError <- err
			return
		}

		for _, data := range recordData {
			if err := stream.Send(&IoT.RecordData{
				ID:        uint32(data.RecordID),
				SensorID:  uint32(data.SensorID),
				DataType:  IoT.SensorDataType(data.DataType),
				Value:     data.Value,
				Timestamp: timestamppb.New(data.CreatedAt),
			}); err != nil {
				streamError <- err
				return
			}
		}

	}(recordDataFilter)
	err := <-streamError

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gRPCModel.CreateGRPCStatus(gRPCModel.NotFoundError{Message: err.Error()})
		} else {
			return gRPCModel.CreateGRPCStatus(gRPCModel.InternalError{Message: err.Error()})
		}
	}
	return gRPCModel.CreateGRPCStatus(gRPCModel.OK{})
}
