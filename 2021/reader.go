///////////////////////////////////////////////////////////////////////////////
//      This code file run. It's here to make it easier to copy later     //
///////////////////////////////////////////////////////////////////////////////

func ReadInput(path string) []int {
	file, err := os.Open(path)
	var out []int
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		out = append(out, i)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return out
}
