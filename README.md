# veda
VEDA Notational System

## Base Understanding
VEDA is an extension to Markdown and as such takes the standard markdown rules [seen here](https://github.com/adam-p/markdown-here/wiki/Markdown-Cheatsheet). This is a living document at this point and should not be considered a final draft of the markup language.

It adds some functionality based upon additional markup.
- `{$tag}`
  - This introduces tagging into the markup and a processor should be able to pull tags into an overall hierarchy
- `@Callout text to be used inline@`
  - This is to be used to allow for inline text to be used as a *cue* for later review.
- `<- Other ## header`

   Text goes here

   `<-`
  - This means the enclosed block of text references a previous header and should be viewed as an *effect* or *subordinate* to it.
- `-> Other ## header`

   Text goes here

   `->`
  - This means the enclosed block of text references a previous header and should be viewed as an *cause* or *central* to it.
- `-! label(optional)`

   Text goes here

   `-!`
  - This is the spot for a *review* or *synthesis* of the notes taken.
- `{description of link}(context(optional)/topic/header)`
  - This will create a link to another document
  - This can be used within subordinate or central blocks to tie notes together.

## File Structure
VEDA is intended to be a method for large scale note taking, helping to tie together disparate collections of notes. This can either be *contextual* or *context-free*, meaning you can have multiple subjects or contexts shared within one folder structure OR you can limit it to one context, simplifying the overall implmentation. One of the benefits to VEDA is that it is a linked system, tying the various components together to better handle large problem sets. A contextual design allows for cross-context tagging and should allow for more complex synthesis of information.

### Contextual
```
\docroot\.
\docroot\toc.vd - Table of Contents
\docroot\tag.vd - Tag Mapping
\docroot\contexts.vd - Context Map
\docroot\summary.vd - Summary Document
\docroot\[context]\toc.vd - Table of Contents
\docroot\[context]\tag.vd - Tag Mapping
\docroot\[context]\summary.vd - Summary Document
\docroot\[context]\[topic].vd - Notes for [topic]
```

### Context-free
```
\docroot\.
\docroot\toc.vd - Table of Contents
\docroot\tag.vd - Tag Mapping
\docroot\summary.vd - Summary Document
\docroot\[topic].vd - Notes for [topic]
```

## Reference Implementation
This repo also includes a reference implementation of the processor to take these inputs and format it into a reasonable representation. The implementation will be written in python.

## Outstanding Tasks
- Tag, ToC, and Summary documents
- Reference Implementation
- Finalization of markup