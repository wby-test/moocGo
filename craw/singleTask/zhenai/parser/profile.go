package parser

import (
	"moocGo/craw/model"
	"moocGo/craw/singleTask/engine"
	"regexp"
	"strconv"
)



var nameRe = regexp.MustCompile(`"nickname":"([^,]+)",`)
var genderRe = regexp.MustCompile(`"genderString":"([^,]+)",`)
var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^>]+)岁</div>`)
var heightRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^>]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^>]+)kg</div>`)
var incomeRe  = regexp.MustCompile(`"salaryString":"([^,]+)",`)
var marriageRe = regexp.MustCompile(`"marriageString":"([^,]+)",`)
var educationRe = regexp.MustCompile(`"educationString":"([^,]+)",`)
var occupationRe = regexp.MustCompile(``)
var hokouRe	 = regexp.MustCompile(`"objectWorkCityString":"([^,]+)",`)
var xingzuoRe	= regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([^>]+座)(.*)</div>`)
var signRe	= regexp.MustCompile(`"introduceContent":"([^,]+)",`)
var houseRe	= regexp.MustCompile(``)
var carRe	= regexp.MustCompile(`<div class="m-btn pink" data-v-bff6f798>([^>]+车[^<]*)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err == nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err == nil {
		profile.Height = height
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err == nil {
		profile.Weight = weight
	}

	profile.Name = "Name:  " + name + "\n"
	profile.Gender = "Gender:   " +  extractString(contents, genderRe) + "\n"
	profile.Income = "Income:   " + extractString(contents, incomeRe) + "\n"
	profile.Marriage = "Marriage:   " + extractString(contents, marriageRe) + "\n"
	profile.Education = "Education:   " + extractString(contents, educationRe) + "\n"
    //profile.Occupation = extractString(contents, occupationRe) + "\n"
	profile.Hokou = "Hokou:   " + extractString(contents, hokouRe) + "\n"
	profile.Xingzuo = "Xingzuo:   " + extractString(contents, xingzuoRe) + "\n"
	profile.Sign = "Sign:   " + extractString(contents, signRe) + "\n"
	//profile.House = extractString(contents, hokouRe) + "\n"
	profile.Car = "Car:   " + extractString(contents, carRe)  + "\n"

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if match != nil {
		return string(match[1])
	} else {
		return ""
	}
}