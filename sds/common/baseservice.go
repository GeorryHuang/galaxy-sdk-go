// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package common

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
	"github.com/XiaoMi/galaxy-sdk-go/sds/errors"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = errors.GoUnusedProtection__
type BaseService interface {  //结构化存储基础接口

  // 获取服务端版本
  GetServerVersion() (r *Version, err error)
  // 检查版本兼容性
  // 
  // Parameters:
  //  - ClientVersion
  ValidateClientVersion(clientVersion *Version) (err error)
  // 获取服务器端当前时间，1970/0/0开始的秒数，可用作ping检查联通性
  GetServerTime() (r int64, err error)
}

//结构化存储基础接口
type BaseServiceClient struct {
  Transport thrift.TTransport
  ProtocolFactory thrift.TProtocolFactory
  InputProtocol thrift.TProtocol
  OutputProtocol thrift.TProtocol
  SeqId int32
}

func NewBaseServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *BaseServiceClient {
  return &BaseServiceClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewBaseServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *BaseServiceClient {
  return &BaseServiceClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// 获取服务端版本
func (p *BaseServiceClient) GetServerVersion() (r *Version, err error) {
  if err = p.sendGetServerVersion(); err != nil { return }
  return p.recvGetServerVersion()
}

func (p *BaseServiceClient) sendGetServerVersion()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getServerVersion", thrift.CALL, p.SeqId); err != nil {
    return
  }
  args := GetServerVersionArgs{
  }
  if err = args.Write(oprot); err != nil {
    return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
    return
  }
  return oprot.Flush()
}


func (p *BaseServiceClient) recvGetServerVersion() (value *Version, err error) {
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
    error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error1 error
    error1, err = error0.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error1
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getServerVersion failed: out of sequence response")
    return
  }
  result := GetServerVersionResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  if result.Se != nil {
  err = result.Se
  return 
  }
  value = result.GetSuccess()
  return
}

// 检查版本兼容性
// 
// Parameters:
//  - ClientVersion
func (p *BaseServiceClient) ValidateClientVersion(clientVersion *Version) (err error) {
  if err = p.sendValidateClientVersion(clientVersion); err != nil { return }
  return p.recvValidateClientVersion()
}

func (p *BaseServiceClient) sendValidateClientVersion(clientVersion *Version)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("validateClientVersion", thrift.CALL, p.SeqId); err != nil {
    return
  }
  args := ValidateClientVersionArgs{
  ClientVersion : clientVersion,
  }
  if err = args.Write(oprot); err != nil {
    return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
    return
  }
  return oprot.Flush()
}


func (p *BaseServiceClient) recvValidateClientVersion() (err error) {
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
    error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error3 error
    error3, err = error2.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error3
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "validateClientVersion failed: out of sequence response")
    return
  }
  result := ValidateClientVersionResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  if result.Se != nil {
  err = result.Se
  return 
  }
  return
}

// 获取服务器端当前时间，1970/0/0开始的秒数，可用作ping检查联通性
func (p *BaseServiceClient) GetServerTime() (r int64, err error) {
  if err = p.sendGetServerTime(); err != nil { return }
  return p.recvGetServerTime()
}

func (p *BaseServiceClient) sendGetServerTime()(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("getServerTime", thrift.CALL, p.SeqId); err != nil {
    return
  }
  args := GetServerTimeArgs{
  }
  if err = args.Write(oprot); err != nil {
    return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
    return
  }
  return oprot.Flush()
}


func (p *BaseServiceClient) recvGetServerTime() (value int64, err error) {
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
    error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error5 error
    error5, err = error4.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error5
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getServerTime failed: out of sequence response")
    return
  }
  result := GetServerTimeResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type BaseServiceProcessor struct {
  processorMap map[string]thrift.TProcessorFunction
  handler BaseService
}

func (p *BaseServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *BaseServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
  processor, ok = p.processorMap[key]
  return processor, ok
}

func (p *BaseServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
  return p.processorMap
}

func NewBaseServiceProcessor(handler BaseService) *BaseServiceProcessor {

  self6 := &BaseServiceProcessor{handler:handler, processorMap:make(map[string]thrift.TProcessorFunction)}
  self6.processorMap["getServerVersion"] = &baseServiceProcessorGetServerVersion{handler:handler}
  self6.processorMap["validateClientVersion"] = &baseServiceProcessorValidateClientVersion{handler:handler}
  self6.processorMap["getServerTime"] = &baseServiceProcessorGetServerTime{handler:handler}
return self6
}

func (p *BaseServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  name, _, seqId, err := iprot.ReadMessageBegin()
  if err != nil { return false, err }
  if processor, ok := p.GetProcessorFunction(name); ok {
    return processor.Process(seqId, iprot, oprot)
  }
  iprot.Skip(thrift.STRUCT)
  iprot.ReadMessageEnd()
  x7 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function " + name)
  oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
  x7.Write(oprot)
  oprot.WriteMessageEnd()
  oprot.Flush()
  return false, x7

}

