package Chapter01

func toLowerCase(s string) string {
	result := make([]rune, len(s)) //int32 的别名，用于处理 Unicode 字符
	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			result[i] = char + 32
		} else {
			result[i] = char
		}
	}
	return string(result)
}
