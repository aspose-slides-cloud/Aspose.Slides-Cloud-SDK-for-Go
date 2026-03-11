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

var animatedShapeIndex int32 = 3

/*
   Test for Get animation
*/
func TestGetAnimation(t *testing.T) {
	var paragraphIndex int32 = 1

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

	animation, _, e := c.SlidesApi.GetAnimation(fileName, slideIndex, nil, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}

	animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &animatedShapeIndex, nil, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}

	animation, _, e = c.SlidesApi.GetAnimation(fileName, slideIndex, &animatedShapeIndex, &paragraphIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
}

/*
   Test for set animation
*/
func TestSetAnimation(t *testing.T) {
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

	dto := slidescloud.NewSlideAnimation()
	effect1 := slidescloud.NewEffect()
	effect1.Type_ = "Blink"
	effect1.ShapeIndex = 2
	effect1.ParagraphIndex = 1

	effect2 := slidescloud.NewEffect()
	effect2.Type_ = "Box"
	effect2.Subtype = "In"
	effect2.PresetClassType = "Entrance"
	effect2.ShapeIndex = 4
	dto.MainSequence = []slidescloud.IEffect{effect1, effect2}
	dto.InteractiveSequences = []slidescloud.IInteractiveSequence{}
	animation, _, e := c.SlidesApi.SetAnimation(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != len(dto.GetMainSequence()) {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", len(dto.GetMainSequence()), len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation effect
*/
func TestCreateAnimationEffect(t *testing.T) {
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

	dto := slidescloud.NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = animatedShapeIndex
	animation, _, e := c.SlidesApi.CreateAnimationEffect(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 2 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 2, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation intercative sequence
*/
func TestCreateAnimationInteractiveSequence(t *testing.T) {
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

	dto := slidescloud.NewInteractiveSequence()
	dto.TriggerShapeIndex = 2
	effect := slidescloud.NewEffect()
	effect.Type_ = "Blast"
	effect.ShapeIndex = animatedShapeIndex
	dto.Effects = []slidescloud.IEffect{effect}
	animation, _, e := c.SlidesApi.CreateAnimationInteractiveSequence(fileName, slideIndex, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 2 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 2, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for create animation interactive sequence effect
*/
func TestCreateAnimationInteractiveSequenceEffect(t *testing.T) {
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

	dto := slidescloud.NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = animatedShapeIndex
	animation, _, e := c.SlidesApi.CreateAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for update animation effect
*/
func TestUpdateAnimationEffect(t *testing.T) {
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

	dto := slidescloud.NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = animatedShapeIndex
	animation, _, e := c.SlidesApi.UpdateAnimationEffect(fileName, slideIndex, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for update animation interactive sequence effect
*/
func TestUpdateAnimationInteractiveSequenceEffect(t *testing.T) {
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

	dto := slidescloud.NewEffect()
	dto.Type_ = "Blast"
	dto.ShapeIndex = animatedShapeIndex
	animation, _, e := c.SlidesApi.UpdateAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, 1, dto, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation
*/
func TestDeleteAnimation(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimation(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation main sequence
*/
func TestDeleteAnimationMainSequence(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimationMainSequence(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation main sequence effect
*/
func TestDeleteAnimationEffect(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimationEffect(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 0 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 0, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequences
*/
func TestDeleteAnimationInteractiveSequences(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequences(fileName, slideIndex, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequence
*/
func TestDeleteAnimationInteractiveSequence(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequence(fileName, slideIndex, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 0 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 0, len(animation.GetInteractiveSequences()))
		return
	}
}

/*
   Test for delete animation interactive sequence effect
*/
func TestDeleteAnimationInteractiveSequenceEffect(t *testing.T) {
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

	animation, _, e := c.SlidesApi.DeleteAnimationInteractiveSequenceEffect(fileName, slideIndex, 1, 1, password, folderName, "")
	if e != nil {
		t.Errorf("Error: %v.", e)
		return
	}
	if len(animation.GetMainSequence()) != 1 {
		t.Errorf("Wrong effect count. Expected not %v but was %v.", 1, len(animation.GetMainSequence()))
		return
	}
	if len(animation.GetInteractiveSequences()) != 1 {
		t.Errorf("Wrong sequence count. Expected not %v but was %v.", 1, len(animation.GetInteractiveSequences()))
		return
	}
}
