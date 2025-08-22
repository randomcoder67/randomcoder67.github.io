# Useful Snippets

Most of these aren't mine, sourced where possible

## Remove Duplicates While Maintaining Order

`cat -n file_name | sort -uk2 | sort -n | cut -f2-`

### Explanation

1. Use `cat -n` to prepend line numbers
2. Use `sort -u` remove duplicate data (`-k2` says 'start at field 2 for sort key')
3. Use `sort -n` to sort by prepended number
4. Use `cut` to remove the line numbering (`-f2-` says 'select field 2 till end')

[Credit to Digital Trauma on StackOverflow](https://stackoverflow.com/a/20639730)
