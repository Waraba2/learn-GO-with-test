# learn-GO-with-test
Test driven learning using  golang

## TDD Test Driven Development
Denis Yu
[Slime] (https://deniseyu.github.io/leveling-up-tdd/) until you got and interface
> Sliming is useful for giving a “skeleton” to your object. Designing an interface and executing logic are 
> two concerns, and sliming tests strategically lets you focus on one at a time.
Coupling and cohesion
### Acceptance Test
### Approval Test

## DRY Do Not Repeat Yourself

## Dependencie Injection

## reading-files
[Discussion on file system interfaces](https://github.com/golang/go/issues/41190)

## blogrenderer(templating)
[SPA](https://en.wikipedia.org/wiki/Single-page_application): single page application.  
[Hotwire](https://hotwired.dev/): to make single page applications dynamic.  
[Mustach](https://mustache.github.io/): templating language.  
[text/templating](https://pkg.go.dev/text/template): Go docs.  
[Go Approval Tests](https://github.com/approvals/go-approval-tests): Go approval tests.  
> ApprovalTests allows for easy testing of larger objects, strings and anything else that can be saved to a
> file (images, sounds, CSV, etc...)
[Emily Blanche](https://twitter.com/emilybache) has an interesting [video where she uses approval tests to add an incredibly extensive set
of tests to a complicated codebase that has zero tests](https://www.youtube.com/watch?v=zyM2Ep28ED8). **"Combinatorial Testing"** is definitely something 
worth looking into.  
[A view moldel](https://stackoverflow.com/questions/11064316/what-is-viewmodel-in-mvc/11074506#11074506): keep in mind while dealing with templates  
[A God object](https://en.wikipedia.org/wiki/God_object)  
### Things Learned 
- What we've learned
- How to create and render HTML templates.
- How to compose templates together and DRY up related markup and help us keep a consistent look and feel.
- How to pass functions into templates, and why you should think twice about it.
- How to write "Approval Tests", which help us test the big ugly output of things like template renderers.
## Generics
A common path I've taken to see if you need to use generics or not:
- One TDD cycle to drive some behaviour
- Another TDD cycle to exercise some other related scenarios
> Hmm, these things look similar - but a little duplication is better than coupling to a bad abstraction
- Sleep on it
- Another TDD cycle
> OK, I'd like to try to see if I can generalise this thing. Thank goodness I am so smart and good-looking 
> because I use TDD, so I can refactor whenever I wish, and the process has helped me understand what 
> behaviour I actually need before designing too much.
- This abstraction feels nice! The tests are still passing, and the code is simpler
- I can now delete a number of tests, I've captured the essence of the behaviour and removed unnecessary detail
## Further Reading
[io.fs](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/)
[John Calhoun's 'Learn Web Development with Go'](https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/) has a number of excellent articles on templating.
[Hotwire](https://hotwired.dev/) - You can use these techniques to create Hotwire web applications. It has 
been built by Basecamp who are primarily a Ruby on Rails shop, but because it is server-side, we can use it 
with Go
