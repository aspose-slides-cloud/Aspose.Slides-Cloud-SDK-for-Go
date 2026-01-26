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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

/*
   Test for header/footer on all slides
*/
func TestSlidesFooter(t *testing.T) {
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

	dto := slidescloud.NewHeaderFooter()
	isFooterVisible := true
	dto.IsFooterVisible = &isFooterVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
	_, _, e = c.SlidesApi.SetPresentationHeaderFooter(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	result, _, e := c.SlidesApi.GetSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on one slide
*/
func TestSlideFooter(t *testing.T) {
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

	dto := slidescloud.NewHeaderFooter()
	isFooterVisible := true
	dto.IsFooterVisible = &isFooterVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
	result, _, e := c.SlidesApi.SetSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsFooterVisible() {
		t.Errorf("Wrong IsFooterVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}

/*
   Test for header/footer on notes slide
*/
func TestNotesSlideFooter(t *testing.T) {
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

	dto := slidescloud.NewNotesSlideHeaderFooter()
	isHeaderVisible := true
	dto.IsHeaderVisible = &isHeaderVisible
	dto.FooterText = "footer"
	isDateTimeVisible := false
	dto.IsDateTimeVisible = &isDateTimeVisible
	result, _, e := c.SlidesApi.SetNotesSlideHeaderFooter(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
	result, _, e = c.SlidesApi.GetNotesSlideHeaderFooter(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if !*result.GetIsHeaderVisible() {
		t.Errorf("Wrong IsHeaderVisible value. Expected to be true.")
		return
	}
	if *result.GetIsDateTimeVisible() {
		t.Errorf("Wrong IsDateTimeVisible value. Expected to be false.")
		return
	}
}