/*
Copyright 2018 The Kubernetes Authors.

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

package v1beta1

import (
	"k8s.io/apimachinery/pkg/conversion"
	"github.com/Angus-F/client-go/pkg/apis/clientauthentication"
)

func Convert_clientauthentication_ExecCredentialSpec_To_v1beta1_ExecCredentialSpec(in *clientauthentication.ExecCredentialSpec, out *ExecCredentialSpec, s conversion.Scope) error {
	// This conversion intentionally omits the Response and Interactive fields, which were only
	// supported in v1alpha1.
	return autoConvert_clientauthentication_ExecCredentialSpec_To_v1beta1_ExecCredentialSpec(in, out, s)
}
