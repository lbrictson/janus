# Admin Config

## General Notes on Admin Config

Modifying the admin configuration is a real time operation that directly effects the operation of your Janus server.
Be careful when changing settings, especially the security ones.

## Security

Janus supports both a built-in authentication mechanism based on email and password as well as single sign on via OAuth2.  The built-in authentication mechanism is enabled by default.

It is highly recommended that you leave the built-in authentication enabled until you have thoroughly tested your SSO config,
disabling the built-in authentication could lock you out of the system if your SSO config is not correct.  To help with this
Janus supports running both configs at the same time.

Admin configs are managed at: `http://<janus-host>/admin` and requires the global admin permission.

| Config Item            | Explanation                                                               |
|------------------------|---------------------------------------------------------------------------|
| Disable Password Login | Removes the ability to authenticate on the login page with email/password |
| Enable SSO             | Adds the SSO login option to the login screen                             |
| SSO Provider           | Your SSO provider, either entra (azuread), google or generic OIDC         |
| SSO Client ID          | Client ID from your identity provider                                     |
| SSO Client Secret      | Client secrets from your identity provider                                |
| Entra Tenant ID        | Only needed for entra, your azure tenant ID                               |
| Discovery URL          | Only neede for generic OIDC, discovery URL                                |

## SMTP

Janus notification channels share a common configuration for SMTP.  This configuration is used to send email notifications on job status updates.

| Config Item | Explanation                                             |
|-------------|---------------------------------------------------------|
| SMTP Server | The SMTP server to use for sending email                |
| SMTP Port   | The port on the SMTP server to use                      |
| SMTP User   | The username to use for SMTP authentication             |
| SMTP Pass   | The password to use for SMTP authentication             |
| SMTP From   | The email address to use as the from address for emails |

## Job Settings

Job settings covers general job settings that impact all jobs.

| Config Item               | Explanation                                                                                               |
|---------------------------|-----------------------------------------------------------------------------------------------------------|
| Default Timeout (Seconds) | When a user creates a new job this is the default value set unless they change it                         |
| Max Concurrent Jobs       | The max number of jobs that can be running on the Janus server, any job started over this limit will fail |

## Data Retention

Data retention settings control how long Janus keeps job history and job output.

| Config Item      | Explanation                                                                                         |
|------------------|-----------------------------------------------------------------------------------------------------|
| Job History Days | The number of days to keep job history in the database, older history will be automatically deleted |