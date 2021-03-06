/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package protoutil_test

import (
	"testing"

	"github.com/Yunpeng-J/fabric-protos-go/common"
	"github.com/Yunpeng-J/HLF-2.2/protoutil"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigGroup(t *testing.T) {
	assert.Equal(t,
		&common.ConfigGroup{
			Groups:   make(map[string]*common.ConfigGroup),
			Values:   make(map[string]*common.ConfigValue),
			Policies: make(map[string]*common.ConfigPolicy),
		},
		protoutil.NewConfigGroup(),
	)
}
