package rules

import (
	"regexp"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
)

// TerraformBackendTypeRule checks whether ...
type PasswordStrengthRuleType struct {
	tflint.DefaultRule
}

/* Configs for the rule itself */

func NewPasswordStrengthRuleType() *PasswordStrengthRuleType {
	return &PasswordStrengthRuleType{}
}

func (r *PasswordStrengthRuleType) Name() string {
	return "password_strength_rule_type"
}

func (r *PasswordStrengthRuleType) Enabled() bool {
	return true
}

func (r *PasswordStrengthRuleType) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *PasswordStrengthRuleType) Link() string {
	return ""
}

/* Rule body */
func (r *PasswordStrengthRuleType) Check(runner tflint.Runner) error {

	var cache = map[string]string{}

	runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {

		// get expression string from file
		eRange := expr.Range()
		eFile := expr.Range().Filename

		file, err := runner.GetFile(eFile)
		if err != nil {
			// todo: return diag instead
			return nil
		}
		plainText := string(file.Bytes[eRange.Start.Byte:eRange.End.Byte])
		plainTextLower := strings.ToLower(plainText)

		// skip if expression not about password
		if !(strings.Contains(plainTextLower, "password") ||
			strings.Contains(plainTextLower, "pass") ||
			strings.Contains(plainTextLower, "pwd")) {
			return nil
		}

		prevVal, exists := cache[plainText]
		if exists {
			return nil
		}

		var value string
		err = runner.EvaluateExpr(expr, &value, &tflint.EvaluateExprOption{
			WantType: &cty.String,
		})
		if err != nil {
			// todo: return diag instead
			return nil
		}
		// write to cache
		cache[plainText] = prevVal

		// skip if expression hard coded, not handled by this rule
		variables := expr.Variables()
		if len(variables) == 0 {
			return nil
		}

		// do strength check
		if len(value) >= 8 &&
			regexp.MustCompile(`[a-z]`).MatchString(value) &&
			regexp.MustCompile(`[0-9]`).MatchString(value) &&
			regexp.MustCompile(`[A-Z]`).MatchString(value) &&
			regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(value) {
			return nil
		}

		runner.EmitIssue(
			r, value+"@"+plainText+" not strong enough", expr.StartRange(),
		)

		return nil
	}))

	return nil
}
