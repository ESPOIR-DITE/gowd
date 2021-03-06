package gowd

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestElement_SetAttributes(t *testing.T) {
	elem := NewElement("div")
	event := EventElement{}
	event.Properties = make(map[string]string)
	event.Properties["value"] = "123"
	event.Properties["id"] = "myID"
	elem.SetAttributes(&event)
	assert.EqualValues(t, event.GetID(), "myID")
	assert.EqualValues(t, event.GetValue(), "123")
	assert.EqualValues(t, event.GetID(), elem.GetID())
	assert.EqualValues(t, event.GetValue(), elem.GetValue())
}

func TestElement_Enable(t *testing.T) {
	elem := NewElement("div")
	elem.Disable()
	val, found := elem.GetAttribute("disabled")
	assert.EqualValues(t, val, "true")
	assert.True(t, found)
	elem.Enable()
	val, found = elem.GetAttribute("disabled")
	assert.EqualValues(t, val, "")
	assert.False(t, found)
}

func TestElement_SetClass(t *testing.T) {
	elem := NewElement("div")
	elem.SetClass("well sunken")
	class, _ := elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well sunken")
	elem.SetClass("upper")
	class, _ = elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well sunken upper")
	elem.UnsetClass("sunken")
	class, _ = elem.GetAttribute("class")
	assert.EqualValues(t, strings.TrimSpace(class), "well  upper")
}

func TestElement_Hide(t *testing.T) {
	em := NewElementMap()
	elem, err := ParseElement(`<div><p id="text">text</p></div>`, em)
	if err != nil {
		t.Fatal(err)
	}
	p := em["text"]
	assert.False(t, p.Hidden)
	testOuput(t, elem, "<div><p id=\"text\">text</p></div>")
	p.Hide()
	assert.True(t, p.Hidden)
	testOuput(t, elem, "<div><!--p--></div>")
	p.SetText("Show me!!")
	p.Show()
	testOuput(t, elem, "<div><p id=\"text\">Show me!!</p></div>")
	elem.RemoveElement(p)
	testOuput(t, elem, "<div></div>")
}

func TestElement_SetValue(t *testing.T) {
	elem := NewElement("div")
	elem.SetValue("hoho")
	assert.Equal(t, "hoho", elem.GetValue())
	testOuput(t, elem, "<div id=\"_div1\" value=\"hoho\"></div>")
}

func TestElement_AutoFocus(t *testing.T) {
	elem := NewElement("div")
	_, exists := elem.GetAttribute("autofocus")
	assert.False(t, exists)
	elem.AutoFocus()
	testOuput(t, elem, "<div id=\"_div1\" autofocus=\"\"></div>")
}

func TestElement_SetElement(t *testing.T) {
	elem := NewElement("div")
	assert.Empty(t, elem.Kids)
	for i := 0; i < 10; i++ {
		elem.AddElement(NewElement("p"))
	}
	assert.Equal(t, 10, len(elem.Kids))
	elem.SetElement(NewElement("a"))
	assert.Equal(t, 1, len(elem.Kids))
	assert.Equal(t, "a", elem.Kids[0].data)
}
