package main

type Cursor struct {
	Disabled bool

	X int
	Y int

	BodyWidth  int
	BodyHeight int
}

func (c *Cursor) SetDimensions(bodyWidth, bodyHeight int) {
	c.BodyWidth = bodyWidth
	c.BodyHeight = bodyHeight
}

func (c *Cursor) SetDisabled(disabled bool) {
	c.Disabled = disabled
}

func (c *Cursor) AbsoluteX() int {
	return c.X + ScrollX
}

func (c *Cursor) AbsoluteY() int {
	return c.Y + ScrollY
}

func (c *Cursor) CurrentLine() string {
	return Lines[c.AbsoluteY()]
}

func (c *Cursor) LineLength() int {
	line := c.CurrentLine()

	return len([]rune(line))
}

func (c *Cursor) CurrentChar() string {
	absX := c.AbsoluteX()

	line := c.CurrentLine()
	runes := []rune(line)

	if absX >= len(runes) {
		return " "
	}

	return string(runes[absX])
}

func (c *Cursor) SetAbsX(x int) {
	ScrollX = 0
	c.X = x

	if c.X < 0 {
		ScrollX += c.X

		if ScrollX < 0 {
			ScrollX = 0
		}

		c.X = 0
	} else {
		maxLine := c.LineLength()
		maxX := maxLine

		if maxX >= c.BodyWidth {
			maxX = c.BodyWidth - 1
		}

		tooFar := c.X - maxX

		if tooFar > 0 {
			ScrollX += tooFar
			c.X = maxX

			if ScrollX > maxLine-maxX {
				ScrollX = maxLine - maxX

				// We were past the end of the line, set the cursor to the end
				c.X = maxX
			}
		}
	}
}

func (c *Cursor) SetAbsY(y int) {
	ScrollY = 0
	c.Y = y

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

	c.SetAbsX(c.AbsoluteX())
}

func (c *Cursor) Up() {
	if c.Disabled {
		return
	}

	c.SetAbsY(c.AbsoluteY() - 1)
}

func (c *Cursor) Down() {
	if c.Disabled {
		return
	}

	c.SetAbsY(c.AbsoluteY() + 1)
}

func (c *Cursor) Left() {
	if c.Disabled {
		return
	}

	if c.AbsoluteX() == 0 {
		ay := c.AbsoluteY()

		if ay > 0 {
			c.Up()

			c.MoveEndOfLine()
		}

		return
	}

	c.SetAbsX(c.AbsoluteX() - 1)
}

func (c *Cursor) Right() {
	if c.Disabled {
		return
	}

	if c.AbsoluteX() == c.LineLength() {
		if c.AbsoluteY() < len(Lines)-1 {
			c.MoveStartOfLine()
		}

		c.Down()

		return
	}

	c.SetAbsX(c.AbsoluteX() + 1)
}

func (c *Cursor) MoveEndOfLine() {
	c.SetAbsX(c.LineLength())
}

func (c *Cursor) MoveStartOfLine() {
	c.SetAbsX(0)
}
