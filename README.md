# Terraform Module Initialiser

## Small CLI tool to initialise a Terraform module

### Usage (macOS):

- Clone repo and compile binary: ``make build``
- Copy combiled binary to path: ``cp ./tfminit /usr/local/bin``
- Initialise Terraform module: ``tfminit <module path>``

### Example: 

```shell
tfminit terraform/modules/new_module
ls -l terraform/modules/new_module/
total 0
-rw-r--r--  1 user  staff  0 Sep  5 09:19 README.md
-rw-r--r--  1 user  staff  0 Sep  5 09:19 main.tf
-rw-r--r--  1 user  staff  0 Sep  5 09:19 outputs.tf
-rw-r--r--  1 user  staff  0 Sep  5 09:19 variables.tf
```

If destination directory does not exist tfminit will create it