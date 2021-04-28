# Terraform IIS Provider

Terraform Provider using the [Microsoft IIS Administration](https://docs.microsoft.com/en-us/IIS-Administration/) API.

# Usage

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
  binding {
    protocol = "http"
    port = 80
    ip_address = "*"
    hostname = "website.domain.com"
  }
}
```
