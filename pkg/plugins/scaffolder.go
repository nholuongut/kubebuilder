/*
Copyright 2020 The Nho Luong DevOps.

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

package plugins

import (
	"sigs.k8s.io/kubebuilder/v4/pkg/machinery"
)

// Scaffolder interface creates files to set up a controller manager
type Scaffolder interface {
	InjectFS(machinery.Filesystem)
	// Scaffold performs the scaffolding
	Scaffold() error
}
