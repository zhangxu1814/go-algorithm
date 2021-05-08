package crawler_log_folder

func minOperations(logs []string) int {
	step := 0
	for _, log := range logs {
		switch log {
		case "../":
			if step > 0 {
				step--
			}
		case "./":
		default:
			step++
		}
	}
	return step
}
