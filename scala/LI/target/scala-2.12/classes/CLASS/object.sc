object Timer {
  var count = 0

  def currentCount():Long = {
    count += 1
    count
  }
}

Timer.currentCount()
Timer.currentCount()