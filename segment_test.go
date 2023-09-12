package hedi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegment_NewSegment(t *testing.T) {
	segment := NewSegment("ISA")
	assert.Equal(t, "ISA", segment.ID)
}

func TestSegment_String(t *testing.T) {
	segment := NewSegment("ISA")
	segment.AddElement(Element{Value: "00"})
	segment.AddElement(Element{Value: "ZZ"})
	assert.Equal(t, "ISA*00*ZZ~", segment.String())
}

func TestSegment_DString(t *testing.T) {
	segment := NewSegment("ISA")
	segment.AddElement(Element{Value: "00"})
	segment.AddElement(Element{Value: "ZZ"})
	assert.Equal(t, "ISA*00*ZZ~", segment.DString(DefaultDelimiters))
}

func TestSegment_GetElement(t *testing.T) {
	segment := NewSegment("ISA")
	segment.AddElement(Element{Value: "00"})
	element, found := segment.GetElement(0)
	assert.True(t, found)
	assert.Equal(t, "00", element.Value)

	_, found = segment.GetElement(1)
	assert.False(t, found)
}

func TestSegment_AddElement(t *testing.T) {
	segment := NewSegment("ISA")
	segment.AddElement(Element{Value: "00"})
	assert.Len(t, segment.Elements, 1)
	assert.Equal(t, "00", segment.Elements[0].Value)
}

func TestSegment_SetElement(t *testing.T) {
	segment := NewSegment("ISA")
	segment.SetElement(1, Element{Value: "ZZ"})
	assert.Len(t, segment.Elements, 2)
	assert.Equal(t, "ZZ", segment.Elements[1].Value)
}
