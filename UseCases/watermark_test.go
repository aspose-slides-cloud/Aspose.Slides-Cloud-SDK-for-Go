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
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

/*
   Test for text watermark on storage
*/
func TestWatermarkTextStorage(t *testing.T) {
	watermarkText := "watermarkText"
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
	getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapeCount := len(getResult.GetShapesLinks()) + 1

	_, e = c.SlidesApi.CreateWatermark(fileName, nil, nil, watermarkText, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.GetShapesLinks()))
		return
	}
	getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.GetName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.GetName())
		return
	}
	if getShapeResult.(slidescloud.IShape).GetText() != watermarkText {
		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, getShapeResult.(slidescloud.IShape).GetText())
		return
	}

	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount-1 != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(getResult.GetShapesLinks()))
		return
	}
}

/*
   Test for text DTO watermark on storage
*/
func TestWatermarkTextDTOStorage(t *testing.T) {
	watermarkText := "watermarkText"
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
	getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapeCount := len(getResult.GetShapesLinks()) + 1

	watermark := slidescloud.NewShape()
	watermark.Text = watermarkText
	_, e = c.SlidesApi.CreateWatermark(fileName, watermark, nil, watermarkText, "", "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.GetShapesLinks()))
		return
	}
	getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.GetName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.GetName())
		return
	}
	if getShapeResult.(slidescloud.IShape).GetText() != watermarkText {
		t.Errorf("Wrong shape text. Expected %v but was %v.", watermarkText, getShapeResult.(slidescloud.IShape).GetText())
		return
	}

	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount-1 != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(getResult.GetShapesLinks()))
		return
	}
}

/*
   Test for image watermark on storage
*/
func TestWatermarkImageStorage(t *testing.T) {
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
	getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapeCount := len(getResult.GetShapesLinks()) + 1

	source, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CreateImageWatermark(fileName, source, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.GetShapesLinks()))
		return
	}
	getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.GetName() != "watermark" {
		t.Errorf("Wrong shape name. Expected %v but was %v.", "watermark", getShapeResult.GetName())
		return
	}

	_, e = c.SlidesApi.DeleteWatermark(fileName, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount-1 != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount-1, len(getResult.GetShapesLinks()))
		return
	}
}

/*
   Test for image DTO watermark on storage
*/
func TestWatermarkImageDTOStorage(t *testing.T) {
	watermarkName := "myWatermark"
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
	getResult, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shapeCount := len(getResult.GetShapesLinks()) + 1

	source, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	watermark := slidescloud.NewPictureFrame()
	fillFormat := slidescloud.NewPictureFill()
	fillFormat.Base64Data = base64.StdEncoding.EncodeToString(source)
	watermark.FillFormat = fillFormat
	watermark.Name = watermarkName
	_, e = c.SlidesApi.CreateImageWatermark(fileName, nil, watermark, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount, len(getResult.GetShapesLinks()))
		return
	}
	getShapeResult, _, e := c.SlidesApi.GetShape(fileName, slideIndex, int32(shapeCount), password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if getShapeResult.GetName() != watermarkName {
		t.Errorf("Wrong shape name. Expected %v but was %v.", watermarkName, getShapeResult.GetName())
		return
	}

	_, e = c.SlidesApi.DeleteWatermark(fileName, watermarkName, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	getResult, _, e = c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeCount-1 != len(getResult.GetShapesLinks()) {
		t.Errorf("Wrong shape count. Expected %v but was %v.", shapeCount - 1, len(getResult.GetShapesLinks()))
		return
	}
}

/*
   Test for text watermark from request
*/
func TestWatermarkTextRequest(t *testing.T) {
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

	postResult, _, e := c.SlidesApi.CreateWatermarkOnline(source, nil, nil, "watermark", "", "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for text DTO watermark from request
*/
func TestWatermarkTextDTORequest(t *testing.T) {
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

	watermark := slidescloud.NewShape()
	watermark.Text = "watermarkText"
	postResult, _, e := c.SlidesApi.CreateWatermarkOnline(source, watermark, nil, "", "", "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for image watermark from request
*/
func TestWatermarkImageRequest(t *testing.T) {
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	watermark, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	postResult, _, e := c.SlidesApi.CreateImageWatermarkOnline(source, watermark, nil, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, "", password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}

/*
   Test for image DTO watermark from request
*/
func TestWatermarkImageDTORequest(t *testing.T) {
	watermarkName := "myWatermark"
	source, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	watermark, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewPictureFrame()
	fillFormat := slidescloud.NewPictureFill()
	fillFormat.Base64Data = base64.StdEncoding.EncodeToString(watermark)
	dto.FillFormat = fillFormat
	dto.Name = watermarkName
	postResult, _, e := c.SlidesApi.CreateImageWatermarkOnline(source, nil, dto, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	postStat, e := os.Stat(postResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(source) == int(postStat.Size()) {
		t.Errorf("Wrong file size. Expected not %v but was %v.", len(source), postStat.Size())
		return
	}

	deleteResult, _, e := c.SlidesApi.DeleteWatermarkOnline(source, watermarkName, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	deleteStat, e := os.Stat(deleteResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if deleteStat.Size() >= postStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", postStat.Size(), deleteStat.Size())
		return
	}
}
