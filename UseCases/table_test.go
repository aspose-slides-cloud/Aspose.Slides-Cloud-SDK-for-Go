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
   Update table cell
*/
func TestUpdateTableCell(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 2
	var cellIndex int32 = 1
	dto := slidescloud.NewTableCell()
	dto.Text = "Test text"

	result, _, e := c.SlidesApi.UpdateTableCell(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), result.GetText())
		return
	}
}

func TestCreateTableRow(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	cell0 := slidescloud.NewTableCell()
	cell0.Text = "Cell 1"
	cell1 := slidescloud.NewTableCell()
	cell1.Text = "Cell 2"
	cell2 := slidescloud.NewTableCell()
	cell2.Text = "Cell 3"
	cell3 := slidescloud.NewTableCell()
	cell3.Text = "Cell 4"

	dto := slidescloud.NewTableRow()
	dto.MinimalHeight = 30
	dto.Cells = []slidescloud.ITableCell{cell0, cell1, cell2, cell3}

	result, _, e := c.SlidesApi.CreateTableRow(fileName, slideIndex, shapeIndex, dto, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetCells()) != len(dto.GetCells()) {
		t.Errorf("Expected %v, but was %v", len(dto.GetCells()), len(result.GetCells()))
		return
	}
}

func TestDeleteTableRow(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 2
	var withAttachedRows bool = true

	result, _, e := c.SlidesApi.DeleteTableRow(fileName, slideIndex, shapeIndex, rowIndex, &withAttachedRows, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetRows()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetRows()))
		return
	}
}

func TestUpdateTableRow(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	dto := slidescloud.NewTableRow()
	dto.SetMinimalHeight(30)

	result, _, e := c.SlidesApi.UpdateTableRow(fileName, slideIndex, shapeIndex, rowIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetMinimalHeight() != dto.GetMinimalHeight() {
		t.Errorf("Expected %v, but was %v", dto.GetMinimalHeight(), result.GetMinimalHeight())
		return
	}
}

func TestMergeTableCells(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1

	dto := slidescloud.NewTableCellMergeOptions()
	dto.FirstRowIndex = 1
	dto.FirstCellIndex = 1
	dto.LastRowIndex = 2
	dto.LastCellIndex = 2

	result, _, e := c.SlidesApi.MergeTableCells(fileName, slideIndex, shapeIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetRows()[0].GetCells()[0].GetColSpan() != 2 {
		t.Errorf("Expected %v, but was %v", 2, result.GetRows()[0].GetCells()[0].GetColSpan())
		return
	}
	if result.GetRows()[0].GetCells()[0].GetRowSpan() != 2 {
		t.Errorf("Expected %v, but was %v", 2, result.GetRows()[0].GetCells()[0].GetRowSpan())
		return
	}
}

func TestSplitTableCellsByWidth(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var cellWidth int32 = 10

	result, _, e := c.SlidesApi.SplitTableCell(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, "SplitByWidth", float64(cellWidth),
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetRows()[0].GetCells()) != 5 {
		t.Errorf("Expected %v, but was %v", 5, len(result.GetRows()[0].GetCells()))
		return
	}
}

func TestSplitTableCellsByHeight(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var cellHeight int32 = 10

	result, _, e := c.SlidesApi.SplitTableCell(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, "SplitByHeight", float64(cellHeight),
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetRows()) != 5 {
		t.Errorf("Expected %v, but was %v", 5, len(result.GetRows()))
		return
	}
}

func TestSplitTableCellsByColSpan(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 3
	var cellIndex int32 = 1
	var colSpan int32 = 1

	result, _, e := c.SlidesApi.SplitTableCell(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, "SplitByColSpan", float64(colSpan),
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetRows()[2].GetCells()[0].GetColSpan() != 0 {
		t.Errorf("Expected %v, but was %v", 0, result.GetRows()[2].GetCells()[0].GetColSpan())
		return
	}
}

func TestSplitTableCellsByRowSpan(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 2
	var cellIndex int32 = 3
	var colSpan int32 = 1

	result, _, e := c.SlidesApi.SplitTableCell(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, "SplitByRowSpan", float64(colSpan),
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetRows()[1].GetCells()[2].GetRowSpan() != 0 {
		t.Errorf("Expected %v, but was %v", 0, result.GetRows()[1].GetCells()[2].GetRowSpan())
		return
	}
}

func TestGetTableCellParagraphs(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1

	result, _, e := c.SlidesApi.GetTableCellParagraphs(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetParagraphLinks()) != 1 {
		t.Errorf("Expected %v, but was %v", 1, len(result.GetParagraphLinks()))
		return
	}
}

func TestGetTableCellParagraph(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1

	result, _, e := c.SlidesApi.GetTableCellParagraph(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetPortionList()))
		return
	}
}

func TestCreateTableCellParagraph(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	portion0 := slidescloud.NewPortion()
	portion0.SetText("Portion 1")
	portion1 := slidescloud.NewPortion()
	portion1.SetText("Portion 2")
	dto := slidescloud.NewParagraph()
	dto.SetPortionList([]slidescloud.IPortion{portion0, portion1})

	result, _, e := c.SlidesApi.CreateTableCellParagraph(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetPortionList()))
		return
	}
}

