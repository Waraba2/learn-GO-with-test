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

## Revisiting arrays with generics
Don't conflate easiness, with simplicity. Doing loops and copy-pasting code is easy, but it's not necessarily 
simple. For more on simple vs easy, watch [Rich Hickey's masterpiece of a talk - Simple Made Easy](https://www.youtube.com/watch?v=SxdOUGdseq4).  
Don't conflate unfamiliarity, with complexity. Fold/reduce may initially sound scary and computer-sciencey 
but all it really is, is an abstraction over a very common operation. Taking a collection, and combining it 
into one item. When you step back, you'll realise you probably do this a lot.  
[The identity element](https://en.wikipedia.org/wiki/Identity_element):
> In mathematics, an identity element, or neutral element, of a binary operation operating on a set is an 
> element of the set which leaves unchanged every element of the set when the operation is applied.
Fold/reduce are pretty universal
The possibilities are endless™️ with Reduce (or Fold). It's a common pattern for a reason, it's not just for  
arithmetic or string concatenation. Try a few other applications.
- Why not mix some color.RGBA into a single colour?
- Total up the number of votes in a poll, or items in a shopping basket.
- More or less anything involving processing a list.
### Resources
Fold is a real fundamental in computer science. Here's some interesting resources if you wish to dig more into it
[Wikipedia: Fold](https://en.wikipedia.org/wiki/Fold)
[A tutorial on the universality and expressiveness of fold](http://www.cs.nott.ac.uk/~pszgmh/fold.pdf)
## Testing fundamentals
### Acceptance test
What are they?
Acceptance tests are a kind of "black-box test". They are sometimes referred to as "functional tests". They should   
exercise the system as a user of the system would.The term "black-box" refers to the idea that the test code has   
no access to the internals of the system, it can only use its public interface and make assertions on the 
behaviours it observes. This means they can only test the system as a whole.  
This is an advantageous trait because it means the tests exercise the system the same as a user would, it can't   
use any special workarounds that could make a test pass, but not actually prove what you need to prove. This is  
similar to the principle of preferring your unit test files to live inside a separate test package, for example,  
package mypkg_test rather than package mypkg.   
#### Benefits of acceptance tests  
- When they pass, you know your entire system behaves how you want it to.
- They are more accurate, quicker, and require less effort than manual testing.
- When written well, they act as accurate, verified documentation of your system. It doesn't fall into the trap
of documentation that diverges from the real behaviour of the system.
- No mocking! It's all real.
#### Potential drawbacks vs unit tests
- They are expensive to write.
- They take longer to run.
- They are dependent on the design of the system.
- When they fail, they typically don't give you a root cause, and can be difficult to debug.
- They don't give you feedback on the internal quality of your system. You could write total garbage and still make an acceptance test pass.
- Not all scenarios are practical to exercise due to the black-box nature.
For this reason, it is foolish to only rely on acceptance tests. They do not have many of the qualities unit 
tests have, and a system with a large number of acceptance tests will tend to suffer in terms of maintenance costs
and poor lead time.
Lead time?
Lead time refers to how long it takes from a commit being merged into your main branch to it being deployed in 
production. This number can vary from weeks and even months for some teams to a matter of minutes. Again, at $WORK,
we value DORA's findings and want to keep our lead time to under 10 minutes.  
A balanced testing approach is required for a reliable system with excellent lead time, and this is usually 
described in terms of the [Test Pyramid](https://martinfowler.com/articles/practical-test-pyramid.html).
## Further Reading
[io.fs](https://benjamincongdon.me/blog/2021/01/21/A-Tour-of-Go-116s-iofs-package/)
[John Calhoun's 'Learn Web Development with Go'](https://www.calhoun.io/intro-to-templates-p1-contextual-encoding/) has a number of excellent articles on templating.
[Hotwire](https://hotwired.dev/) - You can use these techniques to create Hotwire web applications. It has 
been built by Basecamp who are primarily a Ruby on Rails shop, but because it is server-side, we can use it 
with Go  
[Decorator pattern](https://en.wikipedia.org/wiki/Decorator_pattern)  
