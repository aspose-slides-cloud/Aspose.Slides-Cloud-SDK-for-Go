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
	"math"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v26"
)

/*
   Test for get shapes
*/
func TestGetShapes(t *testing.T) {
	var slideIndex int32 = 3

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

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get shapes by type
*/
func TestGetShapesByType(t *testing.T) {
	var slideIndex int32 = 3

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

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "Chart", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get sub-shapes
*/
func TestGetSubshapes(t *testing.T) {
	var slideIndex int32 = 1
	subShape := "4"

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

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", subShape)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, shapes.GetShapesLinks())
		return
	}
}

/*
   Test for get shape
*/
func TestGetShape(t *testing.T) {
	var slideIndex int32 = 3
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

	shape, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if shape.GetType() != "Chart" {
		t.Errorf("Expected %v, but was %v", "Chart", shape.GetType())
		return
	}
}

/*
   Test for get sub-shape
*/
func TestGetSubshape(t *testing.T) {
	var slideIndex int32 = 1
	var shapeIndex int32 = 4
	subShape := "1"

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

	shape, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if shape.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Chart", shape.GetType())
		return
	}
}

/*
   Test for load and save shape
*/
func TestShapeLoadAndSave(t *testing.T) {
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

	dto, _, e := c.SlidesApi.GetShape(fileName, 3, 1, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape, _, e := c.SlidesApi.UpdateShape(fileName, 3, 1, dto, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := shape.(slidescloud.IChart)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add shape
*/
func TestShapeAdd(t *testing.T) {
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
	dto.ShapeType = "Callout1"
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty shape
*/
func TestShapeEmpty(t *testing.T) {
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
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("Shape with undefined type should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty graphical object
*/
func TestGraphicalObjectEmpty(t *testing.T) {
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

	dto := slidescloud.NewGraphicalObject()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("GraphicalObject should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add picture frame
*/
func TestPictureFrameAdd(t *testing.T) {
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

	dto := slidescloud.NewPictureFrame()
	fill := slidescloud.NewPictureFill()
	fill.Base64Data = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAANSURBVBhXY5g+ffp/AAZTAsWGL27gAAAAAElFTkSuQmCC"
	dto.PictureFillFormat = fill
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IPictureFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty picture frame
*/
func TestPictureFrameEmpty(t *testing.T) {
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

	dto := slidescloud.NewPictureFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("PictureFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add audio frame
*/
func TestAudioFrameAdd(t *testing.T) {
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

	dto := slidescloud.NewAudioFrame()
	dto.Base64Data = "bXAzc2FtcGxl"
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IAudioFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty audio frame
*/
func TestAudioFrameEmpty(t *testing.T) {
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

	dto := slidescloud.NewAudioFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("AudioFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add video frame
*/
func TestVideoFrameAdd(t *testing.T) {
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

	dto := slidescloud.NewVideoFrame()
	dto.Base64Data = "bXAzc2FtcGxl"
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IVideoFrame)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty video frame
*/
func TestVideoFrameEmpty(t *testing.T) {
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

	dto := slidescloud.NewVideoFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("VideoFrame with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty OLE object frame
*/
func TestOleObjectFrameEmpty(t *testing.T) {
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

	dto := slidescloud.NewOleObjectFrame()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("OleObjectFrame should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add smart art
*/
func TestSmartArtAdd(t *testing.T) {
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

	dto := slidescloud.NewSmartArt()
	dto.X = 0
	dto.Y = 0
	dto.Width = 300
	dto.Height = 200
	dto.Layout = "BasicProcess"
	dto.QuickStyle = "SimpleFill"
	dto.ColorStyle = "ColoredFillAccent1"
	node1 := slidescloud.NewSmartArtNode()
	node1.Text = "First"
	node1.OrgChartLayout = "Initial"
	subNode1 := slidescloud.NewSmartArtNode()
	subNode1.Text = "SubFirst"
	subNode1.OrgChartLayout = "Initial"
	node1.Nodes = []slidescloud.ISmartArtNode{subNode1}
	node2 := slidescloud.NewSmartArtNode()
	node2.Text = "Second"
	node2.OrgChartLayout = "Initial"
	dto.Nodes = []slidescloud.ISmartArtNode{node1, node2}
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.ISmartArt)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add smart art text formatting
*/
func TestSmartArtTextFormatting(t *testing.T) {
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

	portion := slidescloud.NewPortion()
	portion.SetText("New text")
	portion.SetFontHeight(24)
	portion.SetFontBold("True")
	portion.SetSpacing(3)
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.SetColor("#FFFFFF00")
	portion.FillFormat = fillFormat

	targetNodePath := "1/nodes/2"
	var slideIndex int32 = 7
	var shapeIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	response, _, e := c.SlidesApi.UpdatePortion(fileName, slideIndex, shapeIndex, paragraphIndex,
		portionIndex, portion, password, folderName, "", targetNodePath)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetText() != portion.GetText() {
		t.Errorf("Expected %v, but was %v", portion.GetText(), response.GetText())
		return
	}

	if response.GetFontHeight() != portion.GetFontHeight() {
		t.Errorf("Expected %v, but was %v", portion.GetFontHeight(), response.GetFontHeight())
		return
	}

	if response.GetFontBold() != portion.GetFontBold() {
		t.Errorf("Expected %v, but was %v", portion.GetFontBold(), response.GetFontBold())
		return
	}

	if response.GetSpacing() != portion.GetSpacing() {
		t.Errorf("Expected %v, but was %v", portion.GetSpacing(), response.GetSpacing())
		return
	}

	if response.GetFillFormat().(slidescloud.ISolidFill).GetColor() != fillFormat.GetColor() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetColor(), response.GetFillFormat().(slidescloud.ISolidFill).GetColor())
		return
	}
}

/*
   Test for add empty smart art
*/
func TestSmartArtEmpty(t *testing.T) {
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

	dto := slidescloud.NewSmartArt()
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.ISmartArt)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty chart
   See Chart tests for non-empty chart examples
*/
func TestChartEmpty(t *testing.T) {
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

	dto := slidescloud.NewChart()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("Chart with undefined data should not have been created.")
		return
	}
	if response.StatusCode != 500 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add table
*/
func TestTableAdd(t *testing.T) {
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

	dto := slidescloud.NewTable()
	dto.X = 30
	dto.Y = 20
	dto.Style = "MediumStyle2Accent1"
	row1 := slidescloud.NewTableRow()
	cell11 := slidescloud.NewTableCell()
	cell11.Text = "0.1"
	cell12 := slidescloud.NewTableCell()
	cell12.Text = "0.2"
	cell13 := slidescloud.NewTableCell()
	cell13.Text = "0.3"
	cell14 := slidescloud.NewTableCell()
	cell14.Text = "0.4"
	row1.Cells = []slidescloud.ITableCell{cell11, cell12, cell13, cell14}
	row2 := slidescloud.NewTableRow()
	cell21 := slidescloud.NewTableCell()
	cell21.Text = "1"
	cell22 := slidescloud.NewTableCell()
	cell22.Text = "2-3"
	cell22.ColSpan = 2
	cell22.RowSpan = 2
	cell24 := slidescloud.NewTableCell()
	cell24.Text = "4"
	row2.Cells = []slidescloud.ITableCell{cell21, cell22, cell24}
	row3 := slidescloud.NewTableRow()
	cell31 := slidescloud.NewTableCell()
	cell31.Text = "first"
	cell32 := slidescloud.NewTableCell()
	cell32.Text = "last"
	row3.Cells = []slidescloud.ITableCell{cell31, cell32}
	row4 := slidescloud.NewTableRow()
	cell41 := slidescloud.NewTableCell()
	cell41.Text = "3.1"
	cell42 := slidescloud.NewTableCell()
	cell42.Text = "3.2"
	cell43 := slidescloud.NewTableCell()
	cell43.Text = "3.3"
	cell44 := slidescloud.NewTableCell()
	cell44.Text = "3.4"
	row4.Cells = []slidescloud.ITableCell{cell41, cell42, cell43, cell44}
	row5 := slidescloud.NewTableRow()
	cell51 := slidescloud.NewTableCell()
	cell51.Text = "4.1"
	cell52 := slidescloud.NewTableCell()
	cell52.Text = "4.2"
	cell53 := slidescloud.NewTableCell()
	cell53.Text = "4.3"
	cell54 := slidescloud.NewTableCell()
	cell54.Text = "4.4"
	row5.Cells = []slidescloud.ITableCell{cell51, cell52, cell53, cell54}
	dto.Rows = []slidescloud.ITableRow{row1, row2, row3, row4, row5}
	column1 := slidescloud.NewTableColumn()
	column1.Width = 100
	column2 := slidescloud.NewTableColumn()
	column2.Width = 100
	column3 := slidescloud.NewTableColumn()
	column3.Width = 100
	column4 := slidescloud.NewTableColumn()
	column4.Width = 100
	dto.Columns = []slidescloud.ITableColumn{column1, column2, column3, column4}
	firstRow := true
	dto.FirstRow = &firstRow
	horizontalBanding := true
	dto.HorizontalBanding = &horizontalBanding
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.ITable)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty table
*/
func TestTableEmpty(t *testing.T) {
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

	dto := slidescloud.NewTable()
	_, response, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e == nil {
		t.Errorf("Table with undefinined cell data should not have been created.")
		return
	}
	if response.StatusCode != 400 {
		t.Errorf("Wrong status code. Expected 400 but was %v.", response.StatusCode)
		return
	}
}

/*
   Test for add empty group shape
*/
func TestGroupShapeEmpty(t *testing.T) {
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

	dto := slidescloud.NewGroupShape()
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IGroupShape)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add connector
*/
func TestConnectorAdd(t *testing.T) {
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

	dto := slidescloud.NewConnector()
	dto.ShapeType = "BentConnector3"
	start := slidescloud.NewResourceUri()
	start.Href = "https://api.aspose.cloud/v3.0/slides/myPresentation.pptx/slides/1/shapes/1"
	dto.StartShapeConnectedTo = start
	end := slidescloud.NewResourceUri()
	end.Href = "https://api.aspose.cloud/v3.0/slides/myPresentation.pptx/slides/1/shapes/2"
	dto.EndShapeConnectedTo = end
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IConnector)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for add empty connector
*/
func TestConnectorEmpty(t *testing.T) {
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

	dto := slidescloud.NewConnector()
	result, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, ok := result.(slidescloud.IConnector)
	if !ok {
		t.Errorf("Wrong shape type.")
		return
	}
}

/*
   Test for create shubshape
*/
func TestCreateSubshape(t *testing.T) {
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
	dto.SetShapeType("Rectangle")
	dto.SetX(200)
	dto.SetY(200)
	dto.SetWidth(50)
	dto.SetHeight(50)

	subShape := "4"
	shapeBase, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName,
		"", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shapeBase.GetType() != "Shape" {
		t.Errorf("Expected %v, but was %v", "Shape", shapeBase.GetType())
		return
	}
}

/*
   Test for update shape
*/
func TestUpdateShape(t *testing.T) {
	var shapeIndex int32 = 3

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
	dto.SetWidth(200)
	dto.SetHeight(200)
	fillFormat := slidescloud.NewSolidFill()
	fillFormat.SetColor("#FFF5FF8A")
	dto.FillFormat = fillFormat

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetWidth() != dto.GetWidth() {
		t.Errorf("Expected %v, but was %v", dto.GetWidth(), response.GetWidth())
		return
	}

	if response.GetHeight() != dto.GetHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetHeight(), response.GetHeight())
		return
	}

	if response.GetFillFormat().(slidescloud.ISolidFill).GetColor() != fillFormat.GetColor() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetColor(), response.GetFillFormat().(slidescloud.ISolidFill).GetColor())
		return
	}
}

/*
   Test for update sub-shape
*/
func TestUpdateSubshape(t *testing.T) {
	var shapeIndex int32 = 4
	subShape := "1"

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
	dto.SetWidth(200)
	dto.SetHeight(200)
	fillFormat := slidescloud.NewGradientFill()
	fillFormat.SetShape("Linear")
	fillFormat.SetDirection("FromCorner1")
	gradientStop1 := slidescloud.NewGradientFillStop()
	gradientStop1.Color = "#FFF5FF8A"
	gradientStop1.SetPosition(0)
	gradientStop2 := slidescloud.NewGradientFillStop()
	gradientStop2.Color = "#FFF5FF8A"
	gradientStop2.SetPosition(1)
	fillFormat.SetStops([]slidescloud.IGradientFillStop{gradientStop1, gradientStop2})
	dto.SetFillFormat(fillFormat)

	response, _, e := c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, dto, password, folderName, "", subShape)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetWidth() != dto.GetWidth() {
		t.Errorf("Expected %v, but was %v", dto.GetWidth(), response.GetWidth())
		return
	}

	if response.GetHeight() != dto.GetHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetHeight(), response.GetHeight())
		return
	}

	if response.GetFillFormat().GetType() != fillFormat.GetType() {
		t.Errorf("Expected %v, but was %v", fillFormat.GetType(), response.GetFillFormat().GetType())
		return
	}
}

/*
   Test for delete shapes
*/
func TestDeleteShapes(t *testing.T) {
	var slideIndex int32 = 3

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

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete shapes by indexes
*/
func TestDeleteShapesIndexes(t *testing.T) {
	var slideIndex int32 = 3
	shapesIndexes := []int32{2}

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

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, shapesIndexes, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shapes
*/
func TestDeleteSubshapes(t *testing.T) {
	subShape := "4"

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

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, nil, password, folderName, "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shapes by indexes
*/
func TestDeleteSubshapesIndexes(t *testing.T) {
	subShape := "4"
	shapesIndexes := []int32{2}

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

	response, _, e := c.SlidesApi.DeleteShapes(fileName, slideIndex, shapesIndexes, password, folderName, "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete shape
*/
func TestDeleteShape(t *testing.T) {
	var shapeIndex int32 = 4

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

	response, _, e := c.SlidesApi.DeleteShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for delete sub-shape
*/
func TestDeleteSubshape(t *testing.T) {
	var shapeIndex int32 = 4
	subShape := "1"

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

	response, _, e := c.SlidesApi.DeleteShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(response.GetShapesLinks()))
		return
	}
}

/*
   Test for align shapes
*/
func TestShapesAlign(t *testing.T) {
	var slideIndex int32 = 3
	var shape1Index int32 = 1
	var shape2Index int32 = 2
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

	shape11, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape12, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape11.GetX() == shape12.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape11.GetX(), shape12.GetX())
		return
	}
	if shape11.GetY() == shape12.GetY() {
		t.Errorf("Wrong Y value. Expected not %v but was %v.", shape11.GetY(), shape12.GetY())
		return
	}

	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignTop", nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape21, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape22, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape21.GetX() == shape22.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape21.GetX(), shape22.GetX())
		return
	}
	if math.Abs(shape21.GetY()-shape22.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape21.GetY(), shape22.GetY())
		return
	}

	var alignToSlide bool = true
	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignLeft", &alignToSlide, []int32{1, 2}, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape31, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape1Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape32, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shape2Index, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if math.Abs(shape31.GetX()-shape32.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", shape31.GetX(), shape32.GetX())
		return
	}
	if math.Abs(shape31.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", 0, shape31.GetX())
		return
	}
	if math.Abs(shape31.GetY()-shape32.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape31.GetY(), shape32.GetY())
		return
	}
}

/*
   Test for align shapes in group
*/
func TestShapesAlignGroup(t *testing.T) {
	var shapeIndex int32 = 4
	var subShape1 = "1"
	var subShape2 = "2"
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

	shape11, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape1)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape12, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape2)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape11.GetX() == shape12.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape11.GetX(), shape12.GetX())
		return
	}
	if shape11.GetY() == shape12.GetY() {
		t.Errorf("Wrong Y value. Expected not %v but was %v.", shape11.GetY(), shape12.GetY())
		return
	}

	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignTop", nil, nil, password, folderName, "", "4")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape21, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape1)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape22, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape2)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape21.GetX() == shape22.GetX() {
		t.Errorf("Wrong X value. Expected not %v but was %v.", shape21.GetX(), shape22.GetX())
		return
	}
	if math.Abs(shape21.GetY()-shape22.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape21.GetY(), shape22.GetY())
		return
	}

	var alignToSlide bool = true
	_, _, e = c.SlidesApi.AlignShapes(fileName, slideIndex, "AlignLeft", &alignToSlide, []int32{1, 2}, password, folderName, "", "4")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape31, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape1)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	shape32, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", subShape2)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if math.Abs(shape31.GetX()-shape32.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", shape31.GetX(), shape32.GetX())
		return
	}
	if math.Abs(shape31.GetX()) >= 1 {
		t.Errorf("Wrong X value. Expected %v but was %v.", 0, shape31.GetX())
		return
	}
	if math.Abs(shape31.GetY()-shape32.GetY()) >= 1 {
		t.Errorf("Wrong Y value. Expected %v but was %v.", shape31.GetY(), shape32.GetY())
		return
	}
}

