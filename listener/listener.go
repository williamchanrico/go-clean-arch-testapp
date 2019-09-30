package listener

import (
	"net"
	"os"
	"strconv"

	"github.com/williamchanrico/xtest/log"
)

// Listen is a listener that is aware of socketmaster
func Listen(addr string) (net.Listener, error) {
	var listener net.Listener

	einhornFD := os.Getenv("EINHORN_FDS")
	if einhornFD != "" {
		sock, err := strconv.Atoi(einhornFD)
		if err != nil {
			return nil, err
		}

		log.Debugf("Detected socketmaster, listening on fd: %v", einhornFD)
		file := os.NewFile(uintptr(sock), "listener")
		listener, err = net.FileListener(file)
		if err != nil {
			return nil, err
		}
	}
	if listener != nil {
		return listener, nil
	}

	return net.Listen("tcp4", addr)
}
