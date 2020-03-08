# Code Formatters

## For Golang:

## For Javascript:
We use [Prettier.io](https://prettier.io/docs/en/index.html).

To format with Prettier, make sure to install `prettier` as a command line option.
This assumes you have `nodejs` and `npm` already installed.
```
sudo npm install -g prettier
```

Use this to check if all your files are correctly formatted. CircleCI build verification runs this step.
```
prettier --check "**/*.js"
```

Use this to reformat your code
```
prettier --write "**/*.js"
```

Look [here](https://prettier.io/docs/en/options.html) for more options.
For consistency, we should use the following command to format our code:
```
prettier --tab-width=4 --write "**/*.js"
```
