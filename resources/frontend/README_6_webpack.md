# Webpack
https://webpack.js.org/
https://www.youtube.com/watch?v=5zeXFC_-gMQ

Webpack is a "module bundler". As a developer, you will probably be writing lots of files of javascript, css, etc. However, these files aren't great for publishing websites because a user would have to download a lot of files just to get your website going. Imagine your user having to download all html/js/stylus files in your codebase...

Instead, webpack bundles your files into modules. You can think of modules as Javascript packages, but compacted into one Javascript file. So rather than downloading many Javascript files, a user simply downloads one optimized file to run the website. This also includes CSS. Instead of many CSS files, they only need to download one.

As a developer, this is amazing! On one hand, we can organize our code into different folders and files, but when we need to build the website for users (production mode), webpack does it for us without affecting our codebase organization.

With webpack, you can also do a lot more. One feature I really enjoy is the "watch" feature. The "watch" feature listens to any changes in your codebase and then automatically reloads the browser window so you can quickly see your changes! In other words, every time you save a file (JS, CSS, HTML), webpack will reload the browser so you see your changes immediately!
