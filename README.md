## Change frequency
```
git log --format=format: --name-only | egrep -v '^$' | sort | uniq -c \ 
 | sort -r | head -20
```

## Commits by Author
```
git shortlog -sn
```