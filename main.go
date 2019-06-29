package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the kingRichardKnights function below.
 */
type abn struct {
	a     uint64
	b     uint64
	n     uint64
	ran   int
	angle int
	com   int
}
type sortstruct []abn

func (a sortstruct) Len() int           { return len(a) }
func (a sortstruct) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortstruct) Less(i, j int) bool { return a[i].ran < a[j].ran }

func tranf(k abn, com []uint64) (uint64, uint64) {
	var a, b uint64
	switch k.angle {
	case 0:
		a = com[0] + k.a - 1
		b = com[1] + k.b - 1
		break
	case 1:

		a = com[2] + com[1] - k.a + 1
		b = com[0] + k.b - 1
		break
	case 2:
		a = com[2] + com[0] - k.a + 1
		b = com[2] + com[1] - k.b + 1
		break

	case 3:
		a = com[1] + k.a - 1
		b = com[2] + com[0] - k.b + 1
		break
	}
	return a, b
}
func analyse(k abn, x, y, ksx, ksy, com uint64, xi int) abn {
	//	var dis, ang, ring uint64

	switch k.angle {
	case 0:
		k.a -= x
		k.b -= y

	case 1:
		k.a -= ksy
		k.b -= x
	case 2:
		k.a -= ksx
		k.b -= ksy
	case 3:
		k.a -= y
		k.b -= ksx

	}
	//	println(k.a, k.b, com)
	if k.a == 0 || k.b == 0 || k.a > com+1 || k.b > com+1 {
		return abn{0, 0, 0, 0, 0, 0}
	}
	k.angle = (k.angle + 1) % 4
	k.com++
	return k
}
func kingRichardKnights(n uint64, s uint64, knights []uint64, com [][]uint64) [][]uint64 {

	pos := make([]abn, len(knights))
	result := make([][]uint64, len(knights))
	incom := []uint64{1, 1, n - 1}
	var x, y, ksx, ksy uint64

	guid := 0
	for i := 0; i < len(knights); i++ {
		pos[i].a = (knights[i] / n) + 1
		pos[i].b = (knights[i] % n) + 1
		pos[i].n = knights[i]
		pos[i].ran = i
		pos[i].com = -1
		pos[i].angle = 0

		//	fmt.Fprintf(os.Stdout, "result %d \n", pos[i].ran)
	}

	for ix, v := range com {
		if ix != 0 {
			x = v[0] - com[ix-1][0]
			y = v[1] - com[ix-1][1]
			ksx = com[ix-1][2] - v[2] - x
			ksy = com[ix-1][2] - v[2] - y
		} else {
			x = v[0] - incom[0]
			y = v[1] - incom[1]
			ksx = incom[2] - v[2] - x
			ksy = incom[2] - v[2] - y
		}
		if ix%1000 == 0 {
			log.Println(ix, "   thousand    ",guid,"offset")
		}
		for ii := guid; ii < len(pos); ii++ {
			//	if v[2] > 0 {

			//if (pos[ii].a >= 0 && pos[ii].a <= v[2]) && (pos[ii].b >= 0 && pos[ii].b <= v[2]) {

			sd := analyse(pos[ii], x, y, ksx, ksy, v[2], ix)
			//	fmt.Println(pos[ii])
			if sd.a != 0 {
				pos[ii] = sd

				//a, b := tranf(di					s, ang, ring, vx, vy)
				//
				//fmt.Println(pos)
				//fmt.Println(x, y, ksx, ksy)
				//pos[ii].a = a
				//pos[ii].b = b

			} else {

				sortstruct(pos).Swap(guid, ii)

				guid++

			}
		}

		//guid = 0

		//	}
	}
	sort.Sort(sortstruct(pos))
	for i, v := range pos {
		var a, b uint64

		if v.com == -1 {
			a, b = tranf(v, incom)
		} else {
			a, b = tranf(v, com[v.com])
		}
		//	fmt.Println(pos)
		result[i] = append(result[i], a)
		result[i] = append(result[i], b)
		//fmt.Fprintf(os.Stdout, "result %d , %d , %d\n", v.a, v.b, v.ran)
	}
	return result
}

func main() {
	var knights []uint64
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := uint64(nTemp)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := uint64(sTemp)

	var commands [][]uint64
	for commandsRowItr := 0; commandsRowItr < int(s); commandsRowItr++ {
		commandsRowTemp := strings.Split(readLine(reader), " ")

		var commandsRow []uint64
		for _, commandsRowItem := range commandsRowTemp {
			commandsItemTemp, err := strconv.ParseInt(commandsRowItem, 10, 64)
			checkError(err)
			commandsItem := uint64(commandsItemTemp)
			commandsRow = append(commandsRow, commandsItem)
		}

		if len(commandsRow) != int(3) {
			panic("Bad input")
		}

		commands = append(commands, commandsRow)
	}
	vTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	v := uint64(vTemp)
	for i := uint64(0); i < v; i++ {

		knightstemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		knights = append(knights, uint64(knightstemp))

	}
	//
	start := time.Now()
	result := kingRichardKnights(n, s, knights, commands)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed.String())
	//

	for resultRowItr, rowItem := range result {
		for resultColumnItr, colItem := range rowItem {
			fmt.Fprintf(writer, "%d", colItem)

			if resultColumnItr != len(rowItem)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		if resultRowItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
