/*
Copyright 2019 kaitoy.

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

package strings

// ContainsString checks if the given target string is contained in the given slice.
func ContainsString(slice []string, target string) bool {
	for _, str := range slice {
		if str == target {
			return true
		}
	}
	return false
}

// RemoveString removes the given target string from the given slice.
func RemoveString(slice []string, target string) (result []string) {
	for _, str := range slice {
		if str == target {
			continue
		}
		result = append(result, str)
	}
	return
}
