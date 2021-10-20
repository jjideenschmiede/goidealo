//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of goidealo.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package goidealo

import (
	"errors"
	"strings"
)

// statusCodes
func statusCodes(status string) error {

	// Check each status codes
	switch strings.ReplaceAll(status, " ", "") {
	case "401":
		return errors.New("access is not authorized")
	case "404":
		return errors.New("not found")
	default:
		return nil
	}

}
