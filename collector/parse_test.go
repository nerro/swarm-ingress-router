package collector

import (
	"testing"

	"github.com/docker/docker/api/types/swarm"

	"github.com/tpbowden/swarm-ingress-router/types"
)

type parseTest struct {
	description   string
	swarmServices []swarm.Service
	result        []types.Service
}

var parseTests = []parseTest{
	{
		description: "Parsing a valid service",
		swarmServices: []swarm.Service{
			{
				Spec: swarm.ServiceSpec{
					Annotations: swarm.Annotations{
						Name: "test service",
						Labels: map[string]string{
							"ingress.targetport": "100",
							"ingress.tls":        "true",
							"ingress.forcetls":   "true",
							"ingress.cert":       "a certificate",
							"ingress.key":        "a key",
						},
					},
				},
			},
		},
		result: []types.Service{
			{
				Name:        "test service",
				Port:        100,
				Secure:      true,
				ForceTLS:    true,
				Certificate: "a certificate",
				Key:         "a key",
			},
		},
	},
	{
		description: "Skipping an invalid port number",
		swarmServices: []swarm.Service{
			{
				Spec: swarm.ServiceSpec{
					Annotations: swarm.Annotations{
						Name: "test service",
						Labels: map[string]string{
							"ingress.targetport": "abc",
							"ingress.tls":        "true",
							"ingress.forcetls":   "true",
							"ingress.cert":       "a certificate",
							"ingress.key":        "a key",
						},
					},
				},
			},
		},
		result: []types.Service{},
	},
}

func TestParsingServices(t *testing.T) {
	for _, test := range parseTests {
		parsedServices := parseServices(test.swarmServices)

		for i, res := range parsedServices {
			if test.result[i] != res {
				t.Errorf("Services did not match: expected %v, got %v", test.result[i], res)
			}
		}
	}
}