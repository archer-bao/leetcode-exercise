type automaton struct {
	sign   int
	num    int
	table  map[string][]string
	status string
}

func (a *automaton) constructor() {
	a.num = 0
	a.sign = 1
	a.status = "start"

	// " "(0)	"+/-"(1)	"num"(2)	"other"(3)
	a.table = map[string][]string{
		"start":  {"start", "sign", "number", "end"},
		"sign":   {"end", "end", "number", "end"},
		"number": {"end", "end", "number", "end"},
	}
}

func (a *automaton) getStatusIndex(c byte) int {
	if c == ' ' {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if c-'0' >= 0 && '9'-c <= 9 {
		return 2
	}
	return 3
}

func (a *automaton) get(c byte) {
	a.status = a.table[a.status][a.getStatusIndex(c)]
	if a.status == "number" {
		if a.num > (1<<31-1)/10 {
			a.status = "end"
			if a.sign > 0 {
				a.num = 1<<31 - 1
			} else {
				a.sign = 1
				a.num = -1 << 31
			}
			return
		} else if a.num == (1<<31-1)/10 {
			if int(c-'0') > 7 && a.sign > 0 {
				a.num = 1<<31 - 1
				return
			}
			if int(c-'0') > 8 && a.sign < 0 {
				a.sign = 1
				a.num = -1 << 31
				return
			}
		}

		a.num = a.num*10 + int(c-'0')
	} else if a.status == "sign" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func myAtoi(str string) int {
	a := &automaton{}
	a.constructor()
	for i := 0; i < len(str); i++ {
		if a.status == "end" {
			break
		}
		a.get(str[i])
	}

	return a.sign * a.num
}
