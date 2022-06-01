package Game

import (
	"fmt"
	"math/rand"
	"time"
)

type Pos struct {
	X, Y int
}

type Minesweeper struct {
	Width, Height int
	OpenFields    []Pos
	Mines         []Pos
	Flags         []Pos
	Lost          bool
}

func New(w, h, c int) *Minesweeper {
	m := &Minesweeper{
		Width:  w,
		Height: h,
	}
	m.generateMines(c)

	return m
}

func (m *Minesweeper) generateMines(c int) {
	for i := 0; i < c; i++ {
		m.Mines = append(m.Mines, Pos{
			X: rand.Intn(m.Width),
			Y: rand.Intn(m.Height),
		})
	}
}

func (m *Minesweeper) Open(x, y int) (mineCount int, lose bool, isFlag bool) {
	if m.isFlag(x, y) || m.isOpen(x, y) || m.Lost {
		return -1, false, true
	}

	m.OpenFields = append(m.OpenFields, Pos{x, y})
	if m.isMine(x, y) {
		fmt.Println("You lost!")
		m.Lost = true
		return -1, true, false
	}
	mineCount = m.getNeighbourMines(x, y)

	if mineCount == 0 {
		neighbors := m.getNeighbours(x, y)
		for _, n := range neighbors {
			if !m.isOpen(n.X, n.Y) {
				mineCount, _, _ = m.Open(n.X, n.Y)
			}
		}
	}

	return mineCount, false, false
}

func (m *Minesweeper) ToggleFlag(x, y int) {
	if m.Lost {
		return
	}

	for i, f := range m.Flags { // remove flag
		if f.X == x && f.Y == y {
			m.Flags = append(m.Flags[:i], m.Flags[i+1:]...)
			return
		}
	}

	m.Flags = append(m.Flags, Pos{x, y})

}

func (m *Minesweeper) isMine(x, y int) bool {
	for _, m := range m.Mines {
		if m.X == x && m.Y == y {
			return true
		}
	}
	return false
}

func (m *Minesweeper) isFlag(x, y int) bool {
	for _, f := range m.Flags {
		if f.X == x && f.Y == y {
			return true
		}
	}
	return false
}

func (m *Minesweeper) getNeighbours(x, y int) []Pos {
	neighbors := []Pos{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= m.Width || y+j < 0 || y+j >= m.Height {
				continue
			}
			neighbors = append(neighbors, Pos{x + i, y + j})
		}
	}
	return neighbors
}

func (m *Minesweeper) getNeighbourMines(x, y int) int {
	neighbors := m.getNeighbours(x, y)
	count := 0
	for _, n := range neighbors {
		if m.isMine(n.X, n.Y) {
			count++
		}
	}
	return count
}

func (m *Minesweeper) isOpen(x, y int) bool {
	for _, p := range m.OpenFields {
		if p.X == x && p.Y == y {
			return true
		}
	}
	return false
}

func (m *Minesweeper) Print() string {
	string := ""
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if !m.isOpen(x, y) {
				if m.Lost && m.isMine(x, y) {
					string += "💣  "
				} else if m.isFlag(x, y) {
					string += "🚩  "
				} else {
					string += "🟪  "

				}
			} else if m.isMine(x, y) {
				string += "💣  "
			} else {
				if m.getNeighbourMines(x, y) > 0 {
					string += "" + emoji(m.getNeighbourMines(x, y)) + " "
				} else {
					string += "⬜️  "
				}
			}
		}
		string += "\n "
	}

	return string
}

func emoji(num int) string {
	switch num {
	case 1:
		return "1️⃣"
	case 2:
		return "2️⃣"
	case 3:
		return "3️⃣"
	case 4:
		return "4️⃣"
	case 5:
		return "5️⃣"
	case 6:
		return "6️⃣"
	case 7:
		return "7️⃣"
	case 8:
		return "8️⃣"
	}
	return ""
}

//func main() {
//	fmt.Println("Welcome to Minesweeper!")
//	fmt.Println("You have to open all fields without hitting a mine.")
//	fmt.Println("Use the numbers to open fields and flags to mark fields.")
//	fmt.Println("You have 10 seconds to open a field.")
//	fmt.Println("Good luck!")
//	fmt.Println()
//
//	minesweeper := New(9, 9, 10)
//
//	fmt.Println(minesweeper.Print())
//
//	timer := time.NewTimer(10 * time.Second)
//	go func() {
//		<-timer.C
//		fmt.Println("You lost!")
//		minesweeper.Lost = true
//	}()
//
//	for !minesweeper.Lost {
//		fmt.Print("Enter X and Y: ")
//		var x, y int
//		fmt.Scanf("%d %d", &x, &y)
//		fmt.Println()
//		mineCount, lost, isFlag := minesweeper.Open(x, y)
//		if lost {
//			break
//		}
//		if isFlag {
//			continue
//		}
//		if mineCount > 0 {
//			fmt.Println(emoji(mineCount))
//		}
//		fmt.Println(minesweeper.Print())
//	}
//}

func init() {
	rand.Seed(time.Now().UnixNano())
}
