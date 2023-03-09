# email-authentication
Learning about email authentication protocols and using Golang's net libraries to validate given domains. Currently using MX records, SPF records, and DMARC records
SMFP validation protocols:
1. DNS Mail Exchange (DNS MX)
- These records dictate how an email should be directed to a mail server, and help with load-balancing for mail services. We check this to validate that the mail service is registered with DNS.
2. Sender Policy Framework (SPF). 
- SPF allows the sender to send a DNS TXT entry that essentially validates the IP-domain pair. The receiver looks up the path in the DNS entry and confirms the paths match. However, this requires both a sender and receiver IP address, so hard for high-volume email service providors (ESPs) to run. Does not work with forwarded emails, as the reciever email changes and won't match the path given from the DNS record.
3. Domain-based Message Authentication, Reporting, and Conformance (DMARC).
- Another DNS record, specifies whether a given domain is using SPF, DKIM, or both to handle authentication (used as a common framework to check both protocols). Allows senders to dictate how to handle authentication via 3 policies (none, quarantine (ex spam), and reject). Can be difficult for senders to set-up, however offers regular reports. Can also struggle with forwarding. Checking this is another system to see if an email service has been published to the DNS.

To Do:
    So far this functions as a command-line domain validation tools via DNS look ups. I could try to build this into a valid spam-interface by incorporating BIMI and DKIM verification checks. Would allow for front-end practice as well. Look into other email-authentication tools and see what they do.

