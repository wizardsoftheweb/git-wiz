package main

import "fmt"

func main() {
	// panic(Execute())

	store := NewStoreFromDisk("~/.git-credentials")
	fmt.Println(store)
	for _, site := range store.Sites {
		fmt.Println(site)
	}
	site := store.Get(map[string]string{"host": "api.*.com"})
	fmt.Println(site)
}
