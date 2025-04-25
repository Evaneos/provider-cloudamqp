package vpc

import (
	"fmt"
	"strings"

	// "github.com/crossplane/crossplane-runtime/pkg/logging"
	"github.com/crossplane/upjet/pkg/config"
	// "sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_vpc", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"

		// zl := zap.New(zap.UseDevMode(true))
		// logr := logging.NewLogrLogger(zl.WithName("provider-gcp"))
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			// If we are in GCP, register the vpc id
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
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.VPC",
		}
		r.References["instance_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.Instance",
		}
	})
	p.AddResourceConfigurator("cloudamqp_vpc_gcp_peering", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.Kind = "VpcGCPPeering"

		r.References["vpc_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.VPC",
		}
		r.References["instance_id"] = config.Reference{
			Type: "github.com/evaneos/provider-cloudamqp/apis/cloudamqp/v1alpha1.Instance",
		}
	})
}
