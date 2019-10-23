// Code generated by capnpc-go. DO NOT EDIT.

package captain

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Auth struct{ capnp.Struct }

// Auth_TypeID is the unique identifier for the type Auth.
const Auth_TypeID = 0xd8cdf7e707cbaf27

func NewAuth(s *capnp.Segment) (Auth, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Auth{st}, err
}

func NewRootAuth(s *capnp.Segment) (Auth, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Auth{st}, err
}

func ReadRootAuth(msg *capnp.Message) (Auth, error) {
	root, err := msg.RootPtr()
	return Auth{root.Struct()}, err
}

func (s Auth) String() string {
	str, _ := text.Marshal(0xd8cdf7e707cbaf27, s.Struct)
	return str
}

func (s Auth) NodeId() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Auth) HasNodeId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Auth) NodeIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Auth) SetNodeId(v string) error {
	return s.Struct.SetText(0, v)
}

// Auth_List is a list of Auth.
type Auth_List struct{ capnp.List }

// NewAuth creates a new list of Auth.
func NewAuth_List(s *capnp.Segment, sz int32) (Auth_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Auth_List{l}, err
}

func (s Auth_List) At(i int) Auth { return Auth{s.List.Struct(i)} }

func (s Auth_List) Set(i int, v Auth) error { return s.List.SetStruct(i, v.Struct) }

func (s Auth_List) String() string {
	str, _ := text.MarshalList(0xd8cdf7e707cbaf27, s.List)
	return str
}

// Auth_Promise is a wrapper for a Auth promised by a client call.
type Auth_Promise struct{ *capnp.Pipeline }

func (p Auth_Promise) Struct() (Auth, error) {
	s, err := p.Pipeline.Struct()
	return Auth{s}, err
}

type Ping struct{ capnp.Struct }

// Ping_TypeID is the unique identifier for the type Ping.
const Ping_TypeID = 0x8b3b5f400419ca82

func NewPing(s *capnp.Segment) (Ping, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ping{st}, err
}

func NewRootPing(s *capnp.Segment) (Ping, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Ping{st}, err
}

func ReadRootPing(msg *capnp.Message) (Ping, error) {
	root, err := msg.RootPtr()
	return Ping{root.Struct()}, err
}

func (s Ping) String() string {
	str, _ := text.Marshal(0x8b3b5f400419ca82, s.Struct)
	return str
}

func (s Ping) Nonce() uint32 {
	return s.Struct.Uint32(0)
}

func (s Ping) SetNonce(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Ping_List is a list of Ping.
type Ping_List struct{ capnp.List }

// NewPing creates a new list of Ping.
func NewPing_List(s *capnp.Segment, sz int32) (Ping_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Ping_List{l}, err
}

func (s Ping_List) At(i int) Ping { return Ping{s.List.Struct(i)} }

func (s Ping_List) Set(i int, v Ping) error { return s.List.SetStruct(i, v.Struct) }

func (s Ping_List) String() string {
	str, _ := text.MarshalList(0x8b3b5f400419ca82, s.List)
	return str
}

// Ping_Promise is a wrapper for a Ping promised by a client call.
type Ping_Promise struct{ *capnp.Pipeline }

func (p Ping_Promise) Struct() (Ping, error) {
	s, err := p.Pipeline.Struct()
	return Ping{s}, err
}

type Pong struct{ capnp.Struct }

// Pong_TypeID is the unique identifier for the type Pong.
const Pong_TypeID = 0xc10601ff09f32611

func NewPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func NewRootPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func ReadRootPong(msg *capnp.Message) (Pong, error) {
	root, err := msg.RootPtr()
	return Pong{root.Struct()}, err
}

func (s Pong) String() string {
	str, _ := text.Marshal(0xc10601ff09f32611, s.Struct)
	return str
}

func (s Pong) Nonce() uint32 {
	return s.Struct.Uint32(0)
}

func (s Pong) SetNonce(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Pong_List is a list of Pong.
type Pong_List struct{ capnp.List }

// NewPong creates a new list of Pong.
func NewPong_List(s *capnp.Segment, sz int32) (Pong_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Pong_List{l}, err
}

func (s Pong_List) At(i int) Pong { return Pong{s.List.Struct(i)} }

func (s Pong_List) Set(i int, v Pong) error { return s.List.SetStruct(i, v.Struct) }

func (s Pong_List) String() string {
	str, _ := text.MarshalList(0xc10601ff09f32611, s.List)
	return str
}

// Pong_Promise is a wrapper for a Pong promised by a client call.
type Pong_Promise struct{ *capnp.Pipeline }

func (p Pong_Promise) Struct() (Pong, error) {
	s, err := p.Pipeline.Struct()
	return Pong{s}, err
}

type Announce struct{ capnp.Struct }

// Announce_TypeID is the unique identifier for the type Announce.
const Announce_TypeID = 0xda6b658bc18baf1d

func NewAnnounce(s *capnp.Segment) (Announce, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Announce{st}, err
}

func NewRootAnnounce(s *capnp.Segment) (Announce, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Announce{st}, err
}

func ReadRootAnnounce(msg *capnp.Message) (Announce, error) {
	root, err := msg.RootPtr()
	return Announce{root.Struct()}, err
}

func (s Announce) String() string {
	str, _ := text.Marshal(0xda6b658bc18baf1d, s.Struct)
	return str
}

func (s Announce) EngineId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Announce) SetEngineId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Announce) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Announce) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Announce) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Announce_List is a list of Announce.
type Announce_List struct{ capnp.List }

