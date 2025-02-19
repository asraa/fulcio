/*
Copyright © 2021 Bob Callaway <bcallawa@redhat.com>

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
package oauthflow

import "github.com/coreos/go-oidc/v3/oidc"

func EmailFromIDToken(token *oidc.IDToken) (string, bool, error) {
	// Extract custom claims
	var claims struct {
		Email    string `json:"email"`
		Verified bool   `json:"email_verified"`
	}
	if err := token.Claims(&claims); err != nil {
		return "", false, err
	}
	return claims.Email, claims.Verified, nil
}
