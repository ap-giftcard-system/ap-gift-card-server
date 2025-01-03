package notifications

import (
	"ap-gift-card-server/models"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables instead")
	}
}

// sendEmail is a helper function to send an email with a customized subject and body.
//
// @param recipient string - the recipient's email address
// @param subject string - the subject of the email
// @param body string - the body of the email
//
// @return error - returns an error if the email could not be sent
func sendEmail(recipient string, subject string, body string) error {
	// Load SMTP configuration from environment variables
	apGiftCardEmail := os.Getenv("SMTP_EMAIL")
	apGiftCardEmailPassword := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.titan.email"
	smtpPort := "587"

	if apGiftCardEmail == "" || apGiftCardEmailPassword == "" || smtpHost == "" || smtpPort == "" {
		return fmt.Errorf("SMTP configuration is incomplete. Please check environment variables")
	}

	// Prepare the message
	message := []byte(fmt.Sprintf(
		"From: A&P Gift Card <%s>\r\n"+
			"To: %s\r\n"+
			"Subject: %s\r\n"+
			"\r\n%s",
		apGiftCardEmail, recipient, subject, body,
	))

	// Set up authentication
	auth := smtp.PlainAuth("", apGiftCardEmail, apGiftCardEmailPassword, smtpHost)

	// Send the email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, apGiftCardEmail, []string{recipient}, message)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}

// NotifyGiftCreated sends an email notification to a new ApGiftHolder.
//
// @param giftHolder *models.ApGiftHolder - the gift holder information to include in the email
//
// @return error - returns an error if the email could not be sent
func NotifyGiftCreated(giftHolder *models.ApGiftHolder) error {
	subject := "You've Got a Gift with A&P Nails Art!"
	body := fmt.Sprintf(
		"Hey %s,\n\n"+
        "Exciting news—your gift card has been successfully activated and is now ready to use! 🎉\n\n"+
        "🆔 Gift ID: %s\n"+
        "💳 Gift Amount: $%.2f\n\n"+
        "Thank you for choosing A&P Nails Art! We are dedicated to delivering exceptional service and creating unforgettable experiences that reflect our passion for beauty and care.\n"+
        "If you have any questions or require assistance, please don’t hesitate to get in touch. Our team is here to ensure a seamless experience.\n\n"+
        "We look forward to welcoming you soon!\n\n"+
        "Cheers 🥂\n"+
        "A&P Nails Art\n"+
        "Elevating Beauty, One Nail at a Time\n\n"+
        "--------------------------------------\n"+
        "📍 Visit Us: 308 1st Ave, Coralville, Suite 115, IA, 52241\n"+
        "📞 Call: +1 (319) 883-2322\n"+
        "📸 Instagram: https://instagram.com/ap_nails_arts\n"+
        "🌐 Book Online: https://apnailsart.com/\n"+
        "🕒 Working Hours:\n"+
        "      - Mon-Fri: 8:30 am - 6:00 pm\n"+
        "      - Sat: 8:00 am - 5:00 pm\n"+
        "      - Sun: Closed\n",
		giftHolder.HolderName, giftHolder.BarCode, giftHolder.GiftAmount,
	)

	return sendEmail(giftHolder.HolderEmail, subject, body)
}

// NotifyGiftUpdated sends an email notification to an ApGiftHolder when their gift card details are updated.
//
// @param updatedGiftHolder *models.ApGiftHolder - the updated gift holder information to include in the email
// @param originalGiftHolder *models.ApGiftHolder - the original gift holder information for comparison
//
// @return error - returns an error if the email could not be sent
func NotifyGiftUpdated(updatedGiftHolder, originalGiftHolder *models.ApGiftHolder) error {
	// Construct a clear and professional subject line
	subject := fmt.Sprintf("Your Gift Card Details Have Been Updated - ID: %s", updatedGiftHolder.BarCode)

	// Compare fields and identify changes
	changes := []string{}
	if updatedGiftHolder.HolderName != originalGiftHolder.HolderName {
		changes = append(changes, fmt.Sprintf("👤 Name changed from '%s' to '%s'", originalGiftHolder.HolderName, updatedGiftHolder.HolderName))
	}
	if updatedGiftHolder.HolderEmail != originalGiftHolder.HolderEmail {
		changes = append(changes, fmt.Sprintf("📧 Email changed from '%s' to '%s'", originalGiftHolder.HolderEmail, updatedGiftHolder.HolderEmail))
	}
	if updatedGiftHolder.HolderPhone != originalGiftHolder.HolderPhone {
		changes = append(changes, fmt.Sprintf("📞 Phone number changed from '%s' to '%s'", originalGiftHolder.HolderPhone, updatedGiftHolder.HolderPhone))
	}
	if updatedGiftHolder.GiftAmount != originalGiftHolder.GiftAmount {
		changes = append(changes, fmt.Sprintf("💳 Gift Amount changed from $%.2f to $%.2f", originalGiftHolder.GiftAmount, updatedGiftHolder.GiftAmount))
	}

	// Prepare the body with updates
	body := fmt.Sprintf(
		"Hey %s,\n\n"+
        "We would like to notify you of recent updates to your gift card:\n\n"+
	    "🆔 Gift Card ID: %s\n"+
        "%s\n\n",
		updatedGiftHolder.HolderName, updatedGiftHolder.BarCode, strings.Join(changes, "\n"),
	)

    if updatedGiftHolder.GiftAmount < originalGiftHolder.GiftAmount {
        body += "We hope you had a wonderful experience with A&P Nails Art and look forward to serving you again!\n\n"
    }

    if updatedGiftHolder.GiftAmount > originalGiftHolder.GiftAmount {
        body += "We sincerely appreciate your continued support and valued patronage!\n\n"
    }

    

	// Finish the email body
	body += fmt.Sprintf(
        "Should you need any assistance or have questions, feel free to reach out. We’re always happy to help!\n\n"+
		"Cheers 🥂\n"+
        "A&P Nails Art\n"+
        "Elevating Beauty, One Nail at a Time\n\n"+
        "--------------------------------------\n"+
        "📍 Visit Us: 308 1st Ave, Coralville, Suite 115, IA, 52241\n"+
        "📞 Call Us: +1 (319) 883-2322\n"+
        "📸 Instagram: https://instagram.com/ap_nails_arts\n"+
        "🌐 Book Online: https://apnailsart.com/\n"+
        "🕒 Our Hours:\n"+
        "      - Mon-Fri: 8:30 am - 6:00 pm\n"+
        "      - Sat: 8:00 am - 5:00 pm\n"+
        "      - Sun: Closed\n",
	)

	// Send the email
	return sendEmail(updatedGiftHolder.HolderEmail, subject, body)
}


