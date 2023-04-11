package rules

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformBackendTypeRule checks whether ...
type SecuritySmellsTypeRule struct {
	tflint.DefaultRule
}

/* Configs for the rule itself */

func NewSecuritySmellsTypeRule() *SecuritySmellsTypeRule {
	return &SecuritySmellsTypeRule{}
}

func (r *SecuritySmellsTypeRule) Name() string {
	return "security_smells_type"
}

func (r *SecuritySmellsTypeRule) Enabled() bool {
	return true
}

func (r *SecuritySmellsTypeRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *SecuritySmellsTypeRule) Link() string {
	return ""
}

/* Rule body */

func (r *SecuritySmellsTypeRule) Check(runner tflint.Runner) error {
	// This rule is an example to get attributes of blocks other than resources.
	content, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type: "terraform",
				Body: &hclext.BodySchema{
					Blocks: []hclext.BlockSchema{
						{
							Type:       "backend",
							LabelNames: []string{"type"},
						},
					},
				},
			},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, terraform := range content.Blocks {
		for _, backend := range terraform.Body.Blocks {
			err := runner.EmitIssue(
				r,
				fmt.Sprintf("backend type is %s", backend.Labels[0]),
				backend.DefRange,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
