config {
	varfile = ["terraform.tfvars"]
}

plugin "terraform" {
	enabled=false
}

plugin "smells" {
	enabled=true
	version="0.1.0"
	source="github.com/alanrtao/tflint-ruleset-secsmells"
}