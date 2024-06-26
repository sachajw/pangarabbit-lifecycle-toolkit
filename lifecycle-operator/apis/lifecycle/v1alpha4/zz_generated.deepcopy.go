//go:build !ignore_autogenerated

/*
Copyright 2022.

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

package v1alpha4

import (
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3"
	"github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3/common"
	propagation "go.opentelemetry.io/otel/propagation"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersion) DeepCopyInto(out *KeptnWorkloadVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersion.
func (in *KeptnWorkloadVersion) DeepCopy() *KeptnWorkloadVersion {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkloadVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionList) DeepCopyInto(out *KeptnWorkloadVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KeptnWorkloadVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionList.
func (in *KeptnWorkloadVersionList) DeepCopy() *KeptnWorkloadVersionList {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeptnWorkloadVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionSpec) DeepCopyInto(out *KeptnWorkloadVersionSpec) {
	*out = *in
	in.KeptnWorkloadSpec.DeepCopyInto(&out.KeptnWorkloadSpec)
	if in.TraceId != nil {
		in, out := &in.TraceId, &out.TraceId
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionSpec.
func (in *KeptnWorkloadVersionSpec) DeepCopy() *KeptnWorkloadVersionSpec {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeptnWorkloadVersionStatus) DeepCopyInto(out *KeptnWorkloadVersionStatus) {
	*out = *in
	if in.PreDeploymentTaskStatus != nil {
		in, out := &in.PreDeploymentTaskStatus, &out.PreDeploymentTaskStatus
		*out = make([]v1alpha3.ItemStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentTaskStatus != nil {
		in, out := &in.PostDeploymentTaskStatus, &out.PostDeploymentTaskStatus
		*out = make([]v1alpha3.ItemStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PreDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PreDeploymentEvaluationTaskStatus, &out.PreDeploymentEvaluationTaskStatus
		*out = make([]v1alpha3.ItemStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PostDeploymentEvaluationTaskStatus != nil {
		in, out := &in.PostDeploymentEvaluationTaskStatus, &out.PostDeploymentEvaluationTaskStatus
		*out = make([]v1alpha3.ItemStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.PhaseTraceIDs != nil {
		in, out := &in.PhaseTraceIDs, &out.PhaseTraceIDs
		*out = make(common.PhaseTraceID, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = make(propagation.MapCarrier, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeptnWorkloadVersionStatus.
func (in *KeptnWorkloadVersionStatus) DeepCopy() *KeptnWorkloadVersionStatus {
	if in == nil {
		return nil
	}
	out := new(KeptnWorkloadVersionStatus)
	in.DeepCopyInto(out)
	return out
}
