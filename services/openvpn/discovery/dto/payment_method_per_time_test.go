/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dto

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mysteriumnetwork/node/money"
	"github.com/stretchr/testify/assert"
)

func TestPaymentMethodPerTimeSerialize(t *testing.T) {
	price := money.NewMoney(0.5, money.CurrencyMyst)

	var tests = []struct {
		model        PaymentRate
		expectedJSON string
	}{
		{
			PaymentRate{
				Price:    price,
				Duration: time.Duration(10),
			},
			`{
				"price": {
					"amount": 50000000,
					"currency": "MYST"
				},
				"duration": 10
			}`,
		},
		{
			PaymentRate{},
			`{
				"price": {},
				"duration": 0
			}`,
		},
	}

	for _, test := range tests {
		jsonBytes, err := json.Marshal(test.model)

		assert.Nil(t, err)
		assert.JSONEq(t, test.expectedJSON, string(jsonBytes))
	}
}

func TestPaymentMethodPerTimeUnserialize(t *testing.T) {
	price := money.NewMoney(0.5, money.CurrencyMyst)

	var tests = []struct {
		json          string
		expectedModel PaymentRate
		expectedError error
	}{
		{
			`{
				"price": {
					"amount": 50000000,
					"currency": "MYST"
				},
				"duration": 10
			}`,
			PaymentRate{
				Price:    price,
				Duration: time.Duration(10),
			},
			nil,
		},
		{
			`{
				"price": {},
				"duration": 0
			}`,
			PaymentRate{},
			nil,
		},
		{
			`{}`,
			PaymentRate{},
			nil,
		},
	}

	for _, test := range tests {
		var model PaymentRate
		err := json.Unmarshal([]byte(test.json), &model)

		assert.Equal(t, test.expectedModel, model)
		assert.Equal(t, test.expectedError, err)
	}
}
