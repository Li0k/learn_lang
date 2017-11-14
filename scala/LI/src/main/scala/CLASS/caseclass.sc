abstract class Notification

//case class可基于值直接进行比较
case class Email(sender: String, title: String, body: String) extends Notification

case class SMS(caller: String, message: String) extends Notification

case class VoiceRecording(contactName: String, link: String) extends Notification

def showNotification(notification: Notification): String = {
  notification match {
    case Email(email, title, _) =>
      s"You got an email from $email with title: $title"
    case SMS(number, message) =>
      s"You got an SMS from $number! Message: $message"
    case VoiceRecording(name, link) =>
      s"you received a Voice Recording from $name! Click the link to hear it: $link"
  }
}

val sms1 = SMS("12345", "Are you ok")

val sms2 = SMS("12345", "Are you ok") == sms1 //ture

val someVoiceRecording = VoiceRecording("Tom", "voicerecording.org/id/123")

showNotification(sms1)
showNotification(someVoiceRecording)


//guard
def showImportantNotification(notification: Notification, importantPeopleInfo: Seq[String]): String = {
  notification match {
    case Email(email, _, _) if importantPeopleInfo.contains(email) =>
      "You got an email from special someone!"
    case SMS(number, _) if importantPeopleInfo.contains(number) =>
      "You got an SMS from special someone!"
    case other =>
      showNotification(other) // nothing special, delegate to our original showNotification function
  }
}

val importantPeopleInfo = Seq("867-5309", "jenny@gmail.com")
val importantEmail = Email("jenny@gmail.com", "Drinks tonight?", "I'm free after 5!")
val importantSms = SMS("867-5309", "I'm here! Where are you?")

showImportantNotification(sms1,importantPeopleInfo)

case class Calculator(brand: String, model: String)

val hp20b = Calculator("HP", "20B")
val hp30b = Calculator("HP", "30B")
val other = Calculator("MP","60B")

def caclType(calc :Calculator) = calc match {
  case Calculator("HP", "20B") => "financial"
  case Calculator("HP", "48G") => "scientific"
  case Calculator("HP", "30B") => "business"
//  case Calculator(_, _) => "Calculator:is of unknown type"
  ///捕获这个未知Calculator
  case c@Calculator(_, _) => "Calculator: %s of unknown type".format(c)
}

caclType(hp20b)
caclType(hp30b)
caclType(other)


