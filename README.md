# Golibvirt - The virtualization API
[![Gitter](https://badges.gitter.im/Join Chat.svg)](https://gitter.im/hooklift/golibvirt?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![GoDoc](https://godoc.org/github.com/cloudescape/golibvirt?status.svg)](https://godoc.org/github.com/cloudescape/golibvirt)
[![Build Status](https://travis-ci.org/cloudescape/golibvirt.svg?branch=master)](https://travis-ci.org/cloudescape/golibvirt)

Libvirt is a toolkit to interact with virtualization capabilities
of recent versions of GNU/Linux (and other OSes). (reference: [Libvirt][libvirt_home])

Golibvirt is a set of bindings to the Libvirt C API, which allows to use it from Go

### Virtualization Technologies Supported

   * The [Xen][xen_home] hypervisor on Linux and Solaris hosts.
   * [QEMU][qemu_home] emulator
   * [KVM][kvm_home] Linux hypervisor
   * [LXC][lxc_home] Linux container system
   * [OpenVZ][openvz_home] Linux container system
   * [User Mode Linux][user_mode_linux_home] paravirtualized kernel
   * [VirtualBox][virtualbox_home] hypervisor
   * [VMware ESX and GSX][vmware_home] hypervisors
   * [IBM Power][phyp_home] Hypervisor


### Capabilities

   * Management of virtual machines, virtual networks and storage
   * Remote management using TLS encryption and x509 certificates
   * Remote management authenticating with Kerberos and SASL
   * Local access control using PolicyKit
   * Zero-conf discovery using Avahi multicast-DNS
   * Support for storage on IDE/SCSI/USB disks, FibreChannel, LVM, iSCSI, NFS and filesystems

## Installation
You must have [Go](http://golang.org) and [Libvirt][libvirt_dev] already installed to be able to build Golibvirt.

### Get and install Golibvirt
    $ go get github.com/cloudescape/golibvirt

## Example of use
For now, please take a look at the tests. Further ahead I will be releasing more meaninful examples
    
## API
http://godoc.org/github.com/cloudescape/golibvirt

[libvirt_home]: http://www.libvirt.org
[libvirt_dev]: http://libvirt.org/deployment.html
[xen_home]: http://www.cl.cam.ac.uk/Research/SRG/netos/xen/index.html
[qemu_home]: http://wiki.qemu.org/Index.html
[kvm_home]: http://www.linux-kvm.org
[lxc_home]: http://lxc.sourceforge.net/
[openvz_home]: http://openvz.org/
[user_mode_linux_home]: http://user-mode-linux.sourceforge.net/
[virtualbox_home]: http://www.virtualbox.org/
[vmware_home]: http://www.vmware.com/
[phyp_home]: http://www.ibm.com/developerworks/wikis/display/LinuxP/POWER5+Hypervisor

