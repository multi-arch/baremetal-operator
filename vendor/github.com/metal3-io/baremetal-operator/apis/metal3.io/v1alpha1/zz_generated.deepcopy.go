//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BIOS) DeepCopyInto(out *BIOS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BIOS.
func (in *BIOS) DeepCopy() *BIOS {
	if in == nil {
		return nil
	}
	out := new(BIOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BMCDetails) DeepCopyInto(out *BMCDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BMCDetails.
func (in *BMCDetails) DeepCopy() *BMCDetails {
	if in == nil {
		return nil
	}
	out := new(BMCDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalHost) DeepCopyInto(out *BareMetalHost) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalHost.
func (in *BareMetalHost) DeepCopy() *BareMetalHost {
	if in == nil {
		return nil
	}
	out := new(BareMetalHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalHost) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalHostList) DeepCopyInto(out *BareMetalHostList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BareMetalHost, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalHostList.
func (in *BareMetalHostList) DeepCopy() *BareMetalHostList {
	if in == nil {
		return nil
	}
	out := new(BareMetalHostList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BareMetalHostList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalHostSpec) DeepCopyInto(out *BareMetalHostSpec) {
	*out = *in
	if in.Taints != nil {
		in, out := &in.Taints, &out.Taints
		*out = make([]v1.Taint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.BMC = in.BMC
	if in.RAID != nil {
		in, out := &in.RAID, &out.RAID
		*out = new(RAIDConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(FirmwareConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RootDeviceHints != nil {
		in, out := &in.RootDeviceHints, &out.RootDeviceHints
		*out = new(RootDeviceHints)
		(*in).DeepCopyInto(*out)
	}
	if in.ConsumerRef != nil {
		in, out := &in.ConsumerRef, &out.ConsumerRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(Image)
		(*in).DeepCopyInto(*out)
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.NetworkData != nil {
		in, out := &in.NetworkData, &out.NetworkData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.MetaData != nil {
		in, out := &in.MetaData, &out.MetaData
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.CustomDeploy != nil {
		in, out := &in.CustomDeploy, &out.CustomDeploy
		*out = new(CustomDeploy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalHostSpec.
func (in *BareMetalHostSpec) DeepCopy() *BareMetalHostSpec {
	if in == nil {
		return nil
	}
	out := new(BareMetalHostSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BareMetalHostStatus) DeepCopyInto(out *BareMetalHostStatus) {
	*out = *in
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.HardwareDetails != nil {
		in, out := &in.HardwareDetails, &out.HardwareDetails
		*out = new(HardwareDetails)
		(*in).DeepCopyInto(*out)
	}
	in.Provisioning.DeepCopyInto(&out.Provisioning)
	in.GoodCredentials.DeepCopyInto(&out.GoodCredentials)
	in.TriedCredentials.DeepCopyInto(&out.TriedCredentials)
	in.OperationHistory.DeepCopyInto(&out.OperationHistory)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BareMetalHostStatus.
func (in *BareMetalHostStatus) DeepCopy() *BareMetalHostStatus {
	if in == nil {
		return nil
	}
	out := new(BareMetalHostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPU) DeepCopyInto(out *CPU) {
	*out = *in
	if in.Flags != nil {
		in, out := &in.Flags, &out.Flags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPU.
func (in *CPU) DeepCopy() *CPU {
	if in == nil {
		return nil
	}
	out := new(CPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsStatus) DeepCopyInto(out *CredentialsStatus) {
	*out = *in
	if in.Reference != nil {
		in, out := &in.Reference, &out.Reference
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsStatus.
func (in *CredentialsStatus) DeepCopy() *CredentialsStatus {
	if in == nil {
		return nil
	}
	out := new(CredentialsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDeploy) DeepCopyInto(out *CustomDeploy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDeploy.
func (in *CustomDeploy) DeepCopy() *CustomDeploy {
	if in == nil {
		return nil
	}
	out := new(CustomDeploy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DesiredSettingsMap) DeepCopyInto(out *DesiredSettingsMap) {
	{
		in := &in
		*out = make(DesiredSettingsMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DesiredSettingsMap.
func (in DesiredSettingsMap) DeepCopy() DesiredSettingsMap {
	if in == nil {
		return nil
	}
	out := new(DesiredSettingsMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Firmware) DeepCopyInto(out *Firmware) {
	*out = *in
	out.BIOS = in.BIOS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Firmware.
func (in *Firmware) DeepCopy() *Firmware {
	if in == nil {
		return nil
	}
	out := new(Firmware)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwareConfig) DeepCopyInto(out *FirmwareConfig) {
	*out = *in
	if in.VirtualizationEnabled != nil {
		in, out := &in.VirtualizationEnabled, &out.VirtualizationEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SimultaneousMultithreadingEnabled != nil {
		in, out := &in.SimultaneousMultithreadingEnabled, &out.SimultaneousMultithreadingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SriovEnabled != nil {
		in, out := &in.SriovEnabled, &out.SriovEnabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwareConfig.
func (in *FirmwareConfig) DeepCopy() *FirmwareConfig {
	if in == nil {
		return nil
	}
	out := new(FirmwareConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwareSchema) DeepCopyInto(out *FirmwareSchema) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwareSchema.
func (in *FirmwareSchema) DeepCopy() *FirmwareSchema {
	if in == nil {
		return nil
	}
	out := new(FirmwareSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirmwareSchema) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwareSchemaList) DeepCopyInto(out *FirmwareSchemaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FirmwareSchema, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwareSchemaList.
func (in *FirmwareSchemaList) DeepCopy() *FirmwareSchemaList {
	if in == nil {
		return nil
	}
	out := new(FirmwareSchemaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirmwareSchemaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwareSchemaSpec) DeepCopyInto(out *FirmwareSchemaSpec) {
	*out = *in
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = make(map[string]SettingSchema, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwareSchemaSpec.
func (in *FirmwareSchemaSpec) DeepCopy() *FirmwareSchemaSpec {
	if in == nil {
		return nil
	}
	out := new(FirmwareSchemaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareDetails) DeepCopyInto(out *HardwareDetails) {
	*out = *in
	out.SystemVendor = in.SystemVendor
	out.Firmware = in.Firmware
	if in.NIC != nil {
		in, out := &in.NIC, &out.NIC
		*out = make([]NIC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = make([]Storage, len(*in))
		copy(*out, *in)
	}
	in.CPU.DeepCopyInto(&out.CPU)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareDetails.
func (in *HardwareDetails) DeepCopy() *HardwareDetails {
	if in == nil {
		return nil
	}
	out := new(HardwareDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareRAIDVolume) DeepCopyInto(out *HardwareRAIDVolume) {
	*out = *in
	if in.SizeGibibytes != nil {
		in, out := &in.SizeGibibytes, &out.SizeGibibytes
		*out = new(int)
		**out = **in
	}
	if in.Rotational != nil {
		in, out := &in.Rotational, &out.Rotational
		*out = new(bool)
		**out = **in
	}
	if in.NumberOfPhysicalDisks != nil {
		in, out := &in.NumberOfPhysicalDisks, &out.NumberOfPhysicalDisks
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareRAIDVolume.
func (in *HardwareRAIDVolume) DeepCopy() *HardwareRAIDVolume {
	if in == nil {
		return nil
	}
	out := new(HardwareRAIDVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareSystemVendor) DeepCopyInto(out *HardwareSystemVendor) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareSystemVendor.
func (in *HardwareSystemVendor) DeepCopy() *HardwareSystemVendor {
	if in == nil {
		return nil
	}
	out := new(HardwareSystemVendor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostFirmwareSettings) DeepCopyInto(out *HostFirmwareSettings) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostFirmwareSettings.
func (in *HostFirmwareSettings) DeepCopy() *HostFirmwareSettings {
	if in == nil {
		return nil
	}
	out := new(HostFirmwareSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostFirmwareSettings) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostFirmwareSettingsList) DeepCopyInto(out *HostFirmwareSettingsList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HostFirmwareSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostFirmwareSettingsList.
func (in *HostFirmwareSettingsList) DeepCopy() *HostFirmwareSettingsList {
	if in == nil {
		return nil
	}
	out := new(HostFirmwareSettingsList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HostFirmwareSettingsList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostFirmwareSettingsSpec) DeepCopyInto(out *HostFirmwareSettingsSpec) {
	*out = *in
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(DesiredSettingsMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostFirmwareSettingsSpec.
func (in *HostFirmwareSettingsSpec) DeepCopy() *HostFirmwareSettingsSpec {
	if in == nil {
		return nil
	}
	out := new(HostFirmwareSettingsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostFirmwareSettingsStatus) DeepCopyInto(out *HostFirmwareSettingsStatus) {
	*out = *in
	if in.FirmwareSchema != nil {
		in, out := &in.FirmwareSchema, &out.FirmwareSchema
		*out = new(SchemaReference)
		**out = **in
	}
	if in.Settings != nil {
		in, out := &in.Settings, &out.Settings
		*out = make(SettingsMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.LastUpdated != nil {
		in, out := &in.LastUpdated, &out.LastUpdated
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostFirmwareSettingsStatus.
func (in *HostFirmwareSettingsStatus) DeepCopy() *HostFirmwareSettingsStatus {
	if in == nil {
		return nil
	}
	out := new(HostFirmwareSettingsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	if in.DiskFormat != nil {
		in, out := &in.DiskFormat, &out.DiskFormat
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NIC) DeepCopyInto(out *NIC) {
	*out = *in
	if in.VLANs != nil {
		in, out := &in.VLANs, &out.VLANs
		*out = make([]VLAN, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NIC.
func (in *NIC) DeepCopy() *NIC {
	if in == nil {
		return nil
	}
	out := new(NIC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationHistory) DeepCopyInto(out *OperationHistory) {
	*out = *in
	in.Register.DeepCopyInto(&out.Register)
	in.Inspect.DeepCopyInto(&out.Inspect)
	in.Provision.DeepCopyInto(&out.Provision)
	in.Deprovision.DeepCopyInto(&out.Deprovision)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationHistory.
func (in *OperationHistory) DeepCopy() *OperationHistory {
	if in == nil {
		return nil
	}
	out := new(OperationHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OperationMetric) DeepCopyInto(out *OperationMetric) {
	*out = *in
	in.Start.DeepCopyInto(&out.Start)
	in.End.DeepCopyInto(&out.End)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OperationMetric.
func (in *OperationMetric) DeepCopy() *OperationMetric {
	if in == nil {
		return nil
	}
	out := new(OperationMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreprovisioningImage) DeepCopyInto(out *PreprovisioningImage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreprovisioningImage.
func (in *PreprovisioningImage) DeepCopy() *PreprovisioningImage {
	if in == nil {
		return nil
	}
	out := new(PreprovisioningImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreprovisioningImage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreprovisioningImageList) DeepCopyInto(out *PreprovisioningImageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PreprovisioningImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreprovisioningImageList.
func (in *PreprovisioningImageList) DeepCopy() *PreprovisioningImageList {
	if in == nil {
		return nil
	}
	out := new(PreprovisioningImageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PreprovisioningImageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreprovisioningImageSpec) DeepCopyInto(out *PreprovisioningImageSpec) {
	*out = *in
	if in.AcceptFormats != nil {
		in, out := &in.AcceptFormats, &out.AcceptFormats
		*out = make([]ImageFormat, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreprovisioningImageSpec.
func (in *PreprovisioningImageSpec) DeepCopy() *PreprovisioningImageSpec {
	if in == nil {
		return nil
	}
	out := new(PreprovisioningImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreprovisioningImageStatus) DeepCopyInto(out *PreprovisioningImageStatus) {
	*out = *in
	out.NetworkData = in.NetworkData
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreprovisioningImageStatus.
func (in *PreprovisioningImageStatus) DeepCopy() *PreprovisioningImageStatus {
	if in == nil {
		return nil
	}
	out := new(PreprovisioningImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionStatus) DeepCopyInto(out *ProvisionStatus) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	if in.RootDeviceHints != nil {
		in, out := &in.RootDeviceHints, &out.RootDeviceHints
		*out = new(RootDeviceHints)
		(*in).DeepCopyInto(*out)
	}
	if in.RAID != nil {
		in, out := &in.RAID, &out.RAID
		*out = new(RAIDConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(FirmwareConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CustomDeploy != nil {
		in, out := &in.CustomDeploy, &out.CustomDeploy
		*out = new(CustomDeploy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionStatus.
func (in *ProvisionStatus) DeepCopy() *ProvisionStatus {
	if in == nil {
		return nil
	}
	out := new(ProvisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RAIDConfig) DeepCopyInto(out *RAIDConfig) {
	*out = *in
	if in.HardwareRAIDVolumes != nil {
		in, out := &in.HardwareRAIDVolumes, &out.HardwareRAIDVolumes
		*out = make([]HardwareRAIDVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SoftwareRAIDVolumes != nil {
		in, out := &in.SoftwareRAIDVolumes, &out.SoftwareRAIDVolumes
		*out = make([]SoftwareRAIDVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RAIDConfig.
func (in *RAIDConfig) DeepCopy() *RAIDConfig {
	if in == nil {
		return nil
	}
	out := new(RAIDConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RebootAnnotationArguments) DeepCopyInto(out *RebootAnnotationArguments) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RebootAnnotationArguments.
func (in *RebootAnnotationArguments) DeepCopy() *RebootAnnotationArguments {
	if in == nil {
		return nil
	}
	out := new(RebootAnnotationArguments)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootDeviceHints) DeepCopyInto(out *RootDeviceHints) {
	*out = *in
	if in.Rotational != nil {
		in, out := &in.Rotational, &out.Rotational
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootDeviceHints.
func (in *RootDeviceHints) DeepCopy() *RootDeviceHints {
	if in == nil {
		return nil
	}
	out := new(RootDeviceHints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaReference) DeepCopyInto(out *SchemaReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaReference.
func (in *SchemaReference) DeepCopy() *SchemaReference {
	if in == nil {
		return nil
	}
	out := new(SchemaReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchemaSettingError) DeepCopyInto(out *SchemaSettingError) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchemaSettingError.
func (in *SchemaSettingError) DeepCopy() *SchemaSettingError {
	if in == nil {
		return nil
	}
	out := new(SchemaSettingError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretStatus) DeepCopyInto(out *SecretStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretStatus.
func (in *SecretStatus) DeepCopy() *SecretStatus {
	if in == nil {
		return nil
	}
	out := new(SecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SettingSchema) DeepCopyInto(out *SettingSchema) {
	*out = *in
	if in.AllowableValues != nil {
		in, out := &in.AllowableValues, &out.AllowableValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LowerBound != nil {
		in, out := &in.LowerBound, &out.LowerBound
		*out = new(int)
		**out = **in
	}
	if in.UpperBound != nil {
		in, out := &in.UpperBound, &out.UpperBound
		*out = new(int)
		**out = **in
	}
	if in.MinLength != nil {
		in, out := &in.MinLength, &out.MinLength
		*out = new(int)
		**out = **in
	}
	if in.MaxLength != nil {
		in, out := &in.MaxLength, &out.MaxLength
		*out = new(int)
		**out = **in
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
	if in.ResetRequired != nil {
		in, out := &in.ResetRequired, &out.ResetRequired
		*out = new(bool)
		**out = **in
	}
	if in.Unique != nil {
		in, out := &in.Unique, &out.Unique
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingSchema.
func (in *SettingSchema) DeepCopy() *SettingSchema {
	if in == nil {
		return nil
	}
	out := new(SettingSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in SettingsMap) DeepCopyInto(out *SettingsMap) {
	{
		in := &in
		*out = make(SettingsMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SettingsMap.
func (in SettingsMap) DeepCopy() SettingsMap {
	if in == nil {
		return nil
	}
	out := new(SettingsMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SoftwareRAIDVolume) DeepCopyInto(out *SoftwareRAIDVolume) {
	*out = *in
	if in.SizeGibibytes != nil {
		in, out := &in.SizeGibibytes, &out.SizeGibibytes
		*out = new(int)
		**out = **in
	}
	if in.PhysicalDisks != nil {
		in, out := &in.PhysicalDisks, &out.PhysicalDisks
		*out = make([]RootDeviceHints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SoftwareRAIDVolume.
func (in *SoftwareRAIDVolume) DeepCopy() *SoftwareRAIDVolume {
	if in == nil {
		return nil
	}
	out := new(SoftwareRAIDVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VLAN) DeepCopyInto(out *VLAN) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VLAN.
func (in *VLAN) DeepCopy() *VLAN {
	if in == nil {
		return nil
	}
	out := new(VLAN)
	in.DeepCopyInto(out)
	return out
}
