package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(1997, time.May, 18, 0, 0, 0, 0, time.UTC)
	fmt.Println(t)

	fmt.Println("**************")

	now := time.Now()
	fmt.Println("Now: ", now)

	fmt.Println("Yıl: ", now.Year())
	fmt.Println("Ay: ", now.Month())
	fmt.Println("Gün: ", now.Day())
	fmt.Println("Haftanın Günü: ", now.Weekday())

	fmt.Println("**************")

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Yarın: %v %v %v %v \n", tomorrow.Weekday(), tomorrow.Day(), tomorrow.Month(), tomorrow.Year())

	fmt.Println("**************")

	longFormat := "Saturday, January 2, 2006"
	fmt.Println("Yarın: ", tomorrow.Format(longFormat))

	fmt.Println("\n**************")
	fmt.Println(tomorrow.Format("2006-01-02 15:04:05"))

	fmt.Println("\n**************")
	fmt.Println(tomorrow.Format("02-01-2006"))

	fmt.Println("\n**************")
	x := fmt.Println
	x(now.Before(now))
	x(now.After(t))
	x(now.Equal(now))

	fmt.Println("\n**************")
	diff := now.Sub(t)
	x(diff)
	x(diff.Hours())

	/*fmt.Printf("Şu anki Unix Zamanı: %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("Şu anki Unix Zamanı: %v\n", time.Now().Unix())*/
}
