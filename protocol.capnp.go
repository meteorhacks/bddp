package bddp

// AUTO GENERATED - DO NOT EDIT

import (
	C "github.com/glycerine/go-capnproto"
)

type Message C.Struct
type Message_Which uint16

const (
	MESSAGE_CONNECT     Message_Which = 0
	MESSAGE_CONNECTED   Message_Which = 1
	MESSAGE_FAILED      Message_Which = 2
	MESSAGE_PING        Message_Which = 3
	MESSAGE_PONG        Message_Which = 4
	MESSAGE_SUB         Message_Which = 5
	MESSAGE_UNSUB       Message_Which = 6
	MESSAGE_NOSUB       Message_Which = 7
	MESSAGE_ADDED       Message_Which = 8
	MESSAGE_CHANGED     Message_Which = 9
	MESSAGE_REMOVED     Message_Which = 10
	MESSAGE_READY       Message_Which = 11
	MESSAGE_ADDEDBEFORE Message_Which = 12
	MESSAGE_MOVEDBEFORE Message_Which = 13
	MESSAGE_METHOD      Message_Which = 14
	MESSAGE_RESULT      Message_Which = 15
	MESSAGE_UPDATED     Message_Which = 16
)

func NewMessage(s *C.Segment) Message      { return Message(s.NewStruct(8, 1)) }
func NewRootMessage(s *C.Segment) Message  { return Message(s.NewRootStruct(8, 1)) }
func AutoNewMessage(s *C.Segment) Message  { return Message(s.NewStructAR(8, 1)) }
func ReadRootMessage(s *C.Segment) Message { return Message(s.Root(0).ToStruct()) }
func (s Message) Which() Message_Which     { return Message_Which(C.Struct(s).Get16(0)) }
func (s Message) Connect() ConnectMsg      { return ConnectMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetConnect(v ConnectMsg) {
	C.Struct(s).Set16(0, 0)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Connected() ConnectedMsg { return ConnectedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetConnected(v ConnectedMsg) {
	C.Struct(s).Set16(0, 1)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Failed() FailedMsg { return FailedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetFailed(v FailedMsg) {
	C.Struct(s).Set16(0, 2)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Ping() PingMsg       { return PingMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetPing(v PingMsg)   { C.Struct(s).Set16(0, 3); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Pong() PongMsg       { return PongMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetPong(v PongMsg)   { C.Struct(s).Set16(0, 4); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Sub() SubMsg         { return SubMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetSub(v SubMsg)     { C.Struct(s).Set16(0, 5); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Unsub() UnsubMsg     { return UnsubMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetUnsub(v UnsubMsg) { C.Struct(s).Set16(0, 6); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Nosub() NosubMsg     { return NosubMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetNosub(v NosubMsg) { C.Struct(s).Set16(0, 7); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Added() AddedMsg     { return AddedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetAdded(v AddedMsg) { C.Struct(s).Set16(0, 8); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Changed() ChangedMsg { return ChangedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetChanged(v ChangedMsg) {
	C.Struct(s).Set16(0, 9)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Removed() RemovedMsg { return RemovedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetRemoved(v RemovedMsg) {
	C.Struct(s).Set16(0, 10)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Ready() ReadyMsg     { return ReadyMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetReady(v ReadyMsg) { C.Struct(s).Set16(0, 11); C.Struct(s).SetObject(0, C.Object(v)) }
func (s Message) Addedbefore() AddedBeforeMsg {
	return AddedBeforeMsg(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetAddedbefore(v AddedBeforeMsg) {
	C.Struct(s).Set16(0, 12)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Movedbefore() MovedBeforeMsg {
	return MovedBeforeMsg(C.Struct(s).GetObject(0).ToStruct())
}
func (s Message) SetMovedbefore(v MovedBeforeMsg) {
	C.Struct(s).Set16(0, 13)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Method() MethodMsg { return MethodMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetMethod(v MethodMsg) {
	C.Struct(s).Set16(0, 14)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Result() ResultMsg { return ResultMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetResult(v ResultMsg) {
	C.Struct(s).Set16(0, 15)
	C.Struct(s).SetObject(0, C.Object(v))
}
func (s Message) Updated() UpdatedMsg { return UpdatedMsg(C.Struct(s).GetObject(0).ToStruct()) }
func (s Message) SetUpdated(v UpdatedMsg) {
	C.Struct(s).Set16(0, 16)
	C.Struct(s).SetObject(0, C.Object(v))
}

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Message) MarshalJSON() (bs []byte, err error) { return }

type Message_List C.PointerList

func NewMessageList(s *C.Segment, sz int) Message_List {
	return Message_List(s.NewCompositeList(8, 1, sz))
}
func (s Message_List) Len() int         { return C.PointerList(s).Len() }
func (s Message_List) At(i int) Message { return Message(C.PointerList(s).At(i).ToStruct()) }
func (s Message_List) ToArray() []Message {
	n := s.Len()
	a := make([]Message, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Message_List) Set(i int, item Message) { C.PointerList(s).Set(i, C.Object(item)) }

type ConnectMsg C.Struct

func NewConnectMsg(s *C.Segment) ConnectMsg      { return ConnectMsg(s.NewStruct(0, 3)) }
func NewRootConnectMsg(s *C.Segment) ConnectMsg  { return ConnectMsg(s.NewRootStruct(0, 3)) }
func AutoNewConnectMsg(s *C.Segment) ConnectMsg  { return ConnectMsg(s.NewStructAR(0, 3)) }
func ReadRootConnectMsg(s *C.Segment) ConnectMsg { return ConnectMsg(s.Root(0).ToStruct()) }
func (s ConnectMsg) Session() string             { return C.Struct(s).GetObject(0).ToText() }
func (s ConnectMsg) SetSession(v string)         { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ConnectMsg) Version() string             { return C.Struct(s).GetObject(1).ToText() }
func (s ConnectMsg) SetVersion(v string)         { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s ConnectMsg) Support() C.TextList         { return C.TextList(C.Struct(s).GetObject(2)) }
func (s ConnectMsg) SetSupport(v C.TextList)     { C.Struct(s).SetObject(2, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ConnectMsg) MarshalJSON() (bs []byte, err error) { return }

type ConnectMsg_List C.PointerList

func NewConnectMsgList(s *C.Segment, sz int) ConnectMsg_List {
	return ConnectMsg_List(s.NewCompositeList(0, 3, sz))
}
func (s ConnectMsg_List) Len() int            { return C.PointerList(s).Len() }
func (s ConnectMsg_List) At(i int) ConnectMsg { return ConnectMsg(C.PointerList(s).At(i).ToStruct()) }
func (s ConnectMsg_List) ToArray() []ConnectMsg {
	n := s.Len()
	a := make([]ConnectMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ConnectMsg_List) Set(i int, item ConnectMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type ConnectedMsg C.Struct

func NewConnectedMsg(s *C.Segment) ConnectedMsg      { return ConnectedMsg(s.NewStruct(0, 1)) }
func NewRootConnectedMsg(s *C.Segment) ConnectedMsg  { return ConnectedMsg(s.NewRootStruct(0, 1)) }
func AutoNewConnectedMsg(s *C.Segment) ConnectedMsg  { return ConnectedMsg(s.NewStructAR(0, 1)) }
func ReadRootConnectedMsg(s *C.Segment) ConnectedMsg { return ConnectedMsg(s.Root(0).ToStruct()) }
func (s ConnectedMsg) Session() string               { return C.Struct(s).GetObject(0).ToText() }
func (s ConnectedMsg) SetSession(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ConnectedMsg) MarshalJSON() (bs []byte, err error) { return }

type ConnectedMsg_List C.PointerList

func NewConnectedMsgList(s *C.Segment, sz int) ConnectedMsg_List {
	return ConnectedMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s ConnectedMsg_List) Len() int { return C.PointerList(s).Len() }
func (s ConnectedMsg_List) At(i int) ConnectedMsg {
	return ConnectedMsg(C.PointerList(s).At(i).ToStruct())
}
func (s ConnectedMsg_List) ToArray() []ConnectedMsg {
	n := s.Len()
	a := make([]ConnectedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ConnectedMsg_List) Set(i int, item ConnectedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type FailedMsg C.Struct

func NewFailedMsg(s *C.Segment) FailedMsg      { return FailedMsg(s.NewStruct(0, 1)) }
func NewRootFailedMsg(s *C.Segment) FailedMsg  { return FailedMsg(s.NewRootStruct(0, 1)) }
func AutoNewFailedMsg(s *C.Segment) FailedMsg  { return FailedMsg(s.NewStructAR(0, 1)) }
func ReadRootFailedMsg(s *C.Segment) FailedMsg { return FailedMsg(s.Root(0).ToStruct()) }
func (s FailedMsg) Version() string            { return C.Struct(s).GetObject(0).ToText() }
func (s FailedMsg) SetVersion(v string)        { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s FailedMsg) MarshalJSON() (bs []byte, err error) { return }

type FailedMsg_List C.PointerList

func NewFailedMsgList(s *C.Segment, sz int) FailedMsg_List {
	return FailedMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s FailedMsg_List) Len() int           { return C.PointerList(s).Len() }
func (s FailedMsg_List) At(i int) FailedMsg { return FailedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s FailedMsg_List) ToArray() []FailedMsg {
	n := s.Len()
	a := make([]FailedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s FailedMsg_List) Set(i int, item FailedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type PingMsg C.Struct

func NewPingMsg(s *C.Segment) PingMsg      { return PingMsg(s.NewStruct(0, 1)) }
func NewRootPingMsg(s *C.Segment) PingMsg  { return PingMsg(s.NewRootStruct(0, 1)) }
func AutoNewPingMsg(s *C.Segment) PingMsg  { return PingMsg(s.NewStructAR(0, 1)) }
func ReadRootPingMsg(s *C.Segment) PingMsg { return PingMsg(s.Root(0).ToStruct()) }
func (s PingMsg) Id() string               { return C.Struct(s).GetObject(0).ToText() }
func (s PingMsg) SetId(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s PingMsg) MarshalJSON() (bs []byte, err error) { return }

type PingMsg_List C.PointerList

func NewPingMsgList(s *C.Segment, sz int) PingMsg_List {
	return PingMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s PingMsg_List) Len() int         { return C.PointerList(s).Len() }
func (s PingMsg_List) At(i int) PingMsg { return PingMsg(C.PointerList(s).At(i).ToStruct()) }
func (s PingMsg_List) ToArray() []PingMsg {
	n := s.Len()
	a := make([]PingMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s PingMsg_List) Set(i int, item PingMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type PongMsg C.Struct

func NewPongMsg(s *C.Segment) PongMsg      { return PongMsg(s.NewStruct(0, 1)) }
func NewRootPongMsg(s *C.Segment) PongMsg  { return PongMsg(s.NewRootStruct(0, 1)) }
func AutoNewPongMsg(s *C.Segment) PongMsg  { return PongMsg(s.NewStructAR(0, 1)) }
func ReadRootPongMsg(s *C.Segment) PongMsg { return PongMsg(s.Root(0).ToStruct()) }
func (s PongMsg) Id() string               { return C.Struct(s).GetObject(0).ToText() }
func (s PongMsg) SetId(v string)           { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s PongMsg) MarshalJSON() (bs []byte, err error) { return }

type PongMsg_List C.PointerList

func NewPongMsgList(s *C.Segment, sz int) PongMsg_List {
	return PongMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s PongMsg_List) Len() int         { return C.PointerList(s).Len() }
func (s PongMsg_List) At(i int) PongMsg { return PongMsg(C.PointerList(s).At(i).ToStruct()) }
func (s PongMsg_List) ToArray() []PongMsg {
	n := s.Len()
	a := make([]PongMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s PongMsg_List) Set(i int, item PongMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type SubMsg C.Struct

func NewSubMsg(s *C.Segment) SubMsg      { return SubMsg(s.NewStruct(0, 3)) }
func NewRootSubMsg(s *C.Segment) SubMsg  { return SubMsg(s.NewRootStruct(0, 3)) }
func AutoNewSubMsg(s *C.Segment) SubMsg  { return SubMsg(s.NewStructAR(0, 3)) }
func ReadRootSubMsg(s *C.Segment) SubMsg { return SubMsg(s.Root(0).ToStruct()) }
func (s SubMsg) Id() string              { return C.Struct(s).GetObject(0).ToText() }
func (s SubMsg) SetId(v string)          { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s SubMsg) Name() string            { return C.Struct(s).GetObject(1).ToText() }
func (s SubMsg) SetName(v string)        { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s SubMsg) Params() Param_List      { return Param_List(C.Struct(s).GetObject(2)) }
func (s SubMsg) SetParams(v Param_List)  { C.Struct(s).SetObject(2, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s SubMsg) MarshalJSON() (bs []byte, err error) { return }

type SubMsg_List C.PointerList

func NewSubMsgList(s *C.Segment, sz int) SubMsg_List { return SubMsg_List(s.NewCompositeList(0, 3, sz)) }
func (s SubMsg_List) Len() int                       { return C.PointerList(s).Len() }
func (s SubMsg_List) At(i int) SubMsg                { return SubMsg(C.PointerList(s).At(i).ToStruct()) }
func (s SubMsg_List) ToArray() []SubMsg {
	n := s.Len()
	a := make([]SubMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s SubMsg_List) Set(i int, item SubMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type UnsubMsg C.Struct

func NewUnsubMsg(s *C.Segment) UnsubMsg      { return UnsubMsg(s.NewStruct(0, 1)) }
func NewRootUnsubMsg(s *C.Segment) UnsubMsg  { return UnsubMsg(s.NewRootStruct(0, 1)) }
func AutoNewUnsubMsg(s *C.Segment) UnsubMsg  { return UnsubMsg(s.NewStructAR(0, 1)) }
func ReadRootUnsubMsg(s *C.Segment) UnsubMsg { return UnsubMsg(s.Root(0).ToStruct()) }
func (s UnsubMsg) Id() string                { return C.Struct(s).GetObject(0).ToText() }
func (s UnsubMsg) SetId(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s UnsubMsg) MarshalJSON() (bs []byte, err error) { return }

type UnsubMsg_List C.PointerList

func NewUnsubMsgList(s *C.Segment, sz int) UnsubMsg_List {
	return UnsubMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s UnsubMsg_List) Len() int          { return C.PointerList(s).Len() }
func (s UnsubMsg_List) At(i int) UnsubMsg { return UnsubMsg(C.PointerList(s).At(i).ToStruct()) }
func (s UnsubMsg_List) ToArray() []UnsubMsg {
	n := s.Len()
	a := make([]UnsubMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s UnsubMsg_List) Set(i int, item UnsubMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type NosubMsg C.Struct

func NewNosubMsg(s *C.Segment) NosubMsg      { return NosubMsg(s.NewStruct(0, 2)) }
func NewRootNosubMsg(s *C.Segment) NosubMsg  { return NosubMsg(s.NewRootStruct(0, 2)) }
func AutoNewNosubMsg(s *C.Segment) NosubMsg  { return NosubMsg(s.NewStructAR(0, 2)) }
func ReadRootNosubMsg(s *C.Segment) NosubMsg { return NosubMsg(s.Root(0).ToStruct()) }
func (s NosubMsg) Id() string                { return C.Struct(s).GetObject(0).ToText() }
func (s NosubMsg) SetId(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s NosubMsg) Error() Error              { return Error(C.Struct(s).GetObject(1).ToStruct()) }
func (s NosubMsg) SetError(v Error)          { C.Struct(s).SetObject(1, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s NosubMsg) MarshalJSON() (bs []byte, err error) { return }

type NosubMsg_List C.PointerList

func NewNosubMsgList(s *C.Segment, sz int) NosubMsg_List {
	return NosubMsg_List(s.NewCompositeList(0, 2, sz))
}
func (s NosubMsg_List) Len() int          { return C.PointerList(s).Len() }
func (s NosubMsg_List) At(i int) NosubMsg { return NosubMsg(C.PointerList(s).At(i).ToStruct()) }
func (s NosubMsg_List) ToArray() []NosubMsg {
	n := s.Len()
	a := make([]NosubMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s NosubMsg_List) Set(i int, item NosubMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type AddedMsg C.Struct

func NewAddedMsg(s *C.Segment) AddedMsg      { return AddedMsg(s.NewStruct(0, 3)) }
func NewRootAddedMsg(s *C.Segment) AddedMsg  { return AddedMsg(s.NewRootStruct(0, 3)) }
func AutoNewAddedMsg(s *C.Segment) AddedMsg  { return AddedMsg(s.NewStructAR(0, 3)) }
func ReadRootAddedMsg(s *C.Segment) AddedMsg { return AddedMsg(s.Root(0).ToStruct()) }
func (s AddedMsg) Id() string                { return C.Struct(s).GetObject(0).ToText() }
func (s AddedMsg) SetId(v string)            { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s AddedMsg) Collection() string        { return C.Struct(s).GetObject(1).ToText() }
func (s AddedMsg) SetCollection(v string)    { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s AddedMsg) Fields() C.Object          { return C.Struct(s).GetObject(2) }
func (s AddedMsg) SetFields(v C.Object)      { C.Struct(s).SetObject(2, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s AddedMsg) MarshalJSON() (bs []byte, err error) { return }

type AddedMsg_List C.PointerList

func NewAddedMsgList(s *C.Segment, sz int) AddedMsg_List {
	return AddedMsg_List(s.NewCompositeList(0, 3, sz))
}
func (s AddedMsg_List) Len() int          { return C.PointerList(s).Len() }
func (s AddedMsg_List) At(i int) AddedMsg { return AddedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s AddedMsg_List) ToArray() []AddedMsg {
	n := s.Len()
	a := make([]AddedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s AddedMsg_List) Set(i int, item AddedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type ChangedMsg C.Struct

func NewChangedMsg(s *C.Segment) ChangedMsg      { return ChangedMsg(s.NewStruct(0, 4)) }
func NewRootChangedMsg(s *C.Segment) ChangedMsg  { return ChangedMsg(s.NewRootStruct(0, 4)) }
func AutoNewChangedMsg(s *C.Segment) ChangedMsg  { return ChangedMsg(s.NewStructAR(0, 4)) }
func ReadRootChangedMsg(s *C.Segment) ChangedMsg { return ChangedMsg(s.Root(0).ToStruct()) }
func (s ChangedMsg) Id() string                  { return C.Struct(s).GetObject(0).ToText() }
func (s ChangedMsg) SetId(v string)              { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ChangedMsg) Collection() string          { return C.Struct(s).GetObject(1).ToText() }
func (s ChangedMsg) SetCollection(v string)      { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s ChangedMsg) Fields() C.Object            { return C.Struct(s).GetObject(2) }
func (s ChangedMsg) SetFields(v C.Object)        { C.Struct(s).SetObject(2, v) }
func (s ChangedMsg) Cleared() C.TextList         { return C.TextList(C.Struct(s).GetObject(3)) }
func (s ChangedMsg) SetCleared(v C.TextList)     { C.Struct(s).SetObject(3, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ChangedMsg) MarshalJSON() (bs []byte, err error) { return }

type ChangedMsg_List C.PointerList

func NewChangedMsgList(s *C.Segment, sz int) ChangedMsg_List {
	return ChangedMsg_List(s.NewCompositeList(0, 4, sz))
}
func (s ChangedMsg_List) Len() int            { return C.PointerList(s).Len() }
func (s ChangedMsg_List) At(i int) ChangedMsg { return ChangedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s ChangedMsg_List) ToArray() []ChangedMsg {
	n := s.Len()
	a := make([]ChangedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ChangedMsg_List) Set(i int, item ChangedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type RemovedMsg C.Struct

func NewRemovedMsg(s *C.Segment) RemovedMsg      { return RemovedMsg(s.NewStruct(0, 2)) }
func NewRootRemovedMsg(s *C.Segment) RemovedMsg  { return RemovedMsg(s.NewRootStruct(0, 2)) }
func AutoNewRemovedMsg(s *C.Segment) RemovedMsg  { return RemovedMsg(s.NewStructAR(0, 2)) }
func ReadRootRemovedMsg(s *C.Segment) RemovedMsg { return RemovedMsg(s.Root(0).ToStruct()) }
func (s RemovedMsg) Id() string                  { return C.Struct(s).GetObject(0).ToText() }
func (s RemovedMsg) SetId(v string)              { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s RemovedMsg) Collection() string          { return C.Struct(s).GetObject(1).ToText() }
func (s RemovedMsg) SetCollection(v string)      { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s RemovedMsg) MarshalJSON() (bs []byte, err error) { return }

type RemovedMsg_List C.PointerList

func NewRemovedMsgList(s *C.Segment, sz int) RemovedMsg_List {
	return RemovedMsg_List(s.NewCompositeList(0, 2, sz))
}
func (s RemovedMsg_List) Len() int            { return C.PointerList(s).Len() }
func (s RemovedMsg_List) At(i int) RemovedMsg { return RemovedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s RemovedMsg_List) ToArray() []RemovedMsg {
	n := s.Len()
	a := make([]RemovedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s RemovedMsg_List) Set(i int, item RemovedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type ReadyMsg C.Struct

func NewReadyMsg(s *C.Segment) ReadyMsg      { return ReadyMsg(s.NewStruct(0, 1)) }
func NewRootReadyMsg(s *C.Segment) ReadyMsg  { return ReadyMsg(s.NewRootStruct(0, 1)) }
func AutoNewReadyMsg(s *C.Segment) ReadyMsg  { return ReadyMsg(s.NewStructAR(0, 1)) }
func ReadRootReadyMsg(s *C.Segment) ReadyMsg { return ReadyMsg(s.Root(0).ToStruct()) }
func (s ReadyMsg) Subs() C.TextList          { return C.TextList(C.Struct(s).GetObject(0)) }
func (s ReadyMsg) SetSubs(v C.TextList)      { C.Struct(s).SetObject(0, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ReadyMsg) MarshalJSON() (bs []byte, err error) { return }

type ReadyMsg_List C.PointerList

func NewReadyMsgList(s *C.Segment, sz int) ReadyMsg_List {
	return ReadyMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s ReadyMsg_List) Len() int          { return C.PointerList(s).Len() }
func (s ReadyMsg_List) At(i int) ReadyMsg { return ReadyMsg(C.PointerList(s).At(i).ToStruct()) }
func (s ReadyMsg_List) ToArray() []ReadyMsg {
	n := s.Len()
	a := make([]ReadyMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ReadyMsg_List) Set(i int, item ReadyMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type AddedBeforeMsg C.Struct

func NewAddedBeforeMsg(s *C.Segment) AddedBeforeMsg      { return AddedBeforeMsg(s.NewStruct(0, 4)) }
func NewRootAddedBeforeMsg(s *C.Segment) AddedBeforeMsg  { return AddedBeforeMsg(s.NewRootStruct(0, 4)) }
func AutoNewAddedBeforeMsg(s *C.Segment) AddedBeforeMsg  { return AddedBeforeMsg(s.NewStructAR(0, 4)) }
func ReadRootAddedBeforeMsg(s *C.Segment) AddedBeforeMsg { return AddedBeforeMsg(s.Root(0).ToStruct()) }
func (s AddedBeforeMsg) Id() string                      { return C.Struct(s).GetObject(0).ToText() }
func (s AddedBeforeMsg) SetId(v string)                  { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s AddedBeforeMsg) Collection() string              { return C.Struct(s).GetObject(1).ToText() }
func (s AddedBeforeMsg) SetCollection(v string)          { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s AddedBeforeMsg) Fields() C.Object                { return C.Struct(s).GetObject(2) }
func (s AddedBeforeMsg) SetFields(v C.Object)            { C.Struct(s).SetObject(2, v) }
func (s AddedBeforeMsg) Before() string                  { return C.Struct(s).GetObject(3).ToText() }
func (s AddedBeforeMsg) SetBefore(v string)              { C.Struct(s).SetObject(3, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s AddedBeforeMsg) MarshalJSON() (bs []byte, err error) { return }

type AddedBeforeMsg_List C.PointerList

func NewAddedBeforeMsgList(s *C.Segment, sz int) AddedBeforeMsg_List {
	return AddedBeforeMsg_List(s.NewCompositeList(0, 4, sz))
}
func (s AddedBeforeMsg_List) Len() int { return C.PointerList(s).Len() }
func (s AddedBeforeMsg_List) At(i int) AddedBeforeMsg {
	return AddedBeforeMsg(C.PointerList(s).At(i).ToStruct())
}
func (s AddedBeforeMsg_List) ToArray() []AddedBeforeMsg {
	n := s.Len()
	a := make([]AddedBeforeMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s AddedBeforeMsg_List) Set(i int, item AddedBeforeMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type MovedBeforeMsg C.Struct

func NewMovedBeforeMsg(s *C.Segment) MovedBeforeMsg      { return MovedBeforeMsg(s.NewStruct(0, 3)) }
func NewRootMovedBeforeMsg(s *C.Segment) MovedBeforeMsg  { return MovedBeforeMsg(s.NewRootStruct(0, 3)) }
func AutoNewMovedBeforeMsg(s *C.Segment) MovedBeforeMsg  { return MovedBeforeMsg(s.NewStructAR(0, 3)) }
func ReadRootMovedBeforeMsg(s *C.Segment) MovedBeforeMsg { return MovedBeforeMsg(s.Root(0).ToStruct()) }
func (s MovedBeforeMsg) Id() string                      { return C.Struct(s).GetObject(0).ToText() }
func (s MovedBeforeMsg) SetId(v string)                  { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s MovedBeforeMsg) Collection() string              { return C.Struct(s).GetObject(1).ToText() }
func (s MovedBeforeMsg) SetCollection(v string)          { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s MovedBeforeMsg) Before() string                  { return C.Struct(s).GetObject(2).ToText() }
func (s MovedBeforeMsg) SetBefore(v string)              { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s MovedBeforeMsg) MarshalJSON() (bs []byte, err error) { return }

type MovedBeforeMsg_List C.PointerList

func NewMovedBeforeMsgList(s *C.Segment, sz int) MovedBeforeMsg_List {
	return MovedBeforeMsg_List(s.NewCompositeList(0, 3, sz))
}
func (s MovedBeforeMsg_List) Len() int { return C.PointerList(s).Len() }
func (s MovedBeforeMsg_List) At(i int) MovedBeforeMsg {
	return MovedBeforeMsg(C.PointerList(s).At(i).ToStruct())
}
func (s MovedBeforeMsg_List) ToArray() []MovedBeforeMsg {
	n := s.Len()
	a := make([]MovedBeforeMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s MovedBeforeMsg_List) Set(i int, item MovedBeforeMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type MethodMsg C.Struct

func NewMethodMsg(s *C.Segment) MethodMsg      { return MethodMsg(s.NewStruct(0, 4)) }
func NewRootMethodMsg(s *C.Segment) MethodMsg  { return MethodMsg(s.NewRootStruct(0, 4)) }
func AutoNewMethodMsg(s *C.Segment) MethodMsg  { return MethodMsg(s.NewStructAR(0, 4)) }
func ReadRootMethodMsg(s *C.Segment) MethodMsg { return MethodMsg(s.Root(0).ToStruct()) }
func (s MethodMsg) Method() string             { return C.Struct(s).GetObject(0).ToText() }
func (s MethodMsg) SetMethod(v string)         { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s MethodMsg) Params() Param_List         { return Param_List(C.Struct(s).GetObject(1)) }
func (s MethodMsg) SetParams(v Param_List)     { C.Struct(s).SetObject(1, C.Object(v)) }
func (s MethodMsg) Id() string                 { return C.Struct(s).GetObject(2).ToText() }
func (s MethodMsg) SetId(v string)             { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }
func (s MethodMsg) RandomSeed() C.Object       { return C.Struct(s).GetObject(3) }
func (s MethodMsg) SetRandomSeed(v C.Object)   { C.Struct(s).SetObject(3, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s MethodMsg) MarshalJSON() (bs []byte, err error) { return }

type MethodMsg_List C.PointerList

func NewMethodMsgList(s *C.Segment, sz int) MethodMsg_List {
	return MethodMsg_List(s.NewCompositeList(0, 4, sz))
}
func (s MethodMsg_List) Len() int           { return C.PointerList(s).Len() }
func (s MethodMsg_List) At(i int) MethodMsg { return MethodMsg(C.PointerList(s).At(i).ToStruct()) }
func (s MethodMsg_List) ToArray() []MethodMsg {
	n := s.Len()
	a := make([]MethodMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s MethodMsg_List) Set(i int, item MethodMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type ResultMsg C.Struct

func NewResultMsg(s *C.Segment) ResultMsg      { return ResultMsg(s.NewStruct(0, 3)) }
func NewRootResultMsg(s *C.Segment) ResultMsg  { return ResultMsg(s.NewRootStruct(0, 3)) }
func AutoNewResultMsg(s *C.Segment) ResultMsg  { return ResultMsg(s.NewStructAR(0, 3)) }
func ReadRootResultMsg(s *C.Segment) ResultMsg { return ResultMsg(s.Root(0).ToStruct()) }
func (s ResultMsg) Id() string                 { return C.Struct(s).GetObject(0).ToText() }
func (s ResultMsg) SetId(v string)             { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s ResultMsg) Error() Error               { return Error(C.Struct(s).GetObject(1).ToStruct()) }
func (s ResultMsg) SetError(v Error)           { C.Struct(s).SetObject(1, C.Object(v)) }
func (s ResultMsg) Result() C.Object           { return C.Struct(s).GetObject(2) }
func (s ResultMsg) SetResult(v C.Object)       { C.Struct(s).SetObject(2, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s ResultMsg) MarshalJSON() (bs []byte, err error) { return }

type ResultMsg_List C.PointerList

func NewResultMsgList(s *C.Segment, sz int) ResultMsg_List {
	return ResultMsg_List(s.NewCompositeList(0, 3, sz))
}
func (s ResultMsg_List) Len() int           { return C.PointerList(s).Len() }
func (s ResultMsg_List) At(i int) ResultMsg { return ResultMsg(C.PointerList(s).At(i).ToStruct()) }
func (s ResultMsg_List) ToArray() []ResultMsg {
	n := s.Len()
	a := make([]ResultMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ResultMsg_List) Set(i int, item ResultMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type UpdatedMsg C.Struct

func NewUpdatedMsg(s *C.Segment) UpdatedMsg      { return UpdatedMsg(s.NewStruct(0, 1)) }
func NewRootUpdatedMsg(s *C.Segment) UpdatedMsg  { return UpdatedMsg(s.NewRootStruct(0, 1)) }
func AutoNewUpdatedMsg(s *C.Segment) UpdatedMsg  { return UpdatedMsg(s.NewStructAR(0, 1)) }
func ReadRootUpdatedMsg(s *C.Segment) UpdatedMsg { return UpdatedMsg(s.Root(0).ToStruct()) }
func (s UpdatedMsg) Methods() C.TextList         { return C.TextList(C.Struct(s).GetObject(0)) }
func (s UpdatedMsg) SetMethods(v C.TextList)     { C.Struct(s).SetObject(0, C.Object(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s UpdatedMsg) MarshalJSON() (bs []byte, err error) { return }

type UpdatedMsg_List C.PointerList

func NewUpdatedMsgList(s *C.Segment, sz int) UpdatedMsg_List {
	return UpdatedMsg_List(s.NewCompositeList(0, 1, sz))
}
func (s UpdatedMsg_List) Len() int            { return C.PointerList(s).Len() }
func (s UpdatedMsg_List) At(i int) UpdatedMsg { return UpdatedMsg(C.PointerList(s).At(i).ToStruct()) }
func (s UpdatedMsg_List) ToArray() []UpdatedMsg {
	n := s.Len()
	a := make([]UpdatedMsg, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s UpdatedMsg_List) Set(i int, item UpdatedMsg) { C.PointerList(s).Set(i, C.Object(item)) }

type Param C.Struct

func NewParam(s *C.Segment) Param      { return Param(s.NewStruct(0, 1)) }
func NewRootParam(s *C.Segment) Param  { return Param(s.NewRootStruct(0, 1)) }
func AutoNewParam(s *C.Segment) Param  { return Param(s.NewStructAR(0, 1)) }
func ReadRootParam(s *C.Segment) Param { return Param(s.Root(0).ToStruct()) }
func (s Param) Value() C.Object        { return C.Struct(s).GetObject(0) }
func (s Param) SetValue(v C.Object)    { C.Struct(s).SetObject(0, v) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Param) MarshalJSON() (bs []byte, err error) { return }

type Param_List C.PointerList

func NewParamList(s *C.Segment, sz int) Param_List { return Param_List(s.NewCompositeList(0, 1, sz)) }
func (s Param_List) Len() int                      { return C.PointerList(s).Len() }
func (s Param_List) At(i int) Param                { return Param(C.PointerList(s).At(i).ToStruct()) }
func (s Param_List) ToArray() []Param {
	n := s.Len()
	a := make([]Param, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Param_List) Set(i int, item Param) { C.PointerList(s).Set(i, C.Object(item)) }

type Error C.Struct

func NewError(s *C.Segment) Error      { return Error(s.NewStruct(0, 3)) }
func NewRootError(s *C.Segment) Error  { return Error(s.NewRootStruct(0, 3)) }
func AutoNewError(s *C.Segment) Error  { return Error(s.NewStructAR(0, 3)) }
func ReadRootError(s *C.Segment) Error { return Error(s.Root(0).ToStruct()) }
func (s Error) Error() string          { return C.Struct(s).GetObject(0).ToText() }
func (s Error) SetError(v string)      { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Error) Reason() string         { return C.Struct(s).GetObject(1).ToText() }
func (s Error) SetReason(v string)     { C.Struct(s).SetObject(1, s.Segment.NewText(v)) }
func (s Error) Details() string        { return C.Struct(s).GetObject(2).ToText() }
func (s Error) SetDetails(v string)    { C.Struct(s).SetObject(2, s.Segment.NewText(v)) }

// capn.JSON_enabled == false so we stub MarshallJSON().
func (s Error) MarshalJSON() (bs []byte, err error) { return }

type Error_List C.PointerList

func NewErrorList(s *C.Segment, sz int) Error_List { return Error_List(s.NewCompositeList(0, 3, sz)) }
func (s Error_List) Len() int                      { return C.PointerList(s).Len() }
func (s Error_List) At(i int) Error                { return Error(C.PointerList(s).At(i).ToStruct()) }
func (s Error_List) ToArray() []Error {
	n := s.Len()
	a := make([]Error, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Error_List) Set(i int, item Error) { C.PointerList(s).Set(i, C.Object(item)) }
