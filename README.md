# FORCARD - A Digital Card / Mini website generator.
I still haven't figured out how to implement this actually.

Code Gen? OR just store user preferences in the DB and build a UI upon them. Build pages with componenets, not a code gen actually.

My initial idea involved code_gen I think. The data source was to be included in a simple JSON file.

```json
{
    "site": {
        "name": "The Accountant",
        "meta": "Something about that too",
        "background": "#f9f9f9",
        "font-family": "Open Sans"
    },
    "title": {
        "el": "h1",
        "value": "Accounting Professional"
    },
    "subtitle": {
        "el": "h2",
        value: "Shugi Lior"
    },
    "text": {
        "el": "p",
        value: "My Name Is Lior Shugi, I consult for Fortune 500 companies on all things related to Finances, Big money decisions, internal and external. I'm reliable and on point, I'm literally the last guy you'll ever need and hire",
        "properties": {
            "color": "#010101",
            "background": "#f9f9f9",
            "font-size": "22px"
        },
    },
    "button": {
        "el": "button",
        value: "Book Now",
        properties: {
            "padding":"10px 16px",
            "background": "#bd4c6d",
            "text-align": "center",
        }
    }
}
```

# Challenges

Questions on how to add javascript/CSS:
1. Should Front-End decide and encode the javaScript code based on user
   preferences? eventListeners, interactivity.
2. Should Back-End do this? Escape JavaScript as a file or a script tag
   inside the template?

## Kinds of interactivity
- Links should come from the client as a URL. Easy.
- Buttons?
- Menus
- Animations

## Stylings
- Inline styles +
- Seperate css file (provided by client?) ++
- Include style tags (like Astro/Svelte do)
- 

TODO: Use [Hugo](https://github.com/gohugoio/) as an educatinal resource

Go program should parse the data from a json file and map it into HTML
elements, or astro components and append stylesheets according to user
preferences

Is there a general theme style I'm after?

- Digital Business Cards / Mini Sites
- Full blown landing pages
- Landing pages based on specific Themes, and are customizable only
    according to those specific themes.
- Linkpop-style one pagers for eCommerce.
