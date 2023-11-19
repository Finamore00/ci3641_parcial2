package clientFunctions

import (
	"dataTypeSimulator/dataTypes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtomicCreation(t *testing.T) {
	assert := assert.New(t)

	integer := NewAtomicType("int", 4, 4)
	double := NewAtomicType("double", 8, 8)
	char := NewAtomicType("char", 1, 2)
	boolean := NewAtomicType("bool", 1, 2)

	assert.Equal(integer.Name, "int")
	assert.Equal(integer.AtomicDetails.Size, uint(4))
	assert.Equal(integer.AtomicDetails.Alignment, uint(4))
	assert.Equal(double.Name, "double")
	assert.Equal(double.AtomicDetails.Size, uint(8))
	assert.Equal(double.AtomicDetails.Alignment, uint(8))
	assert.Equal(char.Name, "char")
	assert.Equal(char.AtomicDetails.Size, uint(1))
	assert.Equal(char.AtomicDetails.Alignment, uint(2))
	assert.Equal(boolean.AtomicDetails.Size, uint(1))
	assert.Equal(boolean.AtomicDetails.Alignment, uint(2))

}

func TestStructCreation(t *testing.T) {
	assert := assert.New(t)

	//Tipos miembros del struct
	integer := NewAtomicType("int", 4, 4)
	double := NewAtomicType("double", 8, 8)
	char := NewAtomicType("char", 1, 2)
	boolean := NewAtomicType("bool", 1, 2)

	typeMap := map[string]dataTypes.DataType{
		"integer": integer,
		"double":  double,
		"char":    char,
		"boolean": boolean,
	}

	members := []string{
		"integer",
		"char",
		"char",
		"integer",
		"double",
		"boolean",
	}

	str1 := NewStructType("example", members, &typeMap)

	assert.Equal(str1.Name, "example")
	assert.Equal(str1.StructDetails.Sizes.Regular, uint(25))
	assert.Equal(str1.StructDetails.Sizes.Packed, uint(19))
	assert.Equal(str1.StructDetails.Sizes.Ordered, uint(21))
	assert.Equal(str1.StructDetails.Alignments.Regular, uint(28))
	assert.Equal(str1.StructDetails.Alignments.Packed, uint(20))
	assert.Equal(str1.StructDetails.Alignments.Ordered, uint(24))
	assert.Equal(str1.StructDetails.Wasted.Regular, uint(6))
	assert.Equal(str1.StructDetails.Wasted.Packed, uint(0))
	assert.Equal(str1.StructDetails.Wasted.Ordered, uint(2))

	typeMap["example"] = str1
	members = []string{
		"integer",
		"example",
	}
	//Crear un struct con un struct como miembro
	str2 := NewStructType("example2", members, &typeMap)

	assert.Equal(str2.Name, "example2")
	assert.Equal(uint(53), str2.StructDetails.Sizes.Regular)
	assert.Equal(uint(23), str2.StructDetails.Sizes.Packed)
	assert.Equal(uint(32), str2.StructDetails.Sizes.Ordered)
	assert.Equal(uint(56), str2.StructDetails.Alignments.Regular)
	assert.Equal(uint(24), str2.StructDetails.Alignments.Packed)
	assert.Equal(uint(32), str2.StructDetails.Alignments.Ordered)
}

func TestUnionCreation(t *testing.T) {
	assert := assert.New(t)

	//Tipos miembros del struct
	integer := NewAtomicType("int", 4, 4)
	double := NewAtomicType("double", 8, 8)
	char := NewAtomicType("char", 1, 2)
	boolean := NewAtomicType("bool", 1, 2)
	raro := NewAtomicType("raro", 3, 6)

	members := []string{
		"integer",
		"char",
		"raro",
		"double",
	}

	typeMap := map[string]dataTypes.DataType{
		"integer": integer,
		"double":  double,
		"char":    char,
		"boolean": boolean,
		"raro":    raro,
	}

	un1 := NewUnionType("un1", members, &typeMap)

	assert.Equal("un1", un1.Name)
	assert.Equal(uint(8), un1.UnionDetails.Size.Regular)
	assert.Equal(uint(8), un1.UnionDetails.Size.Packed)
	assert.Equal(uint(8), un1.UnionDetails.Size.Ordered)
	assert.Equal(uint(24), un1.UnionDetails.Alignment.Regular)
	assert.Equal(uint(24), un1.UnionDetails.Alignment.Packed)
	assert.Equal(uint(24), un1.UnionDetails.Alignment.Ordered)

	str := NewStructType("str", []string{
		"char",
		"char",
		"char",
		"char",
		"char",
	}, &typeMap)

	typeMap["str"] = str

	un2 := NewUnionType("un2", []string{
		"double",
		"raro",
		"str",
	}, &typeMap)

	assert.Equal("un2", un2.Name)
	assert.Equal(uint(9), un2.UnionDetails.Size.Regular)
	assert.Equal(uint(8), un2.UnionDetails.Size.Packed)
	assert.Equal(uint(9), un2.UnionDetails.Size.Ordered)
	assert.Equal(uint(24), un2.UnionDetails.Alignment.Regular)
	assert.Equal(uint(24), un2.UnionDetails.Alignment.Packed)
	assert.Equal(uint(24), un2.UnionDetails.Alignment.Ordered)

}
