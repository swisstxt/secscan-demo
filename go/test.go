package main

func main() {
	spin := make(chan any, 1)
	go func() {
		for {
			spin <- true
		}
	}()
	for range spin {
	}
}