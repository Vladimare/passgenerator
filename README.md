# Password Generator

```
passgen [N] [opt]

    N                Password length (default is 16)
    --numbers, -n    Enable numbers (enabled by default)
    --lower, -l      Enable lower case letters (enabled by default)
    --upper, -u      Enable upper case letters (enabled by default)
    --sumbols, -s    Enable special characters (enabled by default)
    --help, -h       Print usage

    If any option is specified then others should be enabled.

    Usage example:
        passgen
        passgen 12
        passgen 10 -nlu
```
