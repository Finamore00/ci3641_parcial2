package clientFunctions

import (
	"dataTypeSimulator/dataTypes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
Funcion invocada al recibir un comando de definición de tipos
atómicos. Se encarga de parsear el comando ingresado, verificar
su correctitud y extrae la información relevante del mismo. Si
hay algún error en el formateo del comando o los argumentos son
incorrectos retorna un error
*/
func ParseAtomicCmd(s string) (string, uint, uint, error) {
	words := strings.Split(s, " ")
	var retVals struct {
		name  string
		size  uint64
		align uint64
		err   error
	}

	if len(words) != 4 {
		retVals.err = errors.New("failed to parse atomic declaration")
		return retVals.name, uint(retVals.size), uint(retVals.align), retVals.err
	}

	retVals.name = words[1]
	retVals.size, retVals.err = strconv.ParseUint(words[2], 10, 32)
	retVals.align, retVals.err = strconv.ParseUint(words[3], 10, 32)

	return retVals.name, uint(retVals.size), uint(retVals.align), retVals.err
}

/*
Función invocada al momento en que se solicita definir un nuevo tipo registro.
Se encarga de parsear el comando ingresado por el usuario, verificar su
correctitud y extraer del mismo la información relevante. Si hay algún error
en el comando ingresado retorna un valor de error.
*/
func ParseStructCmd(s string, dic *map[string]dataTypes.DataType) (string, []string, error) {
	words := strings.Split(s, " ")
	var retVals struct {
		name    string
		members []string
		err     error
	}
	retVals.err = nil

	if len(words) < 3 {
		retVals.err = errors.New("struct declaration has too few args")
		return retVals.name, retVals.members, retVals.err
	}

	//Se extrae el nombre del struct
	retVals.name = words[1]
	retVals.members = []string{} //Inicializar slice
	for _, name := range words[2:] {
		_, ok := (*dic)[name]
		if !ok {
			retVals.err = errors.New("non-existing type passed as struct member")
		}
		retVals.members = append(retVals.members, name)
	}

	return retVals.name, retVals.members, retVals.err
}

/*
Función que se invoca al momento de definir un nuevo registro variante. Se
encarga de parsear el comando ingresado, verificar su correctitud y extraer
del mismo la información relevante. Si hay algún error en el formateo
del comando, o alguno de los tipos ingresados como alternativas de la unión
no está definido, retorna error.
*/
func ParseUnionCmd(s string, dic *map[string]dataTypes.DataType) (string, []string, error) {
	words := strings.Split(s, " ")
	var retVals struct {
		name string
		alts []string
		err  error
	}
	retVals.err = nil

	if len(words) < 3 {
		retVals.err = errors.New("too few arguments in union declaration")
	}

	retVals.name = words[1]
	retVals.alts = []string{}
	for _, ty := range words[2:] {
		_, ok := (*dic)[ty]
		if !ok {
			retVals.err = errors.New("non-existing argument in alternatives list")
		}
		retVals.alts = append(retVals.alts, ty)
	}

	return retVals.name, retVals.alts, retVals.err
}

/*
Función invocada en el momento en que se solicita mostrar los detalles de un tipo.
Se encarga de verificar la correctitud del comando y de extraer la información
relevante del mismo. Si hay algún error en el formateo del comando o los argumentos
son erróneos retorna error.
*/
func ParseShowCmd(s string) (string, error) {
	words := strings.Split(s, " ")

	if len(words) != 2 {
		return "", errors.New("failed to parse show command")
	}

	return words[1], nil
}

/*
Función que recibe el nombre, tamaño de representación, y alineación de un nuevo
tipo atómico y devuelve el objeto DataType que lo representa.
*/
func NewAtomicType(name string, size uint, alignment uint) dataTypes.DataType {
	return dataTypes.DataType{
		Name: name,
		AtomicDetails: &dataTypes.Atomic{
			Size:      size,
			Alignment: alignment,
		},
		StructDetails: nil,
		UnionDetails:  nil,
	}
}

/*
Funcion que recibe el nombre y los tipos posibles de un nuevo tipo union
y devuelve el objeto DataType que lo representa. Se asume que todos los
tipos ingresados como alternativas de la unión existen.
*/
func NewUnionType(name string, alts []string, dic *map[string]dataTypes.DataType) dataTypes.DataType {
	members := []dataTypes.DataType{}

	for _, t := range alts {
		members = append(members, (*dic)[t])
	}

	//Se calculan los detalles de la union para cada configuracion de registros
	regSize, regAlign := unionMemReg(members)
	pacSize, pacAlign := unionMemPacked(members)
	optSize, optAlign := unionMemOpt(members)

	newUnion := dataTypes.DataType{
		Name:          name,
		AtomicDetails: nil,
		StructDetails: nil,
		UnionDetails: &dataTypes.Union{
			Size: struct {
				Regular uint
				Packed  uint
				Ordered uint
			}{
				Regular: regSize,
				Packed:  pacSize,
				Ordered: optSize,
			},
			Alignment: struct {
				Regular uint
				Packed  uint
				Ordered uint
			}{
				Regular: regAlign,
				Packed:  pacAlign,
				Ordered: optAlign,
			},
			TypeAlternatives: alts,
		},
	}

	return newUnion
}

/*
Funcion que retorna el tamaño y alineamiento de un registro variante
cuyos posibles tipos son los hallados en members asumiendo que los
registros disponen sus miembros siguiendo el orden de declaracion
y respetando el alineamiento.
*/
func unionMemReg(members []dataTypes.DataType) (uint, uint) {
	size := uint(0)
	memberAligns := []uint{}
	align := uint(0)

	for _, t := range members {
		//Esencialmente buscar la alternativa de tamaño más grande
		switch {
		case t.AtomicDetails != nil:
			if t.AtomicDetails.Size > size {
				size = t.AtomicDetails.Size
			}
			memberAligns = append(memberAligns, t.AtomicDetails.Alignment)
		case t.StructDetails != nil:
			if t.StructDetails.Sizes.Regular > size {
				size = t.StructDetails.Sizes.Regular
			}
			memberAligns = append(memberAligns, t.StructDetails.Alignments.Regular)
		case t.UnionDetails != nil:
			if t.UnionDetails.Size.Regular > size {
				size = t.UnionDetails.Size.Regular
			}
			memberAligns = append(memberAligns, t.UnionDetails.Alignment.Regular)
		}
	}

	//Se saca el m.c.m de todas las alineaciones de los miembros
	align = memberAligns[0]
	for _, al := range memberAligns {
		align = mcm(align, al)
	}

	return size, align

}

/*
Funcion que retorna el tamaño y alineamiento de un registro variante
cuyos posibles tipos son los hallados en members asumiendo que los
registros disponen sus miembros mediante empaquetado para optimizacion
de espacio.
*/
func unionMemPacked(members []dataTypes.DataType) (uint, uint) {
	size := uint(0)
	align := uint(0)
	memberAligns := []uint{}

	for _, t := range members {
		//Esencialmente se busca el tipo de dato mas grande
		switch {
		case t.AtomicDetails != nil:
			if t.AtomicDetails.Size > size {
				size = t.AtomicDetails.Size
			}
			memberAligns = append(memberAligns, t.AtomicDetails.Alignment)
		case t.StructDetails != nil:
			if t.StructDetails.Sizes.Packed > size {
				size = t.StructDetails.Sizes.Packed
			}
			memberAligns = append(memberAligns, t.StructDetails.Alignments.Packed)
		case t.UnionDetails != nil:
			if t.UnionDetails.Size.Packed > size {
				size = t.UnionDetails.Size.Packed
			}
			memberAligns = append(memberAligns, t.UnionDetails.Alignment.Packed)
		}
	}

	//Se calcula el alineamiento como el m.c.m. de los alineamientos de los miembros
	align = memberAligns[0]
	for _, al := range memberAligns {
		align = mcm(align, al)
	}

	return size, align
}

/*
Funcion que retorna el tamaño y alineamiento de un registro variante
cuyos posibles tipos son los hallados en members. Se asume que los
registros reordenan sus miembros para minimizar el uso de espacio
respetando alineamientos.
*/
func unionMemOpt(members []dataTypes.DataType) (uint, uint) {
	size := uint(0)
	align := uint(0)
	memberAligns := []uint{}

	for _, t := range members {
		//Esencialmente se busca el tipo de dato más grande
		switch {
		case t.AtomicDetails != nil:
			if t.AtomicDetails.Size > size {
				size = t.AtomicDetails.Size
			}
			memberAligns = append(memberAligns, t.AtomicDetails.Alignment)
		case t.StructDetails != nil:
			if t.StructDetails.Sizes.Ordered > size {
				size = t.StructDetails.Sizes.Ordered
			}
			memberAligns = append(memberAligns, t.StructDetails.Alignments.Ordered)
		case t.UnionDetails != nil:
			if t.UnionDetails.Size.Ordered > size {
				size = t.UnionDetails.Size.Ordered
			}
			memberAligns = append(memberAligns, t.UnionDetails.Alignment.Ordered)
		}
	}

	//Se calcula el m.c.m de las alineaciones de los miembros
	align = memberAligns[0]
	for _, al := range memberAligns {
		align = mcm(align, al)
	}

	return size, align
}

/*
Función que recibe el nombre y los tipos miembros de un nuevo tipo struct
y devuelve el objeto DataType que lo representa. Se asume que todos los
tipos ingresados existen.
*/
func NewStructType(name string, members []string, dic *map[string]dataTypes.DataType) dataTypes.DataType {
	//Se obtiene la información de los tipos miembros
	mem := []dataTypes.DataType{}
	for _, ty := range members {
		mem = append(mem, (*dic)[ty])
	}

	//Se calculan los detalles del struct para cada modalidad de agrupacion
	regSize, regAlign, regWaste := structMemReg(mem)
	pacSize, pacAlign, pacWaste := structMemPack(mem)
	optSize, optAlign, optWaste := structMemOpt(mem)
	//Falta por hacer el tamaño óptimo. Seguiré pensando en eso

	//Se crea el nuevo objeto y se retorna
	newStr := dataTypes.DataType{
		Name:          name,
		AtomicDetails: nil,
		StructDetails: &dataTypes.Struct{
			Sizes: struct {
				Regular uint
				Packed  uint
				Ordered uint
			}{
				Regular: regSize,
				Packed:  pacSize,
				Ordered: optSize,
			},
			Alignments: struct {
				Regular uint
				Packed  uint
				Ordered uint
			}{
				Regular: regAlign,
				Packed:  pacAlign,
				Ordered: optAlign,
			},
			Wasted: struct {
				Regular uint
				Packed  uint
				Ordered uint
			}{
				Regular: regWaste,
				Packed:  pacWaste,
				Ordered: optWaste,
			},
			MemberTypes: members,
		},
		UnionDetails: nil,
	}

	return newStr
}

/*
Funcion que calcula el tamaño total de un struct si sus miembros son dispuestos
siguiendo el orden de declaración, y respetando sus alineamientos.
*/
func structMemReg(members []dataTypes.DataType) (uint, uint, uint) {
	size := uint(0)
	align := uint(0)
	waste := uint(0)
	accum := uint(0)

	for _, ty := range members {
		switch {
		case ty.AtomicDetails != nil:
			if gap := size % ty.AtomicDetails.Alignment; gap != 0 {
				size += ty.AtomicDetails.Alignment - gap
			}
			size += ty.AtomicDetails.Size
			accum += ty.AtomicDetails.Size
		case ty.StructDetails != nil:
			if gap := size % ty.StructDetails.Alignments.Regular; gap != 0 {
				size += ty.StructDetails.Alignments.Regular - gap
			}
			size += ty.StructDetails.Sizes.Regular
			accum += ty.StructDetails.Sizes.Regular
		case ty.UnionDetails != nil:
			if gap := size % ty.UnionDetails.Alignment.Regular; gap != 0 {
				size += ty.UnionDetails.Alignment.Regular - gap
			}
			size += ty.UnionDetails.Size.Regular
			accum += ty.UnionDetails.Size.Regular
		}
	}

	waste = size - accum
	if gap := size % 4; gap == 0 {
		align = size
	} else {
		align = size + (4 - gap)
	}

	return size, align, waste
}

/*
Función que calcula el tamaño total de un struct si sus miembros son empaquetados
para optimizar el uso de espacio
*/
func structMemPack(members []dataTypes.DataType) (uint, uint, uint) {
	accum := uint(0)
	align := uint(0)
	for _, ty := range members {
		switch {
		case ty.AtomicDetails != nil:
			accum += ty.AtomicDetails.Size
		case ty.StructDetails != nil:
			accum += ty.StructDetails.Sizes.Packed
		case ty.UnionDetails != nil:
			accum += ty.UnionDetails.Size.Packed
		}
	}

	if gap := accum % 4; gap == 0 {
		align = accum
	} else {
		align = accum + (4 - gap)
	}
	return accum, align, uint(0)
}

/*
Funcion que calcula el tamaño total de un struct si sus miembros son reordenados
para usar el menor espacio posible respetando los alineamientos.
*/
func structMemOpt(members []dataTypes.DataType) (uint, uint, uint) {
	minWaste := uint(4294967295) //Max value of uint
	minSize := uint(4294967295)
	var minAlign uint

	possibleOrderings := permutation(members)

	for _, order := range possibleOrderings {
		s, a, w := structMemReg(order)
		if s < minSize {
			minWaste = w
			minSize = s
			minAlign = a
		}
	}

	return minSize, minAlign, minWaste
}

/*
Funcion auxiliar para obtener todas las permutaciones de un slice
de tipos de datos
*/
func permutation(xs []dataTypes.DataType) (permuts [][]dataTypes.DataType) {
	var rc func([]dataTypes.DataType, int)
	rc = func(a []dataTypes.DataType, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]dataTypes.DataType{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

/*
Funcion que calcula el minimo comun multiplo (m.c.m.) de dos enteros sin signo.
Necesario para alineamiento de registros variantes.
*/
func mcm(a uint, b uint) uint {
	//Obtener máximo común divisor
	x, y := a, b
	for x%y != 0 {
		r := x % y
		x = y
		y = r
	}
	//Ahora y = MCD(a, b)
	return (a * b) / y
}

/*
Funcion que recibe un tipo de dato y muestra en pantalla su información
*/
func ShowDataType(t dataTypes.DataType) {
	fmt.Println("\tNombre del tipo de dato:", t.Name)
	switch {
	case t.AtomicDetails != nil:
		fmt.Println("\tTamaño de representación del dato:", t.AtomicDetails.Size)
		fmt.Println("\tAlineación del dato:", t.AtomicDetails.Alignment)
	case t.StructDetails != nil:
		fmt.Println("\tTamaño de representación del dato:")
		fmt.Println("\t\tDatos sin empaquetar:", t.StructDetails.Sizes.Regular)
		fmt.Println("\t\tDatos empaquetados:", t.StructDetails.Sizes.Packed)
		fmt.Println("\t\tOrdenamiento óptimo:", t.StructDetails.Sizes.Ordered)
		fmt.Println("\tAlineación del dato:")
		fmt.Println("\t\tDatos sin empaquetar:", t.StructDetails.Alignments.Regular)
		fmt.Println("\t\tDatos empaquetados:", t.StructDetails.Alignments.Packed)
		fmt.Println("\t\tOrdenamiento óptimo:", t.StructDetails.Alignments.Ordered)
		fmt.Println("\tEspacio desperdiciado:")
		fmt.Println("\t\tDatos sin empaquetar:", t.StructDetails.Wasted.Regular)
		fmt.Println("\t\tDatos empaquetados:", t.StructDetails.Wasted.Packed)
		fmt.Println("\t\tOrdenamiento óptimo:", t.StructDetails.Wasted.Ordered)
	case t.UnionDetails != nil:
		fmt.Println("\tTamaño de representación de dato:")
		fmt.Println("\t\tDatos sin empaquetar:", t.UnionDetails.Size.Regular)
		fmt.Println("\t\tDatos empaquetados:", t.UnionDetails.Size.Packed)
		fmt.Println("\t\tOrdenamiento óptimo:", t.UnionDetails.Size.Ordered)
		fmt.Println("\tAlineación del dato:")
		fmt.Println("\t\tDatos sin empaquetar:", t.UnionDetails.Alignment.Regular)
		fmt.Println("\t\tDatos empaquetados:", t.UnionDetails.Alignment.Packed)
		fmt.Println("\t\tOrdenamiento óptimo:", t.UnionDetails.Alignment.Ordered)
	}
}

/*
Da la bienvenida al usuario :)
*/
func PrintWelcome() {
	fmt.Println("¡Bienvenido al simulador de alineamiento de tipos!")
	fmt.Println("Para ver comandos disponibles, ver 'help'")
}

/*
Indica al usuario el surgimiento de algún error en el
comando de declaración atómica
*/
func PrintAtomicCmdErr() {
	fmt.Println("Argumentos erróneos o número de argumentos incorrecto.")
	fmt.Println("Para ver uso del comando ATOMICO, ver 'help'")
}

/*
Indica al usuario que cometio un error en el comando de
muestra
*/
func PrintShowCmdErr() {
	fmt.Println("Número de argumentos incorrecto.")
	fmt.Println("Para ver uso del comando MOSTRAR, ver 'help'")
}

/*
Indica la presencia de un error en el comando de declaración de struct
*/
func PrintStructCmdErr() {
	fmt.Println("Declaración de struct mal formateada o alguno de los tipos no existe.")
	fmt.Println("Para ver uso del comando STRUCT, ver 'help'")
}

/*
Indica la presencia de algún error en el comando de declaración de registros variantes
*/
func PrintUnionCmdErr() {
	fmt.Println("Declaración de unión mal formateada o alguno de los tipos ingresados no existe.")
	fmt.Println("Para ver uso del comando UNION, ver 'help'")
}

/*
Indica al usuario que el tipo de dato que solicitó mostrar no está
definido.
*/
func PrintTypeNotFound(name string) {
	fmt.Printf("No existe un tipo de datos de nombre %v.\n", name)
}

/*
Muestra manual de uso al usuario
*/
func PrintHelp() {
	fmt.Println("COMANDOS DISPONIBLES")
	fmt.Println("\ti. ATOMICO <nombre> <representacion> <alineacion>")
	fmt.Println("\t\tDefine un nuevo tipo atómico de nombre <nombre>, que ocupa <representación> bytes y está alineado a <alineacion> bytes")
	fmt.Println("\tii. STRUCT <nombre> [<tipo>]")
	fmt.Println("\t\tDefine un nuevo registro de nombre <nombre>, cuyos campos son de los tipos ingresados en [<tipo>].")
	fmt.Println("\tiii. UNION <nombre> [<tipo>]")
	fmt.Println("\t\tDefine un registro variante de nombre <nombre> y cuyos tipos posibles son aquellos ingresados en [<tipo>]")
	fmt.Println("\tiv. DESCRIBIR <nombre>")
	fmt.Println("\t\tMuestra en pantalla información del tipo de nombre <nombre>")
	fmt.Println("\tv. SALIR")
	fmt.Println("\t\tCulmina la ejecución del programa")
}

/*
Indica al usuario que se cometió un error en el ingreso de un comando
*/
func PrintGenericCmdErr() {
	fmt.Println("Comando ingresado desconocido o mal formateado.")
	fmt.Println("Para ver comandos disponibles y su uso ver 'help'")
}
