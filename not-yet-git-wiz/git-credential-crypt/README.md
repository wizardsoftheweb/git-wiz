# `git-credential-crypt`

This is still in the very early stages of planning and discovery. I'd like to build something that can be used as a git credential store that's also more secure than the always-on-always-plaintext `store` but with more secure permanence than leaving `cache` running in perpetuity.

If you know how the directory is laid out and what to wire up, this has [some of the functionality of `store` and other helpers](https://git-scm.com/docs/api-credentials#_credential_helpers). It can `get`, `store`, and `erase` creds. However, right now, the files need to be manually specified and the functions manually run to do everything. Some of the core functionality of `store` has been drafted but not completed, such as building the default files with they DNE, looking in multiple locations for a requested credential, and erasing a credential in all known locations.
