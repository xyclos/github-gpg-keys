# A simple CLI tool for fetching gpg keys for a user on GitHub

This tool gets all public gpg keys for a user on github from the github api: 
https://developer.github.com/v3/users/gpg_keys/#list-gpg-keys-for-a-user and
saves them to files in the current directory.

The keys can then be easily imported using the gpg cli

## Usage 

```text
Usage of github-gpg-keys
  -s boolean	
        Save keys to current directory
  -t int
    	Client timeout in seconds (default 30)
  -u string
    	Github username to get keys for
```

## Example

```shell script
# first save all the keys to files for a github user
github-gpg-keys -u xyclos -s

# then import them all
for f in *.gpg; do gpg --import $f; done
```

## TODO
- add option to filter keys by email address
- add option to automatically import to gpg
- add option to choose a directory to save to
