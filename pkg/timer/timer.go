package timer

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/JoshVanL/timer-cli/pkg/bar"
)

type Timer struct {
	bar     *bar.Bar
	total   int
	current int
}

func New() (t *Timer) {
	return &Timer{}
}

func (t *Timer) ParseArguments(args []string) (err error) {
	for _, arg := range args {
		if seconds, err := strconv.Atoi(arg); err == nil {
			t.total += seconds
			continue

		}

		seconds, err := t.match(arg, "^[0-9]+(.[0-9]+|)s$")
		if err == nil {
			t.total += int(seconds)
			continue

		}

		minutes, err := t.match(arg, "^[0-9]+(.[0-9]+|)m$")
		if err == nil {
			t.total += int(minutes * 60)
			continue

		}

		hours, err := t.match(arg, "^[0-9]+(.[0-9]+|)h$")
		if err == nil {
			t.total += int(hours * 60 * 60)
			continue
		}

		return fmt.Errorf("could not t.match argument: %s", arg)
	}

	t.bar = bar.New(t.total)
	t.current = t.total
	return nil
}

func (t *Timer) match(str, regex string) (num float64, err error) {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(str)
	if len(match) > 0 {
		num, err = t.getNum(match[0])
		if err != nil {
			return -1, err
		}

		return num, nil
	}

	return -1, errors.New("didn't t.match")
}

func (t *Timer) StartTimer() {
	ticker := time.NewTicker(time.Second)
	for t.current > 0 {
		t.Output()
		<-ticker.C
		t.bar.Increment()
		t.current--
	}
	t.Output()
}

func (t *Timer) getNum(str string) (num float64, err error) {
	str = str[:len(str)-1]
	num, err = strconv.ParseFloat(str, 64)
	if err != nil {
		return -1, fmt.Errorf("failed to convert arg to integer: %v", err)
	}

	return num, nil
}

func (t *Timer) GetString() string {
	hour := t.current / 60 / 60
	min := (t.current / 60) % 60
	second := t.current % 60

	var out string
	out += fmt.Sprintf("%s  ", t.bar.GetString())

	if hour > 0 {
		out += fmt.Sprintf("%dh ", hour)
	}
	if hour > 0 || min > 0 {
		out += fmt.Sprintf("%dm ", min)
	}
	out += fmt.Sprintf("%ds", second)

	return out
}

func (t *Timer) Output() {
	t.Flush()
	fmt.Print(t.GetString())
}

func (t *Timer) Flush() {
	fmt.Print("\r\033[K")
}

func (t *Timer) GetTimes() (times string) {
	hours := t.total / 60 / 60
	minutes := (t.total / 60) % 60
	seconds := t.total % 60

	if hours > 0 {
		times += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		times += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 {
		times += fmt.Sprintf("%ds ", seconds)
	}

	return times
}
