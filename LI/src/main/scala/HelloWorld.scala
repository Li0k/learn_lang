import CLASS.{Point, Point2, Point3}

object HelloWorld {
  def main(args: Array[String]): Unit = {

    println("hello world")

    val point = new Point(1, 3)
    println(point)

    point.move(3, 1)

    println(point)

    val point3 = new Point3
    point3.x = 99
    point3.y = 101 // prints the warning
    println(point3.x)

    def point4 = new Point2

    point4.x = 15
    println(point4.x)

    abstract class AbsIterator {
      type T

      def hasNext: Boolean

      def next(): T
    }

    class StringIterator(s: String) extends AbsIterator {
      type T = Char
      private var i = 0

      def hasNext = i < s.length

      def next = {
        val ch = s charAt i
        i += 1
        ch
      }
    }

    trait RichIterator extends AbsIterator {
      def foreach(f: T => Unit): Unit = while (hasNext) f(next)
    }

    object StringIteratorTest extends App {

      class Iter extends StringIterator(args(0)) with RichIterator

      val iter = new Iter
      iter foreach println
    }

  }
}

