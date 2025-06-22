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
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

var mathSlideIndex     int32 = 2
var mathShapeIndex     int32 = 3
var mathParagraphIndex int32 = 1
var mathPortionIndex   int32 = 1

/*
   Test for Get math
*/
func TestMathGet(t *testing.T) {
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

	portion, _, e := c.SlidesApi.GetPortion(fileName, mathSlideIndex, mathShapeIndex, mathParagraphIndex, mathPortionIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 3 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 3, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[2].(slidescloud.IFractionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for Get empty math
*/
func TestMathGetNull(t *testing.T) {
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

	portion, _, e := c.SlidesApi.GetPortion(fileName, 2, 1, 1, 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() != nil {
		t.Errorf("Error: math paragraph should be nil.")
		return
	}
}

/*
   Test for create math
*/
func TestMathCreate(t *testing.T) {
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
	mathParagraph := slidescloud.NewMathParagraph()
	blockElement := slidescloud.NewBlockElement()
	functionElement := slidescloud.NewFunctionElement()
	limitElement := slidescloud.NewLimitElement()
	text1 := slidescloud.NewTextElement()
	text1.Value = "lim"
	limitElement.Base = text1
	text2 := slidescloud.NewTextElement()
	text2.Value = "x->0"
	limitElement.Limit = text2
	functionElement.Name = limitElement
	fractionElement := slidescloud.NewFractionElement()
	sinusElement := slidescloud.NewFunctionElement()
	text3 := slidescloud.NewTextElement()
	text3.Value = "sin"
	sinusElement.Name = text3
	text4 := slidescloud.NewTextElement()
	text4.Value = "x"
	sinusElement.Base = text4
	fractionElement.Numerator = sinusElement
	text5 := slidescloud.NewTextElement()
	text5.Value = "x"
	fractionElement.Denominator = text5
	functionElement.Base = fractionElement
	blockElement.MathElementList = []slidescloud.IMathElement{functionElement}
	mathParagraph.MathBlockList = []slidescloud.IBlockElement{blockElement}
	dto.MathParagraph = mathParagraph
	portion, _, e := c.SlidesApi.CreatePortion(fileName, 1, 1, 1, dto, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 1 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[0].(slidescloud.IFunctionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for update math
*/
func TestMathUpdate(t *testing.T) {
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
	mathParagraph := slidescloud.NewMathParagraph()
	blockElement := slidescloud.NewBlockElement()
	functionElement := slidescloud.NewFunctionElement()
	limitElement := slidescloud.NewLimitElement()
	text1 := slidescloud.NewTextElement()
	text1.Value = "lim"
	limitElement.Base = text1
	text2 := slidescloud.NewTextElement()
	text2.Value = "x->0"
	limitElement.Limit = text2
	functionElement.Name = limitElement
	fractionElement := slidescloud.NewFractionElement()
	sinusElement := slidescloud.NewFunctionElement()
	text3 := slidescloud.NewTextElement()
	text3.Value = "sin"
	sinusElement.Name = text3
	text4 := slidescloud.NewTextElement()
	text4.Value = "x"
	sinusElement.Base = text4
	fractionElement.Numerator = sinusElement
	text5 := slidescloud.NewTextElement()
	text5.Value = "x"
	fractionElement.Denominator = text5
	functionElement.Base = fractionElement
	blockElement.MathElementList = []slidescloud.IMathElement{functionElement}
	mathParagraph.MathBlockList = []slidescloud.IBlockElement{blockElement}
	dto.MathParagraph = mathParagraph
	portion, _, e := c.SlidesApi.UpdatePortion(fileName, mathSlideIndex, mathShapeIndex, mathParagraphIndex, mathPortionIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if portion.GetMathParagraph() == nil {
		t.Errorf("Error: math paragraph should not be nil.")
		return
	}
	if portion.GetMathParagraph().GetMathBlockList() == nil {
		t.Errorf("Error: math block list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()) != 1 {
		t.Errorf("Wrong math block list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()))
		return
	}
	if portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList() == nil {
		t.Errorf("Error: math element list should not be nil.")
		return
	}
	if len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()) != 1 {
		t.Errorf("Wrong math element list count. Expected %v but was %v.", 1, len(portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()))
		return
	}
	_, ok := portion.GetMathParagraph().GetMathBlockList()[0].GetMathElementList()[0].(slidescloud.IFunctionElement)
	if !ok {
		t.Errorf("Wrong math element type.")
		return
	}
}

/*
   Test for download math
*/
func TestMathDownload(t *testing.T) {
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

	mathMl, _, e := c.SlidesApi.DownloadMathPortion(fileName, mathSlideIndex, mathShapeIndex, mathParagraphIndex, mathPortionIndex, "MathML", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	mathMlStat, e := os.Stat(mathMl.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if mathMlStat.Size() <= 0 {
		t.Errorf("Wrong response size. Expected greater than %v but was %v.", 0, mathMlStat.Size())
		return
	}
}

/*
   Test for download math
*/
func TestMathDownloadNull(t *testing.T) {
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

	_, _, e = c.SlidesApi.DownloadMathPortion(fileName, 1, 1, 1, 1, "MathML", password, folderName, "")
	if e == nil {
		t.Errorf("Error must have been thrown: Cannot convert an ordinary portion to MathML.")
		return
	}
}

/*
   Test for save math
*/
func TestMathSave(t *testing.T) {
	outPath := folderName + "/mathml.xml"
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

	_, e = c.SlidesApi.SaveMathPortion(fileName, mathSlideIndex, mathShapeIndex, mathParagraphIndex, mathPortionIndex, "MathML", outPath, password, folderName, "")
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