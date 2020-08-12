package ligolo

const CertPEM = `-----BEGIN CERTIFICATE-----
MIIFnTCCA4WgAwIBAgIUb8auVNIbbScZ3Si/Q3fclASF3U4wDQYJKoZIhvcNAQEL
BQAwXjELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQHDAZBdXN0
aW4xFDASBgNVBAoMC0RldmVsb3BtZW50MRgwFgYDVQQDDA93d3cuZXhhbXBsZS5j
b20wHhcNMTkwODE2MTQyNDIwWhcNMjAwODE1MTQyNDIwWjBeMQswCQYDVQQGEwJV
UzEOMAwGA1UECAwFVGV4YXMxDzANBgNVBAcMBkF1c3RpbjEUMBIGA1UECgwLRGV2
ZWxvcG1lbnQxGDAWBgNVBAMMD3d3dy5leGFtcGxlLmNvbTCCAiIwDQYJKoZIhvcN
AQEBBQADggIPADCCAgoCggIBAPk94FxUqTYZhdaqVQEIrS/yE7BjO9RMsM1nhZ5n
sM0s0noWMpJmQMGRoOP+uqIFanUT2FHfMe1iSRsmeALOB1hkbZsIGpwS6MTaS0vu
VEVACENJebJz+aNtq5QuUA66C0T+IRvxXolynx+5gWvqoWeq3T3bGOlI4MI8n4Xe
ywGvqVW57M3sp+G7Avs/K0PbmcX4qEC6N+i46oE1Q/siuFz21NeoPpOKgCPMEzPG
P1H9o2JYMp0NODp69HOYBQlPH7z1pvtZca7mS+QJWIGYV9mJ9itN4LmnvypSu6jJ
aavw/s2OGz9lbqpmZQT/3SwB3Mr/lfCdyL03J/JhUJsA4SIG2Xrc/Vu/n9QH8TS7
gzO003wf+bz4PGP6HnIVWLqVpPLOa4mz6KUlcdtiw34/tIjPqobtRRXvo1NVObEF
HRVXDV6CrkZQ+9hFS3lC7GR6i3XieT/j8j9JONz/vvewzhnQ/KyO1MtswEuj8Sbn
1pTJ5dltzG5x3TjcvCwdDk2W+yBzoTqNZtB+FgniVY/FmNFPnhkKQYn86cp5bDIe
5Y2pJ5Nm8uhsi1ILDCrpGPhRUu9/0trxpDKwKcSCuxsb7SrBP/C4iGi5nyi+3A9C
1HRow8bFK8iochM4uAi4s1jYUD8taAooT5vtLkeLUQL/GT6fxwXnb/6P0nf0bR08
+7wFAgMBAAGjUzBRMB0GA1UdDgQWBBRhnTJdK2Ccq18YZkWl37H3M6U9CjAfBgNV
HSMEGDAWgBRhnTJdK2Ccq18YZkWl37H3M6U9CjAPBgNVHRMBAf8EBTADAQH/MA0G
CSqGSIb3DQEBCwUAA4ICAQBCbyButZbs8cAbyPdCGpstVAJvbjQb5FvFPYSGnXHj
mmE7MoYsi+GgL4IfQv62FONDYMMD7LtEgRP+Ab5i23Pficq+rs/cZKlgYuzJyv70
I4xaSk9m4ejohJsUA/l2aCtLXJDg+PCYEa2e8PnlOg2Cely/QXg/izHagnmFBvMu
EfW/9mf0KTASz5GF5C/I/Y3uh/hXvhrpyjksgfyvLu3cJxWeeqh2GaOA6KgQikkR
OfsmFHOrozuyDqLt+GK09zg69l0Tyy4uqgUk2ol8KvmB7fFvKfFxPAap7nwZTwIJ
pB+sdVkRao9schLrn/Df4smUpGcM/Vsq8rVKG+yfIRVfp/PVMwEyDqVcklMo0HF7
syJ6bRZNAdAXIVlP+uvOJYVvrc9tr3mO0hAGP74t/C/uNQbJ91/uPKnh7pKJS8Xb
xgdRKRyGw6FVddJdy7x9UP3KQm5Pb9/v8Af3O7v9ye0VoiZMUT+PpugyIXLdTeSF
7Ti6GyU8oW1CYD3rTTcjIh8POOLBeCVgdhrpo9ATzbmzVg2E+kZcoenU/QuudDM2
vkTkzF4KBmk4UYPPRqgt9UgGMFnkRonOCSVKqJo4JHRnjF+Bh+6VwkjNFLqj8Oyu
suZPV64v0uhwk0oiOcFV1f/7I+l9T+sSElnP9hhSsQHQOulURMRQUwRGhFyX0/Wu
IA==
-----END CERTIFICATE-----`
const KeyPEM = `-----BEGIN PRIVATE KEY-----
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
jWHDR/wbFe00CkzhKvFxOHlilDkQc8bG
-----END PRIVATE KEY-----`
