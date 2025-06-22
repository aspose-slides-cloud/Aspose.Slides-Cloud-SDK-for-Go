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

	slidescloud "github.com/aspose-slides-cloud/aspose-slides-cloud-go/v25"
)

/*
   Test for pipeline
*/
func TestTemplateInput(t *testing.T) {
	file1 := slidescloud.NewRequestInputFile()
	file1.Index = 0

	file2 := slidescloud.NewRequestInputFile()
	file2.Index = 1

	input := slidescloud.NewInput()
	input.TemplateData = file1
	input.Template = file2

	output := slidescloud.NewOutputFile()

	task := slidescloud.NewSave()
	task.Format = "pptx"
	task.Output = output

	pipeline := slidescloud.NewPipeline()
	pipeline.Input = input
	pipeline.Tasks = []slidescloud.ITask{task}

	data1, e := ioutil.ReadFile(localFolder + "/TemplatingCVDataWithBase64.xml")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	data2, e := ioutil.ReadFile(localFolder + "/TemplateCV.pptx")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	files := [][]byte{data1, data2}

	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.Pipeline(pipeline, files)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
