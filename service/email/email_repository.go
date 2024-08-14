package email

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/resend/resend-go/v2"
)

type repository struct {
	db *pgxpool.Pool
}

func NewEmailRespository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) Generateotp(first_name, last_name, email, s_id string) error {

	// get 6 digit otp
	otp := getOTP()
	expires_at := time.Now().Local().Add(time.Second * time.Duration(30)).Unix()

	query := `
	INSERT INTO otp_verification
	(s_id,email,otp,expires_at)
	VALUES ($1,$2,$3,$4) RETURNING id
	`
	var id int

	// add into the database
	if err := r.db.QueryRow(context.Background(), query, s_id, email, otp, expires_at).Scan(&id); err != nil {
		log.Println("1")
		return err
	}

	if id <= 0 {
		log.Println("2")
		return errors.New("can't add otp")
	}

	// send email
	if err := sendMail(first_name, last_name, otp, email); err != nil {
		log.Println("3")
		log.Println("Failed to send email ", email, " ", err.Error())
	}

	log.Println("4")
	return nil
}

func (r *repository) VerifyOTP(otp, email string) error {

	var rotp OTP

	// get otp details
	query := `
	SELECT (s_id,otp,email,expiry_at) FROM otp_verification 
	WHERE email=$1 AND otp=$2
	`
	err := r.db.QueryRow(context.Background(), query, email, otp).Scan(&rotp.SID, &rotp.OTP, &rotp.Expires_at)
	if err != nil {
		return err
	}

	// check otp expiry
	if rotp.Expires_at < int(time.Now().Local().Unix()) {
		return errors.New("otp expired")
	}

	// compare otp
	if otp != rotp.OTP {
		return errors.New("wrong otp")
	}

	return nil
}

func getOTP() string {
	randSource := rand.NewSource(time.Now().Unix())
	randGenerator := rand.New(randSource)
	otp := randGenerator.Intn(1000000)
	return fmt.Sprintf("%06d", otp)
}

func sendMail(first_name, last_name, otp, userEmail string) error {
	var RESEND_API_KEY = os.Getenv("RESEND_API_KEY")

	client := resend.NewClient(RESEND_API_KEY)

	params := &resend.SendEmailRequest{
		From:    "onboarding@resend.dev",
		To:      []string{userEmail},
		Subject: "Email verification",
		Html: `
		<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<title>Email Verification</title>
				<style>
					body {
						font-family: Arial, sans-serif;
						background-color: #f5f5f5;
						margin: 0;
						padding: 0;
						line-height: 1.6;
						color: #333;
					}
					.container {
						width: 400px; /* Fixed width */
						margin: 20px auto;
						padding: 20px;
						background-color: #fff;
						border-radius: 8px;
						box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
					}
					h1 {
						font-size: 24px;
						text-align: center;
						margin-bottom: 20px;
					}
					.btn {
						display: inline-block;
						padding: 10px 20px;
						background-color: #007bff;
						color: white;
						text-decoration: none;
						border-radius: 4px;
						margin-top: 20px;
					}
					.btn:hover {
					  color: white;
						background-color: #0056b3;
					}
					.instructions {
						margin-top: 20px;
						font-size: 14px;
					}
			
				</style>
			</head>
			<body>
				<table style="width: 100%; height: 100%;">
					<tr>
					<td align="center" valign="top">
					<div class="container" style="background-color: #f9f9f9;">
								<h1>Verify Email Address</h1>
								<p>` + first_name + `" "` + last_name + ` , Thank you for becoming a part of Seamless Linkage of Enterprise!</p>
								<p class="instructions">If you did not sign up with us, please ignore this email.</p>
								<hr>
								  <p class="instructions"><h3>` + otp + `</h3></p>
								<hr>
								<p>You have received this email as a registered <br>user of <span style="color:blue;">SLE</span></p>
								<p>&copy; 2024 SLE. All rights reserved.</p>
							</div>
						</td>
					</tr>
				</table>
			</body>
			</html>
			`,
	}
	log.Println("RESEND 1")

	_, err := client.Emails.Send(params)
	if err != nil {
		log.Println("RESEND 2")
		return err
	}

	log.Println("RESEND 3")
	return nil
}