/*
   Test for Get shape geometry paths
*/
func TestShapeGeometryGet(t *testing.T) {
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

	paths, _, e := c.SlidesApi.GetShapeGeometryPath(fileName, 4, 2, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if paths.GetPaths() == nil {
		t.Errorf("Error: paths should not be nil.")
		return
	}
	if len(paths.GetPaths()) != 1 {
		t.Errorf("Wrong paths count. Expected %v but was %v.", 1, len(paths.GetPaths()))
		return
	}
}

/*
   Test for set shape geometry paths
*/
func TestShapeGeometrySet(t *testing.T) {
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

	dto := slidescloud.NewGeometryPaths()
	path := slidescloud.NewGeometryPath()
	startSegment := slidescloud.NewMoveToPathSegment()
	startSegment.X = 0
	startSegment.Y = 0
	line1 := slidescloud.NewLineToPathSegment()
	line1.X = 0
	line1.Y = 200
	line2 := slidescloud.NewLineToPathSegment()
	line2.X = 200
	line2.Y = 300
	line3 := slidescloud.NewLineToPathSegment()
	line3.X = 400
	line3.Y = 200
	line4 := slidescloud.NewLineToPathSegment()
	line4.X = 400
	line4.Y = 0
	endSegment := slidescloud.NewClosePathSegment()
	path.PathData = []slidescloud.IPathSegment{startSegment, line1, line2, line3, line4, endSegment}
	dto.Paths = []slidescloud.IGeometryPath{path}
	shape, _, e := c.SlidesApi.SetShapeGeometryPath(fileName, 4, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if shape == nil {
		t.Errorf("Error: shape should not be nil.")
		return
	}
}

/*
   Test for Zoom Frame add
*/
func TestZoomFrameAdd(t *testing.T) {
	var slideIndex int32 = 3

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

	dto := slidescloud.NewZoomFrame()
	dto.SetX(20)
	dto.SetY(20)
	dto.SetWidth(200)
	dto.SetHeight(100)
	dto.SetTargetSlideIndex(2)

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "ZoomFrame" {
		t.Errorf("Expected %v, but was %v", "ZoomFrame", response.GetType())
		return
	}

	if response.(slidescloud.IZoomFrame).GetTargetSlideIndex() != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.(slidescloud.IZoomFrame).GetTargetSlideIndex())
		return
	}
}

