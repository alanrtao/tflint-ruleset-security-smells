package rules

import (
	"net/url"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
)

// HttpWithoutTLSRule checks whether URLs use HTTP without TLS (i.e., not HTTPS)
type HttpWithoutTLSRule struct {
	tflint.DefaultRule
}

// NewHttpWithoutTLSRule initializes the rule
func NewHttpWithoutTLSRule() *HttpWithoutTLSRule {
	return &HttpWithoutTLSRule{}
}

// Name returns the rule name
func (r *HttpWithoutTLSRule) Name() string {
	return "http_without_tls_rule"
}

// Enabled returns the rule enabled status
func (r *HttpWithoutTLSRule) Enabled() bool {
	return true
}

// Severity returns the rule severity (WARNING in this case)
func (r *HttpWithoutTLSRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *HttpWithoutTLSRule) Link() string {
	return ""
}

// Check is the main function of the rule that checks for non-secure HTTP URLs
func (r *HttpWithoutTLSRule) Check(runner tflint.Runner) error {
	// Walk through all expressions in the Terraform configuration
	runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		// Evaluate the expression and store the result in 'value'
		var value string
		err := runner.EvaluateExpr(expr, &value, &tflint.EvaluateExprOption{
			WantType: &cty.String,
		})
		if err != nil {
			return nil
		}

		// Check if the value contains "http://" substring
		if strings.Contains(value, "http://") {
			// Parse the URL to obtain its scheme (http or https)
			parsedURL, err := url.Parse(value)
			if err != nil {
				return nil
			}

			// If the scheme is "http", emit a warning about using a non-secure connection
			if parsedURL.Scheme == "http" {
				runner.EmitIssue(
					r, "URL '"+value+"' uses HTTP without TLS (HTTPS). Consider using a secure connection.", expr.StartRange(),
				)
			}
		}

		return nil
	}))

	return nil
}
