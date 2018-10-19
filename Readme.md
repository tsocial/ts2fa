# t2s2fa

## Libray

Use the libray in your code that depends on SRE projects

```
	rules := ts2fa.Rules{
		"/some/path": {
			"workspace123":             []string{pre_secret},
			"workspace123-layout237":   []string{},
			"*":                        []string{"some-secret"},
		},
	}

	tfa := ts2fa.New(&ts2fa.Ts2FAConf{
		Rules:     rules
	})

	valid, err := tfa.Verify(obj)
	if err != nil {
		return nil, err
	}

```

## CLI

$ ./ts2fa -email hello@trustingsocial.com -project tessellate

█████████████████████████████████████████████████████████
█████████████████████████████████████████████████████████
████ ▄▄▄▄▄ ██ █▀█ █   ██ ▄▄▄ ▄▄▀█ ▄▄ ▄▄█▄█▀▀ █ ▄▄▄▄▄ ████
████ █   █ █▄▀█▀█ ▄▀▄ ▄▀▄▀  █ ████ ▄█ █████ ▄█ █   █ ████
████ █▄▄▄█ ██▀ ▄█▄▄ ▀▄  ▀█ ▄▄▄  ▄  ▄█  ▄▄ ▀███ █▄▄▄█ ████
████▄▄▄▄▄▄▄█ ▀ █ █▄█ █▄█ ▀ █▄█ ▀▄▀ ▀▄▀ █ █▄▀ █▄▄▄▄▄▄▄████
████ ▄▄▀▄ ▄▄▄ ▀▀ ▄ ▄██ ▄█▀▄▄▄▄▄ ▄██▄▄███ ▀▄▄█ ▄██▀▄▀▀████
████ ▄▀ ▀▄▄▄█▄█▄  ▀█▄ ▀█▄▄█▄ ▀█▀▀▄ ▀ █   ▄▀▄▄▄██ ▄▄▀▄████
████ ██▀▄ ▄▄█ ▀▀▄ █ ▄ █▀▄█▄█▄██▄ ▄█  █▄ ▄█████▄▄ ▄█ ▀████
████▀ █▀ ▄▄ ▀▀▀█ █  ▄   ▄█▄ ██▄ █▀▀▄███▀██▀█▄ █▄▀██ ▄████
████ ▄   █▄    █▀▄ ▄▄▀ ▄█   ▄▄█  ██ ▄▄█▄▄▄▄██▀██▄▄█▀ ████
████▀▀ █ █▄▀▀███▀█▄▄ █▄  █▄▀▀███ ▀▄  █ ▀ █▄▀█ █ ▀▄▄▄▄████
████▄▀█  ▀▄▀ █  ▄█▄█ █▄▄▄▀▄ ▄█▄█▄▄█▄▄▄█  ▄█▄█▀▄ ▄ ██ ████
████  ▀█ ▄▄▄ ▀▀▄█ ▄█▀  ▀█▄ ▄▄▄ ▄███▀█▀ ██▄▄▄ ▄▄▄ █▄ ▄████
████▀███ █▄█ ██ ▄   █  ███ █▄█   ▄█▄▄▄▄  ▄█▀ █▄█ ██  ████
█████▄▀▄ ▄▄▄ ▄ ▄▄ ▀▀█ ▀▄█▀▄   ▄▀▀▀ ▀▀ ▄█▀▀▄▄   ▄▄█▄▄▄████
█████▀▄ ▄█▄██▀ ▄▀ █ ▄ █ █▀ █▄▄▀▄ ▄▄ ▄█▄▀▄▄███▀▄ ████▀████
████▀▄▄█▀█▄ ██ ▀▄██ ▄██▄▄▄ ▀█▄ ▄█▀▀▄▄ ▀ █▀██▀▀ █▄▄▄▄ ████
████▀█▀▀█▄▄▀ █  ▀  ▄█▄ ██ ████ ▄▄▄▄ ▄██▄▄██▀▀▄▀▄█ ▀ ▀████
████▄ ▄▄▄▀▄ ▀▄█▀▄█▄▄▄█▄  ▀██▄ ▄ ▀ ▄▄ █▄▀▀▄ █ ▄▄▄▀▀▄▄ ████
████ ██ █▀▄▀▄ ▄█ █▄█ █▄▄▄▀ ▄▄▀ ▄ ▄▄▄ █▄▄▄▄█▀▀  ██ █▀ ████
█████ ▀▀█▄▄█▄▀██▀ ▀█▀█ ▄  ▄▀█  ▀▄▀▀█  ██▄▀██▀▄ ▄▀██▄ ████
████▄▄▄███▄█ ▄ ██  ██▄ ███ ▄▄▄ ▄ ▄▄█ ██▄ ██▀ ▄▄▄ ▀▀▀▀████
████ ▄▄▄▄▄ █ ▄▀ █ ▀▄▄ ▀█▄▀ █▄█  ▀█ ▄  ▄▀ ▀▄  █▄█ ▀█▄ ████
████ █   █ █▄██ ▄ █ ▄ ██▄▄▄▄ ▄   ▀▄▄▄██▄▄▄█▀ ▄▄  ▀▀ ▀████
████ █▄▄▄█ █ ▄  ▄▄  ▄█ ██  ▄▄▄▄ ███ ▄ ▀████▄ █▄█ ▄▄▀█████
████▄▄▄▄▄▄▄█▄██▄▄▄▄▄█▄▄██▄▄███▄▄▄█▄▄▄█▄▄▄█▄█▄▄█▄███▄▄████
█████████████████████████████████████████████████████████
▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀
Enter OTP:
```
