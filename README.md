# go-log

A la pythonic logging.

Based on the standard `log` module.

Tested using [github.com/stretchr/testify](https://github.com/stretchr/testify).

The `Fatal`s were turned into `panic`s in all the tests. 

coverage: 97.8% of statements _(the actual call to [`logging.Fatal`](log.go#L130)
has not been tested)_ using

```bash
#!/bin/bash
set -e
rm -f cover.out
go test -coverprofile cover.out
sed -e 's,^_.*/\([^/]\+\)$,./\1,' -i cover.out
go tool cover -html=cover.out -o c.html
ls -l c.html
# EOF #
```

# EOF #
