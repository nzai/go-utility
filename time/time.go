package time

import st "time"

func BeginOfDay(_time st.Time) st.Time {
	year, month, day := _time.Date()
	return st.Date(year, month, day, 0, 0, 0, 0, _time.Location())
}
