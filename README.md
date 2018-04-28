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
- `-> label(optional)`
   Text goes here
   `->`
  - This is the spot for a *review* or *synthesis* of the notes taken.

## Reference Implementation
This repo also includes a reference implementation of the processor to take these inputs and format it into a reasonable representation. The implementation will be written in python.

## Outstanding Tasks
- Reference Implementation
- Finalization of markup