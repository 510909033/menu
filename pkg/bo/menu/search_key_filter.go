package menu

type SearchKeyUtil struct{}

func (s *SearchKeyUtil) Filter(query string, list []MenuBO) []MenuBO {
	type One struct {
		IsMatch bool
		Data    string
	}

	l := make([]MenuBO, 0)

	for _, line := range list {
		tmpOne := make([]One, len(line.Title))
		isMatch := false
		for _, v := range query {
			v := string(v)
			for a, b := range line.Title {
				b := string(b)
				if !tmpOne[a].IsMatch {
					tmpOne[a].Data = b
				}

				if !tmpOne[a].IsMatch && v == b {
					tmpOne[a].IsMatch = true
					tmpOne[a].Data = "<em>" + b + "</em>"
					isMatch = true
				}
			} //end for

		} //end for
		if isMatch {
			line.Title = ""
			for _, b := range tmpOne {
				line.Title += b.Data
			}
			l = append(l, line)
		}
	} //end for

	return l
}
