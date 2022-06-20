package main

import (
	"database/sql"
	"errors"
	"log"
)

func literal() {
	// リテラルサンプル
	var a = 1       // イコールの右側の１が定数
	var b = 1 + 100 // 定数の中で演算子を使った計算も可能
	// リテラルサンプル
	log.Println(a, b)
}

func binfType() {
	// bind-type
	var a = 1             // aの型は定数デフォルトの型（ここではint）になる
	var b int32 = 1 + 100 // 明示的に描くことも可能
	function(1000)        // 関数の引数に渡すリテラル定数だが、関数の引数の型に合わせて型が決定
	// bind-type
	log.Println(a, b)
}

func function(i int) {

}

// non-type-content
const (
	a = 1                      // イコールの右側の１が定数
	b = 1 + 2                  // 演算もできる
	c = 223372036854775807 + 1 // uint64を超える数値も定義可能
	d = "Hello world"          // 整数以外に、浮動小数点数や文字列やブール型も扱える
	e = iota + 10              // const宣言でのみiotaも使える
)

// typed-const
type ErrorCode int

const (
	f    int       = 10
	code ErrorCode = 10
)

// constanat-instance
type errDatabase int

func (e errDatabase) Error() string {
	return "Database Error"
}

const (
	ErrDatabase errDatabase = 0
)

func OpenDB(param string) error {
	return ErrDatabase
}

func main() {
	// use-constant-instance
	err := OpenDB("postgres://localhost:5432")
	if err == ErrDatabase {
		log.Fatal("DB接続エラー")
	}
	// use-constant-instance
	// orverride-error-instance
	sql.ErrConnDone = errors.New("エラー乗っ取りもできてしまう")
	// orverride-error-instance
}
