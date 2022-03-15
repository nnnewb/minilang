# minilang

simple s-exp evaluator just for fun.

## features

- support `+`,`-`,`*`,`/`,`printf` operator

## example

```scm
(printf
    "1+2*3=%d"
    (+ 1
        (* 2 3)))
```

save previous content as `a.scm`, then build compiler and compile code we wrote above:

```bash
mage build
build/bin/scc a.scm
clang a.scm.ll
./a.out
```

done!

## dependencies

- [mage](https://github.com/magefile/mage)
- [gocc](https://github.com/goccmack/gocc)

## LICENSE

MIT

