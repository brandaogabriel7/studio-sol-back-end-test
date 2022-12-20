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
	
	Context("Checking that password follows the specified rules", func ()  {
		DescribeTable("minSize, minSpecialChars, noRepeted, minDigit",
			func (password string, minSize, minSpecialChars, minDigit int)  {
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

		DescribeTable("minUppercase, minLowercase, minDigit",
			func (password string, minUppercase, minLowercase, minDigit int)  {
				var resp struct {
					Verify model.Verify
				}
				
				error := c.Post(
					`
					query($password: String!, $minUppercase: Int!, $minLowercase: Int!, $minDigit: Int!) {
						verify(password: $password, rules: [
							{rule: "minUppercase", value: $minUppercase},
							{rule: "minLowercase", value: $minLowercase},
							{rule: "minDigit", value: $minDigit},
						]) {
							verify
							noMatch
						}
					}`,
					&resp,
					client.Var("password", password),
					client.Var("minUppercase", minUppercase),
					client.Var("minLowercase", minLowercase),
					client.Var("minDigit", minDigit),
				)
	
				Expect(error).NotTo(HaveOccurred())
				Expect(resp.Verify.Verify).To(BeTrue())
				Expect(resp.Verify.NoMatch).To(BeEmpty())
			},
			Entry("Test case 1", "OOOOoPpppa1@", 3, 5, 1),
			Entry("Test case 2", "S3nHaForTe!23", 4, 4, 3),
		)
	})

	Context("Checking that password does not follow all the specified rules", func ()  {
		DescribeTable("The following rules fail:",
			func (password string, minSize, minSpecialChars, minDigit, minUppercase, minLowercase int, noMatch []string)  {
				var resp struct {
					Verify model.Verify
				}
				
				error := c.Post(
					`
					query($password: String!, $minSize: Int!, $minSpecialChars: Int!, $minDigit: Int!,$minUppercase: Int!, $minLowercase: Int!) {
						verify(password: $password, rules: [
							{rule: "minSize", value: $minSize},
							{rule: "minSpecialChars", value: $minSpecialChars},
							{rule: "noRepeted", value: 0},
							{rule: "minDigit", value: $minDigit},
							{rule: "minUppercase", value: $minUppercase},
							{rule: "minLowercase", value: $minLowercase},
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
					client.Var("minUppercase", minUppercase),
					client.Var("minLowercase", minLowercase),
				)
	
				Expect(error).NotTo(HaveOccurred())
				Expect(resp.Verify.Verify).To(BeFalse())
				Expect(resp.Verify.NoMatch).NotTo(BeEmpty())
				Expect(resp.Verify.NoMatch).To(Equal(noMatch))
			},
			Entry("minSize", "O0p@", 5, 1, 1, 1, 1, []string{string(validation.MIN_SIZE)}),
			Entry("minSize, minSpecialChars, minDigit", "SenhaForte!2", 20, 4, 2, 1, 1, []string{string(validation.MIN_SIZE), string(validation.MIN_SPECIAL_CHARS), string(validation.MIN_DIGIT)}),
			Entry("noRepeted, minUppercase", "aaaaaaa!2", 3, 1, 1, 3, 1, []string{string(validation.NO_REPETED), string(validation.MIN_UPPERCASE)}),
			Entry("minUppercase, minLowercase", "senha!2", 3, 1, 1, 3, 9, []string{string(validation.MIN_UPPERCASE), string(validation.MIN_LOWERCASE)}),
		)
	})
})
