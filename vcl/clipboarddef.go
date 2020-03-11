//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
)

func SetClipboard(newClipboard IObject) *TClipboard {
	return AsClipboard(Clipboard_SetClipboard(CheckPtr(newClipboard)))
}
