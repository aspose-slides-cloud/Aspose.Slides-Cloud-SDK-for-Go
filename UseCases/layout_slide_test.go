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
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

/*
   Test for layout slides
*/
func TestLayoutSlides(t *testing.T) {
	sourceFile := "TemplateCV.pptx"
	sourcePath := folderName + "/" + sourceFile
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
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + sourceFile, sourcePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	layoutSlides, _, e := c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.GetSlideList()) != 11 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 11, len(layoutSlides.GetSlideList()))
		return
	}
	layoutSlide, _, e := c.SlidesApi.GetLayoutSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.GetName() != "Title Slide" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title Slide", layoutSlide.GetName())
		return
	}

	layoutSlide, _, e = c.SlidesApi.CopyLayoutSlide(fileName, sourcePath, 2, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if layoutSlide.GetName() != "Title and Content" {
		t.Errorf("Wrong layout slide name. Expected %v but was %v.", "Title and Content", layoutSlide.GetName())
		return
	}
	layoutSlides, _, e = c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(layoutSlides.GetSlideList()) != 12 {
		t.Errorf("Wrong layout slide count. Expected %v but was %v.", 12, len(layoutSlides.GetSlideList()))
		return
	}
}

/*
   Test for layout slide shapes
*/
func TestLayoutSlideShapes(t *testing.T) {
	var shapeCount int32 = 6
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

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(shapes.GetShapesLinks()))
		return
	}

	dto := slidescloud.NewShape()
	dto.X = 100
	dto.Y = 100
	dto.Width = 500
	dto.Height = 200
	dto.ShapeType = "Rectangle"
	dto.Text = "New shape"
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "layoutSlide", dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount+1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "layoutSlide", shapeCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "layoutSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}
}

/*
   Test for layout slide paragraphs
*/
func TestLayoutSlideParagraphs(t *testing.T) {
	var shapeIndex int32 = 2
	var paragraphCount int32 = 1
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

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}

	portion := slidescloud.NewPortion()
	portion.Text = "New Paragraph"
	dto := slidescloud.NewParagraph()
	dto.Alignment = "Right"
	dto.PortionList = []slidescloud.IPortion{portion}
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	dto = slidescloud.NewParagraph()
	dto.Alignment = "Center"
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount+1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "layoutSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()))
		return
	}
}

/*
   Test for layout slide portions
*/
func TestLayoutSlidePortions(t *testing.T) {
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 1
	var portionCount int32 = 1
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

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}

	dto := slidescloud.NewPortion()
	dto.Text = "New portion"
	dto.FontBold = "True"
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetText() != dto.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	dto2 := slidescloud.NewPortion()
	dto2.Text = "Updated portion"
	dto2.FontHeight = 22
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount+1, dto2, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetFontBold() != dto.GetFontBold() {
		t.Errorf("Wrong portion font bold. Expected %v but was %v.", dto.GetFontBold(), portion.GetFontBold())
		return
	}
	if portion.GetFontHeight() != dto2.GetFontHeight() {
		t.Errorf("Wrong portion font height. Expected %v but was %v.", dto2.GetFontHeight(), portion.GetFontHeight())
		return
	}
	if portion.GetText() != dto2.GetText() {
		t.Errorf("Wrong portion text. Expected %v but was %v.", dto2.GetText(), portion.GetText())
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, portionCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "layoutSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()))
		return
	}
}

/*
   Test for layout slide animation
*/
func TestLayoutSlideAnimation(t *testing.T) {
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

	animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	effect1 := slidescloud.NewEffect()
	effect1.Type_ = "Blink"
	effect1.ShapeIndex = 2
	effect2 := slidescloud.NewEffect()
	effect2.Type_ = "Appear"
	effect2.ShapeIndex = 3
	dto := slidescloud.NewSlideAnimation()
	dto.MainSequence = []slidescloud.IEffect{effect1, effect2}
	animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	var shapeIndex int32 = 3
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "layoutSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence())-1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence())-1, len(animation.GetMainSequence()))
		return
	}
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}

	animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "layoutSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
}

/*
   Test for deleting layout slide
*/
func TestDeleteUnusedLayoutSlides(t *testing.T) {
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

	layoutSlidesBefore, _, e := c.SlidesApi.GetLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(layoutSlidesBefore.GetSlideList()) != 11 {
		t.Errorf("Expected %v, but was %v", 11, layoutSlidesBefore.GetSlideList())
		return
	}

	layoutSlidesAfter, _, e := c.SlidesApi.DeleteUnusedLayoutSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(layoutSlidesAfter.GetSlideList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, layoutSlidesAfter.GetSlideList())
		return
	}
}

/*
   Test for deleting layout slide from request
*/
func TestDeleteUnusedLayoutSlidesOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	result, _, e := c.SlidesApi.DeleteUnusedLayoutSlidesOnline(source, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if resultStat.Size() == 0 {
		t.Errorf("Wrong response size.")
		return
	}
}
