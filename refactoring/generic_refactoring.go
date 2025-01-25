package refactoring

type IWebInterface interface {
	SetValue(value int)
	SetName(name string)
	GetValue() int
	GetName() string
}

type LAN struct {
	value int
	name  string
}

func (L *LAN) SetValue(value int) {
	L.value = value
}

func (L *LAN) SetName(name string) {
	L.name = name
}

func (L *LAN) GetValue() int {
	return L.value
}

func (L *LAN) GetName() string {
	return L.name
}

type Web struct {
	value int
	name  string
}

func (w *Web) SetValue(value int) {
	w.value = value
}

func (w *Web) SetName(name string) {
	w.name = name
}

func (w *Web) GetValue() int {
	return w.value
}

func (w *Web) GetName() string {
	return w.name
}

func AddValue[T IWebInterface](lan T, web T) int {
	return lan.GetValue() + web.GetValue()
}

func AddValue2(lan IWebInterface, web IWebInterface) int {
	return lan.GetValue() + web.GetValue()
}

func AddValue3(lan IWebInterface, web IWebInterface) int {
	lanVal := lan.(*LAN).value
	webVal := web.(*Web).value
	return lanVal + webVal
}

func AddValue4(lan IWebInterface, web IWebInterface) int {
	lanObj, ok := lan.(*LAN)
	var lanVal int
	if ok {
		lanVal = lanObj.value
	}
	webObj, ok := web.(*Web)
	var webVal int
	if ok {
		webVal = webObj.value
	}
	return lanVal + webVal
}
