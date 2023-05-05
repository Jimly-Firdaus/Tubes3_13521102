# GPT-666

Implementation of String Matching and Regular Expression in the Creation of a Simple ChatGPT.

# Project Description

In this project, we aim to create a simple ChatGPT using string matching and regular expression techniques. The chatbot will be able to understand and respond to user input by matching it against a predefined set of patterns using the KMP (Knuth-Morris-Pratt) and BM (Boyer-Moore) algorithms.

The KMP algorithm is an efficient string matching algorithm that can find all occurrences of a pattern within a text in linear time. It does this by preprocessing the pattern to create a partial match table, which allows it to avoid unnecessary comparisons when a mismatch occurs.

The BM algorithm is another efficient string matching algorithm that can find all occurrences of a pattern within a text in linear time. It does this by preprocessing the pattern to create two heuristics: the bad character rule and the good suffix rule. These heuristics allow it to skip over large sections of the text when a mismatch occurs.

In addition to string matching, we will also use regular expressions to extract relevant information from user input and generate appropriate responses. Regular expressions are powerful tools for pattern matching and text manipulation that can be used to perform complex tasks with just a few lines of code.

Overall, this project will demonstrate how string matching and regular expression techniques can be used to create a simple yet effective chatbot that can understand and respond to user input in a natural and intuitive manner.

# Platform

This project is deployed at [Github Pages](pemuladigital.github.io/#/)

Repository : [Github Pages Repository](https://github.com/Pemuladigital/Pemuladigital.github.io)

Backend webserver: [Netlify](https://iridescent-jalebi-788066.netlify.app)

# Web View

![WebView](https://user-images.githubusercontent.com/72639506/236265283-79b45b5c-ac11-485c-b5f0-1af227f5b762.png)


# How to run locally (just front-end side)

To run this project locally you will need to install:

- node (version ^18 or ^16 or ^14.19, preferrable on node 16)

Clone this repository. Move to the root of this project then,

```bash
npm ci          # for clean install
```

After that on the root of this repository, run:

```bash
quasar dev      # this will open browser at port 9000
```

The development server should start at `localhost:9000` (Do not change the url endpoint!)

# Additional Information

There are GPT Mode!

# Author
NIM  | Nama Anggota
------------- | -------------
13521102  | Jimly Firdaus
13521127  | Marcel Ryan Antony
13521145  | Kenneth Dave Bahana
