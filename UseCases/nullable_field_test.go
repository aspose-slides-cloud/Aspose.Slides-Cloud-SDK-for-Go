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

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

/*
 * Test for nullable properties
 */
func TestDefaultChartValues(t *testing.T) {
	var shapeIndex int32 = 5
	var min1 = 44.3
	var min2 = 12.0
	var max1 = 104.3
	var max2 = 87.0

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

	var dto1 slidescloud.Chart
	dto1.SetType("Chart")
	dto1.SetChartType("Line")
	dto1.SetWidth(400)
	dto1.SetHeight(300)
	hasTitle := true
	dto1.SetHasTitle(&hasTitle)
	var title slidescloud.ChartTitle
	title.SetText("MyTitle")
	dto1.SetTitle(&title)
	var series slidescloud.OneValueSeries
	series.SetType("ClusteredColumn")
	series.SetDataPointType("OneValue")
	series.SetName("Series1")
	var point1 slidescloud.OneValueChartDataPoint
	point1.SetValue(40)
	var point2 slidescloud.OneValueChartDataPoint
	point2.SetValue(50)
	points := make([]slidescloud.IOneValueChartDataPoint, 2)
	points[0] = &point1
	points[1] = &point2
	series.SetDataPoints(points)
	serieses := make([]slidescloud.ISeries, 1)
	serieses[0] = &series
	dto1.SetSeries(serieses)
	var axes slidescloud.Axes
	var axis1 slidescloud.Axis
	isAutomaticMinValue := false
	axis1.SetIsAutomaticMinValue(&isAutomaticMinValue)
	axis1.SetMinValue(min1)
	isAutomaticMaxValue := false
	axis1.SetIsAutomaticMaxValue(&isAutomaticMaxValue)
	axis1.SetMaxValue(max1)
	axes.SetHorizontalAxis(&axis1)
	dto1.SetAxes(&axes)
	dto1.SetX(12)
	dto1.SetY(14)
	_, _, e = c.SlidesApi.CreateShape(fileName, slideIndex, &dto1, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e := c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min1 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

	var dto2 slidescloud.Chart
	dto2.SetType("Chart")
	dto2.SetChartType("Line")
	var axis2 slidescloud.Axis
	axis2.SetMinValue(min2)
	axes.SetHorizontalAxis(&axis2)
	dto2.SetAxes(&axes)
	_, _, e = c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, &dto2, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max1 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}

	var axis3 slidescloud.Axis
	axis3.SetMaxValue(max2)
	axes.SetHorizontalAxis(&axis3)
	_, _, e = c.SlidesApi.UpdateShape(fileName, slideIndex, shapeIndex, &dto2, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	r, _, e = c.SlidesApi.GetShape(fileName, slideIndex, shapeIndex, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMinValue() != min2 {
		t.Errorf("Error: Wrong MinValue.")
		return
	}
	if r.(slidescloud.IChart).GetAxes().GetHorizontalAxis().GetMaxValue() != max2 {
		t.Errorf("Error: Wrong MaxValue.")
		return
	}
}