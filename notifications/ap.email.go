package notifications

import (
	"ap-gift-card-server/models"
	"fmt"
	"log"
	"net/smtp"
	"os"

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

// NotifyNewGift sends an email notification to a new ApGiftHolder.
//
// @param giftHolder *models.ApGiftHolder - the gift holder information to include in the email
//
// @return error - returns an error if the email could not be sent
func NotifyNewGift(giftHolder *models.ApGiftHolder) error {
	subject := "You've Got a Gift with A&P Nails Art!"
	body := fmt.Sprintf(
		"Hey %s,\n\n"+
        "Exciting news—your gift card has been successfully activated and is now ready to use! 🎉\n\n"+
        "🆔 Gift ID: %s\n"+
        "💳 Gift Amount: $%.2f\n\n"+
        "Thank you for choosing A&P Nails Art! We are dedicated to delivering exceptional service and creating unforgettable experiences that reflect our passion for beauty and care.\n"+
        "If you have any questions or require assistance, please don’t hesitate to get in touch. Our team is here to ensure a seamless experience.\n\n"+
        "We look forward to welcoming you soon!\n\n"+
        "Warm regards,\n"+
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

