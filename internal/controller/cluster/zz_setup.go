// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	community "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/community"
	firewall "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/firewall"
	instance "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/instance"
	maintenancewindow "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/maintenancewindow"
	metric "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/metric"
	plugin "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/plugin"
	vpc "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpc"
	vpcgcppeering "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpcgcppeering"
	vpcpeering "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpcpeering"
	providerconfig "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		community.Setup,
		firewall.Setup,
		instance.Setup,
		maintenancewindow.Setup,
		metric.Setup,
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
		maintenancewindow.SetupGated,
		metric.SetupGated,
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
