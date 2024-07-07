## Commit Rules

### Header

every commit messages must have header, these includes:

- **INIT**: for initial
- **NEW**: for adding new feature/function
- **UPDATE**: for updating existing feature/function
- **FIX**: for fixing minor bug
- **MERGE**: for merging branches
- **DONE**: for informing a finished feature, usually happen at the end

for example, you want commit a new feature, type this template:

```
git commit -m "NEW: implement GetbyID function"
```

or perhaps you made the function and template but haven't working on the logic, type these instead:

```
git commit -m "INIT: add GetbyID function"
```

for committing multiple files, use more detailed commit message, like this one:

```
git commit -m "UPDATE: update logic for GetbyID, Create, and Delete users function. Update ENV"
```

the message above is pretty detailed but it is not recommended. Makes sure to split up your recent changes by separating newly updated feature in each commit messages. by following these rules, the message above can be more readable, like this:

```
git commit -m "UPDATE: update logic for GetbyID, Create, and Delete users function"
```

```
git commit -m "UPDATE: update ENV"
```

see the difference? These can be turned into more readable message by committing each feature as separate commit. But the message above is enough.

### Message

commit message should be clear and simple. Also don't forget to reference what are you committing. Here are some examples:

- Correct:

    - `Add user authentication feature`
    - `Fix bug in property listing logic`
    - `Update database schema to include reviews`
    - `Remove deprecated API endpoints`
    - `Improve search functionality performance`

- Incorrect:

    - `adding new feature` (should be imperative: Add new feature)
    - `fixed bug` (should be imperative: Fix bug)
    - `Update database schema to include reviews.` (no period at the end)

thats it, happy coding guys!