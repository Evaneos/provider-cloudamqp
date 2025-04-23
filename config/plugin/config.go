package plugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/crossplane/upjet/pkg/config"
)

const (
	ErrFmtNoAttribute    = "Attribute %s not found"
	ErrFmtUnexpectedType = "Attribute %s is of an unexpected type"
)

func getNameFromFullyQualifiedID(tfstate map[string]any) (string, error) {
	id, ok := tfstate["id"]
	if !ok {
		return "", fmt.Errorf(ErrFmtNoAttribute, "id")
	}
	idStr, ok := id.(string)
	if !ok {
		return "", fmt.Errorf(ErrFmtUnexpectedType, "id")
	}
	words := strings.Split(idStr, ",")
	return words[len(words)-1], nil
}

func getFullyQualifiedIDfunc(ctx context.Context, externalName string, parameters map[string]any, providerConfig map[string]any) (string, error) {
	id, vpcok := parameters["vpc_id"]

	if !vpcok {
		id, instanceok := parameters["instance_id"]
		if !instanceok {
			return "", fmt.Errorf(ErrFmtNoAttribute, "vpc_id")
		}

		return fmt.Sprintf("instance,%s,%s", id, externalName), nil
	}

	return fmt.Sprintf("vpc,%s,%s", id, externalName), nil
}

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("cloudamqp_plugin", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
	})
	p.AddResourceConfigurator("cloudamqp_plugin_community", func(r *config.Resource) {
		r.ShortGroup = "cloudamqp"
		r.ExternalName.GetExternalNameFn = getNameFromFullyQualifiedID
		r.ExternalName.GetIDFn = getFullyQualifiedIDfunc
	})
}
