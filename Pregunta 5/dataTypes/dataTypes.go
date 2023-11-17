package dataTypes

/*
-dataType
|_Atomic
|    |_Name
|    |_Size
|    |_Alignment
|_Struct
|    |_Name
|    |_Members [List of type names]
|_Union
    |_Name
    |_Alternatives [List of type names]
*/

type Atomic struct {
	Size      uint
	Alignment uint
}

type Struct struct {
	Sizes struct {
		Regular uint
		Packed  uint
		Ordered uint
	}
	Alignments struct {
		Regular uint
		Packed  uint
		Ordered uint
	}
	Wasted struct {
		Regular uint
		Packed  uint
		Ordered uint
	}
	MemberTypes []string
}

type Union struct {
	Size struct {
		Regular uint
		Packed  uint
		Ordered uint
	}
	Alignment struct {
		Regular uint
		Packed  uint
		Ordered uint
	}
	TypeAlternatives []string
}

type DataType struct {
	Name          string
	AtomicDetails *Atomic
	StructDetails *Struct
	UnionDetails  *Union
}