type baseServiceProcessorGetServerVersion struct {
  handler BaseService
}

func (p *baseServiceProcessorGetServerVersion) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := GetServerVersionArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getServerVersion", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := GetServerVersionResult{}
var retval *Version
  var err2 error
  if retval, err2 = p.handler.GetServerVersion(); err2 != nil {
  switch v := err2.(type) {
    case *errors.ServiceException:
  result.Se = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getServerVersion: " + err2.Error())
    oprot.WriteMessageBegin("getServerVersion", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  } else {
    result.Success = retval
}
  if err2 = oprot.WriteMessageBegin("getServerVersion", thrift.REPLY, seqId); err2 != nil {
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

type baseServiceProcessorValidateClientVersion struct {
  handler BaseService
}

func (p *baseServiceProcessorValidateClientVersion) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := ValidateClientVersionArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("validateClientVersion", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := ValidateClientVersionResult{}
  var err2 error
  if err2 = p.handler.ValidateClientVersion(args.ClientVersion); err2 != nil {
  switch v := err2.(type) {
    case *errors.ServiceException:
  result.Se = v
    default:
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing validateClientVersion: " + err2.Error())
    oprot.WriteMessageBegin("validateClientVersion", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  }
  }
  if err2 = oprot.WriteMessageBegin("validateClientVersion", thrift.REPLY, seqId); err2 != nil {
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

type baseServiceProcessorGetServerTime struct {
  handler BaseService
}

func (p *baseServiceProcessorGetServerTime) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
  args := GetServerTimeArgs{}
  if err = args.Read(iprot); err != nil {
    iprot.ReadMessageEnd()
    x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
    oprot.WriteMessageBegin("getServerTime", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return false, err
  }

  iprot.ReadMessageEnd()
  result := GetServerTimeResult{}
var retval int64
  var err2 error
  if retval, err2 = p.handler.GetServerTime(); err2 != nil {
    x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing getServerTime: " + err2.Error())
    oprot.WriteMessageBegin("getServerTime", thrift.EXCEPTION, seqId)
    x.Write(oprot)
    oprot.WriteMessageEnd()
    oprot.Flush()
    return true, err2
  } else {
    result.Success = &retval
}
  if err2 = oprot.WriteMessageBegin("getServerTime", thrift.REPLY, seqId); err2 != nil {
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

type GetServerVersionArgs struct {
}

func NewGetServerVersionArgs() *GetServerVersionArgs {
  return &GetServerVersionArgs{}
}

func (p *GetServerVersionArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *GetServerVersionArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getServerVersion_args"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *GetServerVersionArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetServerVersionArgs(%+v)", *p)
}

type GetServerVersionResult struct {
  Success *Version `thrift:"success,0" json:"success"`
  Se *errors.ServiceException `thrift:"se,1" json:"se"`
}

func NewGetServerVersionResult() *GetServerVersionResult {
  return &GetServerVersionResult{}
}

var GetServerVersionResult_Success_DEFAULT *Version
func (p *GetServerVersionResult) GetSuccess() *Version {
  if !p.IsSetSuccess() {
    return GetServerVersionResult_Success_DEFAULT
  }
return p.Success
}
var GetServerVersionResult_Se_DEFAULT *errors.ServiceException
func (p *GetServerVersionResult) GetSe() *errors.ServiceException {
  if !p.IsSetSe() {
    return GetServerVersionResult_Se_DEFAULT
  }
return p.Se
}
func (p *GetServerVersionResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GetServerVersionResult) IsSetSe() bool {
  return p.Se != nil
}

func (p *GetServerVersionResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    case 1:
      if err := p.ReadField1(iprot); err != nil {
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

func (p *GetServerVersionResult)  ReadField0(iprot thrift.TProtocol) error {
  p.Success = &Version{
  Major: 1,

  Patch: "328cff44",
}
  if err := p.Success.Read(iprot); err != nil {
    return fmt.Errorf("%T error reading struct: %s", p.Success, err)
  }
  return nil
}

func (p *GetServerVersionResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Se = &errors.ServiceException{}
  if err := p.Se.Read(iprot); err != nil {
    return fmt.Errorf("%T error reading struct: %s", p.Se, err)
  }
  return nil
}

func (p *GetServerVersionResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getServerVersion_result"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *GetServerVersionResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return fmt.Errorf("%T write field begin error 0:success: %s", p, err); }
    if err := p.Success.Write(oprot); err != nil {
      return fmt.Errorf("%T error writing struct: %s", p.Success, err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 0:success: %s", p, err); }
  }
  return err
}

func (p *GetServerVersionResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetSe() {
    if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
      return fmt.Errorf("%T write field begin error 1:se: %s", p, err); }
    if err := p.Se.Write(oprot); err != nil {
      return fmt.Errorf("%T error writing struct: %s", p.Se, err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 1:se: %s", p, err); }
  }
  return err
}

func (p *GetServerVersionResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetServerVersionResult(%+v)", *p)
}

type ValidateClientVersionArgs struct {
  ClientVersion *Version `thrift:"clientVersion,1" json:"clientVersion"`
}

func NewValidateClientVersionArgs() *ValidateClientVersionArgs {
  return &ValidateClientVersionArgs{}
}

var ValidateClientVersionArgs_ClientVersion_DEFAULT *Version
func (p *ValidateClientVersionArgs) GetClientVersion() *Version {
  if !p.IsSetClientVersion() {
    return ValidateClientVersionArgs_ClientVersion_DEFAULT
  }
return p.ClientVersion
}
func (p *ValidateClientVersionArgs) IsSetClientVersion() bool {
  return p.ClientVersion != nil
}

func (p *ValidateClientVersionArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
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

func (p *ValidateClientVersionArgs)  ReadField1(iprot thrift.TProtocol) error {
  p.ClientVersion = &Version{
  Major: 1,

  Patch: "328cff44",
}
  if err := p.ClientVersion.Read(iprot); err != nil {
    return fmt.Errorf("%T error reading struct: %s", p.ClientVersion, err)
  }
  return nil
}

func (p *ValidateClientVersionArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("validateClientVersion_args"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *ValidateClientVersionArgs) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("clientVersion", thrift.STRUCT, 1); err != nil {
    return fmt.Errorf("%T write field begin error 1:clientVersion: %s", p, err); }
  if err := p.ClientVersion.Write(oprot); err != nil {
    return fmt.Errorf("%T error writing struct: %s", p.ClientVersion, err)
  }
  if err := oprot.WriteFieldEnd(); err != nil {
    return fmt.Errorf("%T write field end error 1:clientVersion: %s", p, err); }
  return err
}

func (p *ValidateClientVersionArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ValidateClientVersionArgs(%+v)", *p)
}

type ValidateClientVersionResult struct {
  Se *errors.ServiceException `thrift:"se,1" json:"se"`
}

func NewValidateClientVersionResult() *ValidateClientVersionResult {
  return &ValidateClientVersionResult{}
}

var ValidateClientVersionResult_Se_DEFAULT *errors.ServiceException
func (p *ValidateClientVersionResult) GetSe() *errors.ServiceException {
  if !p.IsSetSe() {
    return ValidateClientVersionResult_Se_DEFAULT
  }
return p.Se
}
func (p *ValidateClientVersionResult) IsSetSe() bool {
  return p.Se != nil
}

func (p *ValidateClientVersionResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
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

func (p *ValidateClientVersionResult)  ReadField1(iprot thrift.TProtocol) error {
  p.Se = &errors.ServiceException{}
  if err := p.Se.Read(iprot); err != nil {
    return fmt.Errorf("%T error reading struct: %s", p.Se, err)
  }
  return nil
}

func (p *ValidateClientVersionResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("validateClientVersion_result"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *ValidateClientVersionResult) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetSe() {
    if err := oprot.WriteFieldBegin("se", thrift.STRUCT, 1); err != nil {
      return fmt.Errorf("%T write field begin error 1:se: %s", p, err); }
    if err := p.Se.Write(oprot); err != nil {
      return fmt.Errorf("%T error writing struct: %s", p.Se, err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 1:se: %s", p, err); }
  }
  return err
}

func (p *ValidateClientVersionResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("ValidateClientVersionResult(%+v)", *p)
}

type GetServerTimeArgs struct {
}

func NewGetServerTimeArgs() *GetServerTimeArgs {
  return &GetServerTimeArgs{}
}

func (p *GetServerTimeArgs) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
    if err := iprot.Skip(fieldTypeId); err != nil {
      return err
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

func (p *GetServerTimeArgs) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getServerTime_args"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *GetServerTimeArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetServerTimeArgs(%+v)", *p)
}

type GetServerTimeResult struct {
  Success *int64 `thrift:"success,0" json:"success"`
}

func NewGetServerTimeResult() *GetServerTimeResult {
  return &GetServerTimeResult{}
}

var GetServerTimeResult_Success_DEFAULT int64
func (p *GetServerTimeResult) GetSuccess() int64 {
  if !p.IsSetSuccess() {
    return GetServerTimeResult_Success_DEFAULT
  }
return *p.Success
}
func (p *GetServerTimeResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *GetServerTimeResult) Read(iprot thrift.TProtocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return fmt.Errorf("%T read error: %s", p, err)
  }
  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
    }
    if fieldTypeId == thrift.STOP { break; }
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

func (p *GetServerTimeResult)  ReadField0(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI64(); err != nil {
  return fmt.Errorf("error reading field 0: %s", err)
} else {
  p.Success = &v
}
  return nil
}

func (p *GetServerTimeResult) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("getServerTime_result"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *GetServerTimeResult) writeField0(oprot thrift.TProtocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
      return fmt.Errorf("%T write field begin error 0:success: %s", p, err); }
    if err := oprot.WriteI64(int64(*p.Success)); err != nil {
    return fmt.Errorf("%T.success (0) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 0:success: %s", p, err); }
  }
  return err
}

func (p *GetServerTimeResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("GetServerTimeResult(%+v)", *p)
}


