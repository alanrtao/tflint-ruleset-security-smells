## Requirements

-   TFLint v0.40+
-   Go v1.20

## Installation

-   Install [`tflint`](https://github.com/terraform-linters/tflint#installation) and run `tflint --init` within your Terraform project repository

## Usage

-   Move the `.tflint.hcl` file from this repository to your Terraform project repository
    -   Some sections in the config file are purely for demo purposes, you can remove them
    -   Specific rules within the plugin (see [`/rules`](/rules) or alternatively [main.go](main.go)) can be enabled or disabled manually
-   `tflint`
    > For more information refer to [Terraform documentation](https://github.com/terraform-linters/tflint#getting-started)
  - `--no-color` makes the output compatible to plain textlogs such as in Jenkins

## Rules

- No Hardcoded Secret: `variable` files with password-like names should not have default values set
- Password Strength:
  - Any expression containing password-like elements should evaluate to a string that:
    - Is at least 8 characters long
    - Contains capital & lower alphabet, digits, and special characters
- Invalid IP Address Binding: checks whether IP Address is '0.0.0.0'
- HTTP without TLS: checks whether URLs use HTTP without TLS (i.e., not HTTPS)
