package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
)

var _ = Describe("NoRepeted", func() {
	noRepetedStrategy := validation.NoRepetedStrategy{}

	DescribeTable("When password has no consecutive repeated character",
		func (password string) {
			isValid := noRepetedStrategy.IsValid(password, 0)

			Expect(isValid).To(BeTrue())
		},
		Entry("abacate123", "abacate123"),
		Entry("SenhaForte!", "SenhaForte!"),
		Entry("aopa", "aopa"),
	)

	DescribeTable("When password has consecutive repeated characters",
		func (password string) {
			isValid := noRepetedStrategy.IsValid(password, 0)

			Expect(isValid).To(BeFalse())
		},
		Entry("P@ssw0rd", "P@ssw0rd"),
		Entry("Opaaaa123", "Opaaaa123"),
		Entry("1111111", "1111111"),
	)
})
