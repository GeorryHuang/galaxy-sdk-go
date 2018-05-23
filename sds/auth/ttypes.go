// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package auth

import (
	"bytes"
	"fmt"
	"github.com/XiaoMi/galaxy-sdk-go/thrift"
	"github.com/XiaoMi/galaxy-sdk-go/sds/errors"
	"github.com/XiaoMi/galaxy-sdk-go/sds/common"

)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = errors.GoUnusedProtection__
var _ = common.GoUnusedProtection__
var GoUnusedProtection__ int;

//小米存储系统认证信息类型
type UserType int64
const (
  UserType_DEV_XIAOMI_SSO UserType = 1
  UserType_DEV_XIAOMI UserType = 2
  UserType_APP_SECRET UserType = 10
  UserType_APP_ACCESS_TOKEN UserType = 11
  UserType_APP_XIAOMI_SSO UserType = 12
  UserType_APP_ANONYMOUS UserType = 13
)

func (p UserType) String() string {
  switch p {
  case UserType_DEV_XIAOMI_SSO: return "UserType_DEV_XIAOMI_SSO"
  case UserType_DEV_XIAOMI: return "UserType_DEV_XIAOMI"
  case UserType_APP_SECRET: return "UserType_APP_SECRET"
  case UserType_APP_ACCESS_TOKEN: return "UserType_APP_ACCESS_TOKEN"
  case UserType_APP_XIAOMI_SSO: return "UserType_APP_XIAOMI_SSO"
  case UserType_APP_ANONYMOUS: return "UserType_APP_ANONYMOUS"
  }
  return "<UNSET>"
}

func UserTypeFromString(s string) (UserType, error) {
  switch s {
  case "UserType_DEV_XIAOMI_SSO": return UserType_DEV_XIAOMI_SSO, nil 
  case "UserType_DEV_XIAOMI": return UserType_DEV_XIAOMI, nil 
  case "UserType_APP_SECRET": return UserType_APP_SECRET, nil 
  case "UserType_APP_ACCESS_TOKEN": return UserType_APP_ACCESS_TOKEN, nil 
  case "UserType_APP_XIAOMI_SSO": return UserType_APP_XIAOMI_SSO, nil 
  case "UserType_APP_ANONYMOUS": return UserType_APP_ANONYMOUS, nil 
  }
  return UserType(0), fmt.Errorf("not a valid UserType string")
}


func UserTypePtr(v UserType) *UserType { return &v }

//签名使用的HMMAC算法
type MacAlgorithm int64
const (
  MacAlgorithm_HmacMD5 MacAlgorithm = 1
  MacAlgorithm_HmacSHA1 MacAlgorithm = 2
  MacAlgorithm_HmacSHA256 MacAlgorithm = 3
)

func (p MacAlgorithm) String() string {
  switch p {
  case MacAlgorithm_HmacMD5: return "MacAlgorithm_HmacMD5"
  case MacAlgorithm_HmacSHA1: return "MacAlgorithm_HmacSHA1"
  case MacAlgorithm_HmacSHA256: return "MacAlgorithm_HmacSHA256"
  }
  return "<UNSET>"
}

func MacAlgorithmFromString(s string) (MacAlgorithm, error) {
  switch s {
  case "MacAlgorithm_HmacMD5": return MacAlgorithm_HmacMD5, nil 
  case "MacAlgorithm_HmacSHA1": return MacAlgorithm_HmacSHA1, nil 
  case "MacAlgorithm_HmacSHA256": return MacAlgorithm_HmacSHA256, nil 
  }
  return MacAlgorithm(0), fmt.Errorf("not a valid MacAlgorithm string")
}


func MacAlgorithmPtr(v MacAlgorithm) *MacAlgorithm { return &v }

