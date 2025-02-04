//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*

LICENSE

*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenance) DeepCopyInto(out *BackupLocationMaintenance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenance.
func (in *BackupLocationMaintenance) DeepCopy() *BackupLocationMaintenance {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupLocationMaintenance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceList) DeepCopyInto(out *BackupLocationMaintenanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupLocationMaintenance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceList.
func (in *BackupLocationMaintenanceList) DeepCopy() *BackupLocationMaintenanceList {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupLocationMaintenanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceSpec) DeepCopyInto(out *BackupLocationMaintenanceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceSpec.
func (in *BackupLocationMaintenanceSpec) DeepCopy() *BackupLocationMaintenanceSpec {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupLocationMaintenanceStatus) DeepCopyInto(out *BackupLocationMaintenanceStatus) {
	*out = *in
	if in.FullMaintenanceRepoStatus != nil {
		in, out := &in.FullMaintenanceRepoStatus, &out.FullMaintenanceRepoStatus
		*out = make(map[string]RepoMaintenanceStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.QuickMaintenanceRepoStatus != nil {
		in, out := &in.QuickMaintenanceRepoStatus, &out.QuickMaintenanceRepoStatus
		*out = make(map[string]RepoMaintenanceStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupLocationMaintenanceStatus.
func (in *BackupLocationMaintenanceStatus) DeepCopy() *BackupLocationMaintenanceStatus {
	if in == nil {
		return nil
	}
	out := new(BackupLocationMaintenanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExport) DeepCopyInto(out *DataExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExport.
func (in *DataExport) DeepCopy() *DataExport {
	if in == nil {
		return nil
	}
	out := new(DataExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportList) DeepCopyInto(out *DataExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportList.
func (in *DataExportList) DeepCopy() *DataExportList {
	if in == nil {
		return nil
	}
	out := new(DataExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportObjectReference) DeepCopyInto(out *DataExportObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportObjectReference.
func (in *DataExportObjectReference) DeepCopy() *DataExportObjectReference {
	if in == nil {
		return nil
	}
	out := new(DataExportObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataExportSpec) DeepCopyInto(out *DataExportSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataExportSpec.
func (in *DataExportSpec) DeepCopy() *DataExportSpec {
	if in == nil {
		return nil
	}
	out := new(DataExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExportStatus) DeepCopyInto(out *ExportStatus) {
	*out = *in
	if in.RestorePVC != nil {
		in, out := &in.RestorePVC, &out.RestorePVC
		*out = new(v1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExportStatus.
func (in *ExportStatus) DeepCopy() *ExportStatus {
	if in == nil {
		return nil
	}
	out := new(ExportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectInfo) DeepCopyInto(out *ObjectInfo) {
	*out = *in
	out.GroupVersionKind = in.GroupVersionKind
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectInfo.
func (in *ObjectInfo) DeepCopy() *ObjectInfo {
	if in == nil {
		return nil
	}
	out := new(ObjectInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepoMaintenanceStatus) DeepCopyInto(out *RepoMaintenanceStatus) {
	*out = *in
	in.LastRunTimestamp.DeepCopyInto(&out.LastRunTimestamp)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepoMaintenanceStatus.
func (in *RepoMaintenanceStatus) DeepCopy() *RepoMaintenanceStatus {
	if in == nil {
		return nil
	}
	out := new(RepoMaintenanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackup) DeepCopyInto(out *ResourceBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	if in.VolumesInfo != nil {
		in, out := &in.VolumesInfo, &out.VolumesInfo
		*out = make([]*ResourceBackupVolumeInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceBackupVolumeInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ExistingVolumesInfo != nil {
		in, out := &in.ExistingVolumesInfo, &out.ExistingVolumesInfo
		*out = make([]*ResourceRestoreVolumeInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceRestoreVolumeInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackup.
func (in *ResourceBackup) DeepCopy() *ResourceBackup {
	if in == nil {
		return nil
	}
	out := new(ResourceBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackupList) DeepCopyInto(out *ResourceBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackupList.
func (in *ResourceBackupList) DeepCopy() *ResourceBackupList {
	if in == nil {
		return nil
	}
	out := new(ResourceBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackupObjectReference) DeepCopyInto(out *ResourceBackupObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackupObjectReference.
func (in *ResourceBackupObjectReference) DeepCopy() *ResourceBackupObjectReference {
	if in == nil {
		return nil
	}
	out := new(ResourceBackupObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackupProgressStatus) DeepCopyInto(out *ResourceBackupProgressStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*ResourceRestoreResourceInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceRestoreResourceInfo)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackupProgressStatus.
func (in *ResourceBackupProgressStatus) DeepCopy() *ResourceBackupProgressStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceBackupProgressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackupSpec) DeepCopyInto(out *ResourceBackupSpec) {
	*out = *in
	out.ObjRef = in.ObjRef
	out.PVCObjRef = in.PVCObjRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackupSpec.
func (in *ResourceBackupSpec) DeepCopy() *ResourceBackupSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceBackupVolumeInfo) DeepCopyInto(out *ResourceBackupVolumeInfo) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceBackupVolumeInfo.
func (in *ResourceBackupVolumeInfo) DeepCopy() *ResourceBackupVolumeInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceBackupVolumeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExport) DeepCopyInto(out *ResourceExport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	if in.VolumesInfo != nil {
		in, out := &in.VolumesInfo, &out.VolumesInfo
		*out = make([]*ResourceBackupVolumeInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceBackupVolumeInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ExistingVolumesInfo != nil {
		in, out := &in.ExistingVolumesInfo, &out.ExistingVolumesInfo
		*out = make([]*ResourceRestoreVolumeInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceRestoreVolumeInfo)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExport.
func (in *ResourceExport) DeepCopy() *ResourceExport {
	if in == nil {
		return nil
	}
	out := new(ResourceExport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceExport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportList) DeepCopyInto(out *ResourceExportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceExport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportList.
func (in *ResourceExportList) DeepCopy() *ResourceExportList {
	if in == nil {
		return nil
	}
	out := new(ResourceExportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceExportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportObjectReference) DeepCopyInto(out *ResourceExportObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportObjectReference.
func (in *ResourceExportObjectReference) DeepCopy() *ResourceExportObjectReference {
	if in == nil {
		return nil
	}
	out := new(ResourceExportObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceExportSpec) DeepCopyInto(out *ResourceExportSpec) {
	*out = *in
	out.Source = in.Source
	out.Destination = in.Destination
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceExportSpec.
func (in *ResourceExportSpec) DeepCopy() *ResourceExportSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceExportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRestoreResourceInfo) DeepCopyInto(out *ResourceRestoreResourceInfo) {
	*out = *in
	out.ObjectInfo = in.ObjectInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRestoreResourceInfo.
func (in *ResourceRestoreResourceInfo) DeepCopy() *ResourceRestoreResourceInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceRestoreResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRestoreVolumeInfo) DeepCopyInto(out *ResourceRestoreVolumeInfo) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRestoreVolumeInfo.
func (in *ResourceRestoreVolumeInfo) DeepCopy() *ResourceRestoreVolumeInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceRestoreVolumeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceStatus) DeepCopyInto(out *ResourceStatus) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]*ResourceRestoreResourceInfo, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ResourceRestoreResourceInfo)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceStatus.
func (in *ResourceStatus) DeepCopy() *ResourceStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackup) DeepCopyInto(out *VolumeBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackup.
func (in *VolumeBackup) DeepCopy() *VolumeBackup {
	if in == nil {
		return nil
	}
	out := new(VolumeBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupDelete) DeepCopyInto(out *VolumeBackupDelete) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupDelete.
func (in *VolumeBackupDelete) DeepCopy() *VolumeBackupDelete {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupDelete)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackupDelete) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupDeleteList) DeepCopyInto(out *VolumeBackupDeleteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeBackupDelete, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupDeleteList.
func (in *VolumeBackupDeleteList) DeepCopy() *VolumeBackupDeleteList {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupDeleteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackupDeleteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupDeleteSpec) DeepCopyInto(out *VolumeBackupDeleteSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupDeleteSpec.
func (in *VolumeBackupDeleteSpec) DeepCopy() *VolumeBackupDeleteSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupDeleteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupDeleteStatus) DeepCopyInto(out *VolumeBackupDeleteStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupDeleteStatus.
func (in *VolumeBackupDeleteStatus) DeepCopy() *VolumeBackupDeleteStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupDeleteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupList) DeepCopyInto(out *VolumeBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupList.
func (in *VolumeBackupList) DeepCopy() *VolumeBackupList {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupSpec) DeepCopyInto(out *VolumeBackupSpec) {
	*out = *in
	out.BackupLocation = in.BackupLocation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupSpec.
func (in *VolumeBackupSpec) DeepCopy() *VolumeBackupSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeBackupStatus) DeepCopyInto(out *VolumeBackupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeBackupStatus.
func (in *VolumeBackupStatus) DeepCopy() *VolumeBackupStatus {
	if in == nil {
		return nil
	}
	out := new(VolumeBackupStatus)
	in.DeepCopyInto(out)
	return out
}
