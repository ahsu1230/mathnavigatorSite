## CircleCI

https://circleci.com/

CircleCI is a Continuous Delivery (CD) and Continuous Integration (CI) tool. It is often used for automatically building a project every time a new feature is merged into Github.

Usually, if you wanted to build your app after a new feature, you would have to call a set of commands to test your app and then create the production version of your app.

CD/CI tools like CircleCI help you automatically do this and can automatically run tests for you to see if your new feature broke any existing features. This is especially useful for preventing developers from accidentally pushing in changes before passing critical tests.

A CircleCI build triggers whenever a new commit is pushed onto a remote [Github] branch. These tools simply wait for new changes to the codebase (whether merged into master or not) and activate its tests immediately.
