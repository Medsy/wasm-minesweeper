package Game

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Pos struct {
	X, Y int
}

type PosList []Pos

type Minesweeper struct {
	Width, Height int
	OpenFields    PosList
	Mines         PosList
	Flags         PosList
	State         int // -1 lost, 0 playing, 1 won
	tileColor     int // maybe just use a slice of tile colors
}

func New(w, h, c int) *Minesweeper {
	// temp random difficulty
	w, h, c = difficulty(rand.Intn(4))

	m := &Minesweeper{
		Width:  w,
		Height: h,
	}
	m.generateMines(c)
	m.tileColor = rand.Intn(9)

	return m
}

func (m *Minesweeper) generateMines(c int) {
	for i := 0; i < c; {
		x := rand.Intn(m.Width)
		y := rand.Intn(m.Height)

		if !m.isMine(x, y) {
			m.Mines = append(m.Mines, Pos{
				X: x,
				Y: y,
			})

			i++
		}
	}
}

func (m *Minesweeper) Open(x, y int) int {
	mineCount := m.neighbourMineCount(x, y)
	flagCount := m.neighbourFlagCount(x, y)
	neighbors := m.getNeighbours(x, y)

	if m.isFlag(x, y) || m.State != 0 {
		return -1
	}

	if m.isOpen(x, y) {
		if mineCount == flagCount {
			for _, n := range neighbors {
				if !m.isFlag(n.X, n.Y) && !m.isOpen(n.X, n.Y) {
					m.Open(n.X, n.Y)
					if m.isMine(n.X, n.Y) {
						m.State = -1
					}
				}
			}
		}
		return -1
	}

	m.OpenFields = append(m.OpenFields, Pos{x, y})
	if m.isMine(x, y) {
		m.State = -1
		return m.State
	}

	if mineCount == 0 {
		for _, n := range neighbors {
			if !m.isOpen(n.X, n.Y) {
				mineCount = m.Open(n.X, n.Y)
			}
		}
	}

	return mineCount
}

func (m *Minesweeper) ToggleFlag(x, y int) {
	if m.State != 0 || m.isOpen(x, y) {
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

func (m *Minesweeper) neighbourMineCount(x, y int) int {
	neighbors := m.getNeighbours(x, y)
	count := 0
	for _, n := range neighbors {
		if m.isMine(n.X, n.Y) {
			count++
		}
	}
	return count
}

func (m *Minesweeper) neighbourFlagCount(x, y int) int {
	neighbors := m.getNeighbours(x, y)
	count := 0
	for _, n := range neighbors {
		if m.isFlag(n.X, n.Y) {
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
	m.checkWin()
	string := ""
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if !m.isOpen(x, y) {
				if m.State == -1 && m.isMine(x, y) {
					string += "💣  "
				} else if m.isFlag(x, y) {
					string += "🚩  "
				} else {
					string += fmt.Sprintf("%s  ", tileColour(m.tileColor))

				}
			} else if m.isMine(x, y) {
				string += "💣  "
			} else {
				if m.neighbourMineCount(x, y) > 0 {
					string += "" + emoji(m.neighbourMineCount(x, y)) + " "
				} else {
					string += "⬜️  "
				}
			}
		}
		string += "\n "
	}

	return string
}

func (m *Minesweeper) checkWin() {
	m.Flags.sort()
	m.Mines.sort()

	if m.Width*m.Height-len(m.Mines) == len(m.OpenFields) {
		if compare(m.Mines, m.Flags) {
			m.State = 1
		}
	}
}

func difficulty(num int) (int, int, int) {
	switch num {
	case 1:
		return 16, 16, 40
	case 2:
		return 30, 19, 99
	}
	return 10, 10, 10
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

func tileColour(num int) string {
	switch num {
	case 1:
		return "🟥"
	case 2:
		return "🟦"
	case 3:
		return "🟧"
	case 4:
		return "🟪"
	case 5:
		return "🟨"
	case 6:
		return "🟩"
	case 7:
		return "❇️"
	case 8:
		return "🔳"
	case 9:
		return "🔲"

	}
	return "⬛️"
}

func (p *PosList) sort() {
	sort.Slice(*p, func(i, j int) bool {
		if (*p)[i].X == (*p)[j].X {
			return (*p)[i].Y < (*p)[j].Y
		}
		return (*p)[i].X < (*p)[j].X
	})
}

func compare(a, b PosList) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
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
