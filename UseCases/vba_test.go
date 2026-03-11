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
	Create VBA module
*/
func TestCreateVbaModule(t *testing.T) {
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

	dto := slidescloud.NewVbaModule()
	dto.Name = "Module1"
	dto.SourceCode = "Sub Test() MsgBox \"Test\" End Sub"
	reference0 := slidescloud.NewVbaReference()
	reference0.Name = "stdole"
	reference0.LibId = "*\\G{00020430-0000-0000-C000-000000000046}#2.0#0#C:\\Windows\\system32\\stdole2.tlb#OLE Automation"
	reference1 := slidescloud.NewVbaReference()
	reference1.Name = "Office"
	reference1.LibId = "*\\G{2DF8D04C-5BFA-101B-BDE5-00AA0044DE52}#2.0#0#C:\\Program Files\\Common Files\\Microsoft Shared\\OFFICE14\\MSO.DLL#Microsoft Office 14.0 Object Library"
	dto.References = []slidescloud.IVbaReference{reference0, reference1}

	_, _, e = c.SlidesApi.CreateVbaModule(fileName, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
	Delete VBA module
*/
func TestDeleteVbaModule(t *testing.T) {
	macrosFileName := "macros.pptm"
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + macrosFileName, folderName + "/" + macrosFileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var moduleIndex int32 = 1
	_, _, e = c.SlidesApi.DeleteVbaModule(macrosFileName, moduleIndex, "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
	Get VBA module
*/
func TestGetVbaModule(t *testing.T) {
	macrosFileName := "macros.pptm"
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + macrosFileName, folderName + "/" + macrosFileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var moduleIndex int32 = 1
	_, _, e = c.SlidesApi.GetVbaModule(macrosFileName, moduleIndex, "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
	Delete VBA project info
*/
func TestGetVbaProject(t *testing.T) {
	macrosFileName := "macros.pptm"
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + macrosFileName, folderName + "/" + macrosFileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, _, e = c.SlidesApi.GetVbaProject(macrosFileName, "", folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
	Update VBA module
*/
func TestUpdateVbaModule(t *testing.T) {
	macrosFileName := "macros.pptm"
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, e = c.SlidesApi.CopyFile(tempFolderName + "/" + macrosFileName, folderName + "/" + macrosFileName, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	dto := slidescloud.NewVbaModule()
	dto.SourceCode = "Sub Test() MsgBox \"Test\" End Sub"

	var moduleIndex int32 = 1

	_, _, e = c.SlidesApi.UpdateVbaModule(macrosFileName, moduleIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
