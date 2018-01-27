## onessl create client-cert

Generate client certificate pair

### Synopsis

Generate client certificate pair

```
onessl create client-cert [flags]
```

### Options

```
      --cert-dir string       Path to directory where pki files are stored. (default "/home/tamal/go/src/github.com/appscode/onessl/hack/gendocs")
  -h, --help                  help for client-cert
  -o, --organization string   Name of client organizations.
      --overwrite             Overwrite existing cert/key pair
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --analytics                        Send analytical events to Google Guard (default true)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO

* [onessl create](onessl_create.md)	 - create PKI

