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
	a   int32
	b   int32
	n   int32
	ran int
}
type sortstruct []abn

func (a sortstruct) Len() int           { return len(a) }
func (a sortstruct) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortstruct) Less(i, j int) bool { return a[i].ran < a[j].ran }

func tranf(dis, ang, ring, x, y int32) (int32, int32) {
	var a, b int32
	switch ang {
	case 1:
		a = x + dis
		b = y + ring
		break
	case 2:
		a = x + ring
		b = y + ring - dis
		break
	case 3:
		a = x + ring - dis
		b = y
		break

	case 4:
		a = x
		b = y + dis
		break
	}
	return a, b
}
func analyse(k abn, al []int32) (int32, int32, int32, int32, int32) {
	var dis, ang, ring int32
	var v = make([]int32, 3)
	copy(v, al)
	s := v[2] + 1

	for i := int32(0); i < s/2; i++ {
		if k.a == v[0] && k.b <= v[1]+v[2] {
			dis = k.b - v[1]
			ang = 1
			ring = v[2]
			break
		}
		if k.a <= v[0]+v[2] && k.b == v[1]+v[2] {
			dis = k.a - v[0]
			ang = 2
			ring = v[2]
			break
		}
		if k.a == v[0]+v[2] && k.b >= v[1] {
			dis = v[1] + v[2] - k.b
			ang = 3
			ring = v[2]
			break
		}
		if k.a >= v[0] && k.b == v[1] {
			dis = v[0] + v[2] - k.a
			ang = 4
			ring = v[2]
			break
		}
		//fmt.Println("ring ", k, v)
		v[0]++
		v[1]++
		v[2] -= 2
	}

	return dis, ang, ring, v[0], v[1]
}
func kingRichardKnights(n int32, s int32, knights []int32, com [][]int32) [][]int32 {

	pos := make([]abn, len(knights))
	result := make([][]int32, len(knights))
	guid := 0
	for i := 0; i < len(knights); i++ {
		pos[i].a = (knights[i] / n) + 1
		pos[i].b = (knights[i] % n) + 1
		pos[i].n = knights[i]
		pos[i].ran = i
		//	fmt.Fprintf(os.Stdout, "result %d \n", pos[i].ran)
	}
	for _, v := range com {
		for ii := guid; ii < len(pos); ii++ {
			if v[2] > 0 {
				start := time.Now()
				if (pos[ii].a >= v[0] && pos[ii].a <= v[0]+v[2]) && (pos[ii].b >= v[1] && pos[ii].b <= v[1]+v[2]) {
					dis, ang, ring, vx, vy := analyse(pos[ii], v)
					//fmt.Println("ang", dis, ang, ring)
					//	fmt.Println("forw", pos[ii])
					//	fmt.Println("changed", v)

					a, b := tranf(dis, ang, ring, vx, vy)
					//fmt.Println("arg", a, b)
					pos[ii].a = a
					pos[ii].b = b
					elapsed := time.Since(start)
					log.Printf("%d time elapse %s", guid, elapsed)
					///	fmt.Fprintf(os.Stdout, "result %d \n", pos[ii].ran)
					//	fmt.Fprintf(os.Stdout, "resul%d , %d , %d\n", pos[ii].a, pos[ii].b, pos[ii].ran)
				} else {
					//fmt.Fprintf(os.Stdout, "%d , %d \n", guid, ii)
					start := time.Now()
					sortstruct(pos).Swap(guid, ii)
					elapsed := time.Since(start)
					log.Printf("Binomial took %s", elapsed)
					guid++

				}

				guid = 0
			}
		}
	}
	sort.Sort(sortstruct(pos))
	for i, v := range pos {
		result[i] = append(result[i], v.a)
		result[i] = append(result[i], v.b)
		//fmt.Fprintf(os.Stdout, "result %d , %d , %d\n", v.a, v.b, v.ran)
	}
	return result
}

func main() {
	var knights []int32
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)

	//defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	s := int32(sTemp)

	var commands [][]int32
	for commandsRowItr := 0; commandsRowItr < int(s); commandsRowItr++ {
		commandsRowTemp := strings.Split(readLine(reader), " ")

		var commandsRow []int32
		for _, commandsRowItem := range commandsRowTemp {
			commandsItemTemp, err := strconv.ParseInt(commandsRowItem, 10, 64)
			checkError(err)
			commandsItem := int32(commandsItemTemp)
			commandsRow = append(commandsRow, commandsItem)
		}

		if len(commandsRow) != int(3) {
			panic("Bad input")
		}

		commands = append(commands, commandsRow)
	}
	vTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	v := int32(vTemp)
	for i := 0; i < int(v); i++ {

		knightstemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		knights = append(knights, int32(knightstemp))

	}
	result := kingRichardKnights(n, s, knights, commands)

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
