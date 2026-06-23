# gremlins-example — mutator / status ショーケース

> For the English version, see [README.md](README.md).

gremlins v0.6.0 が挿入する全11種類の mutator と、出力される全7種類の結果ステータスを、
独立した Go パッケージと `make` ターゲットで再現・確認できるサンプル。

## 前提

- Go 1.26.2
- gremlins は `go.mod` の `tool` ディレクティブ経由で `go tool gremlins` として実行する。

## 基本

```bash
make test   # 元コードのテストが全てパスすることを確認（gremlins 実行の前提）
make all    # リポジトリ全体に全11 mutator を適用（.gremlins.yaml を使用）
```

## mutator 別（各ターゲットは対象 mutator だけを有効化し、単一の変異に絞る）

| make ターゲット | mutator | 変異例 | 期待結果 |
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

## status 別（全7種類）

| make ターゲット | ステータス | 再現方法 |
|---|---|---|
| `make status-runnable` | RUNNABLE | `--dry-run`（カバー済み変異を実行せず表示） |
| `make status-notcovered` | NOT COVERED | テスト未到達の関数 |
| `make status-killed` | KILLED | 正しいテストが変異を検出 |
| `make status-lived` | LIVED | 弱いテストが変異を見逃す（`Double(0)`） |
| `make status-notviable` | NOT VIABLE（※Go 1.26 では到達不能・下記参照） | 文字列 `+`→`-` でコンパイル不能 |
| `make status-timedout` | TIMED OUT | `i++`→`i--` で無限ループ |
| `make status-skipped` | SKIPPED | `--diff HEAD~1`（直近コミットの差分外にある変異。空 diff では SKIPPED にならない点に注意） |

> ステータスの文字フィルタ `-S` 対応: `l`=LIVED, `c`=NOT COVERED, `t`=TIMED OUT, `k`=KILLED, `v`=NOT VIABLE, `s`=SKIPPED, `r`=RUNNABLE。
>
> **NOT VIABLE についての既知の限界**: gremlins v0.6.0 は `go test` の終了コード2を NOT VIABLE と判定するが、Go 1.26 の `go test` はビルド失敗でも終了コード1を返すため、コンパイル不能な変異も `KILLED` に分類される。`statuses/notviable` は「変異がコンパイル不能になる」概念を正しく実演するが、`make status-notviable`（`-S v`）の出力はこの環境では空になる。NOT VIABLE を実際に観測するには `go test` がビルド失敗で exit 2 を返す旧 Go ツールチェインが必要。

## 出力の読み方

1行の形式は `<STATUS> <MUTATOR_TYPE> at <file>:<line>:<col>`。
例: `KILLED ARITHMETIC_BASE at mutators/arithmetic_base/arithmetic_base.go:6:11`
