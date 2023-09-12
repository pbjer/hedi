package hedi

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSegments_String(t *testing.T) {
	seg := Segments{
		Segment{ID: "ISA", Elements: Elements{{Value: "00"}}},
		Segment{ID: "GS", Elements: Elements{{Value: "PO"}}},
	}

	assert.Equal(t, "ISA*00~GS*PO~", seg.String())
}

func TestSegments_DString(t *testing.T) {
	seg := Segments{
		Segment{ID: "ISA", Elements: Elements{{Value: "00"}}},
		Segment{ID: "GS", Elements: Elements{{Value: "PO"}}},
	}
	delimiters := Delimiters{Element: '|', SubElement: ':', Segment: '~'}

	assert.Equal(t, "ISA|00~GS|PO~", seg.DString(delimiters))
}

func TestSegments_WriteTo(t *testing.T) {
	seg := Segments{
		Segment{ID: "ISA", Elements: Elements{{Value: "00"}}},
		Segment{ID: "GS", Elements: Elements{{Value: "PO"}}},
	}
	buf := bytes.NewBuffer([]byte{})

	n, err := seg.WriteTo(buf)
	assert.NoError(t, err)
	assert.Equal(t, int64(13), n)
	assert.Equal(t, "ISA*00~GS*PO~", buf.String())
}

func TestSegments_DWriteTo(t *testing.T) {
	seg := Segments{
		Segment{ID: "ISA", Elements: Elements{{Value: "00"}}},
		Segment{ID: "GS", Elements: Elements{{Value: "PO"}}},
	}
	buf := bytes.NewBuffer([]byte{})
	delimiters := Delimiters{Element: '|', SubElement: ':', Segment: '~'}

	n, err := seg.DWriteTo(delimiters, buf)
	assert.NoError(t, err)
	assert.Equal(t, int64(13), n)
	assert.Equal(t, "ISA|00~GS|PO~", buf.String())
}

func TestSegments_Last(t *testing.T) {
	seg := Segments{
		Segment{ID: "ISA", Elements: Elements{{Value: "00"}}},
		Segment{ID: "GS", Elements: Elements{{Value: "PO"}}},
	}

	lastSeg, ok := seg.Last()
	assert.True(t, ok)
	assert.Equal(t, "GS", lastSeg.ID)
}
