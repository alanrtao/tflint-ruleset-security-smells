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
	// todo: also check for duplicate password by having an inverse map from password values to password names

	runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {

		// get expression string from file
		exprRange := expr.Range()
		exprFile := expr.Range().Filename

		file, err := runner.GetFile(exprFile)
		if err != nil {
			// todo: return diag instead
			return nil
		}
		exprPlainText := string(file.Bytes[exprRange.Start.Byte:exprRange.End.Byte])
		exprPlainTextLower := strings.ToLower(exprPlainText)

		// skip if expression not about password
		if !(strings.Contains(exprPlainTextLower, "password") ||
			strings.Contains(exprPlainTextLower, "pass") ||
			strings.Contains(exprPlainTextLower, "pwd")) {
			return nil
		}

		prevVal, exists := cache[exprPlainText]
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
		cache[exprPlainText] = prevVal

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
			r, exprPlainText+" must be at least 8 characters long with capitalized and non-capitalized letters, digits, and special characters", expr.StartRange(),
		)

		return nil
	}))

	return nil
}
