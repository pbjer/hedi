package hedi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElement_String(t *testing.T) {
	t.Run("Element with no sub-elements returns value", func(t *testing.T) {
		e := &Element{Value: "00"}
		assert.Equal(t, e.Value, e.String())
	})

	t.Run("Element with sub-elements returns value and delimited sub-elements", func(t *testing.T) {
		e := &Element{Value: "IA", SubElements: []string{"1", "2", "3"}}
		assert.Equal(t, "IA>1>2>3", e.String())
	})
}

func TestElement_DString(t *testing.T) {
	t.Run("Element with no sub-elements returns value", func(t *testing.T) {
		e := &Element{Value: "00"}
		d := Delimiters{SubElement: '+'}
		assert.Equal(t, e.Value, e.DString(d))
	})

	t.Run("Element with sub-elements returns value and delimited sub-elements", func(t *testing.T) {
		e := &Element{Value: "IA", SubElements: []string{"1", "2", "3"}}
		d := Delimiters{SubElement: '+'}
		assert.Equal(t, "IA+1+2+3", e.DString(d))
	})
}

func TestElement_AddSubElement(t *testing.T) {
	e := &Element{Value: "20230911"}
	e.AddSubElement("1200")
	assert.Equal(t, []string{"1200"}, e.SubElements)
}

func TestElements_Last(t *testing.T) {
	// Creating some mock 850-specific elements
	elem1 := Element{Value: "PO1", SubElements: []string{"001"}}
	elem2 := Element{Value: "CTP", SubElements: []string{"USD", "100"}}

	t.Run("Returns last element if exists", func(t *testing.T) {
		elements := Elements{elem1, elem2}
		lastElem, ok := elements.Last()
		assert.True(t, ok)
		assert.Equal(t, "CTP", lastElem.Value)
		assert.Equal(t, []string{"USD", "100"}, lastElem.SubElements)
	})

	t.Run("Returns nil and false if no elements", func(t *testing.T) {
		var elements Elements
		lastElem, ok := elements.Last()
		assert.False(t, ok)
		assert.Nil(t, lastElem)
	})
}
