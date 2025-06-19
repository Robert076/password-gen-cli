# 🚀 password-gen-cli

A cli tool that creates a password based on the flags you provided. You can choose to (not) have special chars, digits, and you can also choose the length of the password.

⚙️ Running the setup (macOS):

1. Build the project

```bash
go build
```

2. Compile and install app in bin

```bash
go install
```

3. Create passwords:

```bash
password-gen-cli generate
```

or longer ones

```bash
password-gen-cli generate -l 36
```

or more complicated ones

```bash
password-gen-cli generate -d
```

or even more complicated ones

```bash
password-gen-cli generate -d -s
```
