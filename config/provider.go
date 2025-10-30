/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/evaneos/provider-cloudamqp/config/instance"
	"github.com/evaneos/provider-cloudamqp/config/maintenance"
	"github.com/evaneos/provider-cloudamqp/config/plugin"
	"github.com/evaneos/provider-cloudamqp/config/security"
	"github.com/evaneos/provider-cloudamqp/config/vpc"
)

const (
	resourcePrefix = "cloudamqp"
	modulePath     = "github.com/evaneos/provider-cloudamqp"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("evaneos.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		instance.Configure,
		vpc.Configure,
		maintenance.Configure,
		plugin.Configure,
		security.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
