package CLASS

class Point(var x: Int, var y: Int) {
  def move(dx: Int, dy: Int): Unit = {
    x = x + dx
    y = y + dy
  }

  override def toString: String = s"($x , $y)"
}

class Point2(var x: Int = 0, var y: Int = 0)

class Point3() {
  private var _x = 0
  private var _y = 0
  private val bound = 100

  //getter
  def x = _x
  //setter
  def x_=(newValue: Int): Unit = {
    if (newValue < bound) _x = newValue else printWarning
  }

  def y = _y

  def y_=(newValue: Int): Unit = {
    if (newValue < bound) _y = newValue else printWarning
  }

  private def printWarning = println("WARNING: Out of bounds")

}

