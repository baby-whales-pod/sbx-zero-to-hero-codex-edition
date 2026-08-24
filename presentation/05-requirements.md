---
marp: true
theme: default
paginate: true
html: true
style: |
  section {
    font-size: 26px;
    padding: 60px 70px;
  }
  section.lead {
    background: #15202b;
    color: #fdfdfb;
  }
  section.lead h1 {
    color: #80bfbf;
    font-size: 52px;
    line-height: 1.15;
  }
  section.lead h2,
  section.lead h3 {
    color: #e8ebda;
    font-weight: normal;
  }
  section.lead a {
    color: #80bfbf;
  }
  h1 {
    color: #15202b;
    border-bottom: 4px solid #80bfbf;
    padding-bottom: 12px;
  }
  h2 {
    color: #0062FF;
  }
  code {
    background: #e8ebda;
    border-radius: 4px;
  }
  pre {
    font-size: 22px;
  }
  section.center {
    text-align: center;
  }
  section.center h1 {
    border: none;
    font-size: 46px;
  }
  .big {
    font-size: 40px;
    color: #4503E0;
    font-weight: bold;
  }
  footer {
    color: #15202b99;
  }
---

<!-- _paginate: false -->

# Requirements - see module `00-00-requirements.md`

- Clone this repository: https://github.com/baby-whales-pod/sbx-zero-to-hero-codex-edition
- Install `sbx`, *In theory, it's already done, but it's better to have the latest version*. 
  - `0.38.0` or `0.39.0`
- GitHub account (and a PAT), *you need a GitHub account for the module `02-01`*
- Docker Hub account (for login with `sbx`)
- OpenAI API key