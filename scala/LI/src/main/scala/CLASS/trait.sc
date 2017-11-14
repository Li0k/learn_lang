trait Iterator[A] {
  def hasNext: Boolean

  def next(): A
}

class IntIerator(to: Int) extends Iterator[Int] {
  private var current = 0

  override def hasNext = current < to

  override def next(): Int = {
    if (hasNext) {
      val t = current
      current += 1
      t
    } else 0
  }
}

val iterator = new IntIerator(10)
iterator.next()
iterator.next()