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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

/*
   Test for Get hyperlink for shape
*/
func TestHyperlinkGetShape(t *testing.T) {
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

	shape, _, e := c.SlidesApi.GetShape(fileName, 2, 2, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape.GetHyperlinkClick() == nil {
		t.Errorf("Error: click hyperlink should not be nil.")
		return
	}
	if shape.GetHyperlinkClick().GetActionType() != "Hyperlink" {
		t.Errorf("Wrong action type. Expected %v but was %v.", "Hyperlink", shape.GetHyperlinkClick().GetActionType())
		return
	}
	if shape.GetHyperlinkMouseOver() != nil {
		t.Errorf("Error: mouse over hyperlink should be nil.")
		return
	}
}

/*
   Test for Get hyperlink for portion
*/
func TestHyperlinkGetPortion(t *testing.T) {
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

	portion, _, e := c.SlidesApi.GetPortion(fileName, 2, 1, 1, 2, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetHyperlinkClick() != nil {
		t.Errorf("Error: click hyperlink should be nil.")
		return
	}
	if portion.GetHyperlinkMouseOver() == nil {
		t.Errorf("Error: mouse over hyperlink should not be nil.")
		return
	}
	if portion.GetHyperlinkMouseOver().GetActionType() != "JumpLastSlide" {
		t.Errorf("Wrong action type. Expected not %v but was %v.", "JumpLastSlide", portion.GetHyperlinkMouseOver().GetActionType())
		return
	}
}

/*
   Test for create hyperlink for shape
*/
func TestHyperlinkCreateShape(t *testing.T) {
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

	shape := slidescloud.NewShape()
	hyperlink := slidescloud.NewHyperlink()
	hyperlink.ActionType = "Hyperlink"
	hyperlink.ExternalUrl = "https://docs.aspose.cloud/slides"
	shape.HyperlinkClick = hyperlink
	updatedShape, _, e := c.SlidesApi.UpdateShape(fileName, 1, 1, shape, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedShape.GetHyperlinkClick() == nil {
		t.Errorf("Error: click hyperlink should not be nil.")
		return
	}
	if updatedShape.GetHyperlinkClick().GetExternalUrl() != shape.GetHyperlinkClick().GetExternalUrl() {
		t.Errorf("Wrong URL. Expected %v but was %v.", shape.GetHyperlinkClick().GetExternalUrl(), updatedShape.GetHyperlinkClick().GetExternalUrl())
		return
	}
}

/*
   Test for create hyperlink for portion
*/
func TestHyperlinkCreatePortion(t *testing.T) {
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

	dto := slidescloud.NewPortion()
	dto.Text = "Link text"
	hyperlink := slidescloud.NewHyperlink()
	hyperlink.ActionType = "JumpLastSlide"
	dto.HyperlinkMouseOver = hyperlink
	updatedPortion, _, e := c.SlidesApi.CreatePortion(fileName, 1, 1, 1, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedPortion.GetHyperlinkMouseOver() == nil {
		t.Errorf("Error: mouse over hyperlink should not be nil.")
		return
	}
	if updatedPortion.GetHyperlinkMouseOver().GetActionType() != "JumpLastSlide" {
		t.Errorf("Wrong action type. Expected %v but was %v.", dto.GetHyperlinkMouseOver().GetActionType(), updatedPortion.GetHyperlinkMouseOver().GetActionType())
		return
	}
}

/*
   Test for delete hyperlink
*/
func TestHyperlinkDelete(t *testing.T) {
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

	shape := slidescloud.NewPictureFrame()
	hyperlink := slidescloud.NewHyperlink()
	isDisabled := true
	hyperlink.IsDisabled = &isDisabled
	shape.HyperlinkClick = hyperlink
	updatedShape, _, e := c.SlidesApi.UpdateShape(fileName, 1, 1, shape, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if updatedShape.GetHyperlinkClick() != nil {
		t.Errorf("Error: click hyperlink should be nil.")
		return
	}
}
