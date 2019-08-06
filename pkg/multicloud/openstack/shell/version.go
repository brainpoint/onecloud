// Copyright 2019 Yunion
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

package shell

import (
	"fmt"

	"yunion.io/x/onecloud/pkg/multicloud/openstack"
	"yunion.io/x/onecloud/pkg/util/shellutils"
)

func init() {
	type VersionOptions struct {
		SERVICE string `help:"Service name" choices:"compute|volume|volumev2|volumev3"`
	}
	shellutils.R(&VersionOptions{}, "version-show", "Show a service version", func(cli *openstack.SRegion, args *VersionOptions) error {
		minVersion, maxVersion, err := cli.GetVersion(args.SERVICE)
		if err != nil {
			return err
		}
		fmt.Printf("min version: %s max version: %s\n", minVersion, maxVersion)
		return nil
	})
}