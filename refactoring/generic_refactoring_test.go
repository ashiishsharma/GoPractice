package refactoring

import "testing"

func TestGenericRefactoring(t *testing.T) {
	var lan IWebInterface
	lan = &LAN{value: 20}

	var web IWebInterface
	web = &Web{value: 25}

	var result = lan.GetValue() + web.GetValue()

	println(result)

	println(AddValue(lan, web))
	println(AddValue2(lan, web))
	println(AddValue3(lan, web))
	println(AddValue4(lan, web))
	//if value != actual {
	//	t.Errorf("expected '%d' but got '%d'", actual, value)
	//}
}
