package templates

import "html/template"

// Generates a Basic Email text template

const PassRecoverTmpl = `Hi {{.Username }},

We've received a password reset from this account.
Please click on the link below to reset your password.

http://localhost:3000/reset-password?token={{.Token}} 

If you did not request a password reset, please ignore this email. {{printf "\n"}}
Thanks,
The Go Server Team`

var PasswordRecoveryEmail = template.Must(template.New("passwordRecovery").Parse(PassRecoverTmpl))
