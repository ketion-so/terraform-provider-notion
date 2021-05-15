# Notion Provider

The [Notion](https://notion.so/) provider is used to interact with the
Notion API resources. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Notion provider
provider "notion" {
  access_token = var.access_token
}
```

## Argument Reference

The following arguments are supported:

* `access_token` - (Required) Access token for Notion API. If not set, the provider will see the `NOTION_ACCESS_TOKEN` environment variable.
