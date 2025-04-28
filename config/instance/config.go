package instance

import (
	"regexp"

	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "instance"
		r.ShortGroup = "cloudamqp"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.VPC",
		}

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if a, ok := attr["url"].(string); ok {
				re := regexp.MustCompile(`amqp://(\w+):(.+)@`)
				matches := re.FindStringSubmatch(a)
				if len(matches) == 3 {
					conn["username"] = []byte(matches[1])
					conn["password"] = []byte(matches[2])
				}
			}

			if a, ok := attr["host"].(string); ok {
				conn["host"] = []byte(a)
			}
			if a, ok := attr["host_internal"].(string); ok {
				conn["host_internal"] = []byte(a)
			}
			if a, ok := attr["vhost"].(string); ok {
				conn["vhost"] = []byte(a)
			}
			if a, ok := attr["vpc_id"].(string); ok {
				conn["vpc_id"] = []byte(a)
			}
			return conn, nil
		}
	})
}
