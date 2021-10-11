package restApi

import (
	IoT "IoT_Service/proto/service"
	"IoT_Service/server/services/restApi/models"
	"IoT_Service/server/services/restApi/transport"
	"encoding/json"
	"io"
	"net/http"
)

func (s *Server) createDevice(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.CreateDevice(s.grpcCtx, &device)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) updateDevice(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.UpdateDevice(s.grpcCtx, &device)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) deleteDevice(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.DeleteDevice(s.grpcCtx, &device)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getDevice(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.GetDevice(s.grpcCtx, &device)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getDevices(w http.ResponseWriter, r *http.Request) {
	stream, err := s.grpcClient.GetDevices(s.grpcCtx, &IoT.Empty{})
	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	streamDone := make(chan bool)

	go func() {
		defer close(streamDone)
		header := false
		var streamEncoder *json.Encoder
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				streamDone <- true
				return
			}
			if err != nil {
				transport.SendError(w, transport.GrpcStatusToHTTPError(err))
				return
			}
			if !header {
				streamEncoder = transport.JsonEncoder(w, http.StatusOK)
				header = true
			}

			transport.SendJSONStream(streamEncoder, res)
		}

	}()
	<-streamDone

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
}

func (s *Server) createSensor(w http.ResponseWriter, r *http.Request) {
	sensor := IoT.Sensor{}

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.CreateSensor(s.grpcCtx, &sensor)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) updateSensor(w http.ResponseWriter, r *http.Request) {
	sensor := IoT.Sensor{}

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.UpdateSensor(s.grpcCtx, &sensor)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) deleteSensor(w http.ResponseWriter, r *http.Request) {
	sensor := IoT.Sensor{}

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.DeleteSensor(s.grpcCtx, &sensor)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getSensor(w http.ResponseWriter, r *http.Request) {
	sensor := IoT.Sensor{}

	if err := json.NewDecoder(r.Body).Decode(&sensor); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.GetSensor(s.grpcCtx, &sensor)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getSensors(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	stream, err := s.grpcClient.GetSensors(s.grpcCtx, &device)
	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	streamDone := make(chan bool)

	go func() {
		defer close(streamDone)
		header := false
		var streamEncoder *json.Encoder
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				streamDone <- true
				return
			}
			if err != nil {
				transport.SendError(w, transport.GrpcStatusToHTTPError(err))
				return
			}
			if !header {
				streamEncoder = transport.JsonEncoder(w, http.StatusOK)
				header = true
			}

			transport.SendJSONStream(streamEncoder, res)
		}

	}()
	<-streamDone

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
}

func (s *Server) createRecord(w http.ResponseWriter, r *http.Request) {
	record := IoT.Record{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.CreateRecord(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) updateRecord(w http.ResponseWriter, r *http.Request) {
	record := IoT.Record{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.UpdateRecord(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) deleteRecord(w http.ResponseWriter, r *http.Request) {
	record := IoT.Record{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.DeleteRecord(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getRecord(w http.ResponseWriter, r *http.Request) {
	record := IoT.Record{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.GetRecord(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getRecords(w http.ResponseWriter, r *http.Request) {
	device := IoT.Device{}

	if err := json.NewDecoder(r.Body).Decode(&device); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	stream, err := s.grpcClient.GetRecords(s.grpcCtx, &device)
	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	streamDone := make(chan bool)

	go func() {
		defer close(streamDone)
		header := false
		var streamEncoder *json.Encoder
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				streamDone <- true
				return
			}
			if err != nil {
				transport.SendError(w, transport.GrpcStatusToHTTPError(err))
				return
			}
			if !header {
				streamEncoder = transport.JsonEncoder(w, http.StatusOK)
				header = true
			}

			transport.SendJSONStream(streamEncoder, res)
		}

	}()
	<-streamDone

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
}

func (s *Server) createRecordData(w http.ResponseWriter, r *http.Request) {
	record := IoT.RecordData{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.CreateRecordData(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}

func (s *Server) getRecordData(w http.ResponseWriter, r *http.Request) {
	dataFilter := IoT.RecordDataFilter{}

	if err := json.NewDecoder(r.Body).Decode(&dataFilter); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	stream, err := s.grpcClient.GetRecordData(s.grpcCtx, &dataFilter)
	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}

	streamDone := make(chan bool)

	go func() {
		defer close(streamDone)
		header := false
		var streamEncoder *json.Encoder
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				streamDone <- true
				return
			}
			if err != nil {
				transport.SendError(w, transport.GrpcStatusToHTTPError(err))
				return
			}
			if !header {
				streamEncoder = transport.JsonEncoder(w, http.StatusOK)
				header = true
			}

			transport.SendJSONStream(streamEncoder, res)
		}

	}()
	<-streamDone

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
}

func (s *Server) deleteRecordData(w http.ResponseWriter, r *http.Request) {
	record := IoT.RecordData{}

	if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
		transport.SendError(w, models.InvalidJSONError{Message: err.Error()})
		return
	}

	resp, err := s.grpcClient.DeleteRecordData(s.grpcCtx, &record)

	if err != nil {
		transport.SendError(w, transport.GrpcStatusToHTTPError(err))
		return
	}
	transport.SendJSON(w, resp, http.StatusOK)
}
