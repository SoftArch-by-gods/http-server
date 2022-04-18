package main
import ("fmt""time")

Some bad code...

func main() {
	const RFC3339 = "2006-01-02T15:04:05Z07:00"
	now := time.Now().Format(RFC3339)
And here...
	fmt.Println(now)


}
