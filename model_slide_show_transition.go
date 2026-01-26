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

package asposeslidescloud
import (
	"encoding/json"
)

// Slide Show Transition.
type ISlideShowTransition interface {

	// Transition Type
	GetType() string
	SetType(newValue string)

	// Advance After
	GetAdvanceAfter() *bool
	SetAdvanceAfter(newValue *bool)

	// Advance After Time
	GetAdvanceAfterTime() int32
	SetAdvanceAfterTime(newValue int32)

	// Advance On Click
	GetAdvanceOnClick() *bool
	SetAdvanceOnClick(newValue *bool)

	// Sound Is Built In
	GetSoundIsBuiltIn() *bool
	SetSoundIsBuiltIn(newValue *bool)

	// Sound Loop
	GetSoundLoop() *bool
	SetSoundLoop(newValue *bool)

	// Sound Mode
	GetSoundMode() string
	SetSoundMode(newValue string)

	// Sound Name
	GetSoundName() string
	SetSoundName(newValue string)

	// Speed
	GetSpeed() string
	SetSpeed(newValue string)

	// Corner Direction.
	GetCornerDirection() string
	SetCornerDirection(newValue string)

	// Eight Direction.
	GetEightDirection() string
	SetEightDirection(newValue string)

	// In/Out Direction.
	GetInOutDirection() string
	SetInOutDirection(newValue string)

	// Has Bounce.
	GetHasBounce() *bool
	SetHasBounce(newValue *bool)

	// Side Direction.
	GetSideDirection() string
	SetSideDirection(newValue string)

	// Pattern.
	GetPattern() string
	SetPattern(newValue string)

	// Left/Right Direction.
	GetLeftRightDirection() string
	SetLeftRightDirection(newValue string)

	// Morph Type.
	GetMorphType() string
	SetMorphType(newValue string)

	// From Black.
	GetFromBlack() *bool
	SetFromBlack(newValue *bool)

	// Orientation Direction.
	GetOrientationDirection() string
	SetOrientationDirection(newValue string)

	// Through Black.
	GetThroughBlack() *bool
	SetThroughBlack(newValue *bool)

	// Orientation.
	GetCornerAndCenterDirection() string
	SetCornerAndCenterDirection(newValue string)

	// Shred Pattern.
	GetShredPattern() string
	SetShredPattern(newValue string)

	// Orientation.
	GetOrientation() string
	SetOrientation(newValue string)

	// Spokes.
	GetSpokes() int32
	SetSpokes(newValue int32)

	// The duration of the slide transition effect in milliseconds. If not set, the duration is determined automatically based on Speed and Type values.
	GetDuration() int32
	SetDuration(newValue int32)
}

type SlideShowTransition struct {

	// Transition Type
	Type_ string `json:"Type,omitempty"`

	// Advance After
	AdvanceAfter *bool `json:"AdvanceAfter"`

	// Advance After Time
	AdvanceAfterTime int32 `json:"AdvanceAfterTime,omitempty"`

	// Advance On Click
	AdvanceOnClick *bool `json:"AdvanceOnClick"`

	// Sound Is Built In
	SoundIsBuiltIn *bool `json:"SoundIsBuiltIn"`

	// Sound Loop
	SoundLoop *bool `json:"SoundLoop"`

	// Sound Mode
	SoundMode string `json:"SoundMode,omitempty"`

	// Sound Name
	SoundName string `json:"SoundName,omitempty"`

	// Speed
	Speed string `json:"Speed,omitempty"`

	// Corner Direction.
	CornerDirection string `json:"CornerDirection,omitempty"`

	// Eight Direction.
	EightDirection string `json:"EightDirection,omitempty"`

	// In/Out Direction.
	InOutDirection string `json:"InOutDirection,omitempty"`

	// Has Bounce.
	HasBounce *bool `json:"HasBounce"`

	// Side Direction.
	SideDirection string `json:"SideDirection,omitempty"`

	// Pattern.
	Pattern string `json:"Pattern,omitempty"`

	// Left/Right Direction.
	LeftRightDirection string `json:"LeftRightDirection,omitempty"`

	// Morph Type.
	MorphType string `json:"MorphType,omitempty"`

	// From Black.
	FromBlack *bool `json:"FromBlack"`

	// Orientation Direction.
	OrientationDirection string `json:"OrientationDirection,omitempty"`

	// Through Black.
	ThroughBlack *bool `json:"ThroughBlack"`

	// Orientation.
	CornerAndCenterDirection string `json:"CornerAndCenterDirection,omitempty"`

	// Shred Pattern.
	ShredPattern string `json:"ShredPattern,omitempty"`

	// Orientation.
	Orientation string `json:"Orientation,omitempty"`

	// Spokes.
	Spokes int32 `json:"Spokes,omitempty"`

	// The duration of the slide transition effect in milliseconds. If not set, the duration is determined automatically based on Speed and Type values.
	Duration int32 `json:"Duration,omitempty"`
}

