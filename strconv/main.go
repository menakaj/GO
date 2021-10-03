package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Container struct {
	index int
	val   string
}

func main() {
	news := []string{
		// "/pet/asd/findByTags.com",
		// "/pet/as/asd/{findByTags}.com",
		// "/pet/{petId}.com",
		// "/pet/{ack}/*.com",
		// "/user/createWithList",
		// "/store/{order}/{orderid}",
		// "/store/{base}nventory",
		// "/store/order/{orderId}",
		// "/user/login{id}",
		// "/{pet1}",
		// "/pet/asd/{findByTags}.com",
		// "/{ept}/pet/",
		// "/pet/findByTags/",
		// "/pet/findByTags/{anc}.com",
		// "/pet/{petId}/uploadImage",
		// "/user/{user}.id",
		// "/user/logout",
		// "/pet/{username}",
		// "/{pet}/{pet1}/{pwr2}",
		// "/pet",
		// "/store/order/{oderId}.com",
		// "/pet/abc/{qe}/on/*",
		// "/store/order",

		// "/pet/ad/ad/findByTags.com",
		// "/store/{order}/orderId",
		"/pet",
		"/pet/{id}",
		"/pet/index.html",
		"/pet/{id}/price?abc={abc}&cde={cde}",
		"/pet/{id}/{price}",
		"/pet/*",
		"/pet/{petId}.com",
		"/pet/pet.{anc}",
		"/pet/{pet}.{anc}",
	}
	for i := range news {
		fmt.Println(generateRegex(news[i]))
	}
	

}

type ByVal []Container

func (a ByVal) Len() int { return len(a) }
func (a ByVal) Less(i, j int) bool {
	// Get the less weighted path.
	// Paths can be in several types.
	// - /pet
	// - /pet/{id}
	// - /pet/index.html
	// - /pet/{id}/price
	// - /pet/*
	// When representing these resources in envoy configuration, they must be ordered correctly.
	// Considerations...
	// The concrete paths are matched first
	// Any path with . character gets higher precidence
	// Precedence decreases when the number of path parameters increses.
	// The wild card path is matched last.

	// Replace all the non symbol characters with empty string ("") Because the alphabatical order is not mandetory.
	// - /pet  				=> /
	// - /pet/{id}			=> //{}
	// - /pet/index.html	=> //.
	// - /pet/{id}/price	=> //{}/
	// - /pet/*				=> //*
	charMatcher := regexp.MustCompile(`[\w\s]`)
	valueI := charMatcher.ReplaceAllString(a[i].val, "")
	valueJ := charMatcher.ReplaceAllString(a[j].val, "")

	// if wildcard is matched for either i or j, it will be returnd as greater.
	wildCardMatcher := regexp.MustCompile(`(\/[*]$)`)
	if wildCardMatcher.Match([]byte(valueI)) || wildCardMatcher.Match([]byte(valueJ)) {
		return !wildCardMatcher.Match([]byte(valueI)) || wildCardMatcher.Match([]byte(valueJ))
	}

	// if the dot is matched (either i or j), the path is considered less than the other one.
	// If both i and j match this at the same time, compare the full path.
	dotMatcher := regexp.MustCompile(`\.`)
	if dotMatcher.Match([]byte(valueI)) && dotMatcher.Match([]byte(valueJ)) {
		return valueI < valueJ
	} else if dotMatcher.Match([]byte(valueI)) {
		return true
	} else if dotMatcher.Match([]byte(valueJ)) {
		return false
	}
	// If non of the above matched, compare the strings.
	return valueI < valueJ
}
func (a ByVal) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func check(q []string) {

	m1 := regexp.MustCompile(`[\w\s]`)
	// m2 := regexp.MustCompile(`([{w+}][.])`)
	var arr []Container
	for x := range q {
		// fmt.Println("is matching .", q[x], m2.Match([]byte(q[x])))
		s := m1.ReplaceAllString(q[x], "")
		count := 0
		c := Container{x, q[x]}
		arr = append(arr, c)
		for i, c := range s {
			if string(c) == "{" || string(c) == "}" {
				count += i + 1
			}
		}
		// fmt.Println(q[x], s, count)
	}
	for a := range q {
		fmt.Println(q[a])
	}
	sort.Sort(ByVal(arr))
	fmt.Println("Sorted>>>>>>>>>>>>>>>>>\n\n")
	for a := range arr {
		fmt.Println(arr[a].val)
	}

}

func generateRegex(fullpath string) string {
	
	pathParaRegex := "([^/]+)"
	wildCardRegex := "((/(.*))*)"
	endRegex := "(\\?([^/]+))?"
	newPath := ""

	m1 := regexp.MustCompile(`{([^}]+)}`)
	a := m1.ReplaceAllString(fullpath, pathParaRegex)
	fmt.Println("Replaced ---- > ",fullpath, " -- ",a)

	if strings.Contains(fullpath, "{") && strings.Contains(fullpath, "}") {
		res1 := strings.Split(fullpath, "/")

		for i, p := range res1 {
			if strings.Contains(p, "{") && strings.Contains(p, "}") {
				startP := strings.Index(p, "{")
				endP := strings.Index(p, "}")
				res1[i] = p[:startP] + pathParaRegex + p[endP+1:]
			}
		}
		newPath = strings.Join(res1[:], "/")

	} else {
		newPath = fullpath
	}

	if strings.HasSuffix(newPath, "/*") {
		newPath = strings.TrimSuffix(newPath, "/*") + wildCardRegex
	}
	fmt.Println("New: ", "^"+a+endRegex+"$")
	fmt.Println("Old: ","^" + newPath + endRegex + "$")


	return "^" + newPath + endRegex + "$"
}
