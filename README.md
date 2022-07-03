# Go language polymorphism with interfaces

This is really a gist rather than a project. 

I had the idea of using several databases simultaneously within a project: I mght use 
Redis as a cache and a SQL and a NoSQL. This might actually be an architecture 'smell' in some instances 
but I became interested in how to do it and flailed a bit. Most examples were just very basic but I couldn't
figure out, or find online, an example of how to use polymorphic interfaces for a generic database initialiser 
function that could use handle two difference DB structures with interface(s).


It turns out that the receiver functions InitSQL/InitNoSQL (poly.go lines 39,43,47) needs to return the interface 'Table' not a struct pointer (i.e. *SQL/*NoSQL). 

I hope this is a productive hint to other Golang users.
