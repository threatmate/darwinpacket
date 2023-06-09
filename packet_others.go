//go:build !darwin
// +build !darwin

package darwinpacket

import (
	"fmt"
	"net"
	"runtime"
	"syscall"
	"time"

	"golang.org/x/net/bpf"
)

// errUnimplemented is returned as appropriate
var errUnimplemented = fmt.Errorf("darwinpacket: not implemented on %s", runtime.GOOS)

func listen(_ *net.Interface, _ int, _ *Config) (*Conn, error) { return nil, errUnimplemented }

func (*Conn) readFrom(_ []byte) (int, net.Addr, error)  { return 0, nil, errUnimplemented }
func (*Conn) writeTo(_ []byte, _ net.Addr) (int, error) { return 0, errUnimplemented }
func (*Conn) setPromiscuous(_ bool) error               { return errUnimplemented }
func (*Conn) stats() (*Stats, error)                    { return nil, errUnimplemented }

type conn struct{}

func (*conn) Close() error                          { return errUnimplemented }
func (*conn) SetDeadline(_ time.Time) error         { return errUnimplemented }
func (*conn) SetReadDeadline(_ time.Time) error     { return errUnimplemented }
func (*conn) SetWriteDeadline(_ time.Time) error    { return errUnimplemented }
func (*conn) SetBPF(_ []bpf.RawInstruction) error   { return errUnimplemented }
func (*conn) SyscallConn() (syscall.RawConn, error) { return nil, errUnimplemented }
