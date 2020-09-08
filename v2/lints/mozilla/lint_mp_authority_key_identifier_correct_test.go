package mozilla

/*
 * ZLint Copyright 2018 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

import (
	"testing"

	"github.com/teamnsrg/zlint/v2/lint"
	"github.com/teamnsrg/zlint/v2/test"
)

func TestAuthorityKeyIdentifier(t *testing.T) {
	testCases := []struct {
		Name           string
		InputFilename  string
		ExpectedResult lint.LintStatus
	}{
		{
			Name:           "Authority key ID includes both the key ID and the issuer's name and serial",
			InputFilename:  "mpAuthorityKeyIdentifierIncorrect.pem",
			ExpectedResult: lint.Error,
		},
		{
			Name:           "Authority key ID includes the key ID",
			InputFilename:  "mpAuthorityKeyIdentifierCorrect.pem",
			ExpectedResult: lint.Pass,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			result := test.TestLint("e_mp_authority_key_identifier_correct", tc.InputFilename)
			if result.Status != tc.ExpectedResult {
				t.Errorf("expected result %v was %v", tc.ExpectedResult, result.Status)
			}
		})
	}
}