func TestUpdateTableCellParagraph(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1
	portion0 := slidescloud.NewPortion()
	portion0.SetText("Portion 1")
	portion1 := slidescloud.NewPortion()
	portion1.SetText("Portion 2")
	dto := slidescloud.NewParagraph()
	dto.SetPortionList([]slidescloud.IPortion{portion0, portion1})

	result, _, e := c.SlidesApi.UpdateTableCellParagraph(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex,
		dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetPortionList()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetPortionList()))
		return
	}
}

func TestDeleteTableCellParagraph(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1

	result, _, e := c.SlidesApi.DeleteTableCellParagraph(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex,
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetParagraphLinks()) != 0 {
		t.Errorf("Expected %v, but was %v", 0, len(result.GetParagraphLinks()))
		return
	}
}

func TestGetTableCellPortions(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1

	result, _, e := c.SlidesApi.GetTableCellPortions(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex,
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetItems()) != 2 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetItems()))
		return
	}
}

func TestGetTableCellPortion(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	result, _, e := c.SlidesApi.GetTableCellPortion(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex, portionIndex,
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetText() != "Header" {
		t.Errorf("Expected %v, but was %v", "Header", result.GetText())
		return
	}
}

func TestCreateTableCellPortion(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1

	dto := slidescloud.NewPortion()
	dto.SetText("Portion 1")

	result, _, e := c.SlidesApi.CreateTableCellPortion(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex, dto,
		password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), result.GetText())
		return
	}
}

func TestUpdateTableCellPortion(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	dto := slidescloud.NewPortion()
	dto.SetText("Portion 1")

	result, _, e := c.SlidesApi.UpdateTableCellPortion(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex,
		portionIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if result.GetText() != dto.GetText() {
		t.Errorf("Expected %v, but was %v", dto.GetText(), result.GetText())
		return
	}
}

func TestDeleteTableCellPortion(t *testing.T) {
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

	var slideIndex int32 = 9
	var shapeIndex int32 = 1
	var rowIndex int32 = 1
	var cellIndex int32 = 1
	var paragraphIndex int32 = 1
	var portionIndex int32 = 1

	dto := slidescloud.NewPortion()
	dto.SetText("Portion 1")

	result, _, e := c.SlidesApi.DeleteTableCellPortion(fileName, slideIndex, shapeIndex, rowIndex, cellIndex, paragraphIndex,
		portionIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(result.GetItems()) != 1 {
		t.Errorf("Expected %v, but was %v", 2, len(result.GetItems()))
		return
	}
}