/*
   Test for Section Zoom Frame add
*/
func TestSectionZoomFrameAdd(t *testing.T) {
	var slideIndex int32 = 3

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

	dto := slidescloud.NewSectionZoomFrame()
	dto.SetX(20)
	dto.SetY(20)
	dto.SetWidth(200)
	dto.SetHeight(100)
	dto.SetTargetSectionIndex(2)

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "SectionZoomFrame" {
		t.Errorf("Expected %v, but was %v", "SectionZoomFrame", response.GetType())
		return
	}

	if response.(slidescloud.ISectionZoomFrame).GetTargetSectionIndex() != 2 {
		t.Errorf("Expected %v, but was %v", 2, response.(slidescloud.ISectionZoomFrame).GetTargetSectionIndex())
		return
	}
}

/*
   Test for OleObjectFrame add by link
*/
func TestOleObjectFrameAddByLink(t *testing.T) {
	var slideIndex int32 = 3

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

	dto := slidescloud.NewOleObjectFrame()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(200)
	dto.SetHeight(200)
	dto.SetLinkPath("OleObject.xlsx")
	dto.SetObjectProgId("Excel.Sheet.8")

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "OleObjectFrame" {
		t.Errorf("Expected %v, but was %v", "OleObjectFrame", response.GetType())
		return
	}

	if response.(slidescloud.IOleObjectFrame).GetLinkPath() != dto.GetLinkPath() {
		t.Errorf("Expected %v, but was %v", dto.GetLinkPath(), response.(slidescloud.IOleObjectFrame).GetLinkPath())
		return
	}
}

