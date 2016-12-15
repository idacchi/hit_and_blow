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
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/shiena/ansicolor"
	"github.com/wsxiaoys/terminal/color"
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
	//	初期化
	//
	initialize()

	//
	//	タイトル表示
	//
	viewTitle()

	//
	//	正解作成
	//
	var ansData = craeteAnsData()

	//
	//	メインループ
	//
	var inputString string
	var i int

	for i = 0; ; {

		var inputData []int

		fmt.Printf("【%d回目】%d桁の数字を入れてね > ", (i + 1), InputLen)

		fmt.Scanln(&inputString)
		//inputString = "1234"

		//	入力文字列検証
		inputData, err := validateSourceText(inputString)

		if nil != err {
			fmt.Println("  error: " + err.Error() + "\n")
			continue
		}

		//	正解と比較
		var hit int
		var blow int

		hit, blow = checkHitAndBlow(inputData, ansData)

		if InputLen <= hit {
			fmt.Println("正解！")
			break
		} else {
			fmt.Printf("HIT: %d, BLOW: %d\n\n", hit, blow)
			i++
		}
	}
}

//******************************************************************************
//	初期化
//******************************************************************************
func initialize() {

	//	乱数初期化
	rand.Seed(time.Now().UnixNano())
}

//******************************************************************************
//	タイトル表示
//******************************************************************************
func viewTitle() {

	//	色を出すテスト
	w := ansicolor.NewAnsiColorWriter(os.Stdout)
	text := "%sforeground %sbold%s %sbackground%s\n"
	fmt.Fprintf(w, text, "\x1b[31m", "\x1b[1m", "\x1b[22m", "\x1b[41;32m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[32m", "\x1b[1m", "\x1b[22m", "\x1b[42;31m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[33m", "\x1b[1m", "\x1b[22m", "\x1b[43;34m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[34m", "\x1b[1m", "\x1b[22m", "\x1b[44;33m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[35m", "\x1b[1m", "\x1b[22m", "\x1b[45;36m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[36m", "\x1b[1m", "\x1b[22m", "\x1b[46;35m", "\x1b[0m")
	fmt.Fprintf(w, text, "\x1b[37m", "\x1b[1m", "\x1b[22m", "\x1b[47;30m", "\x1b[0m")

	color.Fprintf(w, "@{r}foreground @{r!}bold@{|} @{gR}background@{|}\n")
	color.Fprintf(w, "@{g}foreground @{g!}bold@{|} @{rG}background@{|}\n")
	color.Fprintf(w, "@{y}foreground @{y!}bold@{|} @{bY}background@{|}\n")
	color.Fprintf(w, "@{b}foreground @{b!}bold@{|} @{yB}background@{|}\n")
	color.Fprintf(w, "@{m}foreground @{m!}bold@{|} @{cM}background@{|}\n")
	color.Fprintf(w, "@{c}foreground @{c!}bold@{|} @{mC}background@{|}\n")
	color.Fprintf(w, "@{w}foreground @{w!}bold@{|} @{kW}background@{|}\n")

	fmt.Println("")
	fmt.Println("////////////////////////////////")
	fmt.Println("  Hit and Blow!")
	fmt.Println("")
	fmt.Printf("  ※正解は%d桁すべて違う数字です\n", InputLen)
	fmt.Println("////////////////////////////////")
	fmt.Println("")
}

//******************************************************************************
//	正解作成
//******************************************************************************
func craeteAnsData() (ansData []int) {

	//	正解配列は要素の数が固定
	ansData = make([]int, InputLen)

	//	同一の数字を使用しないために既出を保存しておく配列は可変で初期長さは0
	work := make([]int, 0, InputLen)

	for i := 0; i < InputLen; {

		bOk := true

		//	0 ～ 9 の乱数値
		val := rand.Intn(10)

		//	すでに使われているかチェック
		jend := len(work)
		for j := 0; j < jend; j++ {

			if val == work[j] {
				bOk = false
				break
			}
		}

		//	初出だったので使用する
		if bOk {
			ansData[i] = val
			work = append(work, val)
			i++
		}
	}

	return
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

	return
}

//******************************************************************************
//	ヒットとブローを算出
//******************************************************************************
func checkHitAndBlow(iInputData []int, iAnsData []int) (hit int, blow int) {

	var noHitData = make([]int, 0, InputLen)
	var workAnsData = make([]int, 0, InputLen)

	//	HITチェック
	var i int

	for i = 0; i < InputLen; i++ {

		if iInputData[i] == iAnsData[i] {
			hit++
		} else {
			noHitData = append(noHitData, iInputData[i])
			workAnsData = append(workAnsData, iAnsData[i])
		}
	}

	//	すべてHITした場合はここで処理終了
	if 0 >= len(noHitData) {
		return
	}

	//	BLOWチェック
	var iend = len(noHitData)
	var jend = len(workAnsData)

	for i = 0; i < iend; i++ {

		for j := 0; j < jend; j++ {

			if noHitData[i] == workAnsData[j] {
				workAnsData[j] = -1
				blow++
			}
		}
	}

	return
}

//******************************************************************************
//	EOF
//******************************************************************************
