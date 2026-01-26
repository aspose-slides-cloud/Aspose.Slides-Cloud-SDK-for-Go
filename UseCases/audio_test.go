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
   Test for audio caption tracks
*/
func TestAudioCaptionTracks(t *testing.T) {
    var slideIndex int32 = 3
    var shapeIndex int32 = 3
    track1Label := "track1"
    track2Label := "track2"
    track1Data := "WEBVTT\n\n00:00:00.000 --> 00:00:10.000\nCaption 1 text."
    track2Data := "WEBVTT\n\n00:00:00.000 --> 00:00:10.000\nCaption 2 text."

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
	_, _, e = c.SlidesApi.CreateShape(fileName, slideIndex, dto, nil, nil, password, folderName, "", "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	captions, _, e := c.SlidesApi.GetCaptionTracks(fileName, slideIndex, shapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(captions.GetItems()) != 0 {
		t.Errorf("Wrong track count. Expected %v but was %v.", 0, len(captions.GetItems()))
		return
	}

	_, _, e = c.SlidesApi.CreateCaptionTrack(fileName, slideIndex, shapeIndex, track1Label, track1Data, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	_, _, e = c.SlidesApi.CreateCaptionTrack(fileName, slideIndex, shapeIndex, track2Label, track2Data, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	var includeData bool = true
	captions, _, e = c.SlidesApi.GetCaptionTracks(fileName, slideIndex, shapeIndex, &includeData, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(captions.GetItems()) != 2 {
		t.Errorf("Wrong track count. Expected %v but was %v.", 2, len(captions.GetItems()))
		return
	}
	if captions.GetItems()[0].GetLabel() != track1Label {
		t.Errorf("Wrong track label. Expected %v but was %v.", track1Label, captions.GetItems()[0].GetLabel())
		return
	}
	if captions.GetItems()[0].GetDataAsString() != track1Data {
		t.Errorf("Wrong track data. Expected %v but was %v.", track1Data, captions.GetItems()[0].GetDataAsString())
		return
	}
	if captions.GetItems()[1].GetLabel() != track2Label {
		t.Errorf("Wrong track label. Expected %v but was %v.", track2Label, captions.GetItems()[1].GetLabel())
		return
	}
	if captions.GetItems()[1].GetDataAsString() != track2Data {
		t.Errorf("Wrong track data. Expected %v but was %v.", track2Data, captions.GetItems()[1].GetDataAsString())
		return
	}

	_, e = c.SlidesApi.DeleteCaptionTrack(fileName, slideIndex, shapeIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	includeData = false
	captions, _, e = c.SlidesApi.GetCaptionTracks(fileName, slideIndex, shapeIndex, &includeData, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(captions.GetItems()) != 1 {
		t.Errorf("Wrong track count. Expected %v but was %v.", 1, len(captions.GetItems()))
		return
	}
	if captions.GetItems()[0].GetLabel() != track2Label {
		t.Errorf("Wrong track label. Expected %v but was %v.", track2Label, captions.GetItems()[0].GetLabel())
		return
	}
	if captions.GetItems()[0].GetDataAsString() != "" {
		t.Errorf("Wrong track data. Expected %v but was %v.", "", captions.GetItems()[0].GetDataAsString())
		return
	}

	_, e = c.SlidesApi.DeleteCaptionTracks(fileName, slideIndex, shapeIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	includeData = false
	captions, _, e = c.SlidesApi.GetCaptionTracks(fileName, slideIndex, shapeIndex, &includeData, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(captions.GetItems()) != 0 {
		t.Errorf("Wrong track count. Expected %v but was %v.", 0, len(captions.GetItems()))
		return
	}
}
