---
name: pr-format
description: PR formatting conventions for this project
---

PR titles MUST start with the GitHub issue number in the format `#N: `.
For example: `#5: Add Reverse function`

PR descriptions MUST include a `## Checklist` section with these items:
- [ ] Tests pass (`go test ./...`)
- [ ] No vet warnings (`go vet ./...`)
- [ ] New functions have test coverage
