package main

import (
	"github.com/alanrtao/tflint-ruleset-secsmells/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "Security Smells Linter Ruleset",
			Version: "0.2.0",
			Rules: []tflint.Rule{
				rules.NewPasswordStrengthRuleType(),
				rules.NewNoHardcodedSecretRuleType(),
				rules.NewInvalidIPAddressBindingRuleType(),
			},
		},
	})
}
