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
func analyse(k abn, s []uint64, i, u int, com uint64) (uint64, uint64) {
	//	var dis, ang, ring uint64
	k.a -= s[i]
	k.b -= s[u]

	//	println(k.a, k.b, com)
	if k.a <= 0 || k.b <= 0 || k.a > com+1 || k.b > com+1 {
		return 0, 0
	}

	return k.a, k.b
}
func kingRichardKnights(n uint64, s uint64, knights []uint64, com [][]uint64) [][]uint64 {

	pos := make([]abn, len(knights))
	result := make([][]uint64, len(knights))
	tcom := make([][]uint64, len(com))
	incom := []uint64{1, 1, n - 1}
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
		var x, y, ksx, ksy uint64
		if i != 0 {
			x = com[i][0] - com[i-1][0]
			y = com[i][1] - com[i-1][1]
			ksx = com[i-1][2] - com[i][2] - x
			ksy = com[i-1][2] - com[i][2] - y
		} else {
			x = com[i][0] - incom[0]
			y = com[i][1] - incom[1]
			ksx = incom[2] - com[i][2] - x
			ksy = incom[2] - com[i][2] - y
		}
		f := []uint64{x, y, ksx, ksy}

		tcom[i] = append(tcom[i], f...)

	}
	//	sk := time.Now()
	//ee := time.Since(sk)
	//ee = 0
	for ii := 0; ii < len(tcom); ii++ {

		//	angle = 0
		fmt.Println(tcom[ii])
		/*
				for i, v := range tcom {

					//

					//	if v[2] > 0 {com

					//if (pos[ii].a >= 0 && pos[ii].a <= v[2]) && (pos[ii].b >= 0 && pos[ii].b <= v[2]) {
					start := time.Now()

					sd, sb := analyse(pos[ii], v, ind[angle][0], ind[angle][1], com[i][2])
					elapsed := time.Since(start)
					ee += elapsed
					if ii%1000 == 0 {
						log.Println(" Binomial took ", ee.String())
					}
					//	fmt.Println(pos[ii]
					if sd != 0 && len(com)-i > 1 {

						pos[ii].a = sd
						pos[ii].b = sb
						angle = (angle + 1) % 4
						pos[ii].com++
						continue

						//a, b := tranf(di					s, ang, ring, vx, vy)
						//
						//fmt.Println(pos)
						//fmt.Println(x, y, ksx, ksy)
						//pos[ii].a = a
						//pos[ii].b = b

					}
					if len(com)-i == 1 {
						pos[ii].a = sd
						pos[ii].b = sb
						angle = (angle + 1) % 4

						pos[ii].com++
					}
					var a, b uint64
					pos[ii].angle = angle
					if pos[ii].com == -1 {
						a, b = tranf(pos[ii], incom)
					} else {
						a, b = tranf(pos[ii], com[pos[ii].com])
					}
					//	fmt.Println(pos)

					result[ii] = append(result[ii], a)
					result[ii] = append(result[ii], b)
					break
					//sortstruct(pos).Swap(guid, ii)

				}

			}
			//	}

			//	sort.Sort(sortstruct(pos))
			//	for = range pos {
			//		 uint64

			//fmt.Fprintf(os.Stdout, "result %d , %d , %d\n", v.a, v.b, v.ran)
		*/
		//}
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
