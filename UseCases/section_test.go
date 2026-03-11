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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

/*
   Test for Get sections
*/
func TestGetSections(t *testing.T) {
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

	result, _, e := c.SlidesApi.GetSections(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
}

/*
   Test for replace sections
*/
func TestSetSections(t *testing.T) {
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

	dto := slidescloud.NewSections()
	section1 := slidescloud.NewSection()
	section1.Name = "Section1"
	section1.FirstSlideIndex = 1
	section2 := slidescloud.NewSection()
	section2.Name = "Section2"
	section2.FirstSlideIndex = 3
	dto.SectionList = []slidescloud.ISection{section1, section2}
	result, _, e := c.SlidesApi.SetSections(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != len(dto.GetSectionList()) {
		t.Errorf("Wrong section count. Expected %v but was %v.", len(dto.GetSectionList()), len(result.GetSectionList()))
		return
	}
	if len(result.GetSectionList()[0].GetSlideList()) != int(section2.FirstSlideIndex-section1.FirstSlideIndex) {
		t.Errorf("Wrong slide count. Expected %v but was %v.", section2.FirstSlideIndex-section1.FirstSlideIndex, len(result.GetSectionList()[0].GetSlideList()))
		return
	}
}

/*
   Test for create section
*/
func TestCreateSection(t *testing.T) {
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

	result, _, e := c.SlidesApi.CreateSection(fileName, "NewSection", 5, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 4 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 4, len(result.GetSectionList()))
		return
	}
}

/*
   Test for update section
*/
func TestUpdateSection(t *testing.T) {
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

	var sectionIndex int32 = 2
	sectionName := "UpdatedSection"
	result, _, e := c.SlidesApi.UpdateSection(fileName, sectionIndex, sectionName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
	if result.GetSectionList()[sectionIndex-1].GetName() != sectionName {
		t.Errorf("Wrong section name. Expected %v but was %v.", sectionName, result.GetSectionList()[sectionIndex-1].GetName())
		return
	}
}

/*
   Test for move section
*/
func TestMoveSection(t *testing.T) {
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

	result, _, e := c.SlidesApi.MoveSection(fileName, 1, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 3 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 3, len(result.GetSectionList()))
		return
	}
}

/*
   Test for clear sections
*/
func TestClearSections(t *testing.T) {
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

	result, _, e := c.SlidesApi.DeleteSections(fileName, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 0 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 0, len(result.GetSectionList()))
		return
	}
}

/*
   Test for delete sections
*/
func TestDeleteSections(t *testing.T) {
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

	result, _, e := c.SlidesApi.DeleteSections(fileName, []int32{2, 3}, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 1 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 1, len(result.GetSectionList()))
		return
	}
}

/*
   Test for delete section
*/
func TestDeleteSection(t *testing.T) {
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

	result, _, e := c.SlidesApi.DeleteSection(fileName, 2, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetSectionList()) != 2 {
		t.Errorf("Wrong section count. Expected %v but was %v.", 2, len(result.GetSectionList()))
		return
	}
}
