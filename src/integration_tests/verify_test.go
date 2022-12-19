package integration_tests_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/brandaogabriel7/studio-sol-back-end-test/graph"
	"github.com/brandaogabriel7/studio-sol-back-end-test/graph/model"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/factories"
	"github.com/brandaogabriel7/studio-sol-back-end-test/src/strategies/validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Verify", func() {
	c := client.New(handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			PasswordValidationService: factories.GetDefaultPasswordValidationService(),
		}})))
	
	Describe("Checking that password follows the specified rules", func ()  {
		DescribeTable("minSize, minSpecialChars, noRepeted, minDigit",
			func (password string, minSize int, minSpecialChars int, minDigit int)  {
				var resp struct {
					Verify model.Verify
				}
				
				error := c.Post(
					`
					query($password: String!, $minSize: Int!, $minSpecialChars: Int!, $minDigit: Int!) {
						verify(password: $password, rules: [
							{rule: "minSize", value: $minSize},
							{rule: "minSpecialChars", value: $minSpecialChars},
							{rule: "noRepeted", value: 0},
							{rule: "minDigit", value: $minDigit},
						]) {
							verify
							noMatch
						}
					}`,
					&resp,
					client.Var("password", password),
					client.Var("minSize", minSize),
					client.Var("minSpecialChars", minSpecialChars),
					client.Var("minDigit", minDigit),
				)
	
				Expect(error).NotTo(HaveOccurred())
				Expect(resp.Verify.Verify).To(BeTrue())
				Expect(resp.Verify.NoMatch).To(BeEmpty())
			},
			Entry("Test case 1", "Opa1@", 5, 1, 1),
			Entry("Test case 2", "SenhaForte!23", 5, 1, 1),
		)
	})

	Describe("Checking that password does not follow all the specified rules", func ()  {
		DescribeTable("The following rules fail:",
			func (password string, minSize int, minSpecialChars int, minDigit int, noMatch []string)  {
				var resp struct {
					Verify model.Verify
				}
				
				error := c.Post(
					`
					query($password: String!, $minSize: Int!, $minSpecialChars: Int!, $minDigit: Int!) {
						verify(password: $password, rules: [
							{rule: "minSize", value: $minSize},
							{rule: "minSpecialChars", value: $minSpecialChars},
							{rule: "noRepeted", value: 0},
							{rule: "minDigit", value: $minDigit},
						]) {
							verify
							noMatch
						}
					}`,
					&resp,
					client.Var("password", password),
					client.Var("minSize", minSize),
					client.Var("minSpecialChars", minSpecialChars),
					client.Var("minDigit", minDigit),
				)
	
				Expect(error).NotTo(HaveOccurred())
				Expect(resp.Verify.Verify).To(BeFalse())
				Expect(resp.Verify.NoMatch).NotTo(BeEmpty())
				Expect(resp.Verify.NoMatch).To(Equal(noMatch))
			},
			Entry("minSize", "0p@", 5, 1, 1, []string{string(validation.MIN_SIZE)}),
			Entry("minSize, minSpecialChars, minDigit", "SenhaForte!2", 20, 4, 2, []string{string(validation.MIN_SIZE), string(validation.MIN_SPECIAL_CHARS), string(validation.MIN_DIGIT)}),
			Entry("noRepeted", "aaaaaaa!2", 3, 1, 1, []string{string(validation.NO_REPETED)}),
		)
	})
})
