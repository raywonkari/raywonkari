# build README

GitHub silently released a feature, i.e., GitHub Profile READMEs.
All we have to do is, create a new repo matching the name of the username. In my case, it is raywonkari.
Then GitHub will automatically render the README in the profile overview page here: https://github.com/raywonkari

## Self Updating README 

All I have to do is merge PRs created by GH Actions
The go code is executed by Actions, and if the result shows any git diffs, it will create a PR with the changes.
I will have to review and merge them, in order for the README to update.

Also wrote some info on the Go code in my blog post here: https://raywontalks.com/github-profile-readme/
