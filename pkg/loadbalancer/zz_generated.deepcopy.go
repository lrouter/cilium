// +build !ignore_autogenerated

// Copyright 2017-2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package loadbalancer

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Backend) DeepCopyInto(out *Backend) {
	*out = *in
	in.L3n4Addr.DeepCopyInto(&out.L3n4Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Backend.
func (in *Backend) DeepCopy() *Backend {
	if in == nil {
		return nil
	}
	out := new(Backend)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendMaglev) DeepCopyInto(out *BackendMaglev) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendMaglev.
func (in *BackendMaglev) DeepCopy() *BackendMaglev {
	if in == nil {
		return nil
	}
	out := new(BackendMaglev)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackendMeta) DeepCopyInto(out *BackendMeta) {
	*out = *in
	if in.BackendMaglev != nil {
		in, out := &in.BackendMaglev, &out.BackendMaglev
		*out = new(BackendMaglev)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackendMeta.
func (in *BackendMeta) DeepCopy() *BackendMeta {
	if in == nil {
		return nil
	}
	out := new(BackendMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3n4Addr) DeepCopyInto(out *L3n4Addr) {
	clone := in.DeepCopy()
	*out = *clone
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L3n4AddrID) DeepCopyInto(out *L3n4AddrID) {
	*out = *in
	in.L3n4Addr.DeepCopyInto(&out.L3n4Addr)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new L3n4AddrID.
func (in *L3n4AddrID) DeepCopy() *L3n4AddrID {
	if in == nil {
		return nil
	}
	out := new(L3n4AddrID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *L4Addr) DeepCopyInto(out *L4Addr) {
	clone := in.DeepCopy()
	*out = *clone
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaglevElem) DeepCopyInto(out *MaglevElem) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaglevElem.
func (in *MaglevElem) DeepCopy() *MaglevElem {
	if in == nil {
		return nil
	}
	out := new(MaglevElem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in MaglevRing) DeepCopyInto(out *MaglevRing) {
	{
		in := &in
		*out = make(MaglevRing, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaglevRing.
func (in MaglevRing) DeepCopy() MaglevRing {
	if in == nil {
		return nil
	}
	out := new(MaglevRing)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SVC) DeepCopyInto(out *SVC) {
	*out = *in
	in.Frontend.DeepCopyInto(&out.Frontend)
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]Backend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SVC.
func (in *SVC) DeepCopy() *SVC {
	if in == nil {
		return nil
	}
	out := new(SVC)
	in.DeepCopyInto(out)
	return out
}