//第三方身份认证提供方，用于认证应用用户(非开发者)。
//目前提供小米SSO和几种常见OAuth系统
type AppUserAuthProvider int64
const (
  AppUserAuthProvider_XIAOMI_SSO AppUserAuthProvider = 1
  AppUserAuthProvider_XIAOMI_OAUTH AppUserAuthProvider = 2
  AppUserAuthProvider_QQ_OAUTH AppUserAuthProvider = 3
  AppUserAuthProvider_SINA_OAUTH AppUserAuthProvider = 4
  AppUserAuthProvider_RENREN_OAUTH AppUserAuthProvider = 5
  AppUserAuthProvider_WEIXIN_OAUTH AppUserAuthProvider = 6
)

func (p AppUserAuthProvider) String() string {
  switch p {
  case AppUserAuthProvider_XIAOMI_SSO: return "AppUserAuthProvider_XIAOMI_SSO"
  case AppUserAuthProvider_XIAOMI_OAUTH: return "AppUserAuthProvider_XIAOMI_OAUTH"
  case AppUserAuthProvider_QQ_OAUTH: return "AppUserAuthProvider_QQ_OAUTH"
  case AppUserAuthProvider_SINA_OAUTH: return "AppUserAuthProvider_SINA_OAUTH"
  case AppUserAuthProvider_RENREN_OAUTH: return "AppUserAuthProvider_RENREN_OAUTH"
  case AppUserAuthProvider_WEIXIN_OAUTH: return "AppUserAuthProvider_WEIXIN_OAUTH"
  }
  return "<UNSET>"
}

func AppUserAuthProviderFromString(s string) (AppUserAuthProvider, error) {
  switch s {
  case "AppUserAuthProvider_XIAOMI_SSO": return AppUserAuthProvider_XIAOMI_SSO, nil 
  case "AppUserAuthProvider_XIAOMI_OAUTH": return AppUserAuthProvider_XIAOMI_OAUTH, nil 
  case "AppUserAuthProvider_QQ_OAUTH": return AppUserAuthProvider_QQ_OAUTH, nil 
  case "AppUserAuthProvider_SINA_OAUTH": return AppUserAuthProvider_SINA_OAUTH, nil 
  case "AppUserAuthProvider_RENREN_OAUTH": return AppUserAuthProvider_RENREN_OAUTH, nil 
  case "AppUserAuthProvider_WEIXIN_OAUTH": return AppUserAuthProvider_WEIXIN_OAUTH, nil 
  }
  return AppUserAuthProvider(0), fmt.Errorf("not a valid AppUserAuthProvider string")
}


func AppUserAuthProviderPtr(v AppUserAuthProvider) *AppUserAuthProvider { return &v }

type Credential struct {
  TypeA1 *UserType `thrift:"type,1" json:"type"`
  SecretKeyId *string `thrift:"secretKeyId,2" json:"secretKeyId"`
  SecretKey *string `thrift:"secretKey,3" json:"secretKey"`
}

func NewCredential() *Credential {
  return &Credential{}
}

var Credential_TypeA1_DEFAULT UserType
func (p *Credential) GetTypeA1() UserType {
  if !p.IsSetTypeA1() {
    return Credential_TypeA1_DEFAULT
  }
return *p.TypeA1
}
var Credential_SecretKeyId_DEFAULT string
func (p *Credential) GetSecretKeyId() string {
  if !p.IsSetSecretKeyId() {
    return Credential_SecretKeyId_DEFAULT
  }
return *p.SecretKeyId
}
var Credential_SecretKey_DEFAULT string
func (p *Credential) GetSecretKey() string {
  if !p.IsSetSecretKey() {
    return Credential_SecretKey_DEFAULT
  }
return *p.SecretKey
}
func (p *Credential) IsSetTypeA1() bool {
  return p.TypeA1 != nil
}

func (p *Credential) IsSetSecretKeyId() bool {
  return p.SecretKeyId != nil
}

func (p *Credential) IsSetSecretKey() bool {
  return p.SecretKey != nil
}

