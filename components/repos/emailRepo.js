'use strict';

export function sendEmail(message, onSuccess, onFail) {
	if (process.env.NODE_ENV === 'development') {
		sendTestEmail(true, onSuccess, onFail);
	} else {
		sendProdEmail(message, onSuccess, onFail);
	}
}

function sendProdEmail(message, onSuccess, onFail) {
	console.log("Sending prod email...");
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

function sendTestEmail(success, onSuccess, onFail) {
	console.log("Sending test email...");
  if (success && onSuccess) {
    onSuccess();
  } else if (onFail) {
    onFail();
  }
}
