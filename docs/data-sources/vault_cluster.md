---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "hcp_vault_cluster Data Source - terraform-provider-hcp"
subcategory: ""
description: |-
  The cluster data source provides information about an existing HCP Vault cluster.
---

# hcp_vault_cluster (Data Source)

The cluster data source provides information about an existing HCP Vault cluster.

## Example Usage

```terraform
data "hcp_vault_cluster" "example" {
  cluster_id = var.cluster_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **cluster_id** (String) The ID of the HCP Vault cluster.

### Optional

- **id** (String) The ID of this resource.
- **timeouts** (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- **cloud_provider** (String) The provider where the HCP Vault cluster is located.
- **created_at** (String) The time that the Vault cluster was created.
- **hvn_id** (String) The ID of the HVN this HCP Vault cluster is associated to.
- **min_vault_version** (String) The minimum Vault version to use when creating the cluster. If not specified, it is defaulted to the version that is currently recommended by HCP.
- **namespace** (String) The name of the customer namespace this HCP Vault cluster is located in.
- **organization_id** (String) The ID of the organization this HCP Vault cluster is located in.
- **project_id** (String) The ID of the project this HCP Vault cluster is located in.
- **public_endpoint** (Boolean) Denotes that the cluster has a public endpoint. Defaults to false.
- **region** (String) The region where the HCP Vault cluster is located.
- **tier** (String) The tier that the HCP Vault cluster will be provisioned as.  Only 'development' is available at this time.
- **vault_private_endpoint_url** (String) The private URL for the Vault cluster.
- **vault_public_endpoint_url** (String) The public URL for the Vault cluster. This will be empty if `public_endpoint` is `false`.
- **vault_version** (String) The Vault version of the cluster.

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- **default** (String)

