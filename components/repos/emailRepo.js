'use strict';

export function sendEmail(message, onSuccess, onFail) {
	const templateId = "mathnavigatorwebsitecontact";
	const receiverEmail = "andymathnavigator@gmail.com";
	const senderEmail = "anonymous@andymathnavigator.com";
	var templateParams = {
		emailMessage: message
	};

  window.emailjs.send(
    'mailgun',
    templateId,
		templateParams
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
