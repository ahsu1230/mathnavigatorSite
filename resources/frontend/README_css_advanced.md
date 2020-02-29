# Advanced CSS

## Deprecated
The following methods to center horizontally or vertically work, but are not advised.
There have been many new, easier, and more accurate techniques to center things.
 - Using Tables (https://www.w3schools.com/tags/att_table_align.asp)
 - Using CSS float (https://www.w3schools.com/css/css_float.asp)
Instead, use the below techniques to center.

## Centering things Horizontally
(text only)
```
text-align: middle
```

```
display: block
width: ___px
margin: 0 auto
```

```
position: absolute
left: 50%
transform: translateX(-50%)
```

## Centering things Vertically
```
position: absolute
top: 50%
transform: translateY(-50%)
```

Only for aligning images with other images or text
```
display: inline-block
vertical-align: middle
```

## Flexbox
The NEW new
https://internetingishard.com/html-and-css/flexbox/
https://css-tricks.com/snippets/css/a-guide-to-flexbox/
https://css-tricks.com/css-grid-replace-flexbox/
https://hackernoon.com/the-ultimate-css-battle-grid-vs-flexbox-d40da0449faf
