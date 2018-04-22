package main

import (
	"math/rand"
)

type color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

type ColorList struct {
	list []color
}

func NewColorList(num int) *ColorList {
	c := new(ColorList)

	for i := 0; i < num; i++ {
		c.list = append(c.list, color{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))})
	}

	return c
}

func (*ColorList) reverse(yourList ColorList) []color {
	var returnList []color
	for i := len(yourList.list) - 1; i >= 0; i-- {
		returnList = append(returnList, yourList.list[i])
	}
	return returnList
}

func (*ColorList) sortRed(yourList ColorList) []color {
	var returnList []color = yourList.list
	var tempColor color

	for i := 1; i < len(returnList); i++ {
		tempColor = returnList[i]
		j := i - 1
		for j >= 0 && tempColor.r < returnList[j].r {
			returnList[j+1] = returnList[j]
			j = j - 1
		}
		returnList[j+1] = tempColor
	}

	return returnList
}

func (*ColorList) sortGreen(yourList ColorList) []color {
	var returnList []color = yourList.list
	var tempColor color

	for i := 1; i < len(returnList); i++ {
		tempColor = returnList[i]
		j := i - 1
		for j >= 0 && tempColor.g < returnList[j].g {
			returnList[j+1] = returnList[j]
			j = j - 1
		}
		returnList[j+1] = tempColor
	}

	return returnList
}

func (*ColorList) sortBlue(yourList ColorList) []color {
	var returnList []color = yourList.list
	var tempColor color

	for i := 1; i < len(returnList); i++ {
		tempColor = returnList[i]
		j := i - 1
		for j >= 0 && tempColor.b < returnList[j].b {
			returnList[j+1] = returnList[j]
			j = j - 1
		}
		returnList[j+1] = tempColor
	}

	return returnList
}

func (*ColorList) sortAlpha(yourList ColorList) []color {
	var returnList []color = yourList.list
	var tempColor color

	for i := 1; i < len(returnList); i++ {
		tempColor = returnList[i]
		j := i - 1
		for j >= 0 && tempColor.a < returnList[j].a {
			returnList[j+1] = returnList[j]
			j = j - 1
		}
		returnList[j+1] = tempColor
	}

	return returnList
}
