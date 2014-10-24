package helpers

func EmptyStrings(all ...string) bool {
	for i := 0; i < len(all); i++ {
		if all[i] == "" {
			return true
		}
	}
	return false
}
