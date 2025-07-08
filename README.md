# vulnerable-moisture-sensor

## A Go app for a moisture sensor running on Raspberry Pi for exposing two NIST CVE's
https://nvd.nist.gov/vuln/detail/CVE-2024-5433
“The Campbell Scientific CSI Web Server supports a command that will return the most recent file that matches a given expression. A specially crafted expression can lead to a path traversal vulnerability. This command combined with a specially crafted expression allows anonymous, unauthenticated access (allowed by default) by an attacker to files and directories outside of the web server root directory they should be restricted to.”

https://nvd.nist.gov/vuln/detail/CVE-2024-5434
”The Campbell Scientific CSI Web Server stores web authentication credentials in a file with a specific file name. Passwords within that file are stored in a weakly encoded format. There is no known way to remotely access the file unless it has been manually renamed. However, if an attacker were to gain access to the file, passwords could be decoded and reused to gain access.”
