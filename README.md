# packet [![Test Status](https://github.com/threatmate/darwinpacket/workflows/Test/badge.svg)](https://github.com/threatmate/darwinpacket/actions) [![Go Reference](https://pkg.go.dev/badge/github.com/threatmate/darwinpacket.svg)](https://pkg.go.dev/github.com/threatmate/darwinpacket)  [![Go Report Card](https://goreportcard.com/badge/github.com/threatmate/darwinpacket)](https://goreportcard.com/report/github.com/threatmate/darwinpacket)

Package `darwinpacket` provides access to MacOS sockets via BPF.

It provides a `net.PacketConn` interface that can be used, in particular, with `github.com/mdlayher/arp`.

Golang GOOS targets:

* `darwin`

## Usage

```
conn, err := darwinpacket.Listen(iface, protocol, &darwinpacket.Config{})
```

## History

`github.com/mdlayher` has specifically said that `github.com/mdlayher/packet` will only support Linux and
recommended that we create variants for any other operating systems that we are interested in (in this
case, MacOS).

This package is based on `github.com/synfinatic/packet`, but this version specifically limits support
to MacOS in keeping with `github.com/mdlayher`'s design wishes.
