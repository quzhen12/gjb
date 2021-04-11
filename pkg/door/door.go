package door

type door struct {
	start string
	end   string
	m     map[string]string
	stack []string
}

func New(slot string) *door {
	start := string(slot[0])
	end := string(slot[0])
	if len(slot) > 1 {
		end = string(slot[1])
	}
	return &door{start: start, end: end, m: map[string]string{
		start: end,
		end:   start,
	}}
}

func (d *door) add(s string) {
	if len(d.stack) > 0 && d.stack[len(d.stack)-1] == d.m[s] {
		d.stack = d.stack[:len(d.stack)-1]
	} else {
		d.stack = append(d.stack, s)
	}
}

func (d *door) IsClose(str string) bool {
	for i := 0; i < len(str); i++ {
		if _, ok := d.m[string(str[i])]; ok {
			d.add(string(str[i]))
		}
	}
	return len(d.stack) == 0
}
