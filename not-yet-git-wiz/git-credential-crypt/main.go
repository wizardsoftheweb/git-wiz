package git_credentials_store

func main() {
	// CollectGitConfig()
	panic(Execute())
	// ParseSearchInput("")
	// store := NewStoreFromDisk("~/.git-credentials")
	// store.FileName = "~/.git-credentials2"
	// store.Write()
	// fmt.Println(store)
	// for _, site := range store.Sites {
	// 	fmt.Println(site)
	// }
	// site := store.Get(map[string]string{"host": "api.*.com"})
	// fmt.Println(site)
}

func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