func (p *Credential) Read(iprot thrift.TProtocol) error {
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

func (p *Credential)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 1: %s", err)
} else {
  temp := UserType(v)
  p.TypeA1 = &temp
}
  return nil
}

func (p *Credential)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 2: %s", err)
} else {
  p.SecretKeyId = &v
}
  return nil
}

func (p *Credential)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 3: %s", err)
} else {
  p.SecretKey = &v
}
  return nil
}

func (p *Credential) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("Credential"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *Credential) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetTypeA1() {
    if err := oprot.WriteFieldBegin("type", thrift.I32, 1); err != nil {
      return fmt.Errorf("%T write field begin error 1:type: %s", p, err); }
    if err := oprot.WriteI32(int32(*p.TypeA1)); err != nil {
    return fmt.Errorf("%T.type (1) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 1:type: %s", p, err); }
  }
  return err
}

func (p *Credential) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetSecretKeyId() {
    if err := oprot.WriteFieldBegin("secretKeyId", thrift.STRING, 2); err != nil {
      return fmt.Errorf("%T write field begin error 2:secretKeyId: %s", p, err); }
    if err := oprot.WriteString(string(*p.SecretKeyId)); err != nil {
    return fmt.Errorf("%T.secretKeyId (2) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 2:secretKeyId: %s", p, err); }
  }
  return err
}

func (p *Credential) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetSecretKey() {
    if err := oprot.WriteFieldBegin("secretKey", thrift.STRING, 3); err != nil {
      return fmt.Errorf("%T write field begin error 3:secretKey: %s", p, err); }
    if err := oprot.WriteString(string(*p.SecretKey)); err != nil {
    return fmt.Errorf("%T.secretKey (3) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 3:secretKey: %s", p, err); }
  }
  return err
}

func (p *Credential) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("Credential(%+v)", *p)
}

type HttpAuthorizationHeader struct {
  Version string `thrift:"version,1" json:"version"`
  UserType UserType `thrift:"userType,2" json:"userType"`
  SecretKeyId *string `thrift:"secretKeyId,3" json:"secretKeyId"`
  SecretKey *string `thrift:"secretKey,4" json:"secretKey"`
  Signature *string `thrift:"signature,5" json:"signature"`
  Algorithm *MacAlgorithm `thrift:"algorithm,6" json:"algorithm"`
  SignedHeaders *[]string `thrift:"signedHeaders,7" json:"signedHeaders"`
  SupportAccountKey bool `thrift:"supportAccountKey,8" json:"supportAccountKey"`
}

func NewHttpAuthorizationHeader() *HttpAuthorizationHeader {
  return &HttpAuthorizationHeader{
Version: "SDS-V1",

UserType: 13,
}
}

var HttpAuthorizationHeader_Version_DEFAULT string = "SDS-V1"

func (p *HttpAuthorizationHeader) GetVersion() string {
return p.Version
}
var HttpAuthorizationHeader_UserType_DEFAULT UserType = 13

func (p *HttpAuthorizationHeader) GetUserType() UserType {
return p.UserType
}
var HttpAuthorizationHeader_SecretKeyId_DEFAULT string
func (p *HttpAuthorizationHeader) GetSecretKeyId() string {
  if !p.IsSetSecretKeyId() {
    return HttpAuthorizationHeader_SecretKeyId_DEFAULT
  }
return *p.SecretKeyId
}
var HttpAuthorizationHeader_SecretKey_DEFAULT string
func (p *HttpAuthorizationHeader) GetSecretKey() string {
  if !p.IsSetSecretKey() {
    return HttpAuthorizationHeader_SecretKey_DEFAULT
  }
return *p.SecretKey
}
var HttpAuthorizationHeader_Signature_DEFAULT string
func (p *HttpAuthorizationHeader) GetSignature() string {
  if !p.IsSetSignature() {
    return HttpAuthorizationHeader_Signature_DEFAULT
  }
return *p.Signature
}
var HttpAuthorizationHeader_Algorithm_DEFAULT MacAlgorithm
func (p *HttpAuthorizationHeader) GetAlgorithm() MacAlgorithm {
  if !p.IsSetAlgorithm() {
    return HttpAuthorizationHeader_Algorithm_DEFAULT
  }
return *p.Algorithm
}
var HttpAuthorizationHeader_SignedHeaders_DEFAULT []string = []string{
}
func (p *HttpAuthorizationHeader) GetSignedHeaders() []string {
  if !p.IsSetSignedHeaders() {
    return HttpAuthorizationHeader_SignedHeaders_DEFAULT
  }
return *p.SignedHeaders
}
var HttpAuthorizationHeader_SupportAccountKey_DEFAULT bool = false