/*
   Test for OleObjectFrame add embedded
*/
func TestOleObjectFrameAddEmbedded(t *testing.T) {
	var slideIndex int32 = 3

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

	source, e := ioutil.ReadFile(localFolder + "/oleObject.xlsx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewOleObjectFrame()
	dto.SetX(100)
	dto.SetY(100)
	dto.SetWidth(200)
	dto.SetHeight(200)
	dto.SetEmbeddedFileBase64Data(base64.StdEncoding.EncodeToString(source))
	dto.SetEmbeddedFileExtension("xlsx")

	response, _, e := c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if response.GetType() != "OleObjectFrame" {
		t.Errorf("Expected %v, but was %v", "OleObjectFrame", response.GetType())
		return
	}

	if response.(slidescloud.IOleObjectFrame).GetEmbeddedFileBase64Data() == "" {
		t.Errorf("EmbeddedFileBase64Data is not defined.")
		return
	}
}

/*
   Test for GroupShape add
*/
func TestGroupShapeAdd(t *testing.T) {
	var slideIndex int32 = 5
	var subShape = "1"

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

	dto := slidescloud.NewGroupShape()
	c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")

	shape1 := slidescloud.NewShape()
	shape1.SetShapeType("Rectangle")
	shape1.SetX(50)
	shape1.SetY(400)
	shape1.SetWidth(50)
	shape1.SetHeight(50)

	shape2 := slidescloud.NewShape()
	shape2.SetShapeType("Ellipse")
	shape2.SetX(150)
	shape2.SetY(400)
	shape2.SetWidth(50)
	shape2.SetHeight(50)

	shape3 := slidescloud.NewShape()
	shape3.SetShapeType("Triangle")
	shape3.SetX(250)
	shape3.SetY(400)
	shape3.SetWidth(50)
	shape3.SetHeight(50)

	c.SlidesApi.CreateShape(fileName, slideIndex, shape1, nil, nil, password, folderName, "", subShape)
	c.SlidesApi.CreateShape(fileName, slideIndex, shape2, nil, nil, password, folderName, "", subShape)
	c.SlidesApi.CreateShape(fileName, slideIndex, shape3, nil, nil, password, folderName, "", subShape)

	shapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(shapes.GetShapesLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(shapes.GetShapesLinks()))
		return
	}

	subShapes, _, e := c.SlidesApi.GetShapes(fileName, slideIndex, password, folderName, "", "", subShape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(subShapes.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(subShapes.GetShapesLinks()))
		return
	}
}

