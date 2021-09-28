# Terraform IIS Provider

Terraform Provider using the [Microsoft IIS Administration](https://docs.microsoft.com/en-us/IIS-Administration/) API.

# Usage

this is a test

## Setup

```hcl
provider "iis" {
  access_key = "your access key"
  host = "https://localhost:55539"
}
```

## Application pools

```hcl
resource "iis_application_pool" "name" {
  name = "AppPool" // Name of the Application Pool
  runtime = "" // (optional) null or empty for "No Managed Code" or "vX.Y" to specify version
}
```

## Directories

```hcl
resource "iis_file" "name" {
  name = "name.of.your.directory"
  physical_path = "%systemdrive%\\inetpub\\your_app" // full path to directory
  type = "directory" // can also be "file"
  parent = "parent_id" // id of the parent folder
}
```

## Websites

```hcl
data "iis_website" "website-domain-com" {
  name = "website.domain.com"
  physical_path = "%systemdrive%\\inetpub\\your_app" // full path to website folder
  application_pool = iis_application_pool.name.id
  bindings   {
    protocol = "http"
    port = 80
    ip_address = "*"
    hostname = "website.domain.com"
  }
}
```
