package vpc


import (
	"fmt"
	"strings"
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_vpc", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if region, ok := attr["region"].(string); ok && strings.HasPrefix(region, "google-compute-engine") {
				if a, ok := attr["vpc_name"].(string); ok {
					conn["gcp_id"] = []byte(fmt.Sprintf("projects/cloudamqp/global/networks/%s", a))
				}
			}
			return conn, nil
		}
	})
	p.AddResourceConfigurator("cloudamqp_vpc_peering", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "VpcPeering"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudamqp_vpc",
		}
		r.References["instance_id"] = config.Reference{
			TerraformName: "cloudamqp_instance",
		}
	})
	p.AddResourceConfigurator("cloudamqp_vpc_gcp_peering", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "VpcGCPPeering"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "cloudamqp_vpc",
		}
		r.References["instance_id"] = config.Reference{
			TerraformName: "cloudamqp_instance",
		}
	})
}
