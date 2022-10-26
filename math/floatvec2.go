package math

import std_math "math"

type FloatVec2 struct {
  X float64
  Y float64
}

// MANIPULATION w/ OTHER VECS

func (a *FloatVec2) Add(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: a.X + b.X,
    Y: a.Y + b.Y,
  }
}

func (a *FloatVec2) Sub(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: a.X - b.X,
    Y: a.Y - b.Y,
  }
}

func (a *FloatVec2) Mul(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: a.X * b.X,
    Y: a.Y * b.Y,
  }
}

func (a *FloatVec2) Div(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: a.X / b.X,
    Y: a.Y / b.Y,
  }
}

func (a *FloatVec2) Min(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: std_math.Min(a.X, b.X),
    Y: std_math.Min(a.Y, b.Y),
  }
}

func (a *FloatVec2) Max(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: std_math.Max(a.X, b.X),
    Y: std_math.Max(a.Y, b.Y),
  }
}

func (a *FloatVec2) Clamp(mn FloatVec2, mx FloatVec2) FloatVec2 {
  return a.Max(mx).Min(mn)
}

func (a *FloatVec2) Ceil(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: std_math.Ceil(a.X, b.X),
    Y: std_math.Ceil(a.Y, b.Y),
  }
}

func (a *FloatVec2) Round(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: std_math.Round(a.X, b.X),
    Y: std_math.Round(a.Y, b.Y),
  }
}

func (a *FloatVec2) Floor(b FloatVec2) FloatVec2 {
  return FloatVec2 {
    X: std_math.Floor(a.X, b.X),
    Y: std_math.Floor(a.Y, b.Y),
  }
}

func (a *FloatVec2) Hypot() float64 {
  return std_math.Hypot(a.X, a.Y)
}


