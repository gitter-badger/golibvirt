// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>

static void customErrorFunc(void *userdata, virErrorPtr err) {
	//This empty custom error handler function
	//avoids libvirt printing
	//errors in stderr and lets Go take over seamlessly
}

static void setCustomErrorHandler() {
	virSetErrorFunc(NULL, customErrorFunc);
}
*/
import "C"

//virErrorDomain
const (
	VIR_FROM_NONE            = C.VIR_FROM_NONE
	VIR_FROM_XEN             = C.VIR_FROM_XEN             //Error at Xen hypervisor layer
	VIR_FROM_XEND            = C.VIR_FROM_XEND            //Error at connection with xend daemon
	VIR_FROM_XENSTORE        = C.VIR_FROM_XENSTORE        //Error at connection with xen store
	VIR_FROM_SEXPR           = C.VIR_FROM_SEXPR           //Error in the S-Expression code
	VIR_FROM_XML             = C.VIR_FROM_XML             //Error in the XML code
	VIR_FROM_DOM             = C.VIR_FROM_DOM             //Error when operating on a domain
	VIR_FROM_RPC             = C.VIR_FROM_RPC             //Error in the XML-RPC code
	VIR_FROM_CONF            = C.VIR_FROM_CONF            //Error in the configuration file handling
	VIR_FROM_QEMU            = C.VIR_FROM_QEMU            //Error at the QEMU daemon
	VIR_FROM_NET             = C.VIR_FROM_NET             //Error when operating on a network
	VIR_FROM_TEST            = C.VIR_FROM_TEST            //Error from test driver
	VIR_FROM_REMOTE          = C.VIR_FROM_REMOTE          //Error from remote driver
	VIR_FROM_OPENVZ          = C.VIR_FROM_OPENVZ          //Error from OpenVZ driver
	VIR_FROM_XENXM           = C.VIR_FROM_XENXM           //Error at Xen XM layer
	VIR_FROM_STATS_LINUX     = C.VIR_FROM_STATS_LINUX     //Error in the Linux Stats code
	VIR_FROM_LXC             = C.VIR_FROM_LXC             //Error from Linux Container driver
	VIR_FROM_STORAGE         = C.VIR_FROM_STORAGE         //Error from storage driver
	VIR_FROM_NETWORK         = C.VIR_FROM_NETWORK         //Error from network config
	VIR_FROM_DOMAIN          = C.VIR_FROM_DOMAIN          //Error from domain config
	VIR_FROM_UML             = C.VIR_FROM_UML             //Error at the UML driver
	VIR_FROM_NODEDEV         = C.VIR_FROM_NODEDEV         //Error from node device monitor
	VIR_FROM_XEN_INOTIFY     = C.VIR_FROM_XEN_INOTIFY     //Error from xen inotify layer
	VIR_FROM_SECURITY        = C.VIR_FROM_SECURITY        //Error from security framework
	VIR_FROM_VBOX            = C.VIR_FROM_VBOX            //Error from VirtualBox driver
	VIR_FROM_INTERFACE       = C.VIR_FROM_INTERFACE       //Error when operating on an interface
	VIR_FROM_ESX             = C.VIR_FROM_ESX             //Error from ESX driver
	VIR_FROM_PHYP            = C.VIR_FROM_PHYP            //Error from IBM power hypervisor
	VIR_FROM_SECRET          = C.VIR_FROM_SECRET          //Error from secret storage
	VIR_FROM_CPU             = C.VIR_FROM_CPU             //Error from CPU driver
	VIR_FROM_XENAPI          = C.VIR_FROM_XENAPI          //Error from XenAPI
	VIR_FROM_NWFILTER        = C.VIR_FROM_NWFILTER        //Error from network filter driver
	VIR_FROM_HOOK            = C.VIR_FROM_HOOK            //Error from Synchronous hooks
	VIR_FROM_DOMAIN_SNAPSHOT = C.VIR_FROM_DOMAIN_SNAPSHOT //Error from domain snapshot
	VIR_FROM_AUDIT           = C.VIR_FROM_AUDIT           //Error from auditing subsystem
	VIR_FROM_SYSINFO         = C.VIR_FROM_SYSINFO         //Error from sysinfo/SMBIOS
	VIR_FROM_STREAMS         = C.VIR_FROM_STREAMS         //Error from I/O streams
	VIR_FROM_VMWARE          = C.VIR_FROM_VMWARE          //Error from VMware driver
	VIR_FROM_EVENT           = C.VIR_FROM_EVENT           //Error from event loop impl
	VIR_FROM_LIBXL           = C.VIR_FROM_LIBXL           //Error from libxenlight driver
	VIR_FROM_LOCKING         = C.VIR_FROM_LOCKING         //Error from lock manager
	VIR_FROM_HYPERV          = C.VIR_FROM_HYPERV          //Error from Hyper-V driver
	VIR_FROM_CAPABILITIES    = C.VIR_FROM_CAPABILITIES    //Error from capabilities
	VIR_FROM_URI             = C.VIR_FROM_URI             //Error from URI handling
	VIR_FROM_AUTH            = C.VIR_FROM_AUTH            //Error from auth handling
	VIR_FROM_DBUS            = C.VIR_FROM_DBUS            //Error from DBus
	VIR_FROM_PARALLELS       = C.VIR_FROM_PARALLELS       //Error from Parallels
	VIR_FROM_DEVICE          = C.VIR_FROM_DEVICE          //Error from Device
	VIR_FROM_SSH             = C.VIR_FROM_SSH             //Error from libssh2 connection transport
	VIR_FROM_LOCKSPACE       = C.VIR_FROM_LOCKSPACE       //Error from lockspace
	VIR_FROM_INITCTL         = C.VIR_FROM_INITCTL         //Error from initctl device communication
	VIR_FROM_IDENTITY        = C.VIR_FROM_IDENTITY        //Error from identity code
	VIR_FROM_CGROUP          = C.VIR_FROM_CGROUP          //Error from cgroups
	VIR_FROM_ACCESS          = C.VIR_FROM_ACCESS          //Error from access control manager
	VIR_FROM_SYSTEMD         = C.VIR_FROM_SYSTEMD         //Error from systemd code
	//VIR_ERR_DOMAIN_LAST      = C.VIR_ERR_DOMAIN_LAST
)

