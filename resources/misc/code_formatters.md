# Code Formatters

## For Golang:
```
go fmt ./...
```


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

This command will search for custom standards in a `.prettierrc.yaml` file.
```
prettier --write "**/*.js" --config ".prettierrc.yaml"
```
This will re-write all your javascript files in this directory and conform them to the standards described in a `.prettierrc.yaml` file.

In addition, you may use your TextEditor or IDE to help you. Atom and VSCode have Prettier plugins to help you keep track of your formatting.