/*
   Test for import shapes from svg
*/
func TestImportShapesFromSvg(t *testing.T) {
	var slideIndex int32 = 5
	var x int32 = 50
	var y int32 = 50
	var width int32 = 300
	var height int32 = 300

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

	document, e := ioutil.ReadFile(localFolder + "/shapes.svg")
	shapes := []int32{1, 3, 5}
	var groupShape bool = false
	response, _, e := c.SlidesApi.ImportShapesFromSvg(fileName, slideIndex, document, &x, &y, &width, &height, shapes, &groupShape, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetShapesLinks()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, len(response.GetShapesLinks()))
		return
	}
}

func TestCreateSmartArtNode(t *testing.T) {
	var slideIndex int32 = 7
	var smartArtIndex int32 = 1
	newNodeText := "New root node"

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

	response, _, e := c.SlidesApi.CreateSmartArtNode(fileName, slideIndex, smartArtIndex, "", newNodeText, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetNodes()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetNodes()))
		return
	}

	if response.GetNodes()[1].GetText() != newNodeText {
		t.Errorf("Expected %v, but was %v", newNodeText, response.GetNodes()[0].GetText())
		return
	}
}

func TestCreateSmartArtSubNode(t *testing.T) {
	var slideIndex int32 = 7
	var smartArtIndex int32 = 1
	newSubNodeText := "New sub-node"
	subNodePath := "1"
	var position int32 = 1

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

	response, _, e := c.SlidesApi.CreateSmartArtNode(fileName, slideIndex, smartArtIndex, subNodePath, newSubNodeText,
		&position, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetNodes()[0].GetNodes()) != 5 {
		t.Errorf("Expected %v, but was %v", 5, len(response.GetNodes()[0].GetNodes()))
		return
	}

	if response.GetNodes()[0].GetNodes()[0].GetText() != newSubNodeText {
		t.Errorf("Expected %v, but was %v", newSubNodeText, response.GetNodes()[0].GetNodes()[0].GetText())
		return
	}
}

