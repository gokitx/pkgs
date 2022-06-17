package scanx

import (
	"testing"

	"github.com/miekg/dns"
)

func TestDig(t *testing.T) {
	t.Log(Dig("zznq.imipy.com", []uint16{dns.TypeA, dns.TypeMX}, ""))

	t.Log(Dig("zznq.imipy.com", []uint16{dns.TypeA, dns.TypeMX}, "114.114.114.114:"))
}
