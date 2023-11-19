package calcFunctions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpressionCalc(t *testing.T) {
	assert := assert.New(t)
	expPre := map[string]int{
		//Expresiones pre-fijas
		"+ + 5 * 6 1 2": 13, //5 + 6*1 + 2
		"* + 2 3 + 1 4": 25, //(2+3)*(1+4)
		"+ * 7 8 * 2 4": 64, //7*8 + 2*4
		"* 0 + 1 + 2 3": 0,  //0*(1 + 2 + 3)
	}

	expPos := map[string]int{
		//Expresiones post-fijas
		"5 6 1 * + 2 +": 13, //5 + 6*1 + 2
		"2 3 + 1 4 + *": 25, //(2+3)*(1+4)
		"7 8 * 2 4 * +": 64, //7*8 + 2*4
		"0 1 2 + 3 + *": 0,  //0*(1 + 2 + 3)
	}

	for k, v := range expPre {
		assert.Equal(CalculateExprResult(strings.Split(k, " "), PRE), v, "The 2 values should match")
	}

	for k, v := range expPos {
		assert.Equal(CalculateExprResult(strings.Split(k, " "), POST), v, "The 2 values should match")
	}
}

func TestInfixConversion(t *testing.T) {
	assert := assert.New(t)
	expPre := map[string]string{
		//Expresiones pre-fijas
		"+ + 5 * 6 1 2": "5+6*1+2",     //5 + 6*1 + 2
		"* + 2 3 + 1 4": "(2+3)*(1+4)", //(2+3)*(1+4)
		"+ * 7 8 * 2 4": "7*8+2*4",     //7*8 + 2*4
		"* 0 + 1 + 2 3": "0*(1+2+3)",   //0*(1 + 2 + 3)
	}

	expPos := map[string]string{
		//Expresiones post-fijas
		"5 6 1 * + 2 +": "5+6*1+2",     //5 + 6*1 + 2
		"2 3 + 1 4 + *": "(2+3)*(1+4)", //(2+3)*(1+4)
		"7 8 * 2 4 * +": "7*8+2*4",     //7*8 + 2*4
		"0 1 2 + 3 + *": "0*(1+2+3)",   //0*(1 + 2 + 3)
	}

	for k, v := range expPre {
		assert.Equal(ConvertToInfix(strings.Split(k, " "), PRE), v, "The 2 values should match")
	}

	for k, v := range expPos {
		assert.Equal(ConvertToInfix(strings.Split(k, " "), POST), v, "The 2 values should match")
	}
}
