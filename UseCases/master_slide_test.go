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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

/*
   Test for master slides
*/
func TestMasterSlides(t *testing.T) {
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

	masterSlides, _, e := c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 1 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 1, len(masterSlides.GetSlideList()))
		return
	}
	masterSlide, _, e := c.SlidesApi.GetMasterSlide(fileName, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.GetName() != "Office Theme" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Office Theme", masterSlide.GetName())
		return
	}

	masterSlide, _, e = c.SlidesApi.CopyMasterSlide(fileName, sourcePath, 1, "", "", nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if masterSlide.GetName() != "Digital portfolio" {
		t.Errorf("Wrong master slide name. Expected %v but was %v.", "Digital portfolio", masterSlide.GetName())
		return
	}
	masterSlides, _, e = c.SlidesApi.GetMasterSlides(fileName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 2 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 2, len(masterSlides.GetSlideList()))
		return
	}
}

/*
   Test for master slide shapes
*/
func TestMasterSlideShapes(t *testing.T) {
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

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "", "")
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
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "masterSlide", dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount+1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "masterSlide", shapeCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "masterSlide", password, folderName, "", "")
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
   Test for master slide paragraphs
*/
func TestMasterSlideParagraphs(t *testing.T) {
	var shapeIndex int32 = 2
	var paragraphCount int32 = 5
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

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "", "")
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
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "", "")
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
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount+1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount)+1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "masterSlide", shapeIndex, paragraphCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "masterSlide", shapeIndex, password, folderName, "", "")
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
   Test for master slide portions
*/
func TestMasterSlidePortions(t *testing.T) {
	var shapeIndex int32 = 2
	var paragraphIndex int32 = 3
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

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
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
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "", "")
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
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
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
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount+1, dto2, password, folderName, "", "")
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
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount)+1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems())+1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, portionCount+1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "masterSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
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
   Test for master slide animation
*/
func TestMasterSlideAnimation(t *testing.T) {
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

	animation, _, e := c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", nil, nil, password, folderName, "")
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
	animation, _, e = c.SlidesApi.SetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	var shapeIndex int32 = 3
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideAnimationEffect(fileName, slideIndex, "masterSlide", 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence())-1 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", len(dto.GetMainSequence())-1, len(animation.GetMainSequence()))
		return
	}
	animation, _, e = c.SlidesApi.GetSpecialSlideAnimation(fileName, slideIndex, "masterSlide", &shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}

	animation, _, e = c.SlidesApi.DeleteSpecialSlideAnimation(fileName, slideIndex, "masterSlide", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	dto2 := slidescloud.NewInteractiveSequence()
	dto2.TriggerShapeIndex = 1

	effect := slidescloud.NewEffect()
	effect.Type_ = "Fly"
	effect.Subtype = "Bottom"
	effect.PresetClassType = "Entrance"
	effect.ShapeIndex = 3
	effect.TriggerType = "OnClick"
	dto2.Effects = []slidescloud.IEffect{effect}

	animation, _, e = c.SlidesApi.CreateAnimationInteractiveSequence(fileName, 1, dto2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
	}
}

/*
   Test for deleting master slide
*/
func TestDeleteUnusedMasterSlides(t *testing.T) {
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

	var ignorePreserveField bool = true

	masterSlides, _, e := c.SlidesApi.DeleteUnusedMasterSlides(fileName, &ignorePreserveField, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(masterSlides.GetSlideList()) != 1 {
		t.Errorf("Wrong master slide count. Expected %v but was %v.", 1, len(masterSlides.GetSlideList()))
		return
	}
}

/*
   Test for deleting master slide from request
*/
func TestDeleteUnusedMasterSlidesOnline(t *testing.T) {
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

	var ignorePreserveField bool = true

	result, _, e := c.SlidesApi.DeleteUnusedMasterSlidesOnline(source, &ignorePreserveField, password)
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
