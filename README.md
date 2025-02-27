# Commander 

Context dependent command wrapper.

## Usage

``` bash
commander command <name> [args ...]
```

## Config

``` yaml
commander:
  commands:
  - name: check-path
    match_path_regex: .*path/to/match.*
    command: echo matched path
  - name: check-path
    match_command: [ "$MATCH" != "" ]
    command: echo match env
```

See `examples/config.yaml`
