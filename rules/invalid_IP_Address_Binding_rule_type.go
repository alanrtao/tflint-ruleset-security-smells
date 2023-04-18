package rules

import (
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformBackendTypeRule checks whether ...
type InvalidIPAddressBindingRuleType struct {
	tflint.DefaultRule
}

/* Configs for the rule itself */

func NewInvalidIPAddressBindingRuleType() *InvalidIPAddressBindingRuleType {
	return &InvalidIPAddressBindingRuleType{}
}

func (r *InvalidIPAddressBindingRuleType) Name() string {
	return "invalid_IP_Address_rule_type"
}

func (r *InvalidIPAddressBindingRuleType) Enabled() bool {
	return true
}

func (r *InvalidIPAddressBindingRuleType) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *InvalidIPAddressBindingRuleType) Link() string {
	return ""
}

/* Rule body */
func (r *InvalidIPAddressBindingRuleType) Check(runner tflint.Runner) error {
	
	runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		// get expression string from file
		exprRange := expr.Range()
		exprFile := expr.Range().Filename
		
		file, err := runner.GetFile(exprFile)
		if err != nil {
			return nil
		}
		exprPlainText := string(file.Bytes[exprRange.Start.Byte:exprRange.End.Byte])
		exprPlainTextLower := strings.ToLower(exprPlainText)

		// skip if expression not about IP address
		if !(strings.Contains(exprPlainTextLower, "0.0.0.0/0")) {
			return nil
		}

		runner.EmitIssue(
			r, "IP address should not be 0.0.0.0", expr.StartRange(),
		)

		return nil
		
	}))
	return nil
	
}
