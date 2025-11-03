/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	instanceCluster "github.com/evaneos/provider-cloudamqp/config/cluster/instance"
	maintenanceCluster "github.com/evaneos/provider-cloudamqp/config/cluster/maintenance"
	pluginCluster "github.com/evaneos/provider-cloudamqp/config/cluster/plugin"
	securityCluster "github.com/evaneos/provider-cloudamqp/config/cluster/security"
	vpcCluster "github.com/evaneos/provider-cloudamqp/config/cluster/vpc"

	instanceNamespaced "github.com/evaneos/provider-cloudamqp/config/namespaced/instance"
	maintenanceNamespaced "github.com/evaneos/provider-cloudamqp/config/namespaced/maintenance"
	pluginNamespaced "github.com/evaneos/provider-cloudamqp/config/namespaced/plugin"
	securityNamespaced "github.com/evaneos/provider-cloudamqp/config/namespaced/security"
	vpcNamespaced "github.com/evaneos/provider-cloudamqp/config/namespaced/vpc"
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
		instanceCluster.Configure,
		vpcCluster.Configure,
		maintenanceCluster.Configure,
		pluginCluster.Configure,
		securityCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}


// GetProviderNamespaced returns provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("m.evaneos.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		instanceNamespaced.Configure,
		vpcNamespaced.Configure,
		maintenanceNamespaced.Configure,
		pluginNamespaced.Configure,
		securityNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