func NewSlideShowTransition() *SlideShowTransition {
	instance := new(SlideShowTransition)
	return instance
}

func (this *SlideShowTransition) GetType() string {
	return this.Type_
}

func (this *SlideShowTransition) SetType(newValue string) {
	this.Type_ = newValue
}
func (this *SlideShowTransition) GetAdvanceAfter() *bool {
	return this.AdvanceAfter
}

func (this *SlideShowTransition) SetAdvanceAfter(newValue *bool) {
	this.AdvanceAfter = newValue
}
func (this *SlideShowTransition) GetAdvanceAfterTime() int32 {
	return this.AdvanceAfterTime
}

func (this *SlideShowTransition) SetAdvanceAfterTime(newValue int32) {
	this.AdvanceAfterTime = newValue
}
func (this *SlideShowTransition) GetAdvanceOnClick() *bool {
	return this.AdvanceOnClick
}

func (this *SlideShowTransition) SetAdvanceOnClick(newValue *bool) {
	this.AdvanceOnClick = newValue
}
func (this *SlideShowTransition) GetSoundIsBuiltIn() *bool {
	return this.SoundIsBuiltIn
}

func (this *SlideShowTransition) SetSoundIsBuiltIn(newValue *bool) {
	this.SoundIsBuiltIn = newValue
}
func (this *SlideShowTransition) GetSoundLoop() *bool {
	return this.SoundLoop
}

func (this *SlideShowTransition) SetSoundLoop(newValue *bool) {
	this.SoundLoop = newValue
}
func (this *SlideShowTransition) GetSoundMode() string {
	return this.SoundMode
}

func (this *SlideShowTransition) SetSoundMode(newValue string) {
	this.SoundMode = newValue
}
func (this *SlideShowTransition) GetSoundName() string {
	return this.SoundName
}

func (this *SlideShowTransition) SetSoundName(newValue string) {
	this.SoundName = newValue
}
func (this *SlideShowTransition) GetSpeed() string {
	return this.Speed
}

func (this *SlideShowTransition) SetSpeed(newValue string) {
	this.Speed = newValue
}
func (this *SlideShowTransition) GetCornerDirection() string {
	return this.CornerDirection
}

func (this *SlideShowTransition) SetCornerDirection(newValue string) {
	this.CornerDirection = newValue
}
func (this *SlideShowTransition) GetEightDirection() string {
	return this.EightDirection
}

func (this *SlideShowTransition) SetEightDirection(newValue string) {
	this.EightDirection = newValue
}
func (this *SlideShowTransition) GetInOutDirection() string {
	return this.InOutDirection
}

func (this *SlideShowTransition) SetInOutDirection(newValue string) {
	this.InOutDirection = newValue
}
func (this *SlideShowTransition) GetHasBounce() *bool {
	return this.HasBounce
}

func (this *SlideShowTransition) SetHasBounce(newValue *bool) {
	this.HasBounce = newValue
}
func (this *SlideShowTransition) GetSideDirection() string {
	return this.SideDirection
}

func (this *SlideShowTransition) SetSideDirection(newValue string) {
	this.SideDirection = newValue
}
func (this *SlideShowTransition) GetPattern() string {
	return this.Pattern
}

func (this *SlideShowTransition) SetPattern(newValue string) {
	this.Pattern = newValue
}
func (this *SlideShowTransition) GetLeftRightDirection() string {
	return this.LeftRightDirection
}

func (this *SlideShowTransition) SetLeftRightDirection(newValue string) {
	this.LeftRightDirection = newValue
}
func (this *SlideShowTransition) GetMorphType() string {
	return this.MorphType
}

func (this *SlideShowTransition) SetMorphType(newValue string) {
	this.MorphType = newValue
}
func (this *SlideShowTransition) GetFromBlack() *bool {
	return this.FromBlack
}

func (this *SlideShowTransition) SetFromBlack(newValue *bool) {
	this.FromBlack = newValue
}
func (this *SlideShowTransition) GetOrientationDirection() string {
	return this.OrientationDirection
}

