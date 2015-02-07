// +build !windows

package zmq4

/*
#cgo CFLAGS: -I../../clibs/include
#cgo LDFLAGS: -lzmq -L../../clibs/lib

#include <zmq.h>
*/
import "C"

// ZMQ_FD: Retrieve file descriptor associated with the socket
//
// See: http://api.zeromq.org/4-0:zmq-getsockopt#toc24
func (soc *Socket) GetFd() (int, error) {
	return soc.getInt(C.ZMQ_FD)
}
