package main

func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, v := range strs {
		t := [26]int{}
		for _, vv := range v {
			t[vv-'a']++
		}
		mp[t] = append(mp[t], v)
	}
	res:=make([][]string,0,len(mp))
	for _,vv:=range mp{
		res = append(res, vv)
	}
	return res
}
