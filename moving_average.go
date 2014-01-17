package fpd

type MovingAverage struct {
	samples  []*Decimal
	scale    int
	capacity int
}

func NewMovingAverage(capacity int, scale int) *MovingAverage {
	return &MovingAverage{
		scale:    scale,
		capacity: capacity,
	}
}

func (ma *MovingAverage) Append(sample *Decimal) {
	if len(ma.samples) == ma.capacity {
		ma.samples = append(ma.samples[1:], sample)
	} else {
		ma.samples = append(ma.samples, sample)
	}
}

func (ma *MovingAverage) Calculate() *Decimal {
	sum := New(0, ma.scale)
	for _, sample := range ma.samples {
		sum = sum.Add(sample)
	}

	return sum.Div(New(int64(len(ma.samples)), 0))
}

func (ma *MovingAverage) Capacity() int {
	return ma.capacity
}

func (ma *MovingAverage) Size() int {
	return len(ma.samples)
}
