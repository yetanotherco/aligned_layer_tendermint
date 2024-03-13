terraform {
  required_providers {
    hcloud = {
      source = "hetznercloud/hcloud"
      version = "1.42.1"
    }
  }
}

# Set the variable value in *.tfvars file
# or using the -var="hcloud_token=..." CLI option
variable "hcloud_token" {
  sensitive = true # Requires terraform >= 0.14
}

provider "hcloud" {
	token = var.hcloud_token
}
