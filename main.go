package main

import "./liquidmarket"

func main() {
	a := liquidmarket.App{}

	a.Initialize("root", "podsaveamerica", "enhanced-emblem-188503:australia-southeast1:liquidmarket", "liquidmarket", true)

	a.RunLocal(":8080")
}
