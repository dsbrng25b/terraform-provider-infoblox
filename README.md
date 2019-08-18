# Terraform Provider for Infoblox

## Requirements
* [Terraform](https://www.terraform.io/downloads.html) 0.11.x
* [Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)
* [dep](https://github.com/golang/dep)
* Extensible Attributes: If you use Infoblox in a on-prem environment and you don't have the CNA license you have to create the following extensible attributes in the admin GUI under Administration -> Extensible Attributes:
  * `VM Name`
  * `VM ID`
  * `Tenant ID`
  * `CMP Type`
  * `Cloud API Owned`

## Building the Provider

clone repository to ```$GOPATH/src/github.com/infobloxopen/terraform-provider-infoblox```
```
$ mkdir -p $GOPATH/src/github.com/infobloxopen; cd $GOPATH/src/github.com/infobloxopen
$ git clone git@github.com:infobloxopen/terraform-provider-infoblox
```
Enter the provider directory and build the provider
```
$ cd $GOPATH/src/github.com/infobloxopen/terraform-provider-infoblox
$ make build
```

## Using the Provider
If you're building the provider, follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After the build is complete, copy the ```terraform-provider-infoblox``` binary into the same path as your terraform binary.After placing it into your plugins directory, run ```terraform init``` to initialize it.


## Developing the Provider
If you wish to work on the provider, you'll first need Go installed on your machine (version 1.9+ is required). You'll also need to correctly setup a GOPATH, as well as adding ```$GOPATH/bin``` to your ```$PATH```.

To compile the provider, run ```make build```. This will build the provider and put the provider binary in the ```$GOPATH/bin``` directory.
```
$ make build
...
$ $GOPATH/bin/terraform-provider-infoblox
...
```
In order to test the provider, you can simply run make test.
```
$ make test
```
## Features of Provider
* Creation of Network View in NIOS appliance
* Creation &  Deletion of Network in NIOS appliance
* Allocation & Deallocation of IP from a Network
* Association & Disassociation of IP Address for a VM
* Creation and Deletion of A,CNAME,Host and Ptr records.

## Disclaimer
 To use the provider for DNS purposes , a parent (i.e zone) must be created, The plugin doesn't allow creation of zone. 
