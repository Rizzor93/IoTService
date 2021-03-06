// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package IoT

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// IoTServiceClient is the client API for IoTService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IoTServiceClient interface {
	// Device
	CreateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error)
	UpdateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error)
	DeleteDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error)
	GetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error)
	GetDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (IoTService_GetDevicesClient, error)
	// Sensor
	CreateSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error)
	UpdateSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error)
	DeleteSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error)
	GetSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error)
	GetSensors(ctx context.Context, in *Device, opts ...grpc.CallOption) (IoTService_GetSensorsClient, error)
	// Record
	CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error)
	UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error)
	DeleteRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error)
	GetRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
	GetRecords(ctx context.Context, in *Device, opts ...grpc.CallOption) (IoTService_GetRecordsClient, error)
	CreateRecordData(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error)
	GetRecordData(ctx context.Context, in *RecordDataFilter, opts ...grpc.CallOption) (IoTService_GetRecordDataClient, error)
	DeleteRecordData(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error)
}

type ioTServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIoTServiceClient(cc grpc.ClientConnInterface) IoTServiceClient {
	return &ioTServiceClient{cc}
}

func (c *ioTServiceClient) CreateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/CreateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) UpdateDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/UpdateDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) DeleteDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/DeleteDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetDevice(ctx context.Context, in *Device, opts ...grpc.CallOption) (*Device, error) {
	out := new(Device)
	err := c.cc.Invoke(ctx, "/proto.IoTService/GetDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetDevices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (IoTService_GetDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &IoTService_ServiceDesc.Streams[0], "/proto.IoTService/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &ioTServiceGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IoTService_GetDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type ioTServiceGetDevicesClient struct {
	grpc.ClientStream
}

func (x *ioTServiceGetDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ioTServiceClient) CreateSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/CreateSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) UpdateSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/UpdateSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) DeleteSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/DeleteSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetSensor(ctx context.Context, in *Sensor, opts ...grpc.CallOption) (*Sensor, error) {
	out := new(Sensor)
	err := c.cc.Invoke(ctx, "/proto.IoTService/GetSensor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetSensors(ctx context.Context, in *Device, opts ...grpc.CallOption) (IoTService_GetSensorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IoTService_ServiceDesc.Streams[1], "/proto.IoTService/GetSensors", opts...)
	if err != nil {
		return nil, err
	}
	x := &ioTServiceGetSensorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IoTService_GetSensorsClient interface {
	Recv() (*Sensor, error)
	grpc.ClientStream
}

type ioTServiceGetSensorsClient struct {
	grpc.ClientStream
}

func (x *ioTServiceGetSensorsClient) Recv() (*Sensor, error) {
	m := new(Sensor)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ioTServiceClient) CreateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/CreateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) UpdateRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/UpdateRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) DeleteRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/DeleteRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := c.cc.Invoke(ctx, "/proto.IoTService/GetRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetRecords(ctx context.Context, in *Device, opts ...grpc.CallOption) (IoTService_GetRecordsClient, error) {
	stream, err := c.cc.NewStream(ctx, &IoTService_ServiceDesc.Streams[2], "/proto.IoTService/GetRecords", opts...)
	if err != nil {
		return nil, err
	}
	x := &ioTServiceGetRecordsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IoTService_GetRecordsClient interface {
	Recv() (*Record, error)
	grpc.ClientStream
}

type ioTServiceGetRecordsClient struct {
	grpc.ClientStream
}

func (x *ioTServiceGetRecordsClient) Recv() (*Record, error) {
	m := new(Record)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ioTServiceClient) CreateRecordData(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/CreateRecordData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ioTServiceClient) GetRecordData(ctx context.Context, in *RecordDataFilter, opts ...grpc.CallOption) (IoTService_GetRecordDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &IoTService_ServiceDesc.Streams[3], "/proto.IoTService/GetRecordData", opts...)
	if err != nil {
		return nil, err
	}
	x := &ioTServiceGetRecordDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IoTService_GetRecordDataClient interface {
	Recv() (*RecordData, error)
	grpc.ClientStream
}

type ioTServiceGetRecordDataClient struct {
	grpc.ClientStream
}

func (x *ioTServiceGetRecordDataClient) Recv() (*RecordData, error) {
	m := new(RecordData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ioTServiceClient) DeleteRecordData(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.IoTService/DeleteRecordData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IoTServiceServer is the server API for IoTService service.
// All implementations must embed UnimplementedIoTServiceServer
// for forward compatibility
type IoTServiceServer interface {
	// Device
	CreateDevice(context.Context, *Device) (*Empty, error)
	UpdateDevice(context.Context, *Device) (*Empty, error)
	DeleteDevice(context.Context, *Device) (*Empty, error)
	GetDevice(context.Context, *Device) (*Device, error)
	GetDevices(*Empty, IoTService_GetDevicesServer) error
	// Sensor
	CreateSensor(context.Context, *Sensor) (*Empty, error)
	UpdateSensor(context.Context, *Sensor) (*Empty, error)
	DeleteSensor(context.Context, *Sensor) (*Empty, error)
	GetSensor(context.Context, *Sensor) (*Sensor, error)
	GetSensors(*Device, IoTService_GetSensorsServer) error
	// Record
	CreateRecord(context.Context, *Record) (*Empty, error)
	UpdateRecord(context.Context, *Record) (*Empty, error)
	DeleteRecord(context.Context, *Record) (*Empty, error)
	GetRecord(context.Context, *Record) (*Record, error)
	GetRecords(*Device, IoTService_GetRecordsServer) error
	CreateRecordData(context.Context, *RecordData) (*Empty, error)
	GetRecordData(*RecordDataFilter, IoTService_GetRecordDataServer) error
	DeleteRecordData(context.Context, *RecordData) (*Empty, error)
	mustEmbedUnimplementedIoTServiceServer()
}

// UnimplementedIoTServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIoTServiceServer struct {
}

func (UnimplementedIoTServiceServer) CreateDevice(context.Context, *Device) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedIoTServiceServer) UpdateDevice(context.Context, *Device) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDevice not implemented")
}
func (UnimplementedIoTServiceServer) DeleteDevice(context.Context, *Device) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDevice not implemented")
}
func (UnimplementedIoTServiceServer) GetDevice(context.Context, *Device) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevice not implemented")
}
func (UnimplementedIoTServiceServer) GetDevices(*Empty, IoTService_GetDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedIoTServiceServer) CreateSensor(context.Context, *Sensor) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSensor not implemented")
}
func (UnimplementedIoTServiceServer) UpdateSensor(context.Context, *Sensor) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSensor not implemented")
}
func (UnimplementedIoTServiceServer) DeleteSensor(context.Context, *Sensor) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSensor not implemented")
}
func (UnimplementedIoTServiceServer) GetSensor(context.Context, *Sensor) (*Sensor, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSensor not implemented")
}
func (UnimplementedIoTServiceServer) GetSensors(*Device, IoTService_GetSensorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSensors not implemented")
}
func (UnimplementedIoTServiceServer) CreateRecord(context.Context, *Record) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedIoTServiceServer) UpdateRecord(context.Context, *Record) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedIoTServiceServer) DeleteRecord(context.Context, *Record) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedIoTServiceServer) GetRecord(context.Context, *Record) (*Record, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedIoTServiceServer) GetRecords(*Device, IoTService_GetRecordsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecords not implemented")
}
func (UnimplementedIoTServiceServer) CreateRecordData(context.Context, *RecordData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecordData not implemented")
}
func (UnimplementedIoTServiceServer) GetRecordData(*RecordDataFilter, IoTService_GetRecordDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRecordData not implemented")
}
func (UnimplementedIoTServiceServer) DeleteRecordData(context.Context, *RecordData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecordData not implemented")
}
func (UnimplementedIoTServiceServer) mustEmbedUnimplementedIoTServiceServer() {}

// UnsafeIoTServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IoTServiceServer will
// result in compilation errors.
type UnsafeIoTServiceServer interface {
	mustEmbedUnimplementedIoTServiceServer()
}

func RegisterIoTServiceServer(s grpc.ServiceRegistrar, srv IoTServiceServer) {
	s.RegisterService(&IoTService_ServiceDesc, srv)
}

func _IoTService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/CreateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).CreateDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_UpdateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).UpdateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/UpdateDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).UpdateDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_DeleteDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).DeleteDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/DeleteDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).DeleteDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Device)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).GetDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/GetDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).GetDevice(ctx, req.(*Device))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IoTServiceServer).GetDevices(m, &ioTServiceGetDevicesServer{stream})
}

type IoTService_GetDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type ioTServiceGetDevicesServer struct {
	grpc.ServerStream
}

func (x *ioTServiceGetDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _IoTService_CreateSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).CreateSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/CreateSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).CreateSensor(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_UpdateSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).UpdateSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/UpdateSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).UpdateSensor(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_DeleteSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).DeleteSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/DeleteSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).DeleteSensor(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetSensor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sensor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).GetSensor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/GetSensor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).GetSensor(ctx, req.(*Sensor))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetSensors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Device)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IoTServiceServer).GetSensors(m, &ioTServiceGetSensorsServer{stream})
}

