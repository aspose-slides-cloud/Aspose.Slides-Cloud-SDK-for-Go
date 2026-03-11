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
   Test for Get notes slide from storage
*/
func TestNotesSlideGetStorage(t *testing.T) {
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

	result, _, e := c.SlidesApi.GetNotesSlide(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetText()) == 0 {
		t.Errorf("Note text is empty.")
		return
	}
}

/*
   Test for exists notes slide from storage
*/
func TestNotesSlideExistsStorage(t *testing.T) {
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

	result, _, e := c.SlidesApi.NotesSlideExists(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetExists() {
		t.Errorf("Note does not exist.")
		return
	}
}

/*
   Test for download notes slide from storage
*/
func TestNotesSlideDownloadStorage(t *testing.T) {
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

	_, _, e = c.SlidesApi.DownloadNotesSlide(fileName, slideIndex, "png", nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for Get notes slide from request
*/
func TestNotesSlideGetRequest(t *testing.T) {
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
	result, _, e := c.SlidesApi.GetNotesSlideOnline(source, slideIndex, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetText()) == 0 {
		t.Errorf("Note text is empty.")
		return
	}
}

/*
   Test for exists notes slide from request
*/
func TestNotesSlideExistsRequest(t *testing.T) {
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
	result, _, e := c.SlidesApi.NotesSlideExistsOnline(source, slideIndex, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetExists() {
		t.Errorf("Note does not exist.")
		return
	}
}

/*
   Test for download notes slide from request
*/
func TestNotesSlideDownloadRequest(t *testing.T) {
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
	_, _, e = c.SlidesApi.DownloadNotesSlideOnline(source, slideIndex, "png", nil, nil, password, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for notes slide shapes
*/
func TestNotesSlideShapes(t *testing.T) {
	var shapeCount int32 = 3
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

	shapes, _, e := c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "", "")
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
	shape, _, e := c.SlidesApi.CreateSpecialSlideShape(fileName, slideIndex, "notesSlide", dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	dto.Text = "Updated shape"
	shape, _, e = c.SlidesApi.UpdateSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount + 1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.(slidescloud.IShape).GetText() != dto.GetText() {
		t.Errorf("Wrong shape text. Expected %v but was %v.", dto.GetText(), shape.(slidescloud.IShape).GetText())
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount)+1 {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount+1, len(shapes.GetShapesLinks()))
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideShape(fileName, slideIndex, "notesSlide", shapeCount + 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapes, _, e = c.SlidesApi.GetSpecialSlideShapes(fileName, slideIndex, "notesSlide", password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(shapes.GetShapesLinks()) != int(shapeCount) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount + 1, len(shapes.GetShapesLinks()))
		return
	}
}

/*
   Test for notes slide paragraphs
*/
func TestNotesSlideParagraphs(t *testing.T) {
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

	paragraphs, _, e := c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "", "")
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
	paragraph, _, e := c.SlidesApi.CreateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()) + 1)
		return
	}

	dto = slidescloud.NewParagraph()
	dto.Alignment = "Center"
	paragraph, _, e = c.SlidesApi.UpdateSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount + 1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paragraph.GetAlignment() != dto.GetAlignment() {
		t.Errorf("Wrong paragraph alignment. Expected %v but was %v.", dto.GetAlignment(), paragraph.GetAlignment())
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(paragraphs.GetParagraphLinks()) != int(paragraphCount) + 1 {
		t.Errorf("Wrong paragraph count. Expected %v but was %v.", paragraphCount, len(paragraphs.GetParagraphLinks()) + 1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlideParagraph(fileName, slideIndex, "notesSlide", shapeIndex, paragraphCount + 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	paragraphs, _, e = c.SlidesApi.GetSpecialSlideParagraphs(fileName, slideIndex, "notesSlide", shapeIndex, password, folderName, "", "")
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
   Test for notes slide portions
*/
func TestNotesSlidePortions(t *testing.T) {
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

	portions, _, e := c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
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
	portion, _, e := c.SlidesApi.CreateSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, dto, nil, password, folderName, "", "")
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
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()) + 1)
		return
	}

	dto2 := slidescloud.NewPortion()
	dto2.Text = "Updated portion"
	dto2.FontHeight = 22
	portion, _, e = c.SlidesApi.UpdateSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount + 1, dto2, password, folderName, "", "")
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
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(portions.GetItems()) != int(portionCount) + 1 {
		t.Errorf("Wrong portion count. Expected %v but was %v.", portionCount, len(portions.GetItems()) + 1)
		return
	}

	_, _, e = c.SlidesApi.DeleteSpecialSlidePortion(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, portionCount + 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	portions, _, e = c.SlidesApi.GetSpecialSlidePortions(fileName, slideIndex, "notesSlide", shapeIndex, paragraphIndex, password, folderName, "", "")
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
   Test create notes slide
*/
func TestCreateNotesSlide(t *testing.T) {
	notesSlideText := "Notes slides string"

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

	dto := slidescloud.NewNotesSlide()
	dto.SetText(notesSlideText)
	notesSlide, _, e := c.SlidesApi.CreateNotesSlide(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if notesSlide.GetText() != notesSlideText {
		t.Errorf("Notes slides text error. Expected \"%v\" but was \"%v\"", notesSlide.GetText(), notesSlideText)
		return
	}
}

/*
   Test update notes slide
*/
func TestUpdateNotesSlide(t *testing.T) {
	notesSlideText := "Notes slides string"

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

	dto := slidescloud.NewNotesSlide()
	dto.SetText(notesSlideText)
	notesSlide, _, e := c.SlidesApi.UpdateNotesSlide(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if notesSlide.GetText() != notesSlideText {
		t.Errorf("Notes slides text error. Expected \"%v\" but was \"%v\"", notesSlide.GetText(), notesSlideText)
		return
	}
}

/*
   Test for deleting notes slide
*/
func TestDeleteNotesSlide(t *testing.T) {
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

	slide, _, e := c.SlidesApi.DeleteNotesSlide(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if slide.GetNotesSlide() != nil {
		t.Errorf("NoteSlide expected to be nil.")
		return
	}
}