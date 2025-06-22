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
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

/*
   Test for shape line format
*/
func TestShapeFormatLine(t *testing.T) {
	var shapeIndex int32 = 1

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

	dto := slidescloud.NewShape()
	lineFormat := slidescloud.NewLineFormat()
	lineFormat.Style = "ThickThin"
	lineFormat.Width = 7
	lineFormat.DashStyle = "Dash"
	dto.LineFormat = lineFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetLineFormat().GetWidth() != dto.GetLineFormat().GetWidth() {
		t.Errorf("Wrong line width. Expected %v but was %v.", dto.GetLineFormat().GetWidth(), result.GetLineFormat().GetWidth())
		return
	}
}

/*
   Test for shape fill format
*/
func TestShapeFormatFill(t *testing.T) {
	var shapeIndex int32 = 1

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

	dto := slidescloud.NewShape()
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.Color = "#FFFFFF00"
	dto.FillFormat = fillFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	_, ok = result.GetFillFormat().(slidescloud.ISolidFill)
	if !ok {
		t.Errorf("Wrong fill type.")
		return
	}
	if result.GetFillFormat().(slidescloud.ISolidFill).GetColor() != dto.GetFillFormat().(slidescloud.ISolidFill).GetColor() {
		t.Errorf("Wrong fill color. Expected %v but was %v.", dto.GetFillFormat().(slidescloud.ISolidFill).GetColor(), result.GetFillFormat().(slidescloud.ISolidFill).GetColor())
		return
	}
}

/*
   Test for shape picture fill format
*/
func TestShapeFormatPictureFill(t *testing.T) {
	var shapeIndex int32 = 1

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

	dto := slidescloud.NewShape()
	fillFormat := slidescloud.NewPictureFill()
	picture, e := ioutil.ReadFile(localFolder + "/watermark.png")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	fillFormat.Base64Data = base64.StdEncoding.EncodeToString(picture)
	fillFormat.Resolution = 150
	dto.FillFormat = fillFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	_, ok = result.GetFillFormat().(slidescloud.IPictureFill)
	if !ok {
		t.Errorf("Wrong fill type.")
		return
	}
}

/*
   Test for shape effect format
*/
func TestShapeFormatEffect(t *testing.T) {
	var shapeIndex int32 = 1

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

	dto := slidescloud.NewShape()
	effectFormat := slidescloud.NewEffectFormat()
	innerShadow := slidescloud.NewInnerShadowEffect()
	innerShadow.Direction = 35
	innerShadow.BlurRadius = 30
	innerShadow.Distance = 40
	innerShadow.ShadowColor = "#FFFFFF00"
	effectFormat.InnerShadow = innerShadow
	dto.EffectFormat = effectFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetEffectFormat().GetInnerShadow().GetDirection() != dto.GetEffectFormat().GetInnerShadow().GetDirection() {
		t.Errorf("Wrong inner shadow direction. Expected %v but was %v.", dto.GetEffectFormat().GetInnerShadow().GetDirection(), result.GetEffectFormat().GetInnerShadow().GetDirection())
		return
	}
}

/*
   Test for shape 3D format
*/
func TestShapeFormat3D(t *testing.T) {
	var shapeIndex int32 = 1

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

	dto := slidescloud.NewShape()
	threeDFormat := slidescloud.NewThreeDFormat()
	threeDFormat.Depth = 4

	bevelTop := slidescloud.NewShapeBevel()
	bevelTop.BevelType = "Circle"
	bevelTop.Height = 6
	bevelTop.Width = 6
	threeDFormat.BevelTop = bevelTop

	camera := slidescloud.NewCamera()
	camera.CameraType = "OrthographicFront"
	threeDFormat.Camera = camera

	lightRig := slidescloud.NewLightRig()
	lightRig.LightType = "ThreePt"
	lightRig.Direction = "Top"
	threeDFormat.LightRig = lightRig
	dto.ThreeDFormat = threeDFormat
	result, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	result, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok = result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
	if result.GetThreeDFormat().GetDepth() != dto.GetThreeDFormat().GetDepth() {
		t.Errorf("Wrong 3D depth. Expected %v but was %v.", dto.GetThreeDFormat().GetDepth(), result.GetThreeDFormat().GetDepth())
		return
	}
}
