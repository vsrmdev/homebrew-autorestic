# Exec

```bash
autorestic exec [-b, --backend]  [-a, --all] <command> -- [native options]
```

This is avery handy command which enables you to run any native restic command on desired backends. An example would be listing all the snapshots of all your backends:

```bash
autorestic exec -a -- snapshots
```

With `exec` you can basically run every cli command that you would be able to run with the restic cli. It only pre-fills path, key, etc.

> :ToCPrevNext
