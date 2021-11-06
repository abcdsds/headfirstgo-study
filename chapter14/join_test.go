// 패키지는 테스트할 코드와 같은 패키지에 존재해야 한다.
// 반드시 일 필요는 없으나 노출되지 않은 타입이나 함수에 접근하려면 동일한 패키지에 속해야 한다.
package main

import (
	"fmt"
	"testing"
)

func TestTwoElements(t *testing.T) {
	//t.Error("no test written yet")
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	//t.Error("no test written yet")
	list := []string{"apple", "orange", "peer"}
	want := "apple, orange, and peer"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

func TestOneElements(t *testing.T) {
	list := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(list)
	if got != want {
		t.Errorf(errorString(list, got, want))
	}
}

//helper 함수
func errorString(list []string, got string, want string) string{
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

