# notion_user Data source

Provides a Notion user data source.

## Example Usage

```hcl
# Data source for Notion user
data "notion_user" "notion" {
    id        = "masked"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Required) User ID.
* `avatar_url` - (Optional) URL of the avatar.
* `name` - (Optional) Name of the user
* `person` - (Optional) Person object
* `bot` - (Optional) Bot object

## Attribute Reference

* `person` - For person users
    * `email` - Email of the user
* `bot` - For bot users. As for 2021-05-16, no attribute is provider
