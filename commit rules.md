## Commit Rules

### Header

every commit messages must have header, these includes:

- **INIT**: for initial
- **NEW**: for adding new feature/function
- **UPDATE**: for updating existing feature/function
- **BUGFIX**: for fixing minor bug
- **MERGE**: for merging branches
- **DONE**: for informing a finished feature, usually happen at the end

for example, you want commit a new feature, type this template:

```
git commit -m "NEW: implemented GetbyID function"
```

or perhaps you made the function and template but haven't working on the logic, type these instead:

```
git commit -m "INIT: added GetbyID function"
```

for committing multiple files, use more detailed commit message, like this one:

```
git commit -m "UPDATE: updated logic for GetbyID, Create, and Delete users function. Updated ENV"
```

the message above is pretty detailed but it is not recommended. Makes sure to split up your recent changes by separating newly updated feature in each commit messages. by following these rules, the message above can be more readable, like this:

```
git commit -m "UPDATE: updated logic for GetbyID, Create, and Delete users function"
```

```
git commit -m "UPDATE: Updated ENV"
```

see the difference? These can be turned into more readable message by committing each feature as separate commit. But the message above is enough.

thats it, happy coding guys!