/*
 * --------------------------------------------------------------------------------------------------------------------
 * <copyright company="Aspose">
 *   Copyright (c) 2018 Aspose.Slides for Cloud
 * </copyright>
 * <summary>
 *   Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in all
 *  copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *  SOFTWARE.
 * </summary>
 * --------------------------------------------------------------------------------------------------------------------
 */

package asposeslidescloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
)

var (
	testApiClient           *APIClient
	testAsyncApiClient           *APIClient
	testFolderName          = "TempSlidesSDK"
	testFileName            = "test.pptx"
	unprotectedTestFileName = "test-unprotected.pptx"
	changedTestFileName     = "changedtest.pptx"
	testTemplateFileName    = "TemplateCV.pptx"
	testFilePassword        = "password"
	isInitialized           = false
	expectedTestDataVersion = "1"
        testOperationId         = ""
)

func GetOperationId() string {
	if testOperationId == "" {
		c := GetTestSlidesAsyncApiClient()
		source, e := ioutil.ReadFile("TestData/" + testFileName)
		if e != nil {
			return ""
		}
		operationId, _, e := c.SlidesAsyncApi.StartConvert(source, "pdf", testFilePassword, "", "", nil, nil)
		if e != nil {
			return ""
		}
		testOperationId = operationId
		time.Sleep(time.Duration(10) * time.Second)
	}
	return testOperationId
}

func InitializeTest(functionName string, invalidParamName string, invalidParamValue interface{}) error {
	if !isInitialized {
		version := "0"
		c := GetTestSlidesApiClient()
		f, _, e := c.SlidesApi.DownloadFile("TempTests/version.txt", "", "")
		if e == nil {
			data, e := ioutil.ReadFile(f.Name())
			if e == nil {
				version = string(data)
			}
		}
		if version != expectedTestDataVersion {
			files, e := ioutil.ReadDir("TestData")
			if e != nil {
				return e
			}
			for _, file := range files {
				bytes, e := ioutil.ReadFile("TestData/" + file.Name())
				if e != nil {
					return e
				}
				_, _, e = c.SlidesApi.UploadFile("TempTests/"+file.Name(), bytes, "")
				if e != nil {
					return e
				}
			}
			_, _, e = c.SlidesApi.UploadFile("TempTests/version.txt", []byte(expectedTestDataVersion), "")
			if e != nil {
				return e
			}
		}
		isInitialized = true
	}
	files := make(map[string]FileRule)
	for _, rule := range getRules(getTestRules().Files, functionName, invalidParamName, "") {
		fileRule := rule.(FileRule)
		actualName := (untemplatize(fileRule.File, invalidParamValue, invalidParamName)).(string)
		path := "TempSlidesSDK"
		if fileRule.Folder != "" {
			path = untemplatize(fileRule.Folder, invalidParamValue, invalidParamName).(string)
		}
		path = path + "/" + actualName
		fileRule.ActualName = actualName
		files[path] = fileRule
	}

	for path, rule := range files {
		if rule.Action == "Put" {
			c := GetTestSlidesApiClient()
			_, e := c.SlidesApi.CopyFile("TempTests/"+rule.ActualName, path, "", "", "")
			if e != nil {
				return e
			}
		} else if rule.Action == "Delete" {
			c := GetTestSlidesApiClient()
			_, e := c.SlidesApi.DeleteFile(path, "", "")
			if e != nil {
				return e
			}
			_, e = c.SlidesApi.DeleteFolder(path, "", nil)
			if e != nil {
				return e
			}
		}
	}
	return nil
}

var testRules *TestRules

func getTestRules() *TestRules {
	if testRules == nil {
		testRules = &TestRules{}
		file, err := os.Open("testRules.json")
		if err == nil {
			json.NewDecoder(file).Decode(testRules)
		}
	}
	return testRules
}

func GetTestSlidesApiClient() *APIClient {
	if testApiClient == nil {
		cfg := NewConfiguration()
		configFile, err := os.Open("testConfig.json")
		if err == nil {
			json.NewDecoder(configFile).Decode(cfg)
		}
		testApiClient = NewAPIClient(cfg)
	}
	return testApiClient
}

func GetTestSlidesAsyncApiClient() *APIClient {
	if testAsyncApiClient == nil {
		cfg := NewConfiguration()
		configFile, err := os.Open("testConfig.json")
		if err == nil {
			json.NewDecoder(configFile).Decode(cfg)
		}
		cfg.BasePath = cfg.AsyncBasePath
		testAsyncApiClient = NewAPIClient(cfg)
	}
	return testAsyncApiClient
}

func createTestParamValue(functionName string, paramName string, paramType string) interface{} {
	var value interface{}
	value = nil
	ruleType := paramType
	for _, rule := range getRules(getTestRules().Values, functionName, paramName, paramType) {
		valueRule := rule.(ValueRule)
		if valueRule.ValueSet {
			value = valueRule.Value
			if valueRule.Type != "" {
				ruleType = valueRule.Type
			}
		}
	}
	return untemplatize(undefaultize(value, paramType, ruleType), nil, paramName)
}

func undefaultize(value interface{}, paramType string, ruleType string) interface{} {
	if paramType == "int32" {
                v, ok := value.(float64)
                if !ok {
                        return 0
                }
		return int32(v)
	}
	if paramType == "[]int32" {
		var arr = []int32{}
		b, _ := json.Marshal(value)
		json.Unmarshal(b, &arr)
		return arr
	}
	if isModelClass(paramType) {
		valueMap, ok := value.(map[string]interface{})
		if ok {
			if !isModelClass(ruleType) {
				ruleType = paramType
			}
			valueObj := createObjectForMap(ruleType, ruleType, valueMap)
			b, _ := json.Marshal(valueMap)
			json.Unmarshal(b, &valueObj)
			return valueObj
		}
	}
	return value
}

