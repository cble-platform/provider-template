# CBLE Provider Template

## Quick Start

Click the **Use this template** button to create a repository using this template.

Then, clone the repository to get started.

> **Note:** The name of providers should follow the standard format `provider-<name>` (e.g. for OpenStack use `provider-openstack`)

### Edit the Go module name

Modify the [`go.mod`](./go.mod) file to reflect your module name:

```go
module github.com/<your-username-or-org>/provider-<name>

// ...
```

### Rename the provider

To start, rename the provider folder and package name with your custom name (replace `<name>` with your provider's name):

```shell
NAME=<name>
mv example/ $NAME
sed -i "s/^package example/package $NAME/" $NAME/*
```

Update the `cble-metadata.yml` file to reflect this name update:

```yaml
name: provider-<name>
description: <provider description>
author: <Your Name> <github.com/<your-username-or-org>>
version: v0.0.0
# ...
```

### Set the features

Not all providers will support all features of the spec, this is ok! You can toggle the feature support on/off in [`main.go`](./main.go#L57-61):

```go
  // ...
    Features: &cbleGRPC.ProviderFeatures{
      Deploy:  true, // Set to false to disable
      Destroy: true, // Set to false to disable
      Console: true, // Set to false to disable
    },
  // ...
```
