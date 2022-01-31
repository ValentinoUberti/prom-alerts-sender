/*
{
    "status": "firing",
    "labels": {
      "alertname": "",
      "service": "alertmanager-main",
      "severity": "warning"
    },
    "annotations": {
      "summary": "My special summary",
      "message": "My special message",
      "description": "My special description"
    },
    "generatorURL": "",
    "startsAt": "2022-01-26T08:13:37.48164259+01:00",
    "endsAt": "2022-01-26T12:33:37.48164259+01:00"
  }
*/

export default function createJSONMsg(command,result,status,alertname,severity,summary,message,description) {
  var alert = new Object();
  var data = new Object();
  data.command=command
  data.result= result
  alert.status=status;
  alert.labels = new Object();
  alert.labels.alertname = alertname;
  alert.labels.severity = severity;
  alert.annotations = new Object();
  alert.annotations.summary=summary;
  alert.annotations.message=message;
  alert.annotations.description=description;
  data.data=alert;
  var alertString = JSON.stringify(data);
  return alertString;
}










