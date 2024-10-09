# nd-back

**nd-back is the Go-written API by José Vicente del Valle Fayos to serve a minimalist, functional, and secure journal from a Cowboy server**. This API is designed to work in combination with its counterpart, **nd-front**.

Version 7 **includes**:

- Markdown
- On-demand loading
- Form submission protected by dynamic timeout on both client and server sides
- Authentication via JSON Web Tokens
- Colour themes
- Voice reading
- Instant search
- Visit counter
- Comments
- Navigation without page reload

This version **does not include**:

- Multimedia content publishing
- Indexing capability (the applications do not index)

For the API to function correctly, the following **environment variables** need to be defined:

- CORREO_FROM: Stores the outgoing SMTP email address in a String.
- CORREO_MAX_LLAMADAS_TRAMO_1: Stores the maximum number of calls allowed by the contact form in section 1 as a String.
- CORREO_MAX_LLAMADAS_TRAMO_2: Stores the maximum number of calls allowed by the contact form in section 2 as a String.
- CORREO_MAX_LLAMADAS_TRAMO_3: Stores the maximum number of calls allowed by the contact form in section 3 as a String.
- CORREO_MAX_LLAMADAS_TRAMO_4: Stores the maximum number of calls allowed by the contact form in section 4 as a String.
- CORREO_MAX_LLAMADAS_TRAMO_5: Stores the maximum number of calls allowed by the contact form in section 5 as a String.
- CORREO_PASS: Stores the outgoing SMTP server password in a String.
- CORREO_PORT: Stores the outgoing SMTP port number in a String.
- CORREO_TIMEOUT_TRAMO_1: Stores the maximum time between form submissions in section 1 as a String.
- CORREO_TIMEOUT_TRAMO_2: Stores the maximum time between form submissions in section 2 as a String.
- CORREO_TIMEOUT_TRAMO_3: Stores the maximum time between form submissions in section 3 as a String.
- CORREO_TIMEOUT_TRAMO_4: Stores the maximum time between form submissions in section 4 as a String.
- CORREO_TIMEOUT_TRAMO_5: Stores the maximum time between form submissions in section 5 as a String.
- CORREO_TO: Stores the recipient email address of the contact form in a String.
- CORS_DOMINIO_PERMITIDO: Stores the absolute URL of the domain where the application is hosted in a String.
- DATABASE_URL: Contains the database connection URL in a String.
- PORT: Stores the port number from which the web server listens in a String.
- REGISTRAR_ENABLED: Stores either "true" or "false" in a String, indicating whether new user registration is allowed.
- SECRET_JWT: Stores the authentication secret key in a String.
