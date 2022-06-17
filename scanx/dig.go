package scanx

import (
	"fmt"
	"strings"
	"time"

	"github.com/gokitx/pkgs/slicex"
	"github.com/miekg/dns"
)

const dnsResolveServer = "8.8.8.8:53"

var (
	client *dns.Client = &dns.Client{Timeout: time.Second * 5}
)

type Result []Record

type Record struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func (r *Result) Push(answer []dns.RR) {
	for _, an := range answer {
		rec := Record{
			Name: strings.TrimSuffix(an.Header().Name, "."),
		}
		switch record := an.(type) {
		case *dns.A:
			rec.Type = dns.TypeToString[dns.TypeA]
			rec.Value = string(record.A.String())
		case *dns.NS:
			rec.Type = dns.TypeToString[dns.TypeNS]
			rec.Value = string(record.Ns)
		case *dns.CNAME:
			rec.Type = dns.TypeToString[dns.TypeCNAME]
			rec.Value = string(strings.TrimSuffix(record.Target, "."))
		case *dns.MX:
			rec.Type = dns.TypeToString[dns.TypeMX]
			rec.Value = string(record.Mx)
		case *dns.AAAA:
			rec.Type = dns.TypeToString[dns.TypeAAAA]
			rec.Value = string(record.AAAA.String())
		}

		(*r) = append((*r), rec)
	}
}

func (r *Result) RemoveDuplicates() {
	*r = slicex.RevemoRepWithSort(*r)
}

func Dig(name string, requestTypes []uint16, resolver string) (Result, error) {
	if resolver == "" {
		resolver = dnsResolveServer
	}
	s := strings.Split(resolver, ":")
	if len(s) != 2 || s[1] == "" {
		resolver = fmt.Sprintf("%s:53", s[0])
	}

	res := make(Result, 0)
	for i := range requestTypes {
		msg := new(dns.Msg)
		msg.SetQuestion(dns.Fqdn(name), requestTypes[i])
		resp, _, err := client.Exchange(msg, resolver)
		if err != nil {
			return nil, err
		}
		res.Push(resp.Answer)
	}

	res.RemoveDuplicates()
	return res, nil
}
