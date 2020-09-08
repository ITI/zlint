package cabf_br

/*
 * ZLint Copyright 2020 Regents of the University of Michigan
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
	"crypto/dsa"

	"github.com/zmap/zcrypto/x509"
	"github.com/teamnsrg/zlint/v2/lint"
	"github.com/teamnsrg/zlint/v2/util"
)

type dsaTooShort struct{}

func (l *dsaTooShort) Initialize() error {
	return nil
}

func (l *dsaTooShort) CheckApplies(c *x509.Certificate) bool {
	return c.PublicKeyAlgorithm == x509.DSA
}

func (l *dsaTooShort) Execute(c *x509.Certificate) *lint.LintResult {
	dsaKey, ok := c.PublicKey.(*dsa.PublicKey)
	if !ok {
		return &lint.LintResult{Status: lint.NA}
	}
	dsaParams := dsaKey.Parameters
	L := dsaParams.P.BitLen()
	N := dsaParams.Q.BitLen()
	if L >= 2048 && N >= 244 {
		return &lint.LintResult{Status: lint.Pass}
	}
	return &lint.LintResult{Status: lint.Error}
}

func init() {
	lint.RegisterLint(&lint.Lint{
		Name:        "e_dsa_shorter_than_2048_bits",
		Description: "DSA modulus size must be at least 2048 bits",
		Citation:    "BRs: 6.1.5",
		// Refer to BRs: 6.1.5, taking the statement "Before 31 Dec 2010" literally
		Source:        lint.CABFBaselineRequirements,
		EffectiveDate: util.ZeroDate,
		Lint:          &dsaTooShort{},
	})
}
