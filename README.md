# bddp

> ddp for binary data

[DDP](https://github.com/meteor/meteor/blob/devel/packages/ddp/DDP.md) is a protocol used by Meteor for real time data communication. It supports remote procedure calls (RPC) and subscriptions. DDP runs on websockets with sockjs and uses EJSON to serialize data during transmission which is sometimes not good enough for binary data.

BDDP runs on tcp and uses *cap'n proto* to as the data format in order to reduce encoding/decoding time as much as possible.

## TODO

 * implement bddp client/server packages for nodejs and golang
