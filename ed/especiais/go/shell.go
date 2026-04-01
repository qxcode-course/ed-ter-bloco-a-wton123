package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sort"

)

type Pair struct {
	One int
	Two int
}



func occurr(vet []int) []Pair {
	countMap := make(map[int]int)

	for _, v := range vet {
		val := v
		if val < 0 {
			val = -val // usar valor absoluto
		}
		countMap[val]++
	}

	var result []Pair
	for k, v := range countMap {
		result = append(result, Pair{k, v})
	}

	// Ordenar pelo nível de stress
	sort.Slice(result, func(i, j int) bool {
		return result[i].One < result[j].One
	})

	return result
}

//- **times** Uma ou mais pessoas de mesmo stress seguidas formam um time. Gere um vetor compactado indicando o nível de stress e quantas pessoas tem em cada time?
func teams(vet []int) []Pair {
    if len(vet) == 0 {
        return nil
    }
    var result []Pair
    currentValue := vet[0]
    count := 1
    for i := 1; i < len(vet); i++ {
        if vet[i] == currentValue {
            count++
        } else {
            result = append(result, Pair{currentValue, count})
            currentValue = vet[i]
            count = 1
        }
    }
    result = append(result, Pair{currentValue, count})
    return result
}

//- **mnext**: Apresente um mapa colocando 1 nas posições que existem homens ao lado de pelo menos uma mulher.
func mnext(vet []int) []int {
    res := make([]int, len(vet))
    for i := 0; i < len(vet); i++ {
        if vet[i] > 0 { // homem
            // verificar vizinhos
            if (i > 0 && vet[i-1] < 0) || (i < len(vet)-1 && vet[i+1] < 0) {
                res[i] = 1
            }
        }
    }
    return res
}

//- **alone**: quais posições existem homens que não estão do lado de nenhuma mulher?
func alone(vet []int) []int {
    res := make([]int, len(vet))
    for i := 0; i < len(vet); i++ {
        if vet[i] > 0 { // homem
            // homem é "alone" se nenhum vizinho é mulher (<0)
            if (i == 0 || vet[i-1] >= 0) && (i == len(vet)-1 || vet[i+1] >= 0) {
                res[i] = 1
            } else {
                res[i] = 0
            }
        } else {
            res[i] = 0 // mulheres sempre 0
        }
    }
    return res
}

// **couple**: Casais são formados quando quando um homem e uma mulher com o mesmo nível de stress se encontram. Retorne a quantidade de casais que podem ser formados.
func couple(vet []int) int {
    men := make(map[int]int)
    women := make(map[int]int)
    for _, v := range vet {
        if v > 0 { // supondo homem positivo
            men[v]++
        } else { // mulher negativa
            women[-v]++
        }
    }
    count := 0
    for k, v := range men {
        if w, ok := women[k]; ok {
            if v < w {
                count += v
            } else {
                count += w
            }
        }
    }
    return count
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

// **subseq**: Dada uma sequência de valores, procure essa sequência na fila e retorne a primeira posição onde ela começar.
func subseq(vet []int, seq []int) int {
    n, m := len(vet), len(seq)
    for i := 0; i <= n-m; i++ {
        match := true
        for j := 0; j < m; j++ {
            if vet[i+j] != seq[j] {
                match = false
                break
            }
        }
        if match {
            return i
        }
    }
    return -1
}

//- **erase**: dado a lista com os índices de todas as pessoas que saíram da fila, qual a fila resultante?
func erase(vet []int, posList []int) []int {
    toRemove := make(map[int]bool)
    for _, p := range posList {
        toRemove[p] = true
    }
    var result []int
    for i, v := range vet {
        if !toRemove[i] {
            result = append(result, v)
        }
    }
    return result
}

//- **clear**: dado um valor, remova todas as vezes que esse valor aparece na lista.
func clear(vet []int, value int) []int {
    var result []int
    for _, v := range vet {
        if v != value { // apenas adiciona se for diferente do valor a remover
            result = append(result, v)
        }
    }
    return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
