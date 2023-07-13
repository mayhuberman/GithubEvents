package models

import "strconv"

type UrlStars struct {
	Url   string
	Stars int
}

type UrlStarsSlice []UrlStars

func (urlStars UrlStarsSlice) ConvertUrlStarsToStrings() []string {
	strings := make([]string, len(urlStars))

	for i, us := range urlStars {
		strings[i] = us.Url + ": " + strconv.Itoa(us.Stars)
	}

	return strings
}
