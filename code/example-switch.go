// https://cheatsheet.dennyzhang.com/cheatsheet-golang-A4

	for _, c := range flag {
		switch c {
		case 'i':
			syn |= syntax.FoldCase
		}
		case 'j', 'k':
			syn |= syntax.FoldUnCase
		}
	}

	re, err := compile(pattern, syn, true)
	if err != nil {
