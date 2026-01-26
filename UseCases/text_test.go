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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

/*
   Test for Get text
*/
func TestTextGet(t *testing.T) {
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

	var withEmpty bool = true
	result, _, e := c.SlidesApi.GetPresentationTextItems(fileName, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultWithEmpty, _, e := c.SlidesApi.GetPresentationTextItems(fileName, &withEmpty, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResult, _, e := c.SlidesApi.GetSlideTextItems(fileName, slideIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResultWithEmpty, _, e := c.SlidesApi.GetSlideTextItems(fileName, slideIndex, &withEmpty, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetItems()) >= len(resultWithEmpty.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(result.GetItems()), len(resultWithEmpty.GetItems()))
		return
	}
	if len(slideResult.GetItems()) >= len(result.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.GetItems()), len(result.GetItems()))
		return
	}
	if len(slideResult.GetItems()) >= len(slideResultWithEmpty.GetItems()) {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", len(slideResult.GetItems()), len(slideResultWithEmpty.GetItems()))
		return
	}
}

/*
   Test for replace text on storage
*/
func TestTextReplaceStorage(t *testing.T) {
	oldValue := "text"
	newValue := "new_text"
	var ignoreCase bool = true
	var wholeWordsOnly bool = true

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
	result, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultRegex, _, e := c.SlidesApi.ReplacePresentationRegex(fileName, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultWithEmpty, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, &ignoreCase, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultWholeWords, _, e := c.SlidesApi.ReplacePresentationText(fileName, oldValue, newValue, &ignoreCase, &wholeWordsOnly, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResult, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	slideResultWithEmpty, _, e := c.SlidesApi.ReplaceSlideText(fileName, slideIndex, oldValue, newValue, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if result.GetMatches() != resultRegex.GetMatches() {
		t.Errorf("Wrong item count. Expected %v but was %v.", result.GetMatches(), resultRegex.GetMatches())
		return
	}
	if result.GetMatches() >= resultWithEmpty.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", result.GetMatches(), resultWithEmpty.GetMatches())
		return
	}
	if resultWholeWords.GetMatches() >= resultWithEmpty.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", resultWholeWords.GetMatches(), resultWithEmpty.GetMatches())
		return
	}
	if slideResult.GetMatches() >= result.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.GetMatches(), result.GetMatches())
		return
	}
	if slideResult.GetMatches() >= slideResultWithEmpty.GetMatches() {
		t.Errorf("Wrong item count. Expected less than %v but was %v.", slideResult.GetMatches(), slideResultWithEmpty.GetMatches())
		return
	}
}

/*
   Test for replace text on request
*/
func TestTextReplaceRequest(t *testing.T) {
	oldValue := "text"
	newValue := "new_text"
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

	var withEmpty bool = true
	_, _, e = c.SlidesApi.ReplacePresentationTextOnline(source, oldValue, newValue, nil, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplacePresentationRegexOnline(source, oldValue, newValue, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplacePresentationTextOnline(source, oldValue, newValue, &withEmpty, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplaceSlideTextOnline(source, slideIndex, oldValue, newValue, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.ReplaceSlideTextOnline(source, slideIndex, oldValue, newValue, &withEmpty, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for replace text with formatting
*/
func TestReplaceTextFormatting(t *testing.T) {
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1
	oldText := "banana"
	newText := "orange"
	color := "#FFFFA500"

	portion := slidescloud.NewPortion()
	portion.Text = oldText

	portionFormat := slidescloud.NewPortionFormat()
	portionFormat.FontColor = color

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
	_, _, e = c.SlidesApi.CreatePortion(fileName, slideIndex, shapeIndex, paragraphIndex, portion, &portionIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.ReplaceTextFormatting(fileName, oldText, newText, portionFormat, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	updatedPortion, _, e := c.SlidesApi.GetPortion(fileName, slideIndex, shapeIndex, paragraphIndex, portionIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if updatedPortion.GetText() != newText {
		t.Errorf("Wrong portion text. Expected %v but was %v.", newText, updatedPortion.GetText())
		return
	}
	if updatedPortion.GetFontColor() != color {
		t.Errorf("Wrong portion color. Expected %v but was %v.", color, updatedPortion.GetFontColor())
		return
	}
}

/*
   Test for replace text with formatting online
*/
func TestReplaceTextFormattingOnline(t *testing.T) {
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

	portionFormat := slidescloud.NewPortionFormat()
	portionFormat.FontColor = "#FFFFA500"
	_, _, e = c.SlidesApi.ReplaceTextFormattingOnline(source, "orange", "banana", portionFormat, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for highlight shape text
*/
func TestHighlightShapeText(t *testing.T) {
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"
	ignoreCase := false

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

	_, _, e = c.SlidesApi.HighlightShapeText(fileName, slideIndex, shapeIndex, textToHighlight, highlightColor, nil, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}

/*
   Test for highlight shape text with regex
*/
func TestHighlightShapeRegex(t *testing.T) {
	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	highlightRegex := "h.ghl[abci]ght"
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"
	ignoreCase := false
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

	_, _, e = c.SlidesApi.HighlightShapeRegex(fileName, slideIndex, shapeIndex, highlightRegex, highlightColor, &ignoreCase, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}

/*
   Test for highlight presentation text
*/
func TestHighlightPresentationText(t *testing.T) {
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"

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

	ignoreCase := false
	result, _, e := c.SlidesApi.HighlightPresentationText(fileName, textToHighlight, highlightColor, nil, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	ignoreCase = true
	resultIgnoreCase, _, e := c.SlidesApi.HighlightPresentationText(fileName, textToHighlight, highlightColor, nil, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetMatches() != resultIgnoreCase.GetMatches() {
		t.Errorf("Wrong item count. Expected %v but was %v.", result.GetMatches(), resultIgnoreCase.GetMatches())
		return
	}

	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}

/*
   Test for highlight presentation text with regex
*/
func TestHighlightPresentationRegex(t *testing.T) {
	highlightRegex := "h.ghl[abci]ght"
	textToHighlight := "highlight"
	highlightColor := "#FFF5FF8A"
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

	ignoreCase := false
	result, _, e := c.SlidesApi.HighlightPresentationRegex(fileName, highlightRegex, highlightColor, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	ignoreCase = true
	resultIgnoreCase, _, e := c.SlidesApi.HighlightPresentationRegex(fileName, highlightRegex, highlightColor, &ignoreCase, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetMatches() != resultIgnoreCase.GetMatches() {
		t.Errorf("Wrong item count. Expected %v but was %v.", result.GetMatches(), resultIgnoreCase.GetMatches())
		return
	}

	var slideIndex int32 = 6
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	para, _, e := c.SlidesApi.GetParagraph(fileName, slideIndex, shapeIndex, paragraphIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if para.GetPortionList()[0].GetText() == textToHighlight {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[0].GetHighlightColor() == highlightColor {
		t.Errorf("Expected %v, but was %v", "nil", para.GetPortionList()[0].GetHighlightColor())
		return
	}

	if para.GetPortionList()[1].GetText() != textToHighlight {
		t.Errorf("Expected %v, but was %v", textToHighlight, para.GetPortionList()[0].GetText())
		return
	}

	if para.GetPortionList()[1].GetHighlightColor() != highlightColor {
		t.Errorf("Expected %v, but was %v", highlightColor, para.GetPortionList()[0].GetHighlightColor())
		return
	}
}