package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
 * Complete the kingRichardKnights function below.
 */
type abn struct {
	a uint64
	b uint64
	//n     uint64
	//ran   int
	angle int
	com   int
}

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
func analyse(k abn, com []uint64, s int, c uint64) (bool, uint64, uint64) {
	//	var dis, ang, ring uint64
	//k.a -= s[i]
	//k.b -= s[u]
	var a, b uint64
	a = k.a - com[0]
	b = k.b - com[1]

	//	fmt.Println("qes ", a, b)
	if a <= 0 || b <= 0 || a > c+1 || b > c+1 {
		return true, a, b
	}

	return false, a, b
}
func kingRichardKnights(n uint64, s uint64, knights []uint64, com [][]uint64) [][]uint64 {

	pos := make([]abn, len(knights))
	result := make([][]uint64, len(knights))
	tcom := make([][]uint64, len(com))
	incom := []uint64{0, 0}
	//ind := [][]int{{0, 1}, {3, 0}, {2, 3}, {1, 2}}
	//angle := 0
	for i := 0; i < len(knights); i++ {
		pos[i].a = (knights[i] / n) + 1
		pos[i].b = (knights[i] % n) + 1
		//pos[i].n = knights[i]
		//pos[i].ran = i
		pos[i].com = -1
		pos[i].angle = 0

		//	fmt.Fprintf(os.Stdout, "result %d \n", pos[i].ran)
	}

	for i := 0; i < len(com); i++ {
		var x, y, ksx, ksy, res1, res2 uint64
		//	f := []uint64{x, y, ksx, ksy}
		if i != 0 {
			x = com[i][0] - com[i-1][0]
			y = com[i][1] - com[i-1][1]
			ksx = com[i-1][2] - com[i][2] - x
			ksy = com[i-1][2] - com[i][2] - y //com[i][2] - y
			switch i % 4 {
			case 0:
				res1 = tcom[i-1][0] + x
				res2 = tcom[i-1][1] + y
			case 1:
				res1 = tcom[i-1][0] + ksy
				res2 = tcom[i-1][1] + x
			case 2:
				res1 = tcom[i-1][0] + ksx
				res2 = tcom[i-1][1] + ksy
			case 3:
				res1 = tcom[i-1][0] + y
				res2 = tcom[i-1][1] + ksx
			}
		} else {
			x = com[i][0] - 1
			y = com[i][1] - 1
			ksx = n - com[i][2] - x
			ksy = n - com[i][2] - y
			//	f = []uint64{x, y, ksx, ksy}
			res1 = x
			res2 = y
		}
		tcom[i] = append(tcom[i], res1)
		tcom[i] = append(tcom[i], res2)

	}
	//	sk := time.Now()
	//ee := time.Since(sk)
	//ee = 0
	for ii := 0; ii < len(pos); ii++ {
		//fmt.Println(tcom)
		//angle =
		gr := len(com) - 1
		sm := 0
		sgg := 0
		//	var dc, cc uint64
		fmt.Println(tcom)
		//var kf, gf uint64
		for g := 0; g < 4; g++ {
			//	start := time.Now()
			ss := (gr + sm) / 2
			sd, _, _ := analyse(pos[ii], tcom[ss], ss, com[ss][2])
			//fmt.Println(ii, sd, ss)
			//	sd := true
			if sd == true {

				//fmt.Println(kf, gf)
				gr = ss
				sgg = ss

			} else {
				sm = ss

			}

			/*
						fmt.Println(gr, sm)
						fmt.Println(ss)
						if gr-sm == 1 || gr-sm == -1 {
							sd := analyse(pos[ii], com[ss-1])
				var kf, gf uint64
				for g := 0; g < 4; g++ {
					//	start := time.Now()
							fmt.Println(sd)
							if sd == true {
								elem = []int{ii, ss - 1}
								break
							} else {
								break
							}
						}
			*/
		}
		//	fmt.Println(sgg-1, tcom[sgg-1], (sgg-1)%4)
		if sgg == 0 {
			_, pos[ii].a, pos[ii].b = analyse(pos[ii], incom, sgg-1, n)

		} else {
			_, pos[ii].a, pos[ii].b = analyse(pos[ii], tcom[sgg-1], sgg-1, com[sgg-1][2])
		}
		//	fmt.Println(pos[ii])
		//fmt.Println(elem, (elem[1]%4)+1)
		//	elapsed := time.Since(start)
		//	ee += elapsed
		//if ii%1000 == 0 {
		//	log.Println(" Binomial took ", ee.String())

		//	fmt.Println(pos[ii]

		var a, b uint64
		pos[ii].angle = ((sgg - 1) % 4)
		fmt.Println(pos[ii], sgg-1)
		//	fmt.Println("index ", elem[0], "come ", elem[1], "angle ", elem[1]%4)
		if sgg == 0 {

			a, b = tranf(pos[ii], incom)
		} else {
			a, b = tranf(pos[ii], com[sgg])
		}
		//	fmt.Println(pos)

		result[ii] = append(result[ii], a)
		result[ii] = append(result[ii], b)

		//sortstruct(pos).Swap(guid, ii)

	}

	//	}

	//	sort.Sort(sortstruct(pos))
	//	for = range pos {
	//		 uint64

	//fmt.Fprintf(os.Stdout, "result %d , %d , %d\n", v.a, v.b, v.ran)

	//}

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
