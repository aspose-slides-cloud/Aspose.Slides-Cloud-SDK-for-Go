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
   Test for text format 3D
*/
func TestTextFormat3D(t *testing.T) {
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
	dto.ShapeType = "Rectangle"
	dto.X = 100
	dto.Y = 100
	dto.Height = 100
	dto.Width = 200
	dto.Text = "Sample text"

	textFormat := slidescloud.NewTextFrameFormat()
	textFormat.Transform = "ArchUpPour"

	threeDFormat := slidescloud.NewThreeDFormat()
	bevelBottom := slidescloud.NewShapeBevel()
	bevelBottom.BevelType = "Circle"
	bevelBottom.Height = 3.5
	bevelBottom.Width = 3.5
	threeDFormat.BevelBottom = bevelBottom
	bevelTop := slidescloud.NewShapeBevel()
	bevelTop.BevelType = "Circle"
	bevelTop.Height = 4
	bevelTop.Width = 4
	threeDFormat.BevelTop = bevelTop
	threeDFormat.ExtrusionColor = "#FF008000"
	threeDFormat.ExtrusionHeight = 6
	threeDFormat.ContourColor = "#FF25353D"
	threeDFormat.ContourWidth = 1.5
	threeDFormat.Depth = 3
	threeDFormat.Material = "Plastic"
	lightRig := slidescloud.NewLightRig()
	lightRig.LightType = "Balanced"
	lightRig.Direction = "Top"
	lightRig.XRotation = 0
	lightRig.YRotation = 0
	lightRig.ZRotation = 40
	threeDFormat.LightRig = lightRig
	camera := slidescloud.NewCamera()
	camera.CameraType = "PerspectiveContrastingRightFacing"
	threeDFormat.Camera = camera
	textFormat.ThreeDFormat = threeDFormat
	dto.TextFrameFormat = textFormat

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Shape", response.GetType())
		return
	}
}

/*
   Test for text frame format
*/
func TestTextFrameFormatGeneral(t *testing.T) {
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
	dto.ShapeType = "Rectangle"
	dto.X = 100
	dto.Y = 100
	dto.Height = 100
	dto.Width = 200
	dto.Text = "Sample text"

	textFrameFormat := slidescloud.NewTextFrameFormat()
	textFrameFormat.CenterText = "True"
	textFrameFormat.MarginLeft = 2
	textFrameFormat.MarginRight = 2
	textFrameFormat.MarginTop = 2
	textFrameFormat.MarginBottom = 2
	solidFillFormat := slidescloud.NewSolidFill()
	solidFillFormat.Color = "#FF0000"
	defaultParagraphFormat := slidescloud.NewParagraphFormat()
	defaultParagraphFormat.BulletFillFormat = solidFillFormat
	textFrameFormat.DefaultParagraphFormat = defaultParagraphFormat
	dto.TextFrameFormat = textFrameFormat

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Shape", response.GetType())
		return
	}
}
