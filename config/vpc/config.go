package vpc

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_vpc", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
	})
	p.AddResourceConfigurator("cloudamqp_vpc_peering", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "VpcPeering"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.Vpc",
		}
		r.References["instance_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.Instance",
		}
	})
}