//virErrorLevel
const (
	VIR_ERR_NONE    = C.VIR_ERR_NONE
	VIR_ERR_WARNING = C.VIR_ERR_WARNING
	VIR_ERR_ERROR   = C.VIR_ERR_ERROR
)

//virErrorNumber
const (
	VIR_ERR_OK                      = C.VIR_ERR_OK
	VIR_ERR_INTERNAL_ERROR          = C.VIR_ERR_INTERNAL_ERROR          //internal error
	VIR_ERR_NO_MEMORY               = C.VIR_ERR_NO_MEMORY               //memory allocation failure
	VIR_ERR_NO_SUPPORT              = C.VIR_ERR_NO_SUPPORT              //no support for this function
	VIR_ERR_UNKNOWN_HOST            = C.VIR_ERR_UNKNOWN_HOST            //could not resolve hostname
	VIR_ERR_NO_CONNECT              = C.VIR_ERR_NO_CONNECT              //can't connect to hypervisor
	VIR_ERR_INVALID_CONN            = C.VIR_ERR_INVALID_CONN            //invalid connection object
	VIR_ERR_INVALID_DOMAIN          = C.VIR_ERR_INVALID_DOMAIN          //invalid domain object
	VIR_ERR_INVALID_ARG             = C.VIR_ERR_INVALID_ARG             //invalid function argument
	VIR_ERR_OPERATION_FAILED        = C.VIR_ERR_OPERATION_FAILED        //a command to hypervisor failed
	VIR_ERR_GET_FAILED              = C.VIR_ERR_GET_FAILED              //a HTTP GET command to failed
	VIR_ERR_POST_FAILED             = C.VIR_ERR_POST_FAILED             //a HTTP POST command to failed
	VIR_ERR_HTTP_ERROR              = C.VIR_ERR_HTTP_ERROR              //unexpected HTTP error code
	VIR_ERR_SEXPR_SERIAL            = C.VIR_ERR_SEXPR_SERIAL            //failure to serialize an S-Expr
	VIR_ERR_NO_XEN                  = C.VIR_ERR_NO_XEN                  //could not open Xen hypervisor control
	VIR_ERR_XEN_CALL                = C.VIR_ERR_XEN_CALL                //failure doing an hypervisor call
	VIR_ERR_OS_TYPE                 = C.VIR_ERR_OS_TYPE                 //unknown OS type
	VIR_ERR_NO_KERNEL               = C.VIR_ERR_NO_KERNEL               //missing kernel information
	VIR_ERR_NO_ROOT                 = C.VIR_ERR_NO_ROOT                 //missing root device information
	VIR_ERR_NO_SOURCE               = C.VIR_ERR_NO_SOURCE               //missing source device information
	VIR_ERR_NO_TARGET               = C.VIR_ERR_NO_TARGET               //missing target device information
	VIR_ERR_NO_NAME                 = C.VIR_ERR_NO_NAME                 //missing domain name information
	VIR_ERR_NO_OS                   = C.VIR_ERR_NO_OS                   //missing domain OS information
	VIR_ERR_NO_DEVICE               = C.VIR_ERR_NO_DEVICE               //issing domain devices information
	VIR_ERR_NO_XENSTORE             = C.VIR_ERR_NO_XENSTORE             //could not open Xen Store control
	VIR_ERR_DRIVER_FULL             = C.VIR_ERR_DRIVER_FULL             //too many drivers registered
	VIR_ERR_XML_ERROR               = C.VIR_ERR_XML_ERROR               //an XML description is not well formed or broken
	VIR_ERR_DOM_EXIST               = C.VIR_ERR_DOM_EXIST               //the domain already exist
	VIR_ERR_OPERATION_DENIED        = C.VIR_ERR_OPERATION_DENIED        //operation forbidden on read-only connections
	VIR_ERR_OPEN_FAILED             = C.VIR_ERR_OPEN_FAILED             //failed to open a conf file
	VIR_ERR_READ_FAILED             = C.VIR_ERR_READ_FAILED             //failed to read a conf file
	VIR_ERR_PARSE_FAILED            = C.VIR_ERR_PARSE_FAILED            //failed to parse a conf file
	VIR_ERR_CONF_SYNTAX             = C.VIR_ERR_CONF_SYNTAX             //failed to parse the syntax of a conf file
	VIR_ERR_WRITE_FAILED            = C.VIR_ERR_WRITE_FAILED            //failed to write a conf file
	VIR_ERR_XML_DETAIL              = C.VIR_ERR_XML_DETAIL              //detail of an XML error
	VIR_ERR_INVALID_NETWORK         = C.VIR_ERR_INVALID_NETWORK         //invalid network object
	VIR_ERR_NETWORK_EXIST           = C.VIR_ERR_NETWORK_EXIST           //the network already exist
	VIR_ERR_SYSTEM_ERROR            = C.VIR_ERR_SYSTEM_ERROR            //general system call failure
	VIR_ERR_RPC                     = C.VIR_ERR_RPC                     //some sort of RPC error
	VIR_ERR_GNUTLS_ERROR            = C.VIR_ERR_GNUTLS_ERROR            //error from a GNUTLS call
	VIR_WAR_NO_NETWORK              = C.VIR_WAR_NO_NETWORK              //failed to start network
	VIR_ERR_NO_DOMAIN               = C.VIR_ERR_NO_DOMAIN               //domain not found or unexpectedly disappeared
	VIR_ERR_NO_NETWORK              = C.VIR_ERR_NO_NETWORK              //network not found
	VIR_ERR_INVALID_MAC             = C.VIR_ERR_INVALID_MAC             //invalid MAC address
	VIR_ERR_AUTH_FAILED             = C.VIR_ERR_AUTH_FAILED             //authentication failed
	VIR_ERR_INVALID_STORAGE_POOL    = C.VIR_ERR_INVALID_STORAGE_POOL    //invalid storage pool object
	VIR_ERR_INVALID_STORAGE_VOL     = C.VIR_ERR_INVALID_STORAGE_VOL     //invalid storage vol object
	VIR_WAR_NO_STORAGE              = C.VIR_WAR_NO_STORAGE              //failed to start storage
	VIR_ERR_NO_STORAGE_POOL         = C.VIR_ERR_NO_STORAGE_POOL         //storage pool not found
	VIR_ERR_NO_STORAGE_VOL          = C.VIR_ERR_NO_STORAGE_VOL          //storage volume not found
	VIR_WAR_NO_NODE                 = C.VIR_WAR_NO_NODE                 //failed to start node driver
	VIR_ERR_INVALID_NODE_DEVICE     = C.VIR_ERR_INVALID_NODE_DEVICE     //invalid node device object
	VIR_ERR_NO_NODE_DEVICE          = C.VIR_ERR_NO_NODE_DEVICE          //node device not found
	VIR_ERR_NO_SECURITY_MODEL       = C.VIR_ERR_NO_SECURITY_MODEL       //security model not found
	VIR_ERR_OPERATION_INVALID       = C.VIR_ERR_OPERATION_INVALID       //operation is not applicable at this time
	VIR_WAR_NO_INTERFACE            = C.VIR_WAR_NO_INTERFACE            //failed to start interface driver
	VIR_ERR_NO_INTERFACE            = C.VIR_ERR_NO_INTERFACE            //interface driver not running
	VIR_ERR_INVALID_INTERFACE       = C.VIR_ERR_INVALID_INTERFACE       //invalid interface object
	VIR_ERR_MULTIPLE_INTERFACES     = C.VIR_ERR_MULTIPLE_INTERFACES     //more than one matching interface found
	VIR_WAR_NO_NWFILTER             = C.VIR_WAR_NO_NWFILTER             //failed to start nwfilter driver
	VIR_ERR_INVALID_NWFILTER        = C.VIR_ERR_INVALID_NWFILTER        //invalid nwfilter object
	VIR_ERR_NO_NWFILTER             = C.VIR_ERR_NO_NWFILTER             //nw filter pool not found
	VIR_ERR_BUILD_FIREWALL          = C.VIR_ERR_BUILD_FIREWALL          //nw filter pool not found
	VIR_WAR_NO_SECRET               = C.VIR_WAR_NO_SECRET               //failed to start secret storage
	VIR_ERR_INVALID_SECRET          = C.VIR_ERR_INVALID_SECRET          //invalid secret
	VIR_ERR_NO_SECRET               = C.VIR_ERR_NO_SECRET               //secret not found
	VIR_ERR_CONFIG_UNSUPPORTED      = C.VIR_ERR_CONFIG_UNSUPPORTED      //unsupported configuration construct
	VIR_ERR_OPERATION_TIMEOUT       = C.VIR_ERR_OPERATION_TIMEOUT       //timeout occurred during operation
	VIR_ERR_MIGRATE_PERSIST_FAILED  = C.VIR_ERR_MIGRATE_PERSIST_FAILED  //a migration worked, but making the VM persist on the dest host failed
	VIR_ERR_HOOK_SCRIPT_FAILED      = C.VIR_ERR_HOOK_SCRIPT_FAILED      //a synchronous hook script failed
	VIR_ERR_INVALID_DOMAIN_SNAPSHOT = C.VIR_ERR_INVALID_DOMAIN_SNAPSHOT //invalid domain snapshot
	VIR_ERR_NO_DOMAIN_SNAPSHOT      = C.VIR_ERR_NO_DOMAIN_SNAPSHOT      //domain snapshot not found
	VIR_ERR_INVALID_STREAM          = C.VIR_ERR_INVALID_STREAM          //stream pointer not valid
	VIR_ERR_ARGUMENT_UNSUPPORTED    = C.VIR_ERR_ARGUMENT_UNSUPPORTED    //valid API use but unsupported by the given driver
	VIR_ERR_STORAGE_PROBE_FAILED    = C.VIR_ERR_STORAGE_PROBE_FAILED    //storage pool probe failed
	VIR_ERR_STORAGE_POOL_BUILT      = C.VIR_ERR_STORAGE_POOL_BUILT      //storage pool already built
	VIR_ERR_SNAPSHOT_REVERT_RISKY   = C.VIR_ERR_SNAPSHOT_REVERT_RISKY   //force was not requested for a risky domain snapshot revert
	VIR_ERR_OPERATION_ABORTED       = C.VIR_ERR_OPERATION_ABORTED       //operation on a domain was canceled/aborted by user
	VIR_ERR_AUTH_CANCELLED          = C.VIR_ERR_AUTH_CANCELLED          //authentication cancelled
	VIR_ERR_NO_DOMAIN_METADATA      = C.VIR_ERR_NO_DOMAIN_METADATA      //The metadata is not present
	VIR_ERR_MIGRATE_UNSAFE          = C.VIR_ERR_MIGRATE_UNSAFE          //Migration is not safe
	VIR_ERR_OVERFLOW                = C.VIR_ERR_OVERFLOW                //integer overflow
	VIR_ERR_BLOCK_COPY_ACTIVE       = C.VIR_ERR_BLOCK_COPY_ACTIVE       //action prevented by block copy job
	VIR_ERR_OPERATION_UNSUPPORTED   = C.VIR_ERR_OPERATION_UNSUPPORTED   //The requested operation is not supported
	VIR_ERR_SSH                     = C.VIR_ERR_SSH                     //error in ssh transport driver
	VIR_ERR_AGENT_UNRESPONSIVE      = C.VIR_ERR_AGENT_UNRESPONSIVE      //guest agent is unresponsive, not running or not usable
	VIR_ERR_RESOURCE_BUSY           = C.VIR_ERR_RESOURCE_BUSY           //resource is already in use
	VIR_ERR_ACCESS_DENIED           = C.VIR_ERR_ACCESS_DENIED           //operation on the object/resource was denied
	VIR_ERR_DBUS_SERVICE            = C.VIR_ERR_DBUS_SERVICE            //error from a dbus service
	//VIR_ERR_STORAGE_VOL_EXIST       = C.VIR_ERR_STORAGE_VOL_EXIST       //the storage vol already exists
)

type LibvirtError struct {
	ptr     C.virErrorPtr
	Code    int    //The error code, see virErrorNumbers
	Domain  int    //An enum indicating which part of libvirt raised the error see virErrorDomain
	Message string //the full human-readable formatted string of the error
	Level   int    //the error level, usually VIR_ERR_ERROR, though there is room for warnings like VIR_ERR_WARNING
	Str1    string //extra string information
	Str2    string //extra string information
	Str3    string //extra string information
	Int1    int    //extra number information
	Int2    int    //extra number information
}

func init() {
	C.setCustomErrorHandler()
}

func newLibvirtError(err C.virErrorPtr) *LibvirtError {
	if err == nil {
		return nil
	}

	return &LibvirtError{
		err,
		int(err.code),
		int(err.domain),
		C.GoString(err.message),
		int(err.level),
		C.GoString(err.str1),
		C.GoString(err.str2),
		C.GoString(err.str3),
		int(err.int1),
		int(err.int2),
	}
}

func GetLastError() *LibvirtError {
	err := C.virGetLastError()
	defer C.virResetError(err)

	return newLibvirtError(err)
}

func (e *LibvirtError) Error() string {
	return e.Message
}
