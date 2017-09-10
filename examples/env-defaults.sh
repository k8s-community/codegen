#!/bin/bash

# Environment params for example template

# Application name of service (it's validated by "^([a-z]{1}[a-z-0-9]+)$" regexp)
export CODEGEN_APP_NAME="myexample"

# Source directory of template
export CODEGEN_SRC_PATH="./examples/template"

# Destination directory for generated code
export CODEGEN_DEST_PATH="./examples/result"

# Directories (delimiter = ",") that must be skipped by generator
export CODEGEN_SKIP_PATHS=""

# Paths that will be replaced on new paths (it can consist of template data)
export CODEGEN_REPLACE_PATHS="foo/bar.txt:foo/{[(.appName)]}.txt"

# Path of JSON file with template data
export CODEGEN_TEMPLATE_DATA_PATH="./examples/data-defaults.json"
