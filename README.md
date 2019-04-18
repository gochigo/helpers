# Helpers

[![Build Status](https://dev.azure.com/markbates/buffalo/_apis/build/status/gobuffalo.helpers?branchName=master)](https://dev.azure.com/markbates/buffalo/_build/latest?definitionId=49&branchName=master)[![GoDoc](https://godoc.org/github.com/gobuffalo/helpers?status.svg)](https://godoc.org/github.com/gobuffalo/helpers)
---
Note: This file is auto-generated. Do Not Edit
---


## [`bootstrap#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#Form)
<p>Form implements a <code>github.com/gobuffalo/plush</code> helper around the
bootstrap.New function in the <code>github.com/gobuffalo/tags/form/bootstrap</code> package</p>


## [`bootstrap#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#FormFor)
<p>FormFor implements a <code>github.com/gobuffalo/plush</code> helper around the
bootstrap.NewFormFor function in the <code>github.com/gobuffalo/tags/form/bootstrap</code> package</p>


## [`bootstrap#New`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#New)
<p>New returns a map of the helpers within this package.</p>


## [`bootstrap#RemoteForm`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#RemoteForm)
<p>RemoteForm implements a <code>github.com/gobuffalo/plush</code> helper around the
bootstrap.New function in the <code>github.com/gobuffalo/tags/form/bootstrap</code> package</p>


## [`bootstrap#RemoteFormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms/bootstrap#RemoteFormFor)
<p>FormFor implements a <code>github.com/gobuffalo/plush</code> helper around the
bootstrap.NewFormFor function in the <code>github.com/gobuffalo/tags/form/bootstrap</code> package</p>

<pre><code>&lt;%= remoteFormFor(widget, {method: &#34;POST&#34;}) { %&gt;
&lt;% } %&gt;
&lt;form action=&#34;/widgets/b6b0ab24-19ae-4cdd-ad73-c5ecbddd6f91&#34; id=&#34;widget-form&#34; method=&#34;POST&#34;&gt;&lt;input name=&#34;_method&#34; type=&#34;hidden&#34; value=&#34;PUT&#34;&gt;&lt;/form&gt;
</code></pre>


## [`content#ContentFor`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentFor)
<p>ContentFor stores a block of templating code to be re-used later in the template
via the contentOf helper.
An optional map of values can be passed to contentOf,
which are made available to the contentFor block.</p>

<pre><code>&lt;% contentFor(&#34;buttons&#34;) { %&gt;
    &lt;button&gt;hi&lt;/button&gt;
&lt;% } %&gt;
</code></pre>


## [`content#ContentOf`](https://godoc.org/github.com/gobuffalo/helpers/content#ContentOf)
<p>ContentOf retrieves a stored block for templating and renders it.
You can pass an optional map of fields that will be set.</p>

<pre><code>&lt;%= contentOf(&#34;buttons&#34;) %&gt;
&lt;%= contentOf(&#34;buttons&#34;, {&#34;label&#34;: &#34;Click me&#34;}) %&gt;
</code></pre>


## [`content#New`](https://godoc.org/github.com/gobuffalo/helpers/content#New)
<p>New returns a map of the helpers within this package.</p>


## [`debug#Debug`](https://godoc.org/github.com/gobuffalo/helpers/debug#Debug)
<p>Debug by verbosely printing out using &#39;pre&#39; tags.</p>


## [`debug#Inspect`](https://godoc.org/github.com/gobuffalo/helpers/debug#Inspect)
<p>Inspect the interface using the <code>%+v</code> formatter</p>


## [`debug#New`](https://godoc.org/github.com/gobuffalo/helpers/debug#New)
<p>New returns a map of the helpers within this package.</p>


## [`encoders#New`](https://godoc.org/github.com/gobuffalo/helpers/encoders#New)
<p>New returns a map of the helpers within this package.</p>


## [`encoders#Raw`](https://godoc.org/github.com/gobuffalo/helpers/encoders#Raw)
<p>Raw converts a <code>string</code> to a <code>template.HTML</code></p>


## [`encoders#ToJSON`](https://godoc.org/github.com/gobuffalo/helpers/encoders#ToJSON)
<p>ToJSON marshals the interface{} and returns it
as template.HTML</p>


## [`env#New`](https://godoc.org/github.com/gobuffalo/helpers/env#New)
<p>New returns a map of the helpers within this package.</p>


## [`escapes#HTMLEscape`](https://godoc.org/github.com/gobuffalo/helpers/escapes#HTMLEscape)
<p>HTMLEscape will escape a string for HTML</p>


## [`escapes#New`](https://godoc.org/github.com/gobuffalo/helpers/escapes#New)
<p>New returns a map of the helpers within this package.</p>


## [`forms#Form`](https://godoc.org/github.com/gobuffalo/helpers/forms#Form)
<p>Form implements a <code>github.com/gobuffalo/plush</code> helper around the
form.New function in the <code>github.com/gobuffalo/tags/form</code> package</p>


## [`forms#FormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms#FormFor)
<p>FormFor implements a <code>github.com/gobuffalo/plush</code> helper around the
form.NewFormFor function in the <code>github.com/gobuffalo/tags/form</code> package</p>


## [`forms#New`](https://godoc.org/github.com/gobuffalo/helpers/forms#New)
<p>New returns a map of the helpers within this package.</p>


## [`forms#RemoteForm`](https://godoc.org/github.com/gobuffalo/helpers/forms#RemoteForm)
<p>RemoteForm implements a <code>github.com/gobuffalo/plush</code> helper around the
form.New function in the <code>github.com/gobuffalo/tags/form</code> package</p>


## [`forms#RemoteFormFor`](https://godoc.org/github.com/gobuffalo/helpers/forms#RemoteFormFor)
<p>RemoteFormFor implements a <code>github.com/gobuffalo/plush</code> helper around the
form.NewFormFor function in the <code>github.com/gobuffalo/tags/form</code> package</p>

<pre><code>&lt;%= remoteFormFor(widget, {method: &#34;POST&#34;}) { %&gt;
&lt;% } %&gt;
&lt;form action=&#34;/widgets/b6b0ab24-19ae-4cdd-ad73-c5ecbddd6f91&#34; id=&#34;widget-form&#34; method=&#34;POST&#34;&gt;&lt;input name=&#34;_method&#34; type=&#34;hidden&#34; value=&#34;PUT&#34;&gt;&lt;/form&gt;
</code></pre>


## [`inflections#New`](https://godoc.org/github.com/gobuffalo/helpers/inflections#New)
<p>New returns a map of the helpers within this package.</p>


## [`iterators#Between`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Between)
<p>Between will iterate up to, but not including <code>b</code></p>

<pre><code>Between(0,10) // 0,1,2,3,4,5,6,7,8,9
</code></pre>


## [`iterators#GroupBy`](https://godoc.org/github.com/gobuffalo/helpers/iterators#GroupBy)


## [`iterators#New`](https://godoc.org/github.com/gobuffalo/helpers/iterators#New)
<p>New returns a map of the helpers within this package.</p>


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#Next`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Next)


## [`iterators#Range`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Range)


## [`iterators#Until`](https://godoc.org/github.com/gobuffalo/helpers/iterators#Until)
<p>Until will iterate up to, but not including <code>a</code></p>

<pre><code>Until(3) // 0,1,2
</code></pre>


## [`meta#Len`](https://godoc.org/github.com/gobuffalo/helpers/meta#Len)


## [`meta#New`](https://godoc.org/github.com/gobuffalo/helpers/meta#New)
<p>New returns a map of the helpers within this package.</p>


## [`paths#New`](https://godoc.org/github.com/gobuffalo/helpers/paths#New)
<p>New returns a map of the helpers within this package.</p>


## [`paths#PathFor`](https://godoc.org/github.com/gobuffalo/helpers/paths#PathFor)
<p>PathFor takes an <code>interface{}</code>, or a <code>slice</code> of them,
and tries to convert it to a <code>/foos/{id}</code> style URL path.
Rules:</p>

<ul>
<li>if <code>string</code> it is returned as is</li>
<li>if <code>Pathable</code> the <code>ToPath</code> method is returned</li>
<li>if <code>slice</code> or an <code>array</code> each element is run through the helper then joined</li>
<li>if <code>struct</code> the name of the struct, pluralized is used for the name</li>
<li>if <code>Paramable</code> the <code>ToParam</code> method is used to fill the <code>{id}</code> slot</li>
<li>if <code>struct.Slug</code> the slug is used to fill the <code>{id}</code> slot of the URL</li>
<li>if <code>struct.ID</code> the ID is used to fill the <code>{id}</code> slot of the URL</li>
</ul>


## [`tags#CSS`](https://godoc.org/github.com/gobuffalo/helpers/tags#CSS)


## [`tags#Img`](https://godoc.org/github.com/gobuffalo/helpers/tags#Img)


## [`tags#JS`](https://godoc.org/github.com/gobuffalo/helpers/tags#JS)


## [`tags#LinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#LinkTo)


## [`tags#New`](https://godoc.org/github.com/gobuffalo/helpers/tags#New)
<p>New returns a map of the helpers within this package.</p>


## [`tags#RemoteLinkTo`](https://godoc.org/github.com/gobuffalo/helpers/tags#RemoteLinkTo)
<p>RemoteLinkTo creates an AJAXified  tag.</p>

<pre><code>&lt;%= remoteLinkTo(widget, {class: &#34;btn btn-info&#34;, body: &#34;View&#34;}) %&gt;
&lt;a class=&#34;btn btn-info&#34; data-remote=&#34;true&#34; href=&#34;/widgets/b6b0ab24-19ae-4cdd-ad73-c5ecbddd6f91&#34;&gt;View&lt;/a&gt;
</code></pre>


## [`text#Markdown`](https://godoc.org/github.com/gobuffalo/helpers/text#Markdown)
<p>Markdown converts the string into HTML using GitHub flavored markdown.</p>


## [`text#New`](https://godoc.org/github.com/gobuffalo/helpers/text#New)
<p>New returns a map of the helpers within this package.</p>


## [`text#Truncate`](https://godoc.org/github.com/gobuffalo/helpers/text#Truncate)


