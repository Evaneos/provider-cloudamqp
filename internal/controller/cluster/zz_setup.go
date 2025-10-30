// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>package cluster

//
// SPDX-License-Identifier: Apache-2.0

package cluster

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	community "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/community"
	firewall "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/firewall"
	instance "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/instance"
	maintenancewindow "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/maintenancewindow"
	plugin "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/plugin"
	vpc "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpc"
	vpcgcppeering "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpcgcppeering"
	vpcpeering "github.com/evaneos/provider-cloudamqp/internal/controller/cluster/cloudamqp/vpcpeering"
)

// Setup creates all cluster-scoped controllers and adds them to the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		community.Setup,
		firewall.Setup,
		instance.Setup,
		maintenancewindow.Setup,
		plugin.Setup,
		vpc.Setup,
		vpcgcppeering.Setup,
		vpcpeering.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
