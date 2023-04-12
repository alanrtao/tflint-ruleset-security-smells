config {
	varfile = ["terraform.tfvars"]
}

plugin "terraform" {
	enabled=false
}

plugin "secsmells" {
	enabled=true
}