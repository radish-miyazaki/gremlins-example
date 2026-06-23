# gremlins-example — mutator / status showcase

> 日本語版は [README.ja.md](README.ja.md) を参照してください。

A sample that reproduces and verifies every one of the 11 mutators that
gremlins v0.6.0 injects, as well as all 7 result statuses it reports, using
independent Go packages and `make` targets.

## Prerequisites

- Go 1.26.2
- gremlins is run as `go tool gremlins` via the `tool` directive in `go.mod`.

## Basics

```bash
make test   # Verify that all tests on the original code pass (a prerequisite for running gremlins)
make all    # Apply all 11 mutators to the whole repository (uses .gremlins.yaml)
```

## Per mutator (each target enables only the target mutator and isolates a single mutation)

| make target | mutator | example mutation | expected result |
|---|---|---|---|
| `make arithmetic-base` | ARITHMETIC_BASE | `+`→`-` | KILLED |
| `make conditionals-boundary` | CONDITIONALS_BOUNDARY | `<`→`<=` | KILLED |
| `make conditionals-negation` | CONDITIONALS_NEGATION | `==`→`!=` | KILLED |
| `make increment-decrement` | INCREMENT_DECREMENT | `++`→`--` | KILLED |
| `make invert-negatives` | INVERT_NEGATIVES | `-x`→`+x` | KILLED |
| `make invert-assignments` | INVERT_ASSIGNMENTS | `+=`→`-=` | KILLED |
| `make invert-bitwise` | INVERT_BITWISE | `&`→`\|` | KILLED |
| `make invert-bwassign` | INVERT_BWASSIGN | `&=`→`\|=` | KILLED |
| `make invert-logical` | INVERT_LOGICAL | `&&`→`\|\|` | KILLED |
| `make invert-loopctrl` | INVERT_LOOPCTRL | `break`→`continue` | KILLED |
| `make remove-self-assignments` | REMOVE_SELF_ASSIGNMENTS | `a += b`→`a = b` | KILLED |

## Per status (all 7)

| make target | status | how it is reproduced |
|---|---|---|
| `make status-runnable` | RUNNABLE | `--dry-run` (shows covered mutations without executing them) |
| `make status-notcovered` | NOT COVERED | a function not reached by any test |
| `make status-killed` | KILLED | a correct test detects the mutation |
| `make status-lived` | LIVED | a weak test misses the mutation (`Double(0)`) |
| `make status-notviable` | NOT VIABLE (unreachable on Go 1.26 — see below) | string `+`→`-` fails to compile |
| `make status-timedout` | TIMED OUT | `i++`→`i--` produces an infinite loop |
| `make status-skipped` | SKIPPED | `--diff HEAD~1` (a mutation outside the most recent commit's diff; an empty diff does not produce SKIPPED) |

> Status character filter `-S`: `l`=LIVED, `c`=NOT COVERED, `t`=TIMED OUT, `k`=KILLED, `v`=NOT VIABLE, `s`=SKIPPED, `r`=RUNNABLE.
>
> **Known limitation regarding NOT VIABLE**: gremlins v0.6.0 treats exit code 2 from `go test` as NOT VIABLE, but Go 1.26's `go test` returns exit code 1 even on a build failure, so a mutation that fails to compile is also classified as `KILLED`. `statuses/notviable` correctly demonstrates the concept of "a mutation that fails to compile," but the output of `make status-notviable` (`-S v`) is empty in this environment. Observing NOT VIABLE for real requires an older Go toolchain whose `go test` returns exit code 2 on a build failure.

## Reading the output

Each line has the form `<STATUS> <MUTATOR_TYPE> at <file>:<line>:<col>`.
Example: `KILLED ARITHMETIC_BASE at mutators/arithmetic_base/arithmetic_base.go:6:11`
