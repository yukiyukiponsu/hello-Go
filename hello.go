/*
Go言語テスト用プログラム
*/

//実行したいパッケージを定義する。今回はfunc mainを実行したいためpackage mainとしている 
package main

//fmtという名前のパッケージを読み込んでいる（ダブルクオーテーション ("") 必須）
import "fmt"

func main(){
  // ダブルクォーテーション ("") ではなくシングルクオーテーション（''）にするとエラー
 fmt.Print("hello world!\n")
}
