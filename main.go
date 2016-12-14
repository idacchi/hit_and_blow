package main

//	インポート
import "fmt"
import "unicode/utf8"
import "errors"

//	定数
const (
	InputLen = 4 //	入力文字列の長さ
)

//	メイン
func main() {

	viewTitle()

	var src string
	var i int

	//	メインループ
	for i = 0; ; {

		fmt.Printf("【%d回目】%d桁の数字を入れてね > ", (i + 1), InputLen)

		fmt.Scanln(&src)
		fmt.Println("入力>" + src)

		if err := validateSourceText(src); nil != err {

			fmt.Println(err.Error())
			fmt.Println("")

			continue

		} else {

			fmt.Println("いけたね")
			i++

		}

		break
	}
}

//	タイトル表示
func viewTitle() {
	fmt.Println("")
	fmt.Println("////////////////////////////////")
	fmt.Println("  Hit and Blow!")
	fmt.Println("////////////////////////////////")
	fmt.Println("")
}

//	入力文字検証
func validateSourceText(i_Src string) (err error) {

	//	文字数チェック
	l := utf8.RuneCountInString(i_Src)

	if (0 >= l) || (InputLen <= l) {
		err = errors.New("入力文字数が多い")
	}

	//var err error

	//err = errors.New("ほげー")

	return
}
