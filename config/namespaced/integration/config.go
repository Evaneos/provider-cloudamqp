package integration

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_integration_metric", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "IntegrationMetric"

		r.References["instance_id"] = config.Reference{
			TerraformName: "cloudamqp_instance",
		}
	})

	p.AddResourceConfigurator("cloudamqp_integration_log", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "IntegrationLog"

		r.References["instance_id"] = config.Reference{
			TerraformName: "cloudamqp_instance",
		}
	})

	p.AddResourceConfigurator("cloudamqp_integration_metric_prometheus", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "IntegrationPrometheus"

		r.References["instance_id"] = config.Reference{
			TerraformName: "cloudamqp_instance",
		}
	})
}