func (this *SlideShowTransition) SetOrientationDirection(newValue string) {
	this.OrientationDirection = newValue
}
func (this *SlideShowTransition) GetThroughBlack() *bool {
	return this.ThroughBlack
}

func (this *SlideShowTransition) SetThroughBlack(newValue *bool) {
	this.ThroughBlack = newValue
}
func (this *SlideShowTransition) GetCornerAndCenterDirection() string {
	return this.CornerAndCenterDirection
}

func (this *SlideShowTransition) SetCornerAndCenterDirection(newValue string) {
	this.CornerAndCenterDirection = newValue
}
func (this *SlideShowTransition) GetShredPattern() string {
	return this.ShredPattern
}

func (this *SlideShowTransition) SetShredPattern(newValue string) {
	this.ShredPattern = newValue
}
func (this *SlideShowTransition) GetOrientation() string {
	return this.Orientation
}

func (this *SlideShowTransition) SetOrientation(newValue string) {
	this.Orientation = newValue
}
func (this *SlideShowTransition) GetSpokes() int32 {
	return this.Spokes
}

func (this *SlideShowTransition) SetSpokes(newValue int32) {
	this.Spokes = newValue
}
func (this *SlideShowTransition) GetDuration() int32 {
	return this.Duration
}

func (this *SlideShowTransition) SetDuration(newValue int32) {
	this.Duration = newValue
}

