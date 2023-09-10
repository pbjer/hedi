package hedi

type Spec struct {
	Segments SpecSegments
}

type SpecSegments []SpecSegment

type SpecSegment struct {
	ID          string
	Requirement string // (M) mandatory, (O) optional
	Max         int
	Elements    SpecElements
	Loop        SpecLoop
}

type SpecLoop struct {
	Requirement string // (M) mandatory, (O) optional
	Repeat      int    // -1 = inf
	Segments    SpecSegments
}

type SpecElements []SpecElement

type SpecElement struct {
	ID         string
	Type       string
	IsRequired bool
	Min        int
	Max        int
}

var example = Spec{
	Segments: SpecSegments{
		{
			ID:          "ISA",
			Requirement: "Mandatory",
			Max:         1,
			Elements: SpecElements{
				{
					ID:         "ISA-01",
					Type:       "ID",
					IsRequired: true,
					Min:        5,
					Max:        5,
				},
				{
					ID:         "ISA-02",
					Type:       "AN",
					IsRequired: true,
					Min:        10,
					Max:        10,
				},
			},
		},
	},
}