type IoTService_GetSensorsServer interface {
	Send(*Sensor) error
	grpc.ServerStream
}

type ioTServiceGetSensorsServer struct {
	grpc.ServerStream
}

func (x *ioTServiceGetSensorsServer) Send(m *Sensor) error {
	return x.ServerStream.SendMsg(m)
}

func _IoTService_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/CreateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).CreateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/UpdateRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).UpdateRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).DeleteRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/GetRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).GetRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetRecords_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Device)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IoTServiceServer).GetRecords(m, &ioTServiceGetRecordsServer{stream})
}

type IoTService_GetRecordsServer interface {
	Send(*Record) error
	grpc.ServerStream
}

type ioTServiceGetRecordsServer struct {
	grpc.ServerStream
}

func (x *ioTServiceGetRecordsServer) Send(m *Record) error {
	return x.ServerStream.SendMsg(m)
}

func _IoTService_CreateRecordData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).CreateRecordData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/CreateRecordData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).CreateRecordData(ctx, req.(*RecordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _IoTService_GetRecordData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecordDataFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IoTServiceServer).GetRecordData(m, &ioTServiceGetRecordDataServer{stream})
}

type IoTService_GetRecordDataServer interface {
	Send(*RecordData) error
	grpc.ServerStream
}

type ioTServiceGetRecordDataServer struct {
	grpc.ServerStream
}

func (x *ioTServiceGetRecordDataServer) Send(m *RecordData) error {
	return x.ServerStream.SendMsg(m)
}

func _IoTService_DeleteRecordData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoTServiceServer).DeleteRecordData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.IoTService/DeleteRecordData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoTServiceServer).DeleteRecordData(ctx, req.(*RecordData))
	}
	return interceptor(ctx, in, info, handler)
}

// IoTService_ServiceDesc is the grpc.ServiceDesc for IoTService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IoTService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.IoTService",
	HandlerType: (*IoTServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDevice",
			Handler:    _IoTService_CreateDevice_Handler,
		},
		{
			MethodName: "UpdateDevice",
			Handler:    _IoTService_UpdateDevice_Handler,
		},
		{
			MethodName: "DeleteDevice",
			Handler:    _IoTService_DeleteDevice_Handler,
		},
		{
			MethodName: "GetDevice",
			Handler:    _IoTService_GetDevice_Handler,
		},
		{
			MethodName: "CreateSensor",
			Handler:    _IoTService_CreateSensor_Handler,
		},
		{
			MethodName: "UpdateSensor",
			Handler:    _IoTService_UpdateSensor_Handler,
		},
		{
			MethodName: "DeleteSensor",
			Handler:    _IoTService_DeleteSensor_Handler,
		},
		{
			MethodName: "GetSensor",
			Handler:    _IoTService_GetSensor_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _IoTService_CreateRecord_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _IoTService_UpdateRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _IoTService_DeleteRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _IoTService_GetRecord_Handler,
		},
		{
			MethodName: "CreateRecordData",
			Handler:    _IoTService_CreateRecordData_Handler,
		},
		{
			MethodName: "DeleteRecordData",
			Handler:    _IoTService_DeleteRecordData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _IoTService_GetDevices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetSensors",
			Handler:       _IoTService_GetSensors_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRecords",
			Handler:       _IoTService_GetRecords_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetRecordData",
			Handler:       _IoTService_GetRecordData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "IoT.proto",
}