func TestCreateSmartArtSubSubNode(t *testing.T) {
	var slideIndex int32 = 7
	var smartArtIndex int32 = 1
	newSubNodeText := "New sub-sub-node"
	subNodePath := "1/nodes/1"

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

	response, _, e := c.SlidesApi.CreateSmartArtNode(fileName, slideIndex, smartArtIndex, subNodePath, newSubNodeText,
		nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetNodes()[0].GetNodes()[0].GetNodes()) != 1 {
		t.Errorf("Expected %v, but was %v", 5, len(response.GetNodes()[0].GetNodes()[0].GetNodes()))
		return
	}

	if response.GetNodes()[0].GetNodes()[0].GetNodes()[0].GetText() != newSubNodeText {
		t.Errorf("Expected %v, but was %v", newSubNodeText, response.GetNodes()[0].GetNodes()[0].GetNodes()[0].GetText())
		return
	}
}

func TestDeleteSmartArtNode(t *testing.T) {
	var slideIndex int32 = 7
	var smartArtIndex int32 = 2
	var nodeIndex int32 = 1

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

	response, _, e := c.SlidesApi.DeleteSmartArtNode(fileName, slideIndex, smartArtIndex, nodeIndex, "", password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetNodes()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(response.GetNodes()))
		return
	}
}

func TestDeleteSmartArtSubNode(t *testing.T) {
	var slideIndex int32 = 7
	var smartArtIndex int32 = 1
	var nodeIndex int32 = 1
	subNodePath := "2"

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

	response, _, e := c.SlidesApi.DeleteSmartArtNode(fileName, slideIndex, smartArtIndex, nodeIndex, subNodePath, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetNodes()[0].GetNodes()) != 3 {
		t.Errorf("Expected %v, but was %v", 3, response.GetNodes()[0].GetNodes())
		return
	}
}

func TestDownloadShapeFromDto(t *testing.T) {
	shape := slidescloud.NewShape()
	shape.SetShapeType("Rectangle")
	shape.SetWidth(400)
	shape.SetHeight(200)
	shape.SetText("Shape text")

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.DownloadShapeFromDto("png", shape)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}