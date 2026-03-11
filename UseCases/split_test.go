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

package usecasetests

import (
	"archive/zip"
	"io/ioutil"
	"strings"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

/*
   Test for split from storage
*/
func TestSplitStorage(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result1, _, e := c.SlidesApi.Split(fileName, nil, "", nil, nil, nil, nil, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.Split(fileName, nil, "", nil, nil, &from, &to, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result2.GetSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.GetSlides()) <= len(result2.GetSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.GetSlides()), len(result1.GetSlides()))
		return
	}
	url := result1.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for split from request
*/
func TestSplitRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result1, _, e := c.SlidesApi.SplitOnline(source, "png", nil, nil, nil, nil, password, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	zip1, e := zip.OpenReader(result1.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer zip1.Close()
	count1 := 0
	for _, _ = range zip1.File {
		count1++
	}

	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.SplitOnline(source, "png", nil, nil, &from, &to, password, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	zip2, e := zip.OpenReader(result2.Name())
	if e != nil {
		t.Errorf("Failed to open zip: %v.", e)
		return
	}
	defer zip2.Close()
	count2 := 0
	for _, _ = range zip2.File {
		count2++
	}
	if count2 != 2 {
		t.Errorf("Wrong slide count. Expected %v but was %v.", 2, count2)
		return
	}
	if count1 <= count2 {
		t.Errorf("Wrong slide count. Expected greater than %v but was %v.", count2, count1)
		return
	}
}

/*
   Test for split from request to storage
*/
func TestSplitAndSaveRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result1, _, e := c.SlidesApi.SplitAndSaveOnline(source, "png", "", nil, nil, nil, nil, password, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var from int32 = 2
	var to int32 = 3
	result2, _, e := c.SlidesApi.SplitAndSaveOnline(source, "png", "", nil, nil, &from, &to, password, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result2.GetSlides()) != 2 {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result1.GetSlides()) <= len(result2.GetSlides()) {
		t.Errorf("Wrong slide count. Expected greate than %v but was %v.", len(result2.GetSlides()), len(result1.GetSlides()))
		return
	}
	url := result1.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}

/*
   Test for split with options
*/
func TestSplitWithOptions(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	options := slidescloud.NewPdfExportOptions()
	options.JpegQuality = 50
	result, _, e := c.SlidesApi.Split(fileName, options, "", nil, nil, nil, nil, "", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	url := result.GetSlides()[0].GetHref()
	path := url[strings.Index(url, "/storage/file/")+len("/storage/file/"):]
	resultExists, _, e := c.SlidesApi.ObjectExists(path, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", path)
		return
	}
}
