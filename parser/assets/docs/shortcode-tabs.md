---
title: Interactive Tabs
tags: [features, demo]
---
Simulate a tabbed interface for switching between multiple options (e.g. terminal commands for different platforms or package managers).

{{ tabs }}
  {{ tab title="npm" }}
  ```bash
  $ npm install @tamarind/core
  ```
  {{ /tab }}
  {{ tab title="yarn" }}
  ```bash
  $ yarn add @tamarind/core
  ```
  {{ /tab }}
  {{ tab title="pnpm" }}
  ```bash
  $ pnpm add @tamarind/core
  ```
  {{ /tab }}
{{ /tabs }}

### Usage

<pre><code>&#123;&#123; tabs &#125;&#125;
  &#123;&#123; tab title="npm" &#125;&#125;
  ```bash
  $ npm install @tamarind/core
  ```
  &#123;&#123; /tab &#125;&#125;
  &#123;&#123; tab title="yarn" &#125;&#125;
  ```bash
  $ yarn add @tamarind/core
  ```
  &#123;&#123; /tab &#125;&#125;
&#123;&#123; /tabs &#125;&#125;</code></pre>