// NewAnnounce creates a new list of Announce.
func NewAnnounce_List(s *capnp.Segment, sz int32) (Announce_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Announce_List{l}, err
}

func (s Announce_List) At(i int) Announce { return Announce{s.List.Struct(i)} }

func (s Announce_List) Set(i int, v Announce) error { return s.List.SetStruct(i, v.Struct) }

func (s Announce_List) String() string {
	str, _ := text.MarshalList(0xda6b658bc18baf1d, s.List)
	return str
}

// Announce_Promise is a wrapper for a Announce promised by a client call.
type Announce_Promise struct{ *capnp.Pipeline }

func (p Announce_Promise) Struct() (Announce, error) {
	s, err := p.Pipeline.Struct()
	return Announce{s}, err
}

type Request struct{ capnp.Struct }

// Request_TypeID is the unique identifier for the type Request.
const Request_TypeID = 0xe526bff04d4ce1e9

func NewRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Request{st}, err
}

func NewRootRequest(s *capnp.Segment) (Request, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Request{st}, err
}

func ReadRootRequest(msg *capnp.Message) (Request, error) {
	root, err := msg.RootPtr()
	return Request{root.Struct()}, err
}

func (s Request) String() string {
	str, _ := text.Marshal(0xe526bff04d4ce1e9, s.Struct)
	return str
}

func (s Request) EngineId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Request) SetEngineId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Request) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Request) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Request) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

// Request_List is a list of Request.
type Request_List struct{ capnp.List }

// NewRequest creates a new list of Request.
func NewRequest_List(s *capnp.Segment, sz int32) (Request_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return Request_List{l}, err
}

func (s Request_List) At(i int) Request { return Request{s.List.Struct(i)} }

func (s Request_List) Set(i int, v Request) error { return s.List.SetStruct(i, v.Struct) }

func (s Request_List) String() string {
	str, _ := text.MarshalList(0xe526bff04d4ce1e9, s.List)
	return str
}

// Request_Promise is a wrapper for a Request promised by a client call.
type Request_Promise struct{ *capnp.Pipeline }

func (p Request_Promise) Struct() (Request, error) {
	s, err := p.Pipeline.Struct()
	return Request{s}, err
}

type Response struct{ capnp.Struct }

// Response_TypeID is the unique identifier for the type Response.
const Response_TypeID = 0xe06f2535c7dfe01b

func NewResponse(s *capnp.Segment) (Response, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Response{st}, err
}

func NewRootResponse(s *capnp.Segment) (Response, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4})
	return Response{st}, err
}

func ReadRootResponse(msg *capnp.Message) (Response, error) {
	root, err := msg.RootPtr()
	return Response{root.Struct()}, err
}

func (s Response) String() string {
	str, _ := text.Marshal(0xe06f2535c7dfe01b, s.Struct)
	return str
}

func (s Response) EngineId() uint8 {
	return s.Struct.Uint8(0)
}

func (s Response) SetEngineId(v uint8) {
	s.Struct.SetUint8(0, v)
}

func (s Response) EventId() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s Response) HasEventId() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Response) SetEventId(v []byte) error {
	return s.Struct.SetData(0, v)
}

func (s Response) OriginId() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Response) HasOriginId() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Response) OriginIdBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Response) SetOriginId(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Response) TargetIds() (capnp.TextList, error) {
	p, err := s.Struct.Ptr(2)
	return capnp.TextList{List: p.List()}, err
}

func (s Response) HasTargetIds() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Response) SetTargetIds(v capnp.TextList) error {
	return s.Struct.SetPtr(2, v.List.ToPtr())
}

// NewTargetIds sets the targetIds field to a newly
// allocated capnp.TextList, preferring placement in s's segment.
func (s Response) NewTargetIds(n int32) (capnp.TextList, error) {
	l, err := capnp.NewTextList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.TextList{}, err
	}
	err = s.Struct.SetPtr(2, l.List.ToPtr())
	return l, err
}

func (s Response) Payload() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Response) HasPayload() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Response) SetPayload(v []byte) error {
	return s.Struct.SetData(3, v)
}

