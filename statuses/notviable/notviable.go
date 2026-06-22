// Package notviable は NOT VIABLE ステータスのデモ。
// 文字列連結の `+` を `-` に変異させると型エラーでコンパイル不能になる。
package notviable

// Concat は a + b（文字列連結）を返す。`+`→`-` の変異は string では
// コンパイルできず、gremlins は NOT VIABLE と判定する。
func Concat(a, b string) string {
	return a + b
}
