# go-log

A la pythonic logging.

Based on the standard `log` module.

Tested using [github.com/stretchr/testify](https://github.com/stretchr/testify).

The `Fatal`s were turned into `panic`s in all the tests. 

coverage: 100% of statements using

```bash
#!/bin/bash
profile=cover.out                                   # use this coverage profile
html=c.html                                         # make this HTML file
set -e                                              # stop on any error
rm -f "${profile}" "${html}"                        # clear files
go test -coverprofile "${profile}"                  # generate profile
sed -e 's,^_.*/\([^/]\+\)$,./\1,' -i "${profile}"   # make it usable
go tool cover -html="${profile}" -o "${html}"       # generate HTML
rm -f "${profile}"                                  # remove patched profile
ls -l "${html}"                                     # notify on completion
# EOF #
```

# EOF #
