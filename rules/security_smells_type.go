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
	path, err := runner.GetModulePath()
	if err != nil {
		return err
	}
	if !path.IsRoot() {
		// This rule does not evaluate child modules.
		return nil
	}

	body, err := runner.GetModuleContent(&hclext.BodySchema{
		Blocks: []hclext.BlockSchema{
			{
				Type:       "variable",
				LabelNames: []string{"name"},
				Body: &hclext.BodySchema{
					Attributes: []hclext.AttributeSchema{{Name: "type"}},
				},
			},
		},
	}, &tflint.GetModuleContentOption{ExpandMode: tflint.ExpandModeNone})
	if err != nil {
		return err
	}

	for _, variable := range body.Blocks {
		if err := runner.EmitIssue(
			r,
			fmt.Sprintf("`%v` variable detected by plugin", variable.Labels[0]),
			variable.DefRange,
		); err != nil {
			return err
		}
	}

	return nil
}
