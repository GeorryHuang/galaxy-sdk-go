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
var GoUnusedProtection__ int;

//thrift传输协议
type ThriftProtocol int64
const (
  ThriftProtocol_TCOMPACT ThriftProtocol = 0
  ThriftProtocol_TJSON ThriftProtocol = 1
  ThriftProtocol_TBINARY ThriftProtocol = 2
  ThriftProtocol_TBINARYACCELERATED ThriftProtocol = 3
)

func (p ThriftProtocol) String() string {
  switch p {
  case ThriftProtocol_TCOMPACT: return "ThriftProtocol_TCOMPACT"
  case ThriftProtocol_TJSON: return "ThriftProtocol_TJSON"
  case ThriftProtocol_TBINARY: return "ThriftProtocol_TBINARY"
  case ThriftProtocol_TBINARYACCELERATED: return "ThriftProtocol_TBINARYACCELERATED"
  }
  return "<UNSET>"
}

func ThriftProtocolFromString(s string) (ThriftProtocol, error) {
  switch s {
  case "ThriftProtocol_TCOMPACT": return ThriftProtocol_TCOMPACT, nil 
  case "ThriftProtocol_TJSON": return ThriftProtocol_TJSON, nil 
  case "ThriftProtocol_TBINARY": return ThriftProtocol_TBINARY, nil 
  case "ThriftProtocol_TBINARYACCELERATED": return ThriftProtocol_TBINARYACCELERATED, nil 
  }
  return ThriftProtocol(0), fmt.Errorf("not a valid ThriftProtocol string")
}


func ThriftProtocolPtr(v ThriftProtocol) *ThriftProtocol { return &v }

type Version struct {
  Major int32 `thrift:"major,1" json:"major"`
  Minor int32 `thrift:"minor,2" json:"minor"`
  Patch string `thrift:"patch,3" json:"patch"`
  Comments string `thrift:"comments,4" json:"comments"`
}

func NewVersion() *Version {
  return &Version{
Major: 1,

Patch: "328cff44",
}
}

var Version_Major_DEFAULT int32 = 1

func (p *Version) GetMajor() int32 {
return p.Major
}
var Version_Minor_DEFAULT int32 = 0

func (p *Version) GetMinor() int32 {
return p.Minor
}
var Version_Patch_DEFAULT string = "328cff44"

func (p *Version) GetPatch() string {
return p.Patch
}
var Version_Comments_DEFAULT string = ""

func (p *Version) GetComments() string {
return p.Comments
}
func (p *Version) IsSetMajor() bool {
  return p.Major != Version_Major_DEFAULT
}

func (p *Version) IsSetMinor() bool {
  return p.Minor != Version_Minor_DEFAULT
}

func (p *Version) IsSetPatch() bool {
  return p.Patch != Version_Patch_DEFAULT
}

func (p *Version) IsSetComments() bool {
  return p.Comments != Version_Comments_DEFAULT
}

func (p *Version) Read(iprot thrift.TProtocol) error {
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
    case 2:
      if err := p.ReadField2(iprot); err != nil {
        return err
      }
    case 3:
      if err := p.ReadField3(iprot); err != nil {
        return err
      }
    case 4:
      if err := p.ReadField4(iprot); err != nil {
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

func (p *Version)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 1: %s", err)
} else {
  p.Major = v
}
  return nil
}

func (p *Version)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 2: %s", err)
} else {
  p.Minor = v
}
  return nil
}

func (p *Version)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 3: %s", err)
} else {
  p.Patch = v
}
  return nil
}

func (p *Version)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 4: %s", err)
} else {
  p.Comments = v
}
  return nil
}

func (p *Version) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Version"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *Version) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetMajor() {
    if err := oprot.WriteFieldBegin("major", thrift.I32, 1); err != nil {
      return fmt.Errorf("%T write field begin error 1:major: %s", p, err); }
    if err := oprot.WriteI32(int32(p.Major)); err != nil {
    return fmt.Errorf("%T.major (1) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 1:major: %s", p, err); }
  }
  return err
}

func (p *Version) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetMinor() {
    if err := oprot.WriteFieldBegin("minor", thrift.I32, 2); err != nil {
      return fmt.Errorf("%T write field begin error 2:minor: %s", p, err); }
    if err := oprot.WriteI32(int32(p.Minor)); err != nil {
    return fmt.Errorf("%T.minor (2) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 2:minor: %s", p, err); }
  }
  return err
}

func (p *Version) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetPatch() {
    if err := oprot.WriteFieldBegin("patch", thrift.STRING, 3); err != nil {
      return fmt.Errorf("%T write field begin error 3:patch: %s", p, err); }
    if err := oprot.WriteString(string(p.Patch)); err != nil {
    return fmt.Errorf("%T.patch (3) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 3:patch: %s", p, err); }
  }
  return err
}

func (p *Version) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetComments() {
    if err := oprot.WriteFieldBegin("comments", thrift.STRING, 4); err != nil {
      return fmt.Errorf("%T write field begin error 4:comments: %s", p, err); }
    if err := oprot.WriteString(string(p.Comments)); err != nil {
    return fmt.Errorf("%T.comments (4) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 4:comments: %s", p, err); }
  }
  return err
}

func (p *Version) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Version(%+v)", *p)
}

