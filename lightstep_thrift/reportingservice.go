// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package lightstep_thrift

import (
	"bytes"
	"fmt"
	"github.com/lightstep/lightstep-tracer-go/thrift_0_9_2/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ReportingService interface {
	// Parameters:
	//  - Auth
	//  - Request
	Report(auth *Auth, request *ReportRequest) (r *ReportResponse, err error)
}

type ReportingServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewReportingServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ReportingServiceClient {
	return &ReportingServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewReportingServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ReportingServiceClient {
	return &ReportingServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - Auth
//  - Request
func (p *ReportingServiceClient) Report(auth *Auth, request *ReportRequest) (r *ReportResponse, err error) {
	if err = p.sendReport(auth, request); err != nil {
		return
	}
	return p.recvReport()
}

func (p *ReportingServiceClient) sendReport(auth *Auth, request *ReportRequest) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Report", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ReportArgs{
		Auth:    auth,
		Request: request,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ReportingServiceClient) recvReport() (value *ReportResponse, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error13 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error14 error
		error14, err = error13.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error14
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Report failed: out of sequence response")
		return
	}
	result := ReportResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ReportingServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ReportingService
}

func (p *ReportingServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ReportingServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ReportingServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewReportingServiceProcessor(handler ReportingService) *ReportingServiceProcessor {

	self15 := &ReportingServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self15.processorMap["Report"] = &reportingServiceProcessorReport{handler: handler}
	return self15
}

func (p *ReportingServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x16 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x16.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x16

}

type reportingServiceProcessorReport struct {
	handler ReportingService
}

func (p *reportingServiceProcessorReport) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ReportArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Report", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ReportResult{}
	var retval *ReportResponse
	var err2 error
	if retval, err2 = p.handler.Report(args.Auth, args.Request); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Report: "+err2.Error())
		oprot.WriteMessageBegin("Report", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Report", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type ReportArgs struct {
	Auth    *Auth          `thrift:"auth,1" json:"auth"`
	Request *ReportRequest `thrift:"request,2" json:"request"`
}

func NewReportArgs() *ReportArgs {
	return &ReportArgs{}
}

var ReportArgs_Auth_DEFAULT *Auth

func (p *ReportArgs) GetAuth() *Auth {
	if !p.IsSetAuth() {
		return ReportArgs_Auth_DEFAULT
	}
	return p.Auth
}

var ReportArgs_Request_DEFAULT *ReportRequest

func (p *ReportArgs) GetRequest() *ReportRequest {
	if !p.IsSetRequest() {
		return ReportArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *ReportArgs) IsSetAuth() bool {
	return p.Auth != nil
}

func (p *ReportArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *ReportArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ReportArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Auth = &Auth{}
	if err := p.Auth.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Auth, err)
	}
	return nil
}

func (p *ReportArgs) ReadField2(iprot thrift.TProtocol) error {
	p.Request = &ReportRequest{}
	if err := p.Request.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Request, err)
	}
	return nil
}

func (p *ReportArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Report_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ReportArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("auth", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:auth: %s", p, err)
	}
	if err := p.Auth.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Auth, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:auth: %s", p, err)
	}
	return err
}

func (p *ReportArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:request: %s", p, err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Request, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:request: %s", p, err)
	}
	return err
}

func (p *ReportArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReportArgs(%+v)", *p)
}

type ReportResult struct {
	Success *ReportResponse `thrift:"success,0" json:"success"`
}

func NewReportResult() *ReportResult {
	return &ReportResult{}
}

var ReportResult_Success_DEFAULT *ReportResponse

func (p *ReportResult) GetSuccess() *ReportResponse {
	if !p.IsSetSuccess() {
		return ReportResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ReportResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ReportResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *ReportResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &ReportResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *ReportResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Report_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *ReportResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *ReportResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ReportResult(%+v)", *p)
}