func (p *HttpAuthorizationHeader) GetSupportAccountKey() bool {
return p.SupportAccountKey
}
func (p *HttpAuthorizationHeader) IsSetVersion() bool {
  return p.Version != HttpAuthorizationHeader_Version_DEFAULT
}

func (p *HttpAuthorizationHeader) IsSetUserType() bool {
  return p.UserType != HttpAuthorizationHeader_UserType_DEFAULT
}

func (p *HttpAuthorizationHeader) IsSetSecretKeyId() bool {
  return p.SecretKeyId != nil
}

func (p *HttpAuthorizationHeader) IsSetSecretKey() bool {
  return p.SecretKey != nil
}

func (p *HttpAuthorizationHeader) IsSetSignature() bool {
  return p.Signature != nil
}

func (p *HttpAuthorizationHeader) IsSetAlgorithm() bool {
  return p.Algorithm != nil
}

func (p *HttpAuthorizationHeader) IsSetSignedHeaders() bool {
  return p.SignedHeaders != nil
}

func (p *HttpAuthorizationHeader) IsSetSupportAccountKey() bool {
  return p.SupportAccountKey != HttpAuthorizationHeader_SupportAccountKey_DEFAULT
}

func (p *HttpAuthorizationHeader) Read(iprot thrift.TProtocol) error {
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
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
        return err
      }
    case 7:
      if err := p.ReadField7(iprot); err != nil {
        return err
      }
    case 8:
      if err := p.ReadField8(iprot); err != nil {
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

func (p *HttpAuthorizationHeader)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 1: %s", err)
} else {
  p.Version = v
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 2: %s", err)
} else {
  temp := UserType(v)
  p.UserType = temp
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 3: %s", err)
} else {
  p.SecretKeyId = &v
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 4: %s", err)
} else {
  p.SecretKey = &v
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 5: %s", err)
} else {
  p.Signature = &v
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 6: %s", err)
} else {
  temp := MacAlgorithm(v)
  p.Algorithm = &temp
}
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField7(iprot thrift.TProtocol) error {
  _, size, err := iprot.ReadListBegin()
  if err != nil {
    return fmt.Errorf("error reading list begin: %s", err)
  }
  tSlice := make([]string, 0, size)
  p.SignedHeaders =  &tSlice
  for i := 0; i < size; i ++ {
var _elem0 string
    if v, err := iprot.ReadString(); err != nil {
    return fmt.Errorf("error reading field 0: %s", err)
} else {
    _elem0 = v
}
    (*p.SignedHeaders) = append((*p.SignedHeaders), _elem0)
  }
  if err := iprot.ReadListEnd(); err != nil {
    return fmt.Errorf("error reading list end: %s", err)
  }
  return nil
}

func (p *HttpAuthorizationHeader)  ReadField8(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadBool(); err != nil {
  return fmt.Errorf("error reading field 8: %s", err)
} else {
  p.SupportAccountKey = v
}
  return nil
}

