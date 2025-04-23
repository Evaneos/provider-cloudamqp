package instance

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_instance", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "instance"
		r.ShortGroup = "cloudamqp"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.VPC",
		}
	})
}