func getRules(rules interface{}, functionName string, paramName string, paramType string) []ITestRule {
	filteredRules := []ITestRule{}
	ruleArray := reflect.ValueOf(rules)
	for i := 0; i < ruleArray.Len(); i++ {
		rule := ruleArray.Index(i).Interface().(ITestRule)
		if ruleApplies(rule, functionName, paramName, paramType) {
			filteredRules = append(filteredRules, rule)
		}
	}
	return filteredRules
}

func ruleApplies(rule ITestRule, functionName string, paramName string, paramType string) bool {
	return ruleConditionApplies(rule.getMethod(), functionName) &&
		(rule.getInvalid() == nil || (*(rule.getInvalid()) && (paramName != ""))) &&
		ruleConditionApplies(rule.getParameter(), paramName) &&
		ruleTypeApplies(rule.getType(), paramType) &&
		ruleConditionApplies(rule.getLanguage(), "go")
}

func ruleConditionApplies(ruleValue string, actualValue string) bool {
	if ruleValue == "" {
		return true
	}
	if strings.HasPrefix(ruleValue, "/") && strings.HasSuffix(ruleValue, "/") {
		match, _ := regexp.MatchString("(?i)" + ruleValue[1:len(ruleValue) - 1], actualValue)
		return match
	}
	return actualValue != "" && strings.ToLower(ruleValue) == strings.ToLower(actualValue);
}

func ruleTypeApplies(ruleType string, actualType string) bool {
	if ruleType == "" {
		return true
	}
	if actualType == "" {
		return false
	}
        if ruleType == "number" {
            return actualType == "int32" || actualType == "float64"
	}
        if ruleType == "int" {
            return actualType == "int32"
	}
        if ruleType == "bool" {
            return actualType == "bool"
	}
        if ruleType == "int[]" {
            return actualType == "[]int32"
	}
        if ruleType == "stream" {
            return actualType == "[]byte"
	}
        if ruleType == "stream[]" {
            return actualType == "[][]byte"
	}
        if ruleType == "model" {
            return isModelClass(actualType)
	}
	if isModelClass(ruleType) {
		return isSubclass(ruleType, actualType)
	}
	return false
}

func invalidizeTestParamValue(value interface{}, functionName string, paramName string, paramType string) interface{} {
	var invalidValue interface{}
	for _, rule := range getRules(getTestRules().Values, functionName, paramName, paramType) {
		valueRule := rule.(ValueRule)
		if valueRule.InvalidValueSet {
			invalidValue = valueRule.InvalidValue
		}
	}
	invalidValue = undefaultize(invalidValue, paramType, paramType)
	return untemplatize(invalidValue, value, paramName)
}

func assertError(t *testing.T, functionName string, paramName string, paramType string, paramValue interface{}, errorCode int32, e error) {
	if e == nil {
		failed := true
		for _, _ = range getRules(getTestRules().OkToNotFail, functionName, paramName, paramType) {
			failed = false
		}
		if failed {
			t.Errorf("Must have failed.")
			return
		}
	} else {
		var code int32
		message := "Unexpeceted message"
		for _, rule := range getRules(getTestRules().Results, functionName, paramName, paramType) {
			resultRule := rule.(ResultRule)
			if resultRule.Code != 0 {
				code = resultRule.Code
			}
			if resultRule.Message != "" {
				message = resultRule.Message
			}
		}
		if errorCode != code {
			t.Errorf("Unexpected error code: %v.", errorCode)
			return
		}
		if e == nil {
			t.Errorf("Error object is missing.")
			return
		}
		if !strings.Contains(e.Error(), untemplatize(message, paramValue, paramName).(string)) {
			t.Errorf("Unexpected error message: %v.", e)
			return
		}
	}
}

func assertBinaryResponse(file *os.File, t *testing.T) {
	if len(file.Name()) < 0 {
		t.Errorf("Empty or corrupt output file: %v.", file)
	}
}

func untemplatize(template interface{}, value interface{}, name interface{}) interface{} {
	if template != nil && reflect.TypeOf(template).Name() == "string" {
		if strings.HasPrefix(template.(string), "@") {
			fileName := template.(string)[1:]
			if strings.HasPrefix(fileName, "(") && strings.HasSuffix(fileName, ")") {
				files := [][]byte {}
				for _, item := range strings.Split(fileName[1:len(fileName) - 1], ",") {
					itemData, _ := ioutil.ReadFile("TestData/" + item)
					files = append(files, itemData)
				}
				return files
			}
			data, _ := ioutil.ReadFile("TestData/" + fileName)
			return data
		}
		if template.(string) == "#OperationId" {
			return GetOperationId()
		}
		if template.(string) == "#NewId" {
			return "705e4dcb-3ecd-24f3-3a35-3e926e4bded5"
		}
		if value != nil {
			if reflect.TypeOf(value).Name() == "string" {
				return strings.Replace(template.(string), "%v", value.(string), -1)
			}
			if reflect.TypeOf(value).Name() == "int32" {
				return strings.Replace(template.(string), "%v", fmt.Sprint(value), -1)
			}
		}
		if name != nil {
			if reflect.TypeOf(name).Name() == "string" {
				return strings.Replace(template.(string), "%n", name.(string), -1)
			}
			if reflect.TypeOf(name).Name() == "int32" {
				return strings.Replace(template.(string), "%n", fmt.Sprint(name), -1)
			}
		}
	}
	return template
}
