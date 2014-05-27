// Code generated by protoc-gen-go.
// source: pb.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb.proto

It has these top-level messages:
	Error
	AttrRequest
	AttrResponse
	Attr
	ReaddirRequest
	ReaddirResponse
	DirEntry
	ReadlinkRequest
	ReadlinkResponse
	ChmodRequest
	ChmodResponse
	ChownRequest
	ChownResponse
	Time
	UtimeRequest
	UtimeResponse
	TruncateRequest
	TruncateResponse
	LinkRequest
	LinkResponse
	SymlinkRequest
	SymlinkResponse
	MkdirRequest
	MkdirResponse
	MknodRequest
	MknodResponse
	RenameRequest
	RenameResponse
	RmdirRequest
	RmdirResponse
	UnlinkRequest
	UnlinkResponse
	OpenRequest
	OpenResponse
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	WriteRequest
	WriteResponse
	CloseRequest
	CloseResponse
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Error struct {
	// Only one may be set.
	NotExist         *bool   `protobuf:"varint,1,opt,name=not_exist" json:"not_exist,omitempty"`
	ReadOnly         *bool   `protobuf:"varint,3,opt,name=read_only" json:"read_only,omitempty"`
	NotDir           *bool   `protobuf:"varint,4,opt,name=not_dir" json:"not_dir,omitempty"`
	Other            *string `protobuf:"bytes,2,opt,name=other" json:"other,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetNotExist() bool {
	if m != nil && m.NotExist != nil {
		return *m.NotExist
	}
	return false
}

func (m *Error) GetReadOnly() bool {
	if m != nil && m.ReadOnly != nil {
		return *m.ReadOnly
	}
	return false
}

func (m *Error) GetNotDir() bool {
	if m != nil && m.NotDir != nil {
		return *m.NotDir
	}
	return false
}

func (m *Error) GetOther() string {
	if m != nil && m.Other != nil {
		return *m.Other
	}
	return ""
}

type AttrRequest struct {
	// One of name or handle must be set:
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AttrRequest) Reset()         { *m = AttrRequest{} }
func (m *AttrRequest) String() string { return proto.CompactTextString(m) }
func (*AttrRequest) ProtoMessage()    {}

func (m *AttrRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *AttrRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

type AttrResponse struct {
	Attr             *Attr  `protobuf:"bytes,1,opt,name=attr" json:"attr,omitempty"`
	Err              *Error `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AttrResponse) Reset()         { *m = AttrResponse{} }
func (m *AttrResponse) String() string { return proto.CompactTextString(m) }
func (*AttrResponse) ProtoMessage()    {}

func (m *AttrResponse) GetAttr() *Attr {
	if m != nil {
		return m.Attr
	}
	return nil
}

func (m *AttrResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type Attr struct {
	Size             *uint64 `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	AtimeSec         *uint64 `protobuf:"varint,2,opt,name=atime_sec" json:"atime_sec,omitempty"`
	AtimeNano        *uint32 `protobuf:"varint,3,opt,name=atime_nano" json:"atime_nano,omitempty"`
	MtimeSec         *uint64 `protobuf:"varint,4,opt,name=mtime_sec" json:"mtime_sec,omitempty"`
	MtimeNano        *uint32 `protobuf:"varint,5,opt,name=mtime_nano" json:"mtime_nano,omitempty"`
	Mode             *uint32 `protobuf:"varint,6,opt,name=mode" json:"mode,omitempty"`
	Nlink            *uint32 `protobuf:"varint,7,opt,name=nlink" json:"nlink,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Attr) Reset()         { *m = Attr{} }
func (m *Attr) String() string { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()    {}

func (m *Attr) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *Attr) GetAtimeSec() uint64 {
	if m != nil && m.AtimeSec != nil {
		return *m.AtimeSec
	}
	return 0
}

func (m *Attr) GetAtimeNano() uint32 {
	if m != nil && m.AtimeNano != nil {
		return *m.AtimeNano
	}
	return 0
}

func (m *Attr) GetMtimeSec() uint64 {
	if m != nil && m.MtimeSec != nil {
		return *m.MtimeSec
	}
	return 0
}

func (m *Attr) GetMtimeNano() uint32 {
	if m != nil && m.MtimeNano != nil {
		return *m.MtimeNano
	}
	return 0
}

func (m *Attr) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

func (m *Attr) GetNlink() uint32 {
	if m != nil && m.Nlink != nil {
		return *m.Nlink
	}
	return 0
}

type ReaddirRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReaddirRequest) Reset()         { *m = ReaddirRequest{} }
func (m *ReaddirRequest) String() string { return proto.CompactTextString(m) }
func (*ReaddirRequest) ProtoMessage()    {}

func (m *ReaddirRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReaddirResponse struct {
	Err              *Error      `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Entry            []*DirEntry `protobuf:"bytes,2,rep,name=entry" json:"entry,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *ReaddirResponse) Reset()         { *m = ReaddirResponse{} }
func (m *ReaddirResponse) String() string { return proto.CompactTextString(m) }
func (*ReaddirResponse) ProtoMessage()    {}

func (m *ReaddirResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *ReaddirResponse) GetEntry() []*DirEntry {
	if m != nil {
		return m.Entry
	}
	return nil
}

type DirEntry struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Mode             *uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DirEntry) Reset()         { *m = DirEntry{} }
func (m *DirEntry) String() string { return proto.CompactTextString(m) }
func (*DirEntry) ProtoMessage()    {}

func (m *DirEntry) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *DirEntry) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

type ReadlinkRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReadlinkRequest) Reset()         { *m = ReadlinkRequest{} }
func (m *ReadlinkRequest) String() string { return proto.CompactTextString(m) }
func (*ReadlinkRequest) ProtoMessage()    {}

