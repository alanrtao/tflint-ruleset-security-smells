package rules

import (
	"fmt"
	"strings"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformBackendTypeRule checks whether ...
type NoHardcodedSecretRuleType struct {
	tflint.DefaultRule
}

/* Configs for the rule itself */

func NewNoHardcodedSecretRuleType() *NoHardcodedSecretRuleType {
	return &NoHardcodedSecretRuleType{}
}

func (r *NoHardcodedSecretRuleType) Name() string {
	return "no_hardcoded_secret_rule_type"
}

func (r *NoHardcodedSecretRuleType) Enabled() bool {
	return true
}

func (r *NoHardcodedSecretRuleType) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *NoHardcodedSecretRuleType) Link() string {
	return ""
}

/* Rule body */
func (r *NoHardcodedSecretRuleType) Check(runner tflint.Runner) error {
	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "type"}, {Name: "default"}},
				},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return err
	}

	for _, variable := range body.Blocks {
		name := variable.Labels[0]
		nameLower := strings.ToLower(name)
		if !(strings.Contains(nameLower, "password") || strings.Contains(nameLower, "pwd") || strings.Contains(nameLower, "pass")) {
			continue
		}
		for _, attr := range variable.Body.Attributes {
			if strings.Compare(attr.Name, "default") == 0 {
				if err := runner.EmitIssue(
					r,
					fmt.Sprintf("Password `%v` should not have default value", variable.Labels[0]),
					attr.NameRange,
				); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
