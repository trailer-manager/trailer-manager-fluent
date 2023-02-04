package utility

func NvlString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func NvlString2(s *string, d string) string {
	if s == nil {
		return d
	}

	return *s
}

func NvlInteger(i *int) int {
	if i == nil {
		return 0
	}

	return *i
}

func NvlFloat(f *float64) float64 {
	if f == nil {
		return 0.0
	}

	return *f
}

func NvlBool(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}