// Response_List is a list of Response.
type Response_List struct{ capnp.List }

// NewResponse creates a new list of Response.
func NewResponse_List(s *capnp.Segment, sz int32) (Response_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 4}, sz)
	return Response_List{l}, err
}

func (s Response_List) At(i int) Response { return Response{s.List.Struct(i)} }

func (s Response_List) Set(i int, v Response) error { return s.List.SetStruct(i, v.Struct) }

func (s Response_List) String() string {
	str, _ := text.MarshalList(0xe06f2535c7dfe01b, s.List)
	return str
}

// Response_Promise is a wrapper for a Response promised by a client call.
type Response_Promise struct{ *capnp.Pipeline }

func (p Response_Promise) Struct() (Response, error) {
	s, err := p.Pipeline.Struct()
	return Response{s}, err
}

const schema_f71cc0af2f870b3c = "x\xda\xac\x90Oh\x13]\x14\xc5\xcfyw&\xcd\xc7" +
	"\xd7\x06\x86\x99\x8d\xa8teU\xb0j\xd4\x80T!\xd5" +
	"]\xa4B^\xc1\x85+\x09\xc9#\x86\x9671\x7f," +
	"\x05KE\x10+\x11\xac\xe0\xc6\x80\x0b\xc1\x8d \xa4\xb8" +
	"\xecF\xbb\x12\x14\xf7\xe2\xc6\xba\x10q'\xb8\xe8JF" +
	"f\x0a\xa9\xa4A\x04\xdd=f~\xe7\xfc\xee\xbd\xc7\xd7" +
	"8\xad\xb2\xee\xfd\x14\xa0g\xdcTt\xeb\xcd\x1eg\xfa" +
	"\xca\x99\x0et\x86\x8c\xce\xfe\x7f\xe7X\xef\xd5\xbe-8" +
	"#\x80?)]?'\xf1++y0\xf2&\xbe\xff" +
	"\x171\xb51\x8c\xd5\xd2\xf5/'\xec\xa5\x84=\xd8{" +
	";\xf2e\xeb\xdd{x\x99_P\x971\xd1\x96\xae\xbf" +
	"\x94\xb0\x8b\x09\xbb\xbf\xd7\xd9\xe8\x98\xb9\x0f\x03\xbd\xdb\xf0" +
	"#Y\xf7\x9f$\xf0cY\x00\xa3\xbd\x9b\x1f_\xe7\x0e" +
	"\x84\x9b\x83p2\xc5\x0fY\xf7\x93\xd7I:\xe3\x04\xa3" +
	"\xaf\x9ff.~{9\xf1yh\xf5\xa4\xfb\xc2\xcf\xb9" +
	"\xc9~\xee\x02V\xa3V\xa3V\x9e\x9b7GY.\xd5" +
	"m}\xaaX\x13[-\x92\xda\x11\x07p\x08xc'" +
	"\x00\x9d\x16\xea@q\xdc\x86\xb6l\x98\x86b\x1a\xdc\x15" +
	"\x0e\xff\"|\xae-\xad\xab\x03\xe1\xa9\x9dp\xde\x86\x15" +
	"S\xa8p\x14\x8a\xa3C\xd26o\xc3\xb6-\x9b\xb8!" +
	"\xddo8|\x01\xd0\x87\x84\xfa\x94\"\x190\xfe\x96=" +
	"\x0f\xe8#B}Z12\xb6Z\xb3\xa6P\x01\xc0\x14" +
	"\x14S\xe0\xb2\xb9nl\xabP\xe1\x18\x14\xc7v\xabf" +
	"M\xbeY\x0fm3Q\x05}\xd5R\xac\xba!\xd4+" +
	";\xaa\xdb\xb1\xea\xa6P\xdfS\xf4\x14\x03*\xc0\xbb\x1b" +
	"\x83+B\xfdP\xd1\x13\x15P\x00\xef\xc1,\xa0W\x85" +
	"\xfa\xb9\xa2\xe7H@\x07\xf0\x9e\xc5\xf1\xa7B\xbd\xf6\x87" +
	"\x93\x86\x8dZ\xb5f\xb7\x99\xfe\xa1J\x8d\xaai\x15*" +
	"`\x93\x19\xb0(L~e\xc0\xe5ziq>,\xfd" +
	"f\xd1\xf1km\xd3l\xfd\xfb\x93\xfe\x0c\x00\x00\xff\xff" +
	"\xe15\xc9\x97"

func init() {
	schemas.Register(schema_f71cc0af2f870b3c,
		0x8b3b5f400419ca82,
		0xc10601ff09f32611,
		0xd8cdf7e707cbaf27,
		0xda6b658bc18baf1d,
		0xe06f2535c7dfe01b,
		0xe526bff04d4ce1e9)
}
