package notification

import "fmt"

func OtpHtmlEmailTemplate(otpCode string) string {
	var letter = fmt.Sprintf(`
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html dir="ltr" lang="en">
  <head>
    <link
      rel="preload"
      as="image"
      href="https://www.mobarter.com/images/logo/logo.svg"
    />
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
    <meta name="x-apple-disable-message-reformatting" />
    <!--$-->
  </head>
  <div
    style="
      display: none;
      overflow: hidden;
      line-height: 1px;
      opacity: 0;
      max-height: 0;
      max-width: 0;
    "
  >
    Mobarter Email Verification
  </div>
  <body
    style="
      background-color: #ffffff;
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto,
        Oxygen-Sans, Ubuntu, Cantarell, 'Helvetica Neue', sans-serif;
    "
  >
    <table
      align="center"
      width="100%"
      border="0"
      cellpadding="0"
      cellspacing="0"
      role="presentation"
      style="max-width: 560px; margin: 0 auto; padding: 20px 0 48px"
    >
      <tbody>
        <tr style="width: 100%">
          <td>
            <img
              alt="mobarter"
              height="42"
              src="https://www.mobarter.com/images/logo/logo.svg"
              style="
                display: block;
                outline: none;
                border: none;
                text-decoration: none;
                border-radius: 21px;
                width: 42px;
                height: 42px;
              "
              width="42"
            />
            <h1
              style="
                font-size: 24px;
                letter-spacing: -0.5px;
                line-height: 1.3;
                font-weight: 400;
                color: #484848;
                padding: 17px 0 0;
              "
            >
              Mobarter Email Verification
            </h1>

            <p
              style="
                font-size: 15px;
                line-height: 1.4;
                margin: 0 0 15px;
                color: #3c4149;
              "
            >
              We want to make sure it&#x27;s really you. Please enter the
              following verification code when prompted. If you don&#x27;t want
              to create an account, you can ignore this message.
            </p>
            <div
              style="
                display: flex;
                width: 100%;
                align-items: center;
                justify-content: center;
              "
            >
              <code
                style="
                  font-family: monospace;
                  font-weight: 700;
                  padding: 5px 8px;
                  background-color: #dfe1e4;
                  letter-spacing: -0.3px;
                  font-size: 21px;
                  margin-bottom: 20px;
                  border-radius: 4px;
                  color: #3c4149;
                "
                >` + otpCode + `</code
              >
            </div>

            <p
              style="
                font-size: 12px;
                line-height: 1.4;
                margin-top: 20px;
                color: #3c4149;
              "
            >
              Mobarter will never email you and ask you to disclose or verify
              your password credit card or banking account number.
            </p>

            <p
              style="
                font-size: 12px;
                line-height: 1.4;
                margin: 0 0 15px;
                color: #3c4149;
              "
            >
              This message was produced and distributed by Mobarter, Inc. FCT
              Nigeria, Inc.. All rights reserved.
            </p>
            <hr
              style="
                width: 100%;
                border: none;
                border-top: 1px solid #eaeaea;
                border-color: #dfe1e4;
                margin: 42px 0 26px;
              "
            />
            <a
              href="https://mobarter.com"
              style="
                color: #b4becc;
                text-decoration-line: none;
                font-size: 14px;
              "
              target="_blank"
              >Mobarter</a
            >
          </td>
        </tr>
      </tbody>
    </table>
    <!--/$-->
  </body>
</html>

    `)

	return letter
}
