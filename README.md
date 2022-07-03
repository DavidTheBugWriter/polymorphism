# polymorphism
This is really a gist. I had the idea of using several databases simultaneously within a project: I mght use 
Redis as a cache and a SQL and a NoSQL. This might actually be an architecture 'smell' in some instances 
but I became interested in how to do it and flailed a bit. Most examples were just very basic but I couldn't
figure out, or find online, an example of how to use polymorphic interfaces for a generic database initialiser 
function that could use handle two difference DB structures with interface(s).


It turns out that the receiver functions InitSQL/InitNoSQL needs to return the interface 'Table' not a 
struct pointer (i.e. *SQL/*NoSQL). 

I hope this is a productive hint to other Golang users.
