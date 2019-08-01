'use strict';

export function sendEmail(templateId, senderEmail, receiverEmail, message,
	onSuccess, onFail) {
  window.emailjs.send(
    'mailgun',
    templateId,
    {
      senderEmail,
      receiverEmail,
      message
    }
	).then(res => {
    console.log("Email successfully sent!");
		if (onSuccess) {
			onSuccess();
		}
  }).catch(err => {
		console.error("Failed to send email. Error: ", err);
		if (onFail) {
			onFail();
		}
	});
}

export function sendTestEmail(onSuccess, onFail, success) {
  if (success && onSuccess) {
    onSuccess();
  } else if (onFail) {
    onFail();
  }
}
