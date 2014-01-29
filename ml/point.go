//
//  point.go
//
//  Created by Hicham Bouabdallah
//  Copyright (c) 2012 SimpleRocket LLC
//
//  Permission is hereby granted, free of charge, to any person
//  obtaining a copy of this software and associated documentation
//  files (the "Software"), to deal in the Software without
//  restriction, including without limitation the rights to use,
//  copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the
//  Software is furnished to do so, subject to the following
//  conditions:
//
//  The above copyright notice and this permission notice shall be
//  included in all copies or substantial portions of the Software.
//
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
//  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
//  OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
//  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
//  HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//  WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
//  FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
//  OTHER DEALINGS IN THE SOFTWARE.
//

package ml

import "github.com/hishboy/gocommons/lang"
import "math"
import "fmt"

type Point struct {
	items *lang.ArrayList
}

func NewPoint(items []float64) *Point {
	self := &Point{}
	self.items = lang.NewArrayList()
	for i :=0; i < len(items); i++ {
		self.items.Add(items[i])
	}
	return self
}

func (self *Point) Items() *lang.ArrayList {
	return self.items
}

func (self *Point) DistanceFromPoint(otherPoint *Point) float64 {
	// FIXME: hicham - should return error if array size doesn't match
	if (self.items.Len() != otherPoint.items.Len()) {
		panic(fmt.Sprintf("item ", self.items.ToSlice(), " doesn't match ", otherPoint.items.ToSlice(), " size"))
	}
	
	total := 0.0 
	for i := 0; i < self.items.Len(); i++ {
		thisCoordinate := self.items.Get(i).(float64)
		otherCoordinate := otherPoint.items.Get(i).(float64)
		total = total + math.Pow(thisCoordinate-otherCoordinate, 2)
	}
	
	return math.Sqrt(total)
}