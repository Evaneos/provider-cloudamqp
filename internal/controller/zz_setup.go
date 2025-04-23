// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	community "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/community"
	instance "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/instance"
	maintenancewindow "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/maintenancewindow"
	plugin "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/plugin"
	vpc "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/vpc"
	vpcpeering "github.com/evaneos/provider-cloudamqp/internal/controller/cloudamqp/vpcpeering"
	providerconfig "github.com/evaneos/provider-cloudamqp/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		community.Setup,
		instance.Setup,
		maintenancewindow.Setup,
		plugin.Setup,
		vpc.Setup,
		vpcpeering.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
