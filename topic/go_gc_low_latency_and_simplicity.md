# The Go Blog: Go GC: Prioritizing low latency and simplicity

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

In a tri-color collector, every object is either white, grey, or black and we view the heap as a graph of connected 
objects. At the start of a GC cycle all objects are white. The GC visits all roots, which are objects directly accessible
by the application such as globals and things on the stack, and colors these grey. The GC then chooses a grey object,
blackens it, and then scans it for pointers to other objects. When this scan finds a pointer to a white object, it turns
that object grey. This process repeats until there are no more grey objects. At this point, white objects are known to 
be unreachable and can be reused.

Of course the devil is in the details. When do we start a GC cycle? What metrics do we use to make that decision? How 
should the GC interact with the Go scheduler? How do we pause a mutator thread long enough to scan its stack?  How do 
we represent white, grey, and black so we can efficiently find and scan grey objects? How do we know where the roots 
are? How do we know where in an object pointers are located? How do we minimize memory fragmentation? How do we deal 
with cache performance issues? How big should the heap be? And on and on, some related to allocation, some to finding 
reachable objects, some related to scheduling, but many related to performance. Low-level discussions of each of these
areas are beyond the scope of this blog post.
