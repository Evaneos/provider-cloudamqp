package maintenance

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_maintenance_window", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "MaintenanceWindow"
		r.References["instance_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.Instance",
		}
	})
}
