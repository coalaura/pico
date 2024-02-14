package main

type Cursor struct {
	Disabled bool

	X int
	Y int

	BodyWidth  int
	BodyHeight int

	maxRememberedX int // epic
}

func (c *Cursor) SetDimensions(bodyWidth, bodyHeight int) {
	c.BodyWidth = bodyWidth
	c.BodyHeight = bodyHeight
}

func (c *Cursor) SetDisabled(disabled bool) {
	c.Disabled = disabled
}

func (c *Cursor) AbsoluteX() int {
	return c.X
}

func (c *Cursor) AbsoluteY() int {
	return c.Y + ScrollY
}

func (c *Cursor) CurrentLine() string {
	return Lines[c.AbsoluteY()]
}

func (c *Cursor) MaxX() int {
	line := c.CurrentLine()

	return len([]rune(line))
}

func (c *Cursor) ClampX() {
	maxX := c.MaxX()

	if c.X > maxX {
		c.X = maxX
	} else if c.X < 0 {
		c.X = 0
	}
}

func (c *Cursor) SetX(x int) {
	c.X = x
	c.maxRememberedX = c.X

	c.ClampX()
}

func (c *Cursor) SetY(y int) {
	c.Y = y

	if c.maxRememberedX > c.X {
		c.X = c.maxRememberedX
	}

	if c.Y < 0 {
		ScrollY += c.Y

		if ScrollY < 0 {
			ScrollY = 0
		}

		c.Y = 0
	} else {
		maxLines := len(Lines) - 1
		maxBodyY := c.BodyHeight

		if maxBodyY > maxLines {
			maxBodyY = maxLines
		}

		tooFar := c.Y - maxBodyY

		if tooFar > 0 {
			ScrollY += tooFar
			c.Y = maxBodyY

			if ScrollY > maxLines-maxBodyY {
				ScrollY = maxLines - maxBodyY

				// We were past the end of the file, set the cursor to the end
				c.MoveEndOfLine()
			}
		}
	}

	c.ClampX()
}

func (c *Cursor) Up() {
	if c.Disabled {
		return
	}

	c.SetY(c.Y - 1)
}

func (c *Cursor) Down() {
	if c.Disabled {
		return
	}

	c.SetY(c.Y + 1)
}

func (c *Cursor) Left() {
	if c.Disabled {
		return
	}

	if c.X == 0 {
		ay := c.AbsoluteY()

		if ay > 0 {
			c.Up()

			c.MoveEndOfLine()
		}

		return
	}

	c.SetX(c.X - 1)
}

func (c *Cursor) Right() {
	if c.Disabled {
		return
	}

	if c.X == c.MaxX() {
		if c.AbsoluteY() < len(Lines)-1 {
			c.SetX(0)
		}

		c.Down()

		return
	}

	c.SetX(c.X + 1)
}

func (c *Cursor) MoveEndOfLine() {
	line := Lines[c.AbsoluteY()]

	c.SetX(len([]rune(line)))
}
