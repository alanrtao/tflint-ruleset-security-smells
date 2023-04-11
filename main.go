package main

import (
	"github.com/alanrtao/security-smells-ruleset/rules"
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "Security Smells Linter Ruleset",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				// rules.NewAwsInstanceExampleTypeRule(),
				// rules.NewAwsS3BucketExampleLifecycleRule(),
				// rules.NewGoogleComputeSSLPolicyRule(),
				// rules.NewTerraformBackendTypeRule(),
				rules.NewSecuritySmellsTypeRule(),
			},
		},
	})
}
