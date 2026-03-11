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
	"io/ioutil"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

/*
   Test for merge from storage
*/
func TestMerge(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	fileNamePdf := "test.pdf"
	path2 := folderName + "/" + fileName2
	pathPdf := folderName + "/" + fileNamePdf
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
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + fileName2, path2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + fileNamePdf, pathPdf, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewPresentationsMergeRequest()
	request.PresentationPaths = []string{path2, pathPdf}
	_, _, e = c.SlidesApi.Merge(fileName, request, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for ordered merge from storage
*/
func TestOrderedMerge(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	path2 := folderName + "/" + fileName2
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
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + fileName2, path2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation := slidescloud.NewPresentationToMerge()
	presentation.Path = path2
	presentation.Slides = []int32{2, 1}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation}
	_, _, e = c.SlidesApi.OrderedMerge(fileName, request, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge from request
*/
func TestMergeOnline(t *testing.T) {
	source1, e := ioutil.ReadFile(localFolder + "/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile(localFolder + "/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.MergeOnline([][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for merge and save from request
*/
func TestMergeAndSaveOnline(t *testing.T) {
	source1, e := ioutil.ReadFile(localFolder + "/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile(localFolder + "/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	outPath := folderName + "/out.pptx"
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.MergeAndSaveOnline(outPath, [][]byte{source1, source2}, nil, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultExists, _, e := c.SlidesApi.ObjectExists(outPath, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*resultExists.GetExists() {
		t.Errorf("File %v does not exist on the storage.", outPath)
		return
	}
}

/*
   Test for ordered merge from request
*/
func TestMergeOnlineWithRequest(t *testing.T) {
	source1, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source2, e := ioutil.ReadFile(localFolder + "/test-unprotected.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = password
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Path = "file2"
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}
	_, _, e = c.SlidesApi.MergeOnline([][]byte{source1, source2}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for combined merge from request
*/
func TestMergeOnlineCombined(t *testing.T) {
	fileName2 := "test-unprotected.pptx"
	path2 := folderName + "/" + fileName2
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + fileName2, path2, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Path = "file1"
	presentation1.Password = password
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Source = "Storage"
	presentation2.Path = path2
	presentation2.Slides = []int32{1, 2}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}
	_, _, e = c.SlidesApi.MergeOnline([][]byte{source}, request, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for
*/
func TestMergeOnlineUrl(t *testing.T) {
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

	request := slidescloud.NewOrderedMergeRequest()
	presentation1 := slidescloud.NewPresentationToMerge()
	presentation1.Source = "Storage"
	presentation1.Path = filePath
	presentation1.Slides = []int32{1, 2}
	presentation1.Password = password
	presentation2 := slidescloud.NewPresentationToMerge()
	presentation2.Source = "Url"
	presentation2.Path = "https://www.dropbox.com/scl/fi/nz3yt2wpg4eugiw12bz3g/info.pptx?rlkey=wju4vbi6h46q8733tzzvfgufz&e=1&st=pqt66x86&dl=1"
	presentation2.Slides = []int32{1}
	request.Presentations = []slidescloud.IPresentationToMerge{presentation1, presentation2}

	_, _, e = c.SlidesApi.MergeOnline(nil, request, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
