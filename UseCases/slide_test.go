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
	"strings"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

/*
   Test for get slides
*/
func TestGetSlides(t *testing.T) {
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

	response, _, e := c.SlidesApi.GetSlides(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != 9 {
		t.Errorf("Expected %v, but was %v", 9, len(response.GetSlideList()))
		return
	}
}

/*
   Test for get slide
*/
func TestGetSlide(t *testing.T) {
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

	_, _, e = c.SlidesApi.GetSlide(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for create slide
*/
func TestCreateSlide(t *testing.T) {
	var position int32 = 1
	var slidesCount int = 9

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

	response, _, e := c.SlidesApi.CreateSlide(fileName, "layoutSlides/3", &position, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount + 1 {
		t.Errorf("Expected %v, but was %v", slidesCount + 1, len(response.GetSlideList()))
		return
	}

	response, _, e = c.SlidesApi.CreateSlide(fileName, "layoutSlides/3", nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount + 2 {
		t.Errorf("Expected %v, but was %v", slidesCount + 2, len(response.GetSlideList()))
		return
	}
}

/*
   Test for copy slide
*/
func TestCopySlide(t *testing.T) {
	var slidesCount int = 9

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

	response, _, e := c.SlidesApi.CopySlide(fileName, 3, nil, "", "", "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount + 1 {
		t.Errorf("Expected %v, but was %v", slidesCount + 1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for copy slide from source
*/
func TestCopySlideFromSource(t *testing.T) {
	sourceFileName := "TemplateCV.pptx"
	sourceFilePath := folderName + "/" + sourceFileName
	var slidesCount int = 9
	var position int32 = 1

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

	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + sourceFileName, sourceFilePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	response, _, e := c.SlidesApi.CopySlide(fileName, slideIndex, &position, sourceFilePath, "", "", password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount + 1 {
		t.Errorf("Expected %v, but was %v", slidesCount + 1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for move slide
*/
func TestMoveSlide(t *testing.T) {
	var slidesCount int = 9
	var position int32 = 2

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

	response, _, e := c.SlidesApi.MoveSlide(fileName, slideIndex, position, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount {
		t.Errorf("Expected %v, but was %v", slidesCount, len(response.GetSlideList()))
		return
	}
}

/*
   Test for reorder slides
*/
func TestReorderSlides(t *testing.T) {
	var slidesCount int = 9

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

	oldPositions := []int32{1, 2, 3, 4, 5, 6}
	newPositions := []int32{6, 5, 4, 3, 2, 1}

	response, _, e := c.SlidesApi.ReorderSlides(fileName, oldPositions, newPositions, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount {
		t.Errorf("Expected %v, but was %v", slidesCount, len(response.GetSlideList()))
		return
	}
}

/*
   Test for update slide
*/
func TestUpdateSlide(t *testing.T) {
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

	layoutSlide := slidescloud.NewResourceUri()
	layoutSlide.Href = "layoutSlides/3"
	dto := slidescloud.NewSlide()
	dto.LayoutSlide = layoutSlide

	response, _, e := c.SlidesApi.UpdateSlide(fileName, slideIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if !strings.Contains(response.GetLayoutSlide().GetHref(), "layoutSlides/3") {
		t.Errorf("Expected %v, but was %v", "contains \"layoutSlides/3\"", "\""+response.GetLayoutSlide().GetHref()+"\"")
		return
	}
}

/*
   Test for set slide transition
*/
func TestSetSlideTransition(t *testing.T) {
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

	transition := slidescloud.NewSlideShowTransition()
	transition.Type_ = "Circle"
	transition.Speed = "Medium"
	dto := slidescloud.NewSlide()
	dto.SlideShowTransition = transition

	response, _, e := c.SlidesApi.UpdateSlide(fileName, slideIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if !strings.Contains(response.GetSlideShowTransition().GetType(), dto.SlideShowTransition.GetType()) {
		t.Errorf("Expected %v, but was %v", dto.SlideShowTransition.GetType(), response.GetSlideShowTransition().GetType())
		return
	}
}

/*
   Test for delete slide
*/
func TestDeleteSlides(t *testing.T) {
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

	response, _, e := c.SlidesApi.DeleteSlides(fileName, nil, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for delete slides by indexes
*/
func TestDeleteSlideByIndexes(t *testing.T) {
	var slidesCount int = 9

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

	indexes := []int32{1, 3, 5}
	response, _, e := c.SlidesApi.DeleteSlides(fileName, indexes, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount-len(indexes) {
		t.Errorf("Expected %v, but was %v", slidesCount-len(indexes), len(response.GetSlideList()))
		return
	}
}

/*
   Test for delete slide
*/
func TestDeleteSlide(t *testing.T) {
	var slidesCount int = 9

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

	response, _, e := c.SlidesApi.DeleteSlide(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetSlideList()) != slidesCount - 1 {
		t.Errorf("Expected %v, but was %v", slidesCount - 1, len(response.GetSlideList()))
		return
	}
}

/*
   Test for get background
*/
func TestGetBackground(t *testing.T) {
	var slideIndex int32 = 5

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

	response, _, e := c.SlidesApi.GetBackground(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for set background
*/
func TestSetBackground(t *testing.T) {
	color := "#FFF5FF8A"
	var slideIndex int32 = 5

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

	dto := slidescloud.NewSlideBackground()
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = color
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.SetBackground(fileName, slideIndex, dto, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}

	if response.GetFillFormat().(slidescloud.ISolidFill).GetColor() != color {
		t.Errorf("Expected %v, but was %v", color, response.GetFillFormat().(slidescloud.ISolidFill).GetColor())
		return
	}
}

/*
   Test for set background color
*/
func TestSetBackgroundColor(t *testing.T) {
	color := "#FFF5FF8A"

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

	response, _, e := c.SlidesApi.SetBackgroundColor(fileName, slideIndex, color, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "Solid" {
		t.Errorf("Expected %v, but was %v", "Solid", response.GetFillFormat().GetType())
		return
	}

	if response.GetFillFormat().(slidescloud.ISolidFill).GetColor() != color {
		t.Errorf("Expected %v, but was %v", color, response.GetFillFormat().(slidescloud.ISolidFill).GetColor())
		return
	}
}

/*
   Test for delete background
*/
func TestDeleteBackground(t *testing.T) {
	var slideIndex int32 = 5

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

	response, _, e := c.SlidesApi.DeleteBackground(fileName, slideIndex, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetFillFormat().GetType() != "NoFill" {
		t.Errorf("Expected %v, but was %v", "NoFill", response.GetFillFormat().GetType())
		return
	}
}