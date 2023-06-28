package user

func TypeToString(uts ...TYPE) (types []string) {
	for _, t := range uts {
		types = append(types, t.String())
	}
	return
}
