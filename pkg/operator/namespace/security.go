/*
Copyright © 2018 inwinSTACK Inc

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

package namespace

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *Controller) updateSecurity(namespace string, sourceAddresses []string) error {
	secs, err := c.blendedset.InwinstackV1().Securities(namespace).List(metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, sec := range secs.Items {
		sec.Spec.SourceAddresses = sourceAddresses
		if _, err := c.blendedset.InwinstackV1().Securities(namespace).Update(&sec); err != nil {
			return err
		}
	}
	return nil
}
