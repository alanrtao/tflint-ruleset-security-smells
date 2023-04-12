config {
	varfile = ["terraform.tfvars"]
}

plugin "terraform" {
	enabled=false
}

plugin "smells" {
	enabled=true
}