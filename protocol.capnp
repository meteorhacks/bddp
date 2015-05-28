using Go = import "../../glycerine/go-capnproto/go.capnp";

@0xce79b00365034eeb;
$Go.package("bddp");
$Go.import("github.com/glycerine/go-capnproto/capnpc-go");

using Id = UInt64;

// type 0
struct ConnectMsg {
  session @0 :Text;
  version @1 :Text;
  support @2 :List(Text);
}

// type 1
struct ConnectedMsg {
  session @0 :Text;
}

// type 2
struct FailedMsg {
  version @0 :Text;
}

// type 3
struct PingMsg {
  id @0 :Text;
}

// type 4
struct PongMsg {
  id @0 :Text;
}

// type 5
struct SubMsg {
  id @0 :Text;
  name @1 :Text;
  params @2 :Data;
}

// type 6
struct UnsubMsg {
  id @0 :Text;
}

// type 7
struct NosubMsg {
  id @0 :Text;
  error @1 :Error;
}

// type 8
struct AddedMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
}

// type 9
struct ChangedMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
  cleared @3 :List(Text);
}

// type 10
struct RemovedMsg {
  id @0 :Text;
  collection @1 :Text;
}

// type 11
struct ReadyMsg {
  subs @0 :List(Text);
}

// type 12
struct AddedBeforeMsg {
  id @0 :Text;
  collection @1 :Text;
  fields @2 :Data;
  before @3 :Text;
}

// type 13
struct MovedBeforeMsg {
  id @0 :Text;
  collection @1 :Text;
  before @2 :Text;
}

// type 14
struct MethodMsg {
  method @0 :Text;
  params @1 :Data;
  id @2 :Text;
  randomSeed @3 :Data;
}

// type 15
struct ResultMsg {
  id @0 :Text;
  error @1 :Error;
  result @2 :Data;
}

// type 16
struct UpdatedMsg {
  methods @0 :List(Text);
}

struct Error {
  error @0 :Text;
  reason @1 :Text;
  details @2 :Text;
}
