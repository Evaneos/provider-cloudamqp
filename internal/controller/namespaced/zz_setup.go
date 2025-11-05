// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	community "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/community"
	firewall "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/firewall"
	instance "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/instance"
	integrationlog "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/integrationlog"
	integrationmetric "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/integrationmetric"
	integrationprometheus "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/integrationprometheus"
	maintenancewindow "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/maintenancewindow"
	plugin "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/plugin"
	vpc "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/vpc"
	vpcgcppeering "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/vpcgcppeering"
	vpcpeering "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/cloudamqp/vpcpeering"
	providerconfig "github.com/evaneos/provider-cloudamqp/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		community.Setup,
		firewall.Setup,
		instance.Setup,
		integrationlog.Setup,
		integrationmetric.Setup,
		integrationprometheus.Setup,
		maintenancewindow.Setup,
		plugin.Setup,
		vpc.Setup,
		vpcgcppeering.Setup,
		vpcpeering.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		community.SetupGated,
		firewall.SetupGated,
		instance.SetupGated,
		integrationlog.SetupGated,
		integrationmetric.SetupGated,
		integrationprometheus.SetupGated,
		maintenancewindow.SetupGated,
		plugin.SetupGated,
		vpc.SetupGated,
		vpcgcppeering.SetupGated,
		vpcpeering.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
