using Go = import "../../glycerine/go-capnproto/go.capnp";

@0xce79b00365034eeb;
$Go.package("bddp");
$Go.import("github.com/glycerine/go-capnproto/capnpc-go");

using Id = UInt64;

struct Message {
  union {
    connect @0 :ConnectMsg;
    connected @1 :ConnectedMsg;
    failed @2 :FailedMsg;
    ping @3 :PingMsg;
    pong @4 :PongMsg;
    sub @5 :SubMsg;
    unsub @6 :UnsubMsg;
    nosub @7 :NosubMsg;
    added @8 :AddedMsg;
    changed @9 :ChangedMsg;
    removed @10 :RemovedMsg;
    ready @11 :ReadyMsg;
    addedbefore @12 :AddedBeforeMsg;
    movedbefore @13 :MovedBeforeMsg;
    method @14 :MethodMsg;
    result @15 :ResultMsg;
    updated @16 :UpdatedMsg;
  }
}

struct ConnectMsg {
  session @0 :Text;
  version @1 :Text;
  support @2 :List(Text);
}

struct ConnectedMsg {
  session @0 :Text;
}

struct FailedMsg {
  version @0 :Text;
}

struct PingMsg {
  id @0 :Text;
}

struct PongMsg {
  id @0 :Text;
}

struct SubMsg {
  id @0 :Text;
  name @1 :Text;
  params @2 :Data;
}

struct UnsubMsg {
  id @0 :Text;
}

struct NosubMsg {
  id @0 :Text;
  error @1 :Error;
}

struct AddedMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
}

struct ChangedMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
  cleared @3 :List(Text);
}

struct RemovedMsg {
  id @0 :Text;
  collection @1 :Text;
}

struct ReadyMsg {
  subs @0 :List(Text);
}

struct AddedBeforeMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
  before @3 :Text;
}

struct MovedBeforeMsg {
  id @0 :Text;
  collection @1 :Text;
  before @2 :Text;
}

struct MethodMsg {
  method @0 :Text;
  params @1 :Data;
  id @2 :Text;
  randomSeed @3 :Data;
}

struct ResultMsg {
  id @0 :Text;
  error @1 :Error;
  result @2 :Data;
}

struct UpdatedMsg {
  methods @0 :List(Text);
}

struct Error {
  error @0 :Text;
  reason @1 :Text;
  details @2 :Text;
}
