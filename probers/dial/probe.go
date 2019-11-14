package dial

import (
	"net"
)

func (p *dialConfig) probe() error {
	if conn, err := net.Dial(p.network, p.address); err != nil {
		p.healthy = false
		return err
	} else {
		conn.Close()
		p.healthy = true
		return nil
	}
}
