package tui

import (
	"fmt"
	"strconv"
	"strings"
)

type MenuAction interface {
	Run() int
}

type returnCodes struct {
	ReturnToMain int
	Back         int
	Quit         int
}
type menu struct {
	opts            map[string]MenuAction
	orderedOpts     []string
	header          string
	title           string
	menuIndent      string
	StatusCodes     returnCodes
	currentKeyAlpha string
	keyType         string
	currentKeyNum   int
}

func NewMenu(title string) *menu {
	opts := make(map[string]MenuAction)
	opts["q"] = &menu{}
	return &menu{
		title:           title,
		opts:            opts,
		orderedOpts:     make([]string, 0),
		header:          "",
		menuIndent:      "\t\t",
		StatusCodes:     returnCodes{ReturnToMain: 9, Back: 0, Quit: 99},
		currentKeyAlpha: "A",
		currentKeyNum:   1,
		//Whether to use letters or numbers for menu keys
		keyType: "Alpha",
	}
}

func (m *menu) SetHeader(text string) {
	m.header = text
}

func (m *menu) SetMenuIndent(s string) {
	bits := []byte(s)
	for i := 0; i < len(bits); i++ {
		if bits[i] != 32 && bits[i] != 9 {
			bits[i] = 32
		}
	}
	m.menuIndent = string(bits)
}

func (m *menu) printFormattedLine(s string) {
	fmt.Printf("%s%s\n", m.menuIndent, s)
}

func (m *menu) AddOption(flag, option string, next MenuAction) {
	m.opts[strings.ToLower(flag)] = next
	s := fmt.Sprintf("%s: %s", flag, option)
	m.orderedOpts = append(m.orderedOpts, s)
}

func (m *menu) AddSubOption(flag, option string, next MenuAction) {
	m.opts[strings.ToLower(flag)] = next
	s := fmt.Sprintf(" [+] %s: %s", flag, option)
	m.orderedOpts = append(m.orderedOpts, s)
}

func (m *menu) AddBackOption() {
	m.opts["r"] = &menu{}
	s := fmt.Sprintf("%s: %s", "R", "Back")
	m.orderedOpts = append(m.orderedOpts, s)
}

func (m *menu) AddBlank() {
	m.orderedOpts = append(m.orderedOpts, "")
}

func (m *menu) listen() int {
	for {
		choice := ""
		for ok := false; !ok; _, ok = m.opts[strings.ToLower(choice)] {
			fmt.Printf("%sSelection: ", m.menuIndent)
			fmt.Scanln(&choice)
			fmt.Print("\u001b[1A\u001b[2K") //reset line
		}
		if strings.ToLower(choice) == "r" {
			return m.StatusCodes.Back
		}
		if strings.ToLower(choice) == "rr" {
			return m.StatusCodes.ReturnToMain
		}
		if strings.ToLower(choice) == "q" {
			return m.StatusCodes.Quit
		}
		// for i := 0; i < len(m.orderedOpts)+3; i++ {
		// 	fmt.Print("\u001b[1A\u001b[2K") //reset line
		// }
		fmt.Println("\u001b[H\u001b[2J")
		fmt.Println(m.header)

		ret_code := m.opts[strings.ToLower(choice)].Run()
		if ret_code == m.StatusCodes.ReturnToMain {
			return m.StatusCodes.ReturnToMain
		}
		if ret_code == m.StatusCodes.Quit {
			return m.StatusCodes.Quit
		}
		if ret_code == m.StatusCodes.Back {
			return 1
		}

	}
}

func (m *menu) Run() int {
	for {
		fmt.Print("\u001b[H\u001b[2J")
		fmt.Println(m.header)
		fmt.Print("\n\n")
		t := fmt.Sprintf(">>> %s <<<\n\n", m.title)
		m.printFormattedLine(t)
		for _, key := range m.orderedOpts {
			m.printFormattedLine(key)
		}

		fmt.Print("\n")
		m.printFormattedLine("Q: Quit\n")
		ret_code := m.listen()
		if ret_code == m.StatusCodes.Back {
			return m.StatusCodes.Back
		}
		if ret_code == m.StatusCodes.ReturnToMain {
			return m.StatusCodes.ReturnToMain
		}
		if ret_code == m.StatusCodes.Quit {
			return m.StatusCodes.Quit
		}
	}

}

type menuOption struct {
	current    int
	additional int
	optionType string
}

// Return the next number or Letter in the sequence
func (m *menuOption) Next() string {
	m.current++
	if m.optionType == "number" {
		return strconv.Itoa(m.current)
	}
	// Convert the int to a byte so that string(b) will return a Letter
	b := byte(m.current)
	response := string(b)

	if m.additional > 0 {
		response += strconv.Itoa(m.additional)
	}

	// If Letter == Z (90) set it back to the byte before A (65)
	// Increment the cycle by 1
	if m.current == 90 {
		m.current = 64
		m.additional++
	}
	return response
}

func NewMenuOptionManager(optionType string) menuOption {
	m := menuOption{optionType: optionType, current: 0}
	if optionType == "letter" {
		m.current = 64
	}
	return m
}
