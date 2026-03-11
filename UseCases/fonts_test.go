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
	"os"
	"testing"

	slidescloud "github.com/aspose-slides-cloud/Aspose.Slides-Cloud-SDK-for-Go/v26"
)

var fontName = "Calibri"

/*
   Test for get available fonts
*/
func TestGetAvailableFonts(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	response, _, e := c.SlidesApi.GetAvailableFonts("", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(response.GetList()) <= 1 {
		t.Errorf("Expected greater thatn %v, but was %v", 1, len(response.GetList()))
		return
	}
	isCustom := response.GetList()[0].GetIsCustom()
	if isCustom != nil && *isCustom {
		t.Errorf("Expected false, but was true")
		return
	}
}

/*
   Test for get fonts
*/
func TestGetFonts(t *testing.T) {
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

	response, _, e := c.SlidesApi.GetFonts(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 4 {
		t.Errorf("Expected %v, but was %v", 4, len(response.GetList()))
		return
	}
}

/*
   Test for get fonts
*/
func TestGetFontsOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	document, e := ioutil.ReadFile(localTestFile)

	response, _, e := c.SlidesApi.GetFontsOnline(document, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	if len(response.GetList()) != 4 {
		t.Errorf("Expected %v, but was %v", 4, len(response.GetList()))
		return
	}
}

/*
   Test for set embedded font
*/
func TestSetEmbeddedFont(t *testing.T) {
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

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
}

/*
   Test for set embedded font online
*/
func TestSetEmbeddedFontOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	document, e := ioutil.ReadFile(localTestFile)

	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for set embedded fonts
*/
func TestSetEmbeddedFonts(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fonts := []string{ fontName }
	var onlyUsed bool = false
	response, _, e := c.SlidesApi.SetEmbeddedFonts(fileName, nil, fonts, &onlyUsed, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
	fontName := "Calibri"
	if response.GetList()[2].GetFontName() != fontName {
		t.Errorf("Expected %v, but was %v", fontName, response.GetList()[2].GetFontName())
		return
	}
}

/*
   Test for set embedded font from request
*/
func TestSetEmbeddedFontsFromRequest(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	fontFile, e := ioutil.ReadFile(localFolder + "/calibri.ttf")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	_, e = c.SlidesApi.CopyFile(tempFilePath, filePath, "", "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false
	response, _, e := c.SlidesApi.SetEmbeddedFonts(fileName, [][]byte{ fontFile }, nil, &onlyUsed, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
	fontName := "Calibri"
	if response.GetList()[2].GetFontName() != fontName {
		t.Errorf("Expected %v, but was %v", fontName, response.GetList()[2].GetFontName())
		return
	}
}

/*
   Test for set embedded font from request
*/
func TestSetEmbeddedFontsOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fonts := []string{ fontName }
	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontsOnline(document, nil, fonts, &onlyUsed, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for set embedded font from request
*/
func TestSetEmbeddedFontFromRequestOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	fontFile, e := ioutil.ReadFile(localFolder + "/calibri.ttf")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false
	_, _, e = c.SlidesApi.SetEmbeddedFontsOnline(document, [][]byte{ fontFile }, nil, &onlyUsed, password)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for compress embedded fonts
*/
func TestCompressEmbeddedFonts(t *testing.T) {
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

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}

        //In a real world example, you would rather get the same result by calling SetEmbeddedFont with onlyUsed = true
	_, e = c.SlidesApi.CompressEmbeddedFonts(fileName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for compress embedded fonts online
*/
func TestCompressEmbeddedFontsOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false
	result, _, e := c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	resultStat, e := os.Stat(result.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

        //In a real world example, you would rather get the same result by calling SetEmbeddedFont with onlyUsed = true
	compressedResult, _, e := c.SlidesApi.CompressEmbeddedFontsOnline(document, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	compressedResultStat, e := os.Stat(compressedResult.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if compressedResultStat.Size() >= resultStat.Size() {
		t.Errorf("Wrong response size. Expected less than %v but was %v.", resultStat.Size(), compressedResultStat.Size())
		return
	}
}

/*
   Test for set embedded font
*/
func TestDeleteEmbeddedFont(t *testing.T) {
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

	var onlyUsed bool = false

	response, _, e := c.SlidesApi.SetEmbeddedFont(fileName, fontName, &onlyUsed, password, folderName, "", "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}

	response, _, e = c.SlidesApi.DeleteEmbeddedFont(fileName, fontName, password, folderName, "")

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded = response.GetList()[0].GetIsEmbedded()
	if isEmbedded != nil && *isEmbedded {
		t.Errorf("Expected %v, but was %v", false, isEmbedded)
		return
	}
}

/*
   Test for delete embedded font online
*/
func TestDeleteEmbeddedFontOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	var onlyUsed bool = false
	embedded, _, e := c.SlidesApi.SetEmbeddedFontOnline(document, fontName, &onlyUsed, password, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	embeddedDocument, e := ioutil.ReadFile(embedded.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	unembedded, _, e := c.SlidesApi.DeleteEmbeddedFontOnline(embeddedDocument, fontName, password)

	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	embeddedStat, e := os.Stat(embedded.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	unembeddedStat, e := os.Stat(unembedded.Name())
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if unembeddedStat.Size() >= embeddedStat.Size() {
		t.Errorf("Wrong file size. Expected less than %v but was %v.", embeddedStat.Size(), unembeddedStat.Size())
		return
	}
}

/*
   Test for replace font
*/
func TestReplaceFont(t *testing.T) {
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

	sourceFontName := "Calibri"
	targetFontName := "Times New Roman"
	var embed bool = true
	response, _, e := c.SlidesApi.ReplaceFont(fileName, sourceFontName, targetFontName, &embed, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	isEmbedded := response.GetList()[2].GetIsEmbedded()
	if isEmbedded == nil || !*isEmbedded {
		t.Errorf("Expected %v, but was %v", true, response.GetList()[2].GetIsEmbedded())
		return
	}
	if response.GetList()[2].GetFontName() != targetFontName {
		t.Errorf("Expected %v, but was %v", targetFontName, response.GetList()[2].GetFontName())
		return
	}
}

/*
   Test for replace font online
*/
func TestReplaceFontOnline(t *testing.T) {
	c, e := GetApiClient()
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	document, e := ioutil.ReadFile(localTestFile)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}

	sourceFontName := "Calibri"
	targetFontName := "Times New Roman"
	var embed bool = true
	_, _, e = c.SlidesApi.ReplaceFontOnline(document, sourceFontName, targetFontName, &embed, password, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}

/*
   Test for font substitution
*/
func TestFontSubstitution(t *testing.T) {
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

	targetFontName := "Times New Roman"

	fontRule1 := slidescloud.NewFontSubstRule()
	fontRule1.SetSourceFont("Arial")
	fontRule1.SetTargetFont(targetFontName)
	notFoundOnly := false
	fontRule1.SetNotFoundOnly(&notFoundOnly)
	fontRule2 := slidescloud.NewFontSubstRule()
	fontRule2.SetSourceFont("Calibri")
	fontRule2.SetTargetFont(targetFontName)

	exportOptions := slidescloud.NewImageExportOptions()
	exportOptions.SetFontSubstRules([]slidescloud.IFontSubstRule{fontRule1, fontRule2})

	_, _, e = c.SlidesApi.DownloadPresentation(fileName, "PNG", exportOptions, password, folderName, "", "", nil)
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
}
