# Provider CloudAMQP

`provider-cloudamqp` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
CloudAMQP API.

## Getting Started

Install the provider by using the following command after changing the image tag
to the [latest release](https://marketplace.upbound.io/providers/evaneos/provider-cloudamqp):
```
up ctp provider install evaneos/provider-cloudamqp:v0.1.0
```

Alternatively, you can use declarative installation:
```
cat <<EOF | kubectl apply -f -
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudamqp
spec:
  package: evaneos/provider-cloudamqp:v0.1.0
EOF
```

Notice that in this example Provider resource is referencing ControllerConfig with debug enabled.

You can see the API reference [here](https://doc.crds.dev/github.com/evaneos/provider-cloudamqp).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/evaneos/provider-cloudamqp/issues).

## Upjet v2 : Structure multi-scope (Cluster/Namespaced)

Ce provider est compatible Upjet v2 et Crossplane v2 :
- Les ressources sont disponibles en deux scopes : cluster (`cloudamqp.evaneos.com`) et namespaced (`cloudamqp.m.evaneos.com`).
- Les CRDs générés sont présents dans `package/crds/` pour chaque scope.
- Les exemples YAML sont fournis pour les deux scopes (voir dossier `examples/`).
- Pour utiliser les ressources namespaced, appliquez les CRDs `cloudamqp.m.evaneos.com_*` et utilisez le champ `metadata.namespace` dans vos manifests.

Pour plus de détails sur la migration et la structure, voir `upgrade-resume.md` et la documentation officielle :
https://github.com/crossplane/upjet/blob/main/docs/upjet-v2-upgrade.md
