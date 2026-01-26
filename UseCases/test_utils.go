package usecasetests

import (
	"encoding/json"
	"io/ioutil"
	"os"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

var (
	testApiClient      *slidescloud.APIClient
	testAsyncApiClient *slidescloud.APIClient
	isInitialized            = false
	expectedTestDataVersion  = "1"
	tempFolderName           = "TempTests"
	folderName               = "TempSlidesSDK"
	fileName                 = "test.pptx"
	tempFilePath             = tempFolderName + "/" + fileName
	filePath                 = folderName + "/" + fileName
	localFolder              = "../TestData"
	localTestFile            = localFolder + "/" + fileName
	slideIndex         int32 = 1
	password                 = "password"
)

func GetApiClient() (*slidescloud.APIClient, error) {
	if testApiClient == nil {
		cfg := slidescloud.NewConfiguration()
		configFile, err := os.Open("testConfig.json")
		if err != nil {
			return nil, err
		}
		json.NewDecoder(configFile).Decode(cfg)
		testApiClient = slidescloud.NewAPIClient(cfg)
		err = InitializeTesting()
		if err != nil {
			return nil, err
		}
	}
	return testApiClient, nil
}

func GetAsyncApiClient() (*slidescloud.APIClient, error) {
	_, err := GetApiClient()
	if err != nil {
		return nil, err
	}
	if testAsyncApiClient == nil {
		cfg := slidescloud.NewConfiguration()
		configFile, err := os.Open("testConfig.json")
		if err != nil {
			return nil, err
		}
		json.NewDecoder(configFile).Decode(cfg)
		cfg.BasePath = cfg.AsyncBasePath
		testAsyncApiClient = slidescloud.NewAPIClient(cfg)
	}
	return testAsyncApiClient, nil
}

func InitializeTesting() error {
	if !isInitialized {
		version := "0"
		f, _, e := testApiClient.SlidesApi.DownloadFile("TempTests/version.txt", "", "")
		if e == nil {
			data, e := ioutil.ReadFile(f.Name())
			if e == nil {
				version = string(data)
			}
		}
		if version != expectedTestDataVersion {
			files, e := ioutil.ReadDir("../TestData")
			if e != nil {
				return e
			}
			for _, file := range files {
				bytes, e := ioutil.ReadFile("../TestData/" + file.Name())
				if e != nil {
					return e
				}
				_, _, e = testApiClient.SlidesApi.UploadFile("TempTests/" + file.Name(), bytes, "")
				if e != nil {
					return e
				}
			}
			_, _, e = testApiClient.SlidesApi.UploadFile("TempTests/version.txt", []byte(expectedTestDataVersion), "")
			if e != nil {
				return e
			}
		}
		isInitialized = true
	}
	return nil
}