func (this *SlideShowTransition) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	
	if valType, ok := GetMapValue(objMap, "type"); ok {
		if valType != nil {
			var valueForType string
			err = json.Unmarshal(*valType, &valueForType)
			if err != nil {
				var valueForTypeInt int32
				err = json.Unmarshal(*valType, &valueForTypeInt)
				if err != nil {
					return err
				}
				this.Type_ = string(valueForTypeInt)
			} else {
				this.Type_ = valueForType
			}
		}
	}
	
	if valAdvanceAfter, ok := GetMapValue(objMap, "advanceAfter"); ok {
		if valAdvanceAfter != nil {
			var valueForAdvanceAfter *bool
			err = json.Unmarshal(*valAdvanceAfter, &valueForAdvanceAfter)
			if err != nil {
				return err
			}
			this.AdvanceAfter = valueForAdvanceAfter
		}
	}
	
	if valAdvanceAfterTime, ok := GetMapValue(objMap, "advanceAfterTime"); ok {
		if valAdvanceAfterTime != nil {
			var valueForAdvanceAfterTime int32
			err = json.Unmarshal(*valAdvanceAfterTime, &valueForAdvanceAfterTime)
			if err != nil {
				return err
			}
			this.AdvanceAfterTime = valueForAdvanceAfterTime
		}
	}
	
	if valAdvanceOnClick, ok := GetMapValue(objMap, "advanceOnClick"); ok {
		if valAdvanceOnClick != nil {
			var valueForAdvanceOnClick *bool
			err = json.Unmarshal(*valAdvanceOnClick, &valueForAdvanceOnClick)
			if err != nil {
				return err
			}
			this.AdvanceOnClick = valueForAdvanceOnClick
		}
	}
	
	if valSoundIsBuiltIn, ok := GetMapValue(objMap, "soundIsBuiltIn"); ok {
		if valSoundIsBuiltIn != nil {
			var valueForSoundIsBuiltIn *bool
			err = json.Unmarshal(*valSoundIsBuiltIn, &valueForSoundIsBuiltIn)
			if err != nil {
				return err
			}
			this.SoundIsBuiltIn = valueForSoundIsBuiltIn
		}
	}
	
	if valSoundLoop, ok := GetMapValue(objMap, "soundLoop"); ok {
		if valSoundLoop != nil {
			var valueForSoundLoop *bool
			err = json.Unmarshal(*valSoundLoop, &valueForSoundLoop)
			if err != nil {
				return err
			}
			this.SoundLoop = valueForSoundLoop
		}
	}
	
	if valSoundMode, ok := GetMapValue(objMap, "soundMode"); ok {
		if valSoundMode != nil {
			var valueForSoundMode string
			err = json.Unmarshal(*valSoundMode, &valueForSoundMode)
			if err != nil {
				var valueForSoundModeInt int32
				err = json.Unmarshal(*valSoundMode, &valueForSoundModeInt)
				if err != nil {
					return err
				}
				this.SoundMode = string(valueForSoundModeInt)
			} else {
				this.SoundMode = valueForSoundMode
			}
		}
	}
	
	if valSoundName, ok := GetMapValue(objMap, "soundName"); ok {
		if valSoundName != nil {
			var valueForSoundName string
			err = json.Unmarshal(*valSoundName, &valueForSoundName)
			if err != nil {
				return err
			}
			this.SoundName = valueForSoundName
		}
	}
	
	if valSpeed, ok := GetMapValue(objMap, "speed"); ok {
		if valSpeed != nil {
			var valueForSpeed string
			err = json.Unmarshal(*valSpeed, &valueForSpeed)
			if err != nil {
				var valueForSpeedInt int32
				err = json.Unmarshal(*valSpeed, &valueForSpeedInt)
				if err != nil {
					return err
				}
				this.Speed = string(valueForSpeedInt)
			} else {
				this.Speed = valueForSpeed
			}
		}
	}
	
	if valCornerDirection, ok := GetMapValue(objMap, "cornerDirection"); ok {
		if valCornerDirection != nil {
			var valueForCornerDirection string
			err = json.Unmarshal(*valCornerDirection, &valueForCornerDirection)
			if err != nil {
				var valueForCornerDirectionInt int32
				err = json.Unmarshal(*valCornerDirection, &valueForCornerDirectionInt)
				if err != nil {
					return err
				}
				this.CornerDirection = string(valueForCornerDirectionInt)
			} else {
				this.CornerDirection = valueForCornerDirection
			}
		}
	}
	
	if valEightDirection, ok := GetMapValue(objMap, "eightDirection"); ok {
		if valEightDirection != nil {
			var valueForEightDirection string
			err = json.Unmarshal(*valEightDirection, &valueForEightDirection)
			if err != nil {
				var valueForEightDirectionInt int32
				err = json.Unmarshal(*valEightDirection, &valueForEightDirectionInt)
				if err != nil {
					return err
				}
				this.EightDirection = string(valueForEightDirectionInt)
			} else {
				this.EightDirection = valueForEightDirection
			}
		}
	}
	
	if valInOutDirection, ok := GetMapValue(objMap, "inOutDirection"); ok {
		if valInOutDirection != nil {
			var valueForInOutDirection string
			err = json.Unmarshal(*valInOutDirection, &valueForInOutDirection)
			if err != nil {
				var valueForInOutDirectionInt int32
				err = json.Unmarshal(*valInOutDirection, &valueForInOutDirectionInt)
				if err != nil {
					return err
				}
				this.InOutDirection = string(valueForInOutDirectionInt)
			} else {
				this.InOutDirection = valueForInOutDirection
			}
		}
	}
	
	if valHasBounce, ok := GetMapValue(objMap, "hasBounce"); ok {
		if valHasBounce != nil {
			var valueForHasBounce *bool
			err = json.Unmarshal(*valHasBounce, &valueForHasBounce)
			if err != nil {
				return err
			}
			this.HasBounce = valueForHasBounce
		}
	}
	
	if valSideDirection, ok := GetMapValue(objMap, "sideDirection"); ok {
		if valSideDirection != nil {
			var valueForSideDirection string
			err = json.Unmarshal(*valSideDirection, &valueForSideDirection)
			if err != nil {
				var valueForSideDirectionInt int32
				err = json.Unmarshal(*valSideDirection, &valueForSideDirectionInt)
				if err != nil {
					return err
				}
				this.SideDirection = string(valueForSideDirectionInt)
			} else {
				this.SideDirection = valueForSideDirection
			}
		}
	}
	
	if valPattern, ok := GetMapValue(objMap, "pattern"); ok {
		if valPattern != nil {
			var valueForPattern string
			err = json.Unmarshal(*valPattern, &valueForPattern)
			if err != nil {
				var valueForPatternInt int32
				err = json.Unmarshal(*valPattern, &valueForPatternInt)
				if err != nil {
					return err
				}
				this.Pattern = string(valueForPatternInt)
			} else {
				this.Pattern = valueForPattern
			}
		}
	}
	
	if valLeftRightDirection, ok := GetMapValue(objMap, "leftRightDirection"); ok {
		if valLeftRightDirection != nil {
			var valueForLeftRightDirection string
			err = json.Unmarshal(*valLeftRightDirection, &valueForLeftRightDirection)
			if err != nil {
				var valueForLeftRightDirectionInt int32
				err = json.Unmarshal(*valLeftRightDirection, &valueForLeftRightDirectionInt)
				if err != nil {
					return err
				}
				this.LeftRightDirection = string(valueForLeftRightDirectionInt)
			} else {
				this.LeftRightDirection = valueForLeftRightDirection
			}
		}
	}
	
	if valMorphType, ok := GetMapValue(objMap, "morphType"); ok {
		if valMorphType != nil {
			var valueForMorphType string
			err = json.Unmarshal(*valMorphType, &valueForMorphType)
			if err != nil {
				var valueForMorphTypeInt int32
				err = json.Unmarshal(*valMorphType, &valueForMorphTypeInt)
				if err != nil {
					return err
				}
				this.MorphType = string(valueForMorphTypeInt)
			} else {
				this.MorphType = valueForMorphType
			}
		}
	}
	
	if valFromBlack, ok := GetMapValue(objMap, "fromBlack"); ok {
		if valFromBlack != nil {
			var valueForFromBlack *bool
			err = json.Unmarshal(*valFromBlack, &valueForFromBlack)
			if err != nil {
				return err
			}
			this.FromBlack = valueForFromBlack
		}
	}
	
	if valOrientationDirection, ok := GetMapValue(objMap, "orientationDirection"); ok {
		if valOrientationDirection != nil {
			var valueForOrientationDirection string
			err = json.Unmarshal(*valOrientationDirection, &valueForOrientationDirection)
			if err != nil {
				var valueForOrientationDirectionInt int32
				err = json.Unmarshal(*valOrientationDirection, &valueForOrientationDirectionInt)
				if err != nil {
					return err
				}
				this.OrientationDirection = string(valueForOrientationDirectionInt)
			} else {
				this.OrientationDirection = valueForOrientationDirection
			}
		}
	}
	
	if valThroughBlack, ok := GetMapValue(objMap, "throughBlack"); ok {
		if valThroughBlack != nil {
			var valueForThroughBlack *bool
			err = json.Unmarshal(*valThroughBlack, &valueForThroughBlack)
			if err != nil {
				return err
			}
			this.ThroughBlack = valueForThroughBlack
		}
	}
	
	if valCornerAndCenterDirection, ok := GetMapValue(objMap, "cornerAndCenterDirection"); ok {
		if valCornerAndCenterDirection != nil {
			var valueForCornerAndCenterDirection string
			err = json.Unmarshal(*valCornerAndCenterDirection, &valueForCornerAndCenterDirection)
			if err != nil {
				var valueForCornerAndCenterDirectionInt int32
				err = json.Unmarshal(*valCornerAndCenterDirection, &valueForCornerAndCenterDirectionInt)
				if err != nil {
					return err
				}
				this.CornerAndCenterDirection = string(valueForCornerAndCenterDirectionInt)
			} else {
				this.CornerAndCenterDirection = valueForCornerAndCenterDirection
			}
		}
	}
	
	if valShredPattern, ok := GetMapValue(objMap, "shredPattern"); ok {
		if valShredPattern != nil {
			var valueForShredPattern string
			err = json.Unmarshal(*valShredPattern, &valueForShredPattern)
			if err != nil {
				var valueForShredPatternInt int32
				err = json.Unmarshal(*valShredPattern, &valueForShredPatternInt)
				if err != nil {
					return err
				}
				this.ShredPattern = string(valueForShredPatternInt)
			} else {
				this.ShredPattern = valueForShredPattern
			}
		}
	}
	
	if valOrientation, ok := GetMapValue(objMap, "orientation"); ok {
		if valOrientation != nil {
			var valueForOrientation string
			err = json.Unmarshal(*valOrientation, &valueForOrientation)
			if err != nil {
				var valueForOrientationInt int32
				err = json.Unmarshal(*valOrientation, &valueForOrientationInt)
				if err != nil {
					return err
				}
				this.Orientation = string(valueForOrientationInt)
			} else {
				this.Orientation = valueForOrientation
			}
		}
	}
	
	if valSpokes, ok := GetMapValue(objMap, "spokes"); ok {
		if valSpokes != nil {
			var valueForSpokes int32
			err = json.Unmarshal(*valSpokes, &valueForSpokes)
			if err != nil {
				return err
			}
			this.Spokes = valueForSpokes
		}
	}
	
	if valDuration, ok := GetMapValue(objMap, "duration"); ok {
		if valDuration != nil {
			var valueForDuration int32
			err = json.Unmarshal(*valDuration, &valueForDuration)
			if err != nil {
				return err
			}
			this.Duration = valueForDuration
		}
	}

	return nil
}
