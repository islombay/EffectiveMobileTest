package adduser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAge(name, ageUrl string) (int, error) {
	res, err := http.Get(fmt.Sprintf("%s/?name=%s", ageUrl, name))
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	resBodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var resObj ResponseAgeService

	if err := json.Unmarshal(resBodyByte, &resObj); err != nil {
		return 0, err
	}
	return resObj.Age, nil
}

func GetGender(name, genderUrl string) (string, error) {
	res, err := http.Get(fmt.Sprintf("%s/?name=%s", genderUrl, name))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	resBodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var resObj ResponseGenderService

	if err := json.Unmarshal(resBodyByte, &resObj); err != nil {
		return "", err
	}
	return resObj.Gender, nil
}

func GetNationality(name, natUrl string) (ResponseNationalityCountryService, error) {
	var maxVal ResponseNationalityCountryService
	res, err := http.Get(fmt.Sprintf("%s/?name=%s", natUrl, name))
	if err != nil {
		return maxVal, err
	}
	defer res.Body.Close()

	resBodyByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return maxVal, err
	}

	var resObj ResponseNationalityService

	if err := json.Unmarshal(resBodyByte, &resObj); err != nil {
		return maxVal, err
	}

	if len(resObj.Country) > 0 {
		maxVal = resObj.Country[0]
	}
	for _, obj := range resObj.Country {
		if obj.Probability > maxVal.Probability {
			maxVal = obj
		}
	}
	return maxVal, nil
}
