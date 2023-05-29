package offer

// 日程表
type MyCalendar struct {
	Time []*Time
}

type Time struct {
	Start int
	End   int
}

func Constructor10() MyCalendar {
	return MyCalendar{
		Time: []*Time{},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for _, t := range this.Time {
		if start <= t.End || end > t.Start {
			return false
		}
	}

	t := &Time{
		Start: start,
		End:   end,
	}
	this.Time = append(this.Time, t)
	return true
}