func (p *HttpAuthorizationHeader) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("HttpAuthorizationHeader"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := p.writeField5(oprot); err != nil { return err }
  if err := p.writeField6(oprot); err != nil { return err }
  if err := p.writeField7(oprot); err != nil { return err }
  if err := p.writeField8(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *HttpAuthorizationHeader) writeField1(oprot thrift.TProtocol) (err error) {
  if p.IsSetVersion() {
    if err := oprot.WriteFieldBegin("version", thrift.STRING, 1); err != nil {
      return fmt.Errorf("%T write field begin error 1:version: %s", p, err); }
    if err := oprot.WriteString(string(p.Version)); err != nil {
    return fmt.Errorf("%T.version (1) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 1:version: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField2(oprot thrift.TProtocol) (err error) {
  if p.IsSetUserType() {
    if err := oprot.WriteFieldBegin("userType", thrift.I32, 2); err != nil {
      return fmt.Errorf("%T write field begin error 2:userType: %s", p, err); }
    if err := oprot.WriteI32(int32(p.UserType)); err != nil {
    return fmt.Errorf("%T.userType (2) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 2:userType: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetSecretKeyId() {
    if err := oprot.WriteFieldBegin("secretKeyId", thrift.STRING, 3); err != nil {
      return fmt.Errorf("%T write field begin error 3:secretKeyId: %s", p, err); }
    if err := oprot.WriteString(string(*p.SecretKeyId)); err != nil {
    return fmt.Errorf("%T.secretKeyId (3) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 3:secretKeyId: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetSecretKey() {
    if err := oprot.WriteFieldBegin("secretKey", thrift.STRING, 4); err != nil {
      return fmt.Errorf("%T write field begin error 4:secretKey: %s", p, err); }
    if err := oprot.WriteString(string(*p.SecretKey)); err != nil {
    return fmt.Errorf("%T.secretKey (4) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 4:secretKey: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetSignature() {
    if err := oprot.WriteFieldBegin("signature", thrift.STRING, 5); err != nil {
      return fmt.Errorf("%T write field begin error 5:signature: %s", p, err); }
    if err := oprot.WriteString(string(*p.Signature)); err != nil {
    return fmt.Errorf("%T.signature (5) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 5:signature: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField6(oprot thrift.TProtocol) (err error) {
  if p.IsSetAlgorithm() {
    if err := oprot.WriteFieldBegin("algorithm", thrift.I32, 6); err != nil {
      return fmt.Errorf("%T write field begin error 6:algorithm: %s", p, err); }
    if err := oprot.WriteI32(int32(*p.Algorithm)); err != nil {
    return fmt.Errorf("%T.algorithm (6) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 6:algorithm: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField7(oprot thrift.TProtocol) (err error) {
  if p.IsSetSignedHeaders() {
    if err := oprot.WriteFieldBegin("signedHeaders", thrift.LIST, 7); err != nil {
      return fmt.Errorf("%T write field begin error 7:signedHeaders: %s", p, err); }
    if err := oprot.WriteListBegin(thrift.STRING, len(*p.SignedHeaders)); err != nil {
      return fmt.Errorf("error writing list begin: %s", err)
    }
    for _, v := range *p.SignedHeaders {
      if err := oprot.WriteString(string(v)); err != nil {
      return fmt.Errorf("%T. (0) field write error: %s", p, err) }
    }
    if err := oprot.WriteListEnd(); err != nil {
      return fmt.Errorf("error writing list end: %s", err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 7:signedHeaders: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) writeField8(oprot thrift.TProtocol) (err error) {
  if p.IsSetSupportAccountKey() {
    if err := oprot.WriteFieldBegin("supportAccountKey", thrift.BOOL, 8); err != nil {
      return fmt.Errorf("%T write field begin error 8:supportAccountKey: %s", p, err); }
    if err := oprot.WriteBool(bool(p.SupportAccountKey)); err != nil {
    return fmt.Errorf("%T.supportAccountKey (8) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 8:supportAccountKey: %s", p, err); }
  }
  return err
}

func (p *HttpAuthorizationHeader) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("HttpAuthorizationHeader(%+v)", *p)
}

type OAuthInfo struct {
  XiaomiAppId string `thrift:"xiaomiAppId,1" json:"xiaomiAppId"`
  AppUserAuthProvider AppUserAuthProvider `thrift:"appUserAuthProvider,2" json:"appUserAuthProvider"`
  AccessToken *string `thrift:"accessToken,3" json:"accessToken"`
  OpenId *string `thrift:"openId,4" json:"openId"`
  MacKey *string `thrift:"macKey,5" json:"macKey"`
  MacAlgorithm *string `thrift:"macAlgorithm,6" json:"macAlgorithm"`
}

func NewOAuthInfo() *OAuthInfo {
  return &OAuthInfo{}
}


func (p *OAuthInfo) GetXiaomiAppId() string {
return p.XiaomiAppId
}

func (p *OAuthInfo) GetAppUserAuthProvider() AppUserAuthProvider {
return p.AppUserAuthProvider
}
var OAuthInfo_AccessToken_DEFAULT string
func (p *OAuthInfo) GetAccessToken() string {
  if !p.IsSetAccessToken() {
    return OAuthInfo_AccessToken_DEFAULT
  }
return *p.AccessToken
}
var OAuthInfo_OpenId_DEFAULT string
func (p *OAuthInfo) GetOpenId() string {
  if !p.IsSetOpenId() {
    return OAuthInfo_OpenId_DEFAULT
  }
return *p.OpenId
}
var OAuthInfo_MacKey_DEFAULT string
func (p *OAuthInfo) GetMacKey() string {
  if !p.IsSetMacKey() {
    return OAuthInfo_MacKey_DEFAULT
  }
return *p.MacKey
}
var OAuthInfo_MacAlgorithm_DEFAULT string
func (p *OAuthInfo) GetMacAlgorithm() string {
  if !p.IsSetMacAlgorithm() {
    return OAuthInfo_MacAlgorithm_DEFAULT
  }
return *p.MacAlgorithm
}
func (p *OAuthInfo) IsSetAccessToken() bool {
  return p.AccessToken != nil
}

func (p *OAuthInfo) IsSetOpenId() bool {
  return p.OpenId != nil
}

func (p *OAuthInfo) IsSetMacKey() bool {
  return p.MacKey != nil
}

func (p *OAuthInfo) IsSetMacAlgorithm() bool {
  return p.MacAlgorithm != nil
}

func (p *OAuthInfo) Read(iprot thrift.TProtocol) error {
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
    case 5:
      if err := p.ReadField5(iprot); err != nil {
        return err
      }
    case 6:
      if err := p.ReadField6(iprot); err != nil {
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

func (p *OAuthInfo)  ReadField1(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 1: %s", err)
} else {
  p.XiaomiAppId = v
}
  return nil
}

func (p *OAuthInfo)  ReadField2(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadI32(); err != nil {
  return fmt.Errorf("error reading field 2: %s", err)
} else {
  temp := AppUserAuthProvider(v)
  p.AppUserAuthProvider = temp
}
  return nil
}

func (p *OAuthInfo)  ReadField3(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 3: %s", err)
} else {
  p.AccessToken = &v
}
  return nil
}

func (p *OAuthInfo)  ReadField4(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 4: %s", err)
} else {
  p.OpenId = &v
}
  return nil
}

func (p *OAuthInfo)  ReadField5(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 5: %s", err)
} else {
  p.MacKey = &v
}
  return nil
}

func (p *OAuthInfo)  ReadField6(iprot thrift.TProtocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return fmt.Errorf("error reading field 6: %s", err)
} else {
  p.MacAlgorithm = &v
}
  return nil
}

func (p *OAuthInfo) Write(oprot thrift.TProtocol) error {
  if err := oprot.WriteStructBegin("OAuthInfo"); err != nil {
    return fmt.Errorf("%T write struct begin error: %s", p, err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := p.writeField2(oprot); err != nil { return err }
  if err := p.writeField3(oprot); err != nil { return err }
  if err := p.writeField4(oprot); err != nil { return err }
  if err := p.writeField5(oprot); err != nil { return err }
  if err := p.writeField6(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return fmt.Errorf("write field stop error: %s", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return fmt.Errorf("write struct stop error: %s", err) }
  return nil
}

func (p *OAuthInfo) writeField1(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("xiaomiAppId", thrift.STRING, 1); err != nil {
    return fmt.Errorf("%T write field begin error 1:xiaomiAppId: %s", p, err); }
  if err := oprot.WriteString(string(p.XiaomiAppId)); err != nil {
  return fmt.Errorf("%T.xiaomiAppId (1) field write error: %s", p, err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return fmt.Errorf("%T write field end error 1:xiaomiAppId: %s", p, err); }
  return err
}

func (p *OAuthInfo) writeField2(oprot thrift.TProtocol) (err error) {
  if err := oprot.WriteFieldBegin("appUserAuthProvider", thrift.I32, 2); err != nil {
    return fmt.Errorf("%T write field begin error 2:appUserAuthProvider: %s", p, err); }
  if err := oprot.WriteI32(int32(p.AppUserAuthProvider)); err != nil {
  return fmt.Errorf("%T.appUserAuthProvider (2) field write error: %s", p, err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return fmt.Errorf("%T write field end error 2:appUserAuthProvider: %s", p, err); }
  return err
}

func (p *OAuthInfo) writeField3(oprot thrift.TProtocol) (err error) {
  if p.IsSetAccessToken() {
    if err := oprot.WriteFieldBegin("accessToken", thrift.STRING, 3); err != nil {
      return fmt.Errorf("%T write field begin error 3:accessToken: %s", p, err); }
    if err := oprot.WriteString(string(*p.AccessToken)); err != nil {
    return fmt.Errorf("%T.accessToken (3) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 3:accessToken: %s", p, err); }
  }
  return err
}

func (p *OAuthInfo) writeField4(oprot thrift.TProtocol) (err error) {
  if p.IsSetOpenId() {
    if err := oprot.WriteFieldBegin("openId", thrift.STRING, 4); err != nil {
      return fmt.Errorf("%T write field begin error 4:openId: %s", p, err); }
    if err := oprot.WriteString(string(*p.OpenId)); err != nil {
    return fmt.Errorf("%T.openId (4) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 4:openId: %s", p, err); }
  }
  return err
}

func (p *OAuthInfo) writeField5(oprot thrift.TProtocol) (err error) {
  if p.IsSetMacKey() {
    if err := oprot.WriteFieldBegin("macKey", thrift.STRING, 5); err != nil {
      return fmt.Errorf("%T write field begin error 5:macKey: %s", p, err); }
    if err := oprot.WriteString(string(*p.MacKey)); err != nil {
    return fmt.Errorf("%T.macKey (5) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 5:macKey: %s", p, err); }
  }
  return err
}

func (p *OAuthInfo) writeField6(oprot thrift.TProtocol) (err error) {
  if p.IsSetMacAlgorithm() {
    if err := oprot.WriteFieldBegin("macAlgorithm", thrift.STRING, 6); err != nil {
      return fmt.Errorf("%T write field begin error 6:macAlgorithm: %s", p, err); }
    if err := oprot.WriteString(string(*p.MacAlgorithm)); err != nil {
    return fmt.Errorf("%T.macAlgorithm (6) field write error: %s", p, err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return fmt.Errorf("%T write field end error 6:macAlgorithm: %s", p, err); }
  }
  return err
}

func (p *OAuthInfo) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("OAuthInfo(%+v)", *p)
}

