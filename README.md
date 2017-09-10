# Code generator

## Description
A code generator can be useful to stop to do repetative work for service code generation. It can create service code from template.

[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/k8s-community/code-generator/issues)

### How to use a code generator:

#### 1. Create JSON file with data for template 
Example:

    {
        "app": {
            "desc": "This is just example of template"
        },
        
        "owner": {
            "name": "Foo Bar",
            "email": "foo@bar.com"
        }
    }

#### 2. Set environment params
Example:

    # Application name of service (it's validated by "^([a-z]{1}[a-z-0-9]+)$" regexp)
    export CODEGEN_APP_NAME="myexample"

    # Source directory of template
    export CODEGEN_SRC_PATH="./examples/template"

    # Destination directory for generated code
    export CODEGEN_DEST_PATH="./examples/result"

    # Directory(-s) (delimiter is ",") that must be skipped by a generator
    export CODEGEN_SKIP_PATHS=""
    
    # Paths that will be replaced on a new paths (it can consist of template data as you can see below)
    export CODEGEN_REPLACE_PATHS="foo/bar.txt:foo/{[(.appName)]}.txt"

    # Path of JSON file with template data
    export CODEGEN_TEMPLATE_DATA_PATH="./examples/data-defaults.json"

    # Left delimiter of template constructions, vars, ...
    CODEGEN_LEFT_DELIM?="{[("
    
    # Right delimiter of template constructions, vars, ...
    CODEGEN_RIGHT_DELIM?=")]}"

#### 3. Build and run a code generator
    
    make build && make run

#### 4. Finish
A generated code is ready to use it.

### Example (see ./examples)

#### 1. Create JSON file with data for template 
Get and update ``./examples/data-defaults.json`` file. There is a template in ``/.examples/template`` directory.

#### 2. Set environment params
Extract defaults params from file:

    source ./examples/env-defaults.sh

#### 3. Build and run a code generator
    
    make build && make run

#### 4. Finish
Go to ``./examples/result`` directory. There is a result of code generation from template.

## Contributors

All the contributors are welcome. If you would like to be the contributor please accept some rules:
- The pull requests will be accepted only in `develop` branch
- All modifications or additions should be tested

Thank you for your understanding!

## License

[MIT Public License](https://github.com/k8s-community/cluster-deploy/blob/master/LICENSE)

## Author Information

Kubernets Community [k8s-community](https://github.com/k8s-community)