func (m *ReadlinkRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type ReadlinkResponse struct {
	Err              *Error  `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Target           *string `protobuf:"bytes,2,opt,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReadlinkResponse) Reset()         { *m = ReadlinkResponse{} }
func (m *ReadlinkResponse) String() string { return proto.CompactTextString(m) }
func (*ReadlinkResponse) ProtoMessage()    {}

func (m *ReadlinkResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *ReadlinkResponse) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

type ChmodRequest struct {
	// One of name or handle must be set:
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	Mode             *uint32 `protobuf:"varint,3,req,name=mode" json:"mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChmodRequest) Reset()         { *m = ChmodRequest{} }
func (m *ChmodRequest) String() string { return proto.CompactTextString(m) }
func (*ChmodRequest) ProtoMessage()    {}

func (m *ChmodRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChmodRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *ChmodRequest) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

type ChmodResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChmodResponse) Reset()         { *m = ChmodResponse{} }
func (m *ChmodResponse) String() string { return proto.CompactTextString(m) }
func (*ChmodResponse) ProtoMessage()    {}

func (m *ChmodResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type ChownRequest struct {
	// One of name or handle must be set:
	Name   *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Handle *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	// This can set either the userid or the groupid, or both,
	// depending on what's set. The server should send both the numeric
	// and named version of the user and/or group, for the client to
	// determine the mapping, since the two machines will likely have
	// different sets of users.
	Uid              *uint32 `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	Gid              *uint32 `protobuf:"varint,4,opt,name=gid" json:"gid,omitempty"`
	User             *string `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Group            *string `protobuf:"bytes,6,opt,name=group" json:"group,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChownRequest) Reset()         { *m = ChownRequest{} }
func (m *ChownRequest) String() string { return proto.CompactTextString(m) }
func (*ChownRequest) ProtoMessage()    {}

func (m *ChownRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChownRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *ChownRequest) GetUid() uint32 {
	if m != nil && m.Uid != nil {
		return *m.Uid
	}
	return 0
}

func (m *ChownRequest) GetGid() uint32 {
	if m != nil && m.Gid != nil {
		return *m.Gid
	}
	return 0
}

func (m *ChownRequest) GetUser() string {
	if m != nil && m.User != nil {
		return *m.User
	}
	return ""
}

func (m *ChownRequest) GetGroup() string {
	if m != nil && m.Group != nil {
		return *m.Group
	}
	return ""
}

type ChownResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ChownResponse) Reset()         { *m = ChownResponse{} }
func (m *ChownResponse) String() string { return proto.CompactTextString(m) }
func (*ChownResponse) ProtoMessage()    {}

func (m *ChownResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type Time struct {
	// Like a Go time.Time.
	Sec              *int64 `protobuf:"varint,1,req,name=sec" json:"sec,omitempty"`
	Nsec             *int32 `protobuf:"varint,2,opt,name=nsec" json:"nsec,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}

func (m *Time) GetSec() int64 {
	if m != nil && m.Sec != nil {
		return *m.Sec
	}
	return 0
}

func (m *Time) GetNsec() int32 {
	if m != nil && m.Nsec != nil {
		return *m.Nsec
	}
	return 0
}

type UtimeRequest struct {
	// One of name or handle must be set:
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	Atime            *Time   `protobuf:"bytes,3,opt,name=atime" json:"atime,omitempty"`
	Mtime            *Time   `protobuf:"bytes,4,opt,name=mtime" json:"mtime,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UtimeRequest) Reset()         { *m = UtimeRequest{} }
func (m *UtimeRequest) String() string { return proto.CompactTextString(m) }
func (*UtimeRequest) ProtoMessage()    {}

func (m *UtimeRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UtimeRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *UtimeRequest) GetAtime() *Time {
	if m != nil {
		return m.Atime
	}
	return nil
}

func (m *UtimeRequest) GetMtime() *Time {
	if m != nil {
		return m.Mtime
	}
	return nil
}

type UtimeResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *UtimeResponse) Reset()         { *m = UtimeResponse{} }
func (m *UtimeResponse) String() string { return proto.CompactTextString(m) }
func (*UtimeResponse) ProtoMessage()    {}

func (m *UtimeResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type TruncateRequest struct {
	// One of name or handle must be set:
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	Size             *uint64 `protobuf:"varint,3,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TruncateRequest) Reset()         { *m = TruncateRequest{} }
func (m *TruncateRequest) String() string { return proto.CompactTextString(m) }
func (*TruncateRequest) ProtoMessage()    {}

func (m *TruncateRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *TruncateRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *TruncateRequest) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type TruncateResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TruncateResponse) Reset()         { *m = TruncateResponse{} }
func (m *TruncateResponse) String() string { return proto.CompactTextString(m) }
func (*TruncateResponse) ProtoMessage()    {}

func (m *TruncateResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type LinkRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Target           *string `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *LinkRequest) Reset()         { *m = LinkRequest{} }
func (m *LinkRequest) String() string { return proto.CompactTextString(m) }
func (*LinkRequest) ProtoMessage()    {}

func (m *LinkRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *LinkRequest) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

type LinkResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *LinkResponse) Reset()         { *m = LinkResponse{} }
func (m *LinkResponse) String() string { return proto.CompactTextString(m) }
func (*LinkResponse) ProtoMessage()    {}

func (m *LinkResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type SymlinkRequest struct {
	Value            *string `protobuf:"bytes,1,req,name=value" json:"value,omitempty"`
	Name             *string `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SymlinkRequest) Reset()         { *m = SymlinkRequest{} }
func (m *SymlinkRequest) String() string { return proto.CompactTextString(m) }
func (*SymlinkRequest) ProtoMessage()    {}

func (m *SymlinkRequest) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *SymlinkRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type SymlinkResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *SymlinkResponse) Reset()         { *m = SymlinkResponse{} }
func (m *SymlinkResponse) String() string { return proto.CompactTextString(m) }
func (*SymlinkResponse) ProtoMessage()    {}

func (m *SymlinkResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type MkdirRequest struct {
	Name             *string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Mode             *uint32 `protobuf:"varint,2,opt,name=mode" json:"mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MkdirRequest) Reset()         { *m = MkdirRequest{} }
func (m *MkdirRequest) String() string { return proto.CompactTextString(m) }
func (*MkdirRequest) ProtoMessage()    {}

func (m *MkdirRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MkdirRequest) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

type MkdirResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MkdirResponse) Reset()         { *m = MkdirResponse{} }
func (m *MkdirResponse) String() string { return proto.CompactTextString(m) }
func (*MkdirResponse) ProtoMessage()    {}

func (m *MkdirResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type MknodRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Mode             *uint32 `protobuf:"varint,2,req,name=mode" json:"mode,omitempty"`
	Dev              *uint32 `protobuf:"varint,3,opt,name=dev" json:"dev,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MknodRequest) Reset()         { *m = MknodRequest{} }
func (m *MknodRequest) String() string { return proto.CompactTextString(m) }
func (*MknodRequest) ProtoMessage()    {}

func (m *MknodRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *MknodRequest) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

func (m *MknodRequest) GetDev() uint32 {
	if m != nil && m.Dev != nil {
		return *m.Dev
	}
	return 0
}

type MknodResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MknodResponse) Reset()         { *m = MknodResponse{} }
func (m *MknodResponse) String() string { return proto.CompactTextString(m) }
func (*MknodResponse) ProtoMessage()    {}

func (m *MknodResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type RenameRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Target           *string `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RenameRequest) Reset()         { *m = RenameRequest{} }
func (m *RenameRequest) String() string { return proto.CompactTextString(m) }
func (*RenameRequest) ProtoMessage()    {}

func (m *RenameRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *RenameRequest) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

type RenameResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RenameResponse) Reset()         { *m = RenameResponse{} }
func (m *RenameResponse) String() string { return proto.CompactTextString(m) }
func (*RenameResponse) ProtoMessage()    {}

func (m *RenameResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type RmdirRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RmdirRequest) Reset()         { *m = RmdirRequest{} }
func (m *RmdirRequest) String() string { return proto.CompactTextString(m) }
func (*RmdirRequest) ProtoMessage()    {}

func (m *RmdirRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type RmdirResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RmdirResponse) Reset()         { *m = RmdirResponse{} }
func (m *RmdirResponse) String() string { return proto.CompactTextString(m) }
func (*RmdirResponse) ProtoMessage()    {}

func (m *RmdirResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type UnlinkRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UnlinkRequest) Reset()         { *m = UnlinkRequest{} }
func (m *UnlinkRequest) String() string { return proto.CompactTextString(m) }
func (*UnlinkRequest) ProtoMessage()    {}

func (m *UnlinkRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type UnlinkResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *UnlinkResponse) Reset()         { *m = UnlinkResponse{} }
func (m *UnlinkResponse) String() string { return proto.CompactTextString(m) }
func (*UnlinkResponse) ProtoMessage()    {}

func (m *UnlinkResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type OpenRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Flags            *uint32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OpenRequest) Reset()         { *m = OpenRequest{} }
func (m *OpenRequest) String() string { return proto.CompactTextString(m) }
func (*OpenRequest) ProtoMessage()    {}

func (m *OpenRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *OpenRequest) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

type OpenResponse struct {
	Err              *Error  `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OpenResponse) Reset()         { *m = OpenResponse{} }
func (m *OpenResponse) String() string { return proto.CompactTextString(m) }
func (*OpenResponse) ProtoMessage()    {}

func (m *OpenResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *OpenResponse) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

type CreateRequest struct {
	Name             *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Flags            *uint32 `protobuf:"varint,2,opt,name=flags" json:"flags,omitempty"`
	Mode             *uint32 `protobuf:"varint,3,opt,name=mode" json:"mode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}

func (m *CreateRequest) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *CreateRequest) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func (m *CreateRequest) GetMode() uint32 {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return 0
}

type CreateResponse struct {
	Err              *Error  `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Handle           *uint64 `protobuf:"varint,2,opt,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}

func (m *CreateResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *CreateResponse) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

type ReadRequest struct {
	Handle           *uint64 `protobuf:"varint,1,req,name=handle" json:"handle,omitempty"`
	Offset           *uint64 `protobuf:"varint,2,req,name=offset" json:"offset,omitempty"`
	Size             *uint64 `protobuf:"varint,3,req,name=size" json:"size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}

func (m *ReadRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *ReadRequest) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *ReadRequest) GetSize() uint64 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

type ReadResponse struct {
	// One will be set:
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}

func (m *ReadResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *ReadResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteRequest struct {
	Handle           *uint64 `protobuf:"varint,1,req,name=handle" json:"handle,omitempty"`
	Offset           *uint64 `protobuf:"varint,2,req,name=offset" json:"offset,omitempty"`
	Data             []byte  `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

func (m *WriteRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

func (m *WriteRequest) GetOffset() uint64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *WriteRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type WriteResponse struct {
	Err              *Error  `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Written          *uint64 `protobuf:"varint,2,opt,name=written" json:"written,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}

func (m *WriteResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *WriteResponse) GetWritten() uint64 {
	if m != nil && m.Written != nil {
		return *m.Written
	}
	return 0
}

type CloseRequest struct {
	Handle           *uint64 `protobuf:"varint,1,req,name=handle" json:"handle,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}

func (m *CloseRequest) GetHandle() uint64 {
	if m != nil && m.Handle != nil {
		return *m.Handle
	}
	return 0
}

type CloseResponse struct {
	Err              *Error `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}

func (m *CloseResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func init() {
}
