config {
	varfile = ["terraform.tfvars"]
}

plugin "terraform" {
	enabled=false
}

plugin "secsmells" {
	enabled = true
	version = "0.1.0"
	source = "github.com/alanrtao/tflint-ruleset-secsmells"
	signing_key = <<-KEY
	-----BEGIN PGP PUBLIC KEY BLOCK-----
	
	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
	Lbr9m4+OH4vpgVrGnFGk8K2kKTHKl8I=
	=5INJ
	-----END PGP PUBLIC KEY BLOCK-----
	KEY
}