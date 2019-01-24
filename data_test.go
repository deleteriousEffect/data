/*
Data. Easy-ish storage unit utility.
Copyright © 2019 Hayley Swimelar

This program is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 2 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, see <http://www.gnu.org/licenses/>.
*/

package data

import "testing"

func TestValue(t *testing.T) {
	var tableTests = []struct {
		in  Amount
		out float64
	}{
		{NewByte(2000), 2000},
		{NewByte(10), 10},
		{NewByte(4234), 4234},
		{NewByte(-10000), -10000},
		{NewByte(KiB), 1024},
		{NewByte(0), 0},

		{NewKibiByte(512), .5},
		{NewKibiByte(KiB), 1},
		{NewKibiByte(MiB), 1024},
		{NewKibiByte(MiB * 2), 2048},
		{NewKibiByte(ByteSize(float64(MiB) * float64(2.5))), 2560},
		{NewKibiByte(0), 0},
	}

	for _, tt := range tableTests {
		if tt.in.Value() != tt.out {
			t.Errorf("expected: %f, got: %f, from %+v", tt.out, tt.in.Value(), tt.in)
		}
	}
}

func TestInclusiveSize(t *testing.T) {
	var tableTests = []struct {
		in  Amount
		out int64
	}{
		{NewByte(2000), 2000},
		{NewByte(10), 10},
		{NewByte(4234), 4234},
		{NewByte(-10000), -10000},
		{NewByte(KiB), 1024},
		{NewByte(ByteSize(NewKibiByte(KiB).To(B))), 1024},
		{NewByte(0), 0},

		{NewKibiByte(512), 1},
		{NewKibiByte(KiB + 1), 2},
		{NewKibiByte(KiB - 1), 1},
		{NewKibiByte(0), 0},
		{NewKibiByte((KiB + 1) * -1), -2},
		{NewKibiByte((KiB - 1) * -1), -1},
	}

	for _, tt := range tableTests {
		if tt.in.InclusiveSize() != tt.out {
			t.Errorf("expected: %d, got: %d, from %+v", tt.out, tt.in.InclusiveSize(), tt.in)
		}
	}
}

func TestString(t *testing.T) {
	var tableTests = []struct {
		in  Amount
		out string
	}{
		{NewByte(10), "10.0 B"},
		{NewByte(0), "0.0 B"},

		{NewKibiByte(KiB), "1.0 KiB"},
		{NewKibiByte(0), "0.0 B"},

		{NewKibiByte(GiB), "1.0 GiB"},
		{NewKibiByte(TiB), "1.0 TiB"},
		{NewKibiByte(EiB), "1.0 EiB"},

		{NewByte(GiB), "1.0 GiB"},
		{NewByte(TiB), "1.0 TiB"},
		{NewByte(EiB), "1.0 EiB"},
	}

	for _, tt := range tableTests {
		if tt.in.String() != tt.out {
			t.Errorf("expected: %s, got: %s, from %+v", tt.out, tt.in, tt.in)
		}
	}
}

func TestTo(t *testing.T) {
	var tableTests = []struct {
		in   Amount
		base ByteSize
		out  float64
	}{

		{NewKibiByte(B), B, 1},
		{NewKibiByte(KiB), KiB, 1},
		{NewKibiByte(MiB), MiB, 1},
		{NewKibiByte(MiB), KiB, 1024},
		{NewKibiByte(GiB), GiB, 1},
	}

	for _, tt := range tableTests {
		if tt.in.To(tt.base) != tt.out {
			t.Errorf("expected: %f, got: %f, from %+v", tt.out, tt.in.Value(), tt.in)
		}
	}
}
