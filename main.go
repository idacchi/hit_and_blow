//******************************************************************************
//	ヒットアンドブロー
//
//	-comment-
//	GO言語の練習用に作ったやつ
//
//	-note-
//	2016/12/13	新規作成
//******************************************************************************
package main

//******************************************************************************
//	インポート
//******************************************************************************
import (
	"errors"
	"fmt"
	"strconv"
	"unicode/utf8"
)

//******************************************************************************
//	定数
//******************************************************************************
const (
	InputLen = 4 //	入力文字列の長さ
)

//******************************************************************************
//	メイン
//******************************************************************************
func main() {

	//
	//	タイトル表示
	//
	viewTitle()

	var inputString string
	var i int

	//
	//	メインループ
	//
	for i = 0; ; {

		var inputData []int

		fmt.Printf("【%d回目】%d桁の数字を入れてね > ", (i + 1), InputLen)

		fmt.Scanln(&inputString)
		//src = "1234"
		fmt.Println("入力>" + inputString)

		//	入力文字列検証
		inputData, err := validateSourceText(inputString)

		if nil != err {
			fmt.Println(err.Error())
			fmt.Println("")
			continue
		} else {
			fmt.Println("いけたね")
			fmt.Println(inputData)
			i++
		}

		break
	}
}

//******************************************************************************
//	タイトル表示
//******************************************************************************
func viewTitle() {

	fmt.Println("")
	fmt.Println("////////////////////////////////")
	fmt.Println("  Hit and Blow!")
	fmt.Println("////////////////////////////////")
	fmt.Println("")
}

//******************************************************************************
//	入力文字検証
//******************************************************************************
func validateSourceText(iSrc string) (inputData []int, err error) {

	//
	//	文字数チェック
	//
	srcLen := utf8.RuneCountInString(iSrc)

	if 0 >= srcLen {
		err = errors.New("入力が正しくない")
		return
	} else if InputLen < srcLen {
		err = errors.New("入力文字数が多い")
		return
	}

	//
	//	すべての文字が数字かチェック＆戻りの配列に値をセット
	//
	inputData = make([]int, InputLen)

	for i := 0; i < srcLen; i++ {

		var char = iSrc[i:(i + 1)]
		var work int
		var errWork error

		//	数字へ変換
		work, errWork = strconv.Atoi(char)

		if nil == errWork {
			inputData[i] = work
		} else {
			err = errors.New("数字以外が含まれている (" + char + ")")
			return
		}
	}

	//err = errors.New("ほげー")

	return